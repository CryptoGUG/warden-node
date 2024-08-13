import { chains } from "chain-registry";
import * as ethers from "ethers";
// import ethers from "ethers";
import { fromBech32, fromHex, toBech32 } from "@cosmjs/encoding";
import { IWeb3Wallet } from "@walletconnect/web3wallet";
import type { PendingRequestTypes, ProposalTypes } from "@walletconnect/types";

import {
	getSdkError,
	buildApprovedNamespaces,
	BuildApprovedNamespacesParams,
} from "@walletconnect/utils";
import { AddressType } from "@wardenprotocol/wardenjs/codegen/warden/warden/v1beta3/key";
import { base64FromBytes } from "@wardenprotocol/wardenjs/codegen/helpers";
import type { AddressResponse } from "@wardenprotocol/wardenjs/codegen/warden/warden/v1beta3/query";

import { getClient } from "@/hooks/useClient";
import type { CommonActions } from "@/utils/common";
import { RemoteMessageType, type RemoteState } from "./types";
import { useEthereumTx } from "@/hooks/useEthereumTx";
import { getProviderByChainId } from "@/lib/eth";
import { env } from "@/env";
import useRequestSignature from "@/hooks/useRequestSignature";
import { fixAddress } from "../modals/ReceiveAssets";
import { COSMOS_CHAINS } from "../assets/hooks";

export function decodeRemoteMessage(
	message: Uint8Array,
): CommonActions<RemoteState> {
	const type = message[0];
	const data = message.slice(1);

	switch (type) {
		case RemoteMessageType.Ready:
			return { type: "ready", payload: Boolean(data[0]) };
		case RemoteMessageType.Data:
			return { type: "data", payload: data };
		case RemoteMessageType.Success:
			return {
				type: "metadata",
				payload: Buffer.from(data).toString("utf-8"),
			};
		case RemoteMessageType.Error:
			return {
				type: "error",
				payload: Buffer.from(data).toString("utf-8"),
			};
		default:
			throw new Error(`Unknown message type: ${type}`);
	}
}

export function encodeRemoteMessage(
	action: CommonActions<RemoteState>,
): Uint8Array {
	switch (action.type) {
		case "ready":
			const ready = action.payload as boolean;
			return new Uint8Array([RemoteMessageType.Ready, Number(ready)]);
		case "data":
			const data = action.payload as Uint8Array;
			return new Uint8Array([RemoteMessageType.Data, ...data]);
		case "metadata":
			const metadata = action.payload as string;
			return new Uint8Array([
				RemoteMessageType.Success,
				...Buffer.from(metadata, "utf-8"),
			]);
		case "error":
			const error = action.payload as string;

			return new Uint8Array([
				RemoteMessageType.Error,
				...Buffer.from(error, "utf-8"),
			]);
		default:
			throw new Error(`Action type not implemented: ${action.type}`);
	}
}

export async function rejectSession(w: IWeb3Wallet, id: number) {
	try {
		const session = await w.rejectSession({
			id,
			reason: getSdkError("USER_REJECTED_METHODS"),
		});
		console.log("session proposal rejected. Session:", session);
	} catch (e) {
		console.error("Failed to reject session", e);
	}
}

export async function approveSession(
	w: IWeb3Wallet,
	spaceId: string,
	proposal: ProposalTypes.Struct,
	addresses?: AddressResponse[],
) {
	const { id, relays, optionalNamespaces, requiredNamespaces } = proposal;

	const requestedNamespaces: BuildApprovedNamespacesParams["supportedNamespaces"] =
		{
			...optionalNamespaces,
			...requiredNamespaces,
			// fixme types
		} as any;

	if (requestedNamespaces.eip155) {
		const eth = addresses?.filter(
			(a) => a.type === AddressType.ADDRESS_TYPE_ETHEREUM,
		);

		if (!eth?.length) {
			throw new Error(`No Ethereum addresses found for space ${spaceId}`);
		}

		const chains = requestedNamespaces.eip155.chains ?? [];
		const accounts: string[] = [];

		for (const chain of chains) {
			for (const address of eth) {
				accounts.push(`${chain}:${address.address}`);
			}
		}

		requestedNamespaces.eip155.accounts = accounts;
	}

	if (requestedNamespaces.cosmos) {
		const cosmos = addresses?.filter(
			(a) => a.type === AddressType.ADDRESS_TYPE_OSMOSIS,
		);

		if (!cosmos?.length) {
			throw new Error(`No Osmosis addresses found for space ${spaceId}`);
		}

		const chains = requestedNamespaces.cosmos.chains ?? [];
		const accounts: string[] = [];

		for (const chain of chains) {
			for (const address of cosmos) {
				accounts.push(`${chain}:${address.address}`);
			}
		}

		requestedNamespaces.cosmos.accounts = accounts;
	}

	const namespaces = buildApprovedNamespaces({
		proposal,
		supportedNamespaces: requestedNamespaces,
	});

	try {
		const session = await w.approveSession({
			id,
			relayProtocol: relays[0].protocol,
			namespaces,
		});

		localStorage.setItem(
			`WALLETCONNECT_SESSION_WS_${session.topic}`,
			spaceId,
		);

		console.log("session proposal approved. Session:", session);
	} catch (e) {
		console.error("Failed to approve session", e);
	}
}

async function findKeyByAddress(spaceId: string, address: string) {
	const client = await getClient();
	const queryKeys = client.warden.warden.v1beta3.keysBySpaceId;

	const res = await queryKeys({
		spaceId: BigInt(spaceId),
		deriveAddresses: [
			AddressType.ADDRESS_TYPE_ETHEREUM,
			AddressType.ADDRESS_TYPE_OSMOSIS,
		],
	});

	return res.keys?.find((key) =>
		key.addresses
			?.map((w) => w.address?.toLowerCase())
			.includes(address.toLowerCase()),
	)?.key;
}

interface EthParams {
	gas: string;
	value: string;
	from: string;
	to: string;
	data: string;
}

const isValidEthParams = (params: any): params is EthParams => {
	if (!params) {
		return false;
	}

	return (
		typeof params.gas === "string" &&
		(!params.value || typeof params.value === "string") &&
		typeof params.from === "string" &&
		typeof params.to === "string" &&
		typeof params.data === "string"
	);
};

export async function approveRequest({
	w,
	req,
	eth,
	cosm,
}: {
	w?: IWeb3Wallet | null;
	req: PendingRequestTypes.Struct;
	eth: ReturnType<typeof useEthereumTx>;
	cosm: ReturnType<typeof useRequestSignature>;
}) {
	if (!w) {
		throw new Error("WalletConnect not initialized");
	}

	const topic = req.topic;

	try {
		// todo rename
		const wsAddr = localStorage.getItem(
			`WALLETCONNECT_SESSION_WS_${topic}`,
		);

		if (!wsAddr) {
			throw new Error(
				`Unknown space address for session topic: ${topic}`,
			);
		}

		let response = null;
		const [chainType, chainId] = req.params.chainId.split(":");

		switch (req.params.request.method) {
			case "personal_sign": {
				if (chainType !== "eip155") {
					throw new Error(`Unsupported chain type: ${chainType}`);
				}

				const address = req.params.request.params[1];
				const key = await findKeyByAddress(wsAddr, address);

				if (!key) {
					console.error("Unknown address", address);
					return;
				}

				// prepare message
				const msg = fromHex(req.params.request.params[0].slice(2));
				const text = new TextDecoder().decode(msg);

				const message =
					"\x19Ethereum Signed Message:\n" + text.length + text;

				const hash = ethers.keccak256(
					Uint8Array.from(Buffer.from(message, "utf-8")),
				);

				// send signature request to Warden Protocol and wait response
				const sig = await eth.signRaw(key.id, ethers.getBytes(hash));

				if (!sig) {
					return;
				}

				response = {
					result: ethers.hexlify(sig),
					id: req.id,
					jsonrpc: "2.0",
				};

				break;
			}
			case "eth_sendTransaction": {
				if (chainType !== "eip155") {
					throw new Error(`Unsupported chain type: ${chainType}`);
				}

				const txParam = req.params.request.params[0];
				console.log({ txParam });

				if (!isValidEthParams(txParam)) {
					throw new Error("Invalid transaction parameters");
				}

				const key = await findKeyByAddress(wsAddr, txParam.from);

				if (!key) {
					throw new Error(`Unknown address ${txParam.from}`);
				}

				const provider = getProviderByChainId(chainId);
				const nonce = await provider.getTransactionCount(txParam.from);
				const feeData = await provider.getFeeData();

				const tx = ethers.Transaction.from({
					type: 2,
					chainId: Number(chainId),
					maxPriorityFeePerGas: feeData.maxPriorityFeePerGas,
					maxFeePerGas: feeData.maxFeePerGas,
					nonce,
					to: txParam.to,
					value: txParam.value,
					gasLimit: txParam.gas,
					data: txParam.data,
				});

				const signedTx = await eth.signEthereumTx(key.id, tx);

				if (!signedTx) {
					return;
				}

				await provider.broadcastTransaction(signedTx.serialized);

				response = {
					result: signedTx.hash,
					id: req.id,
					jsonrpc: "2.0",
				};
				break;
			}
			case "eth_signTypedData_v4": {
				const from = req.params.request.params[0];
				const key = await findKeyByAddress(wsAddr, from);
				if (!key) {
					throw new Error(`Unknown address ${from}`);
				}
				const data = JSON.parse(req.params.request.params[1]);

				// ethers.TypedDataEncoder tries to determine the
				// primaryType automatically, but it fails because we
				// have multiple "roots" in the DAG: one is
				// EIP712Domain, one is specified in
				// `data.primaryType` (e.g. "PermitSingle" for
				// Uniswap, "dYdX", ...).
				// I split the types into two objects and manually
				// create two different encoders.
				const typesWithoutDomain = { ...data.types };
				delete typesWithoutDomain.EIP712Domain;
				const domainEncoder = new ethers.TypedDataEncoder({
					EIP712Domain: data.types.EIP712Domain,
				});
				const messageEncoder = new ethers.TypedDataEncoder(
					typesWithoutDomain,
				);

				// In short, we need to sign:
				//   sign(keccak256("\x19\x01" ‖ domainSeparator ‖ hashStruct(message)))
				//
				// See EIP-712 for the definition of the message to be signed.
				// https://eips.ethereum.org/EIPS/eip-712#definition-of-domainseparator
				const domainSeparator = domainEncoder.hashStruct(
					"EIP712Domain",
					data.domain,
				);
				const message = messageEncoder.hashStruct(
					data.primaryType,
					data.message,
				);
				const toSign = ethers.keccak256(
					ethers.concat([
						ethers.getBytes("0x1901"),
						ethers.getBytes(domainSeparator),
						ethers.getBytes(message),
					]),
				);

				const signature = await eth.signRaw(
					key.id,
					ethers.getBytes(toSign),
				);

				if (!signature) {
					return;
				}

				response = {
					result: ethers.hexlify(signature),
					id: req.id,
					jsonrpc: "2.0",
				};
				break;
			}
			case "cosmos_getAccounts": {
				const { chainId: _chainId } = req.params;
				const [chainType, chainId] = _chainId.split(":");

				if (chainType !== "cosmos") {
					throw new Error(`Unsupported chain type: ${chainType}`);
				}

				const chain = chains.find((c) => c.chain_id === chainId);

				if (!chain) {
					throw new Error(`Unknown chain id ${chainId}`);
				}

				const client = await getClient();
				const queryKeys = client.warden.warden.v1beta3.keysBySpaceId;

				const res = await queryKeys({
					spaceId: BigInt(wsAddr),
					deriveAddresses: [AddressType.ADDRESS_TYPE_OSMOSIS],
				});

				const addresses = res.keys.flatMap((key) =>
					key.addresses.map((addr) => ({
						...fixAddress(addr, chain.chain_name),
						publicKey: key.key.publicKey,
					})),
				);

				response = {
					result: addresses?.map(({ address, publicKey }) => ({
						address,
						algo: "secp256k1",
						pubkey: base64FromBytes(publicKey),
					})),
					id: req.id,
					jsonrpc: "2.0",
				};

				break;
			}
			case "cosmos_signAmino": {
				const { request, chainId: _chainId } = req.params;
				const [chainType, chainId] = _chainId.split(":");

				if (chainType !== "cosmos") {
					throw new Error(`Unsupported chain type: ${chainType}`);
				}

				const chain = chains.find((c) => c.chain_id === chainId);

				if (!chain) {
					throw new Error(`Unknown chain id ${chainId}`);
				}

				const {
					signerAddress,
					signDoc,
				}: {
					signerAddress: string;
					signDoc: any;
				} = request.params;

				const key = await findKeyByAddress(
					wsAddr,
					toBech32(
						chain.bech32_prefix,
						fromBech32(signerAddress).data,
					),
				);

				if (!key) {
					throw new Error(`Unknown address ${signerAddress}`);
				}

				const feeAmount = COSMOS_CHAINS.find(
					({ chainName }) => (chainName = chain.chain_name),
				)?.feeAmount;

				// fixme testnet hack
				if (signDoc?.fee) {
					if (!signDoc.fee?.amount.length && feeAmount) {
						// todo look into fee_tokens
						const denom = chain.fees?.fee_tokens[0]?.denom;

						if (!denom) {
							throw new Error("Missing fee denom");
						}

						signDoc.fee.amount = [{ denom, amount: feeAmount }];
					}
				}

				let signature = await cosm.requestSignature(
					key.id,
					[env.aminoAnalyzerContract],
					Uint8Array.from(
						JSON.stringify(signDoc)
							.split("")
							.map((c) => c.charCodeAt(0)),
					),
				);

				if (signature?.length === 65) {
					signature = signature.slice(0, -1);
				}

				if (signature?.length !== 64) {
					throw new Error("unexpected signature length");
				}

				response = {
					jsonrpc: "2.0",
					id: req.id,
					result: {
						signed: signDoc,
						signature: {
							signature: base64FromBytes(signature),
							pub_key: {
								type: "tendermint/PubKeySecp256k1", // hardcoded value
								value: base64FromBytes(key.publicKey),
							},
						},
					},
				};

				break;
			}
			default:
				throw new Error(
					`Unknown or unsupported method: ${req.params.request.method}`,
				);
		}

		await w!.respondSessionRequest({ topic, response });
	} catch (error) {
		console.error(error);

		await w!.respondSessionRequest({
			topic,
			response: {
				jsonrpc: "2.0",
				id: req.id,
				error: {
					code: 1,
					message: `${error}`,
				},
			},
		});
	}
}
