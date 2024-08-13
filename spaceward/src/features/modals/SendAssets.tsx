import clsx from "clsx";
import type { TransactionReceipt } from "ethers";
import { useCallback, useContext, useMemo, useState } from "react";
import { AddressType } from "@wardenprotocol/wardenjs/codegen/warden/warden/v1beta3/key";
import { Icons } from "@/components/ui/icons-assets";
import type { TransferParams } from "./types";
import { bigintToFixed, bigintToFloat } from "@/lib/math";
import { useEthereumTx } from "@/hooks/useEthereumTx";
import SignRequestDialog from "@/components/SignRequestDialog";
import { SignRequesterState } from "@/hooks/useRequestSignature";
import { TxBuild, buildTransaction } from "./util";
import { COSMOS_CHAINS, useAssetQueries } from "../assets/hooks";
import { useSpaceId } from "@/hooks/useSpaceId";
import useFiatConversion from "@/hooks/useFiatConversion";
import { numRestrict } from "../staking/util";
import { useKeychainSigner } from "@/hooks/useKeychainSigner";
import { NetworkIcons, TokenIcons } from "@/components/ui/icons-crypto";
import { AssetPlaceholder } from "../assets/AssetRow";
import { validateAddress } from "../intents/AddAddressModal";
import { SigningStargateClient } from "@cosmjs/stargate";
import { walletContext } from "@cosmos-kit/react-lite";
import { useModalState } from "./state";
import KeySelector from "./KeySelector";

export default function SendAssetsModal({
	// address,
	chainName,
	token,
	type,
	keyResponse: key,
}: TransferParams) {
	const { walletManager } = useContext(walletContext);
	const { setData: setModal } = useModalState();
	const { formatter, fiatConversion } = useFiatConversion();

	const { spaceId } = useSpaceId();
	const { queryBalances } = useAssetQueries(spaceId);

	const results = queryBalances
		.filter((result) => result.data?.key?.key?.id === key?.key?.id)
		.flatMap(({ data, refetch }) => {
			return (data?.results ?? []).map((item) => ({ ...item, refetch }));
		});

	const noAssets = results.every((result) => !result.balance);

	const _selectedToken = useMemo(
		() =>
			results.find(
				(data) => data.chainName === chainName && data.token === token,
			) ?? results[0],
		[token, chainName, results],
	);

	// fixme types
	const selectedToken = _selectedToken as typeof _selectedToken | undefined;

	const [pending, setPending] = useState(false);
	const [receipt, setReceipt] = useState<TransactionReceipt>();
	const [amount, setAmount] = useState("");
	const [destinationAddress, setDestinationAddress] = useState("");
	const amountNum = parseFloat(amount);

	const amountWarning =
		amountNum && selectedToken
			? Number.isFinite(amountNum)
				? bigintToFloat(selectedToken.balance, selectedToken.decimals) <
					amountNum
				: true
			: false;

	const addressWarning =
		destinationAddress && selectedToken
			? !validateAddress(destinationAddress, [
					selectedToken.type.startsWith("eip155:") ? "eth" : "bech32",
				]).ok
			: false;

	const { signEthereumTx, ...eth } = useEthereumTx();

	const keys = useMemo(() => (key ? [key] : []), [key]);

	const { signer, ...cosm } = useKeychainSigner({
		keys,
		chainName: selectedToken?.chainName,
	});

	const req =
		type === AddressType.ADDRESS_TYPE_ETHEREUM
			? eth
			: type === AddressType.ADDRESS_TYPE_OSMOSIS
				? cosm
				: null;

	const state =
		(req?.state === SignRequesterState.IDLE && pending
			? SignRequesterState.BROADCAST_SIGN_REQUEST
			: req?.state) ?? SignRequesterState.IDLE;

	const error = req ? req.error : "Wrong address type"; // todo tx broadcast error

	const reset = useCallback(() => {
		req?.reset();
		setReceipt(undefined);

		if (receipt) {
			setModal({ type: undefined, params: undefined });
		}
	}, [receipt, req]);

	async function submit() {
		if (!key || !selectedToken) {
			return;
		}

		const { address, chainName } = selectedToken;
		setPending(true);

		try {
			const txBuild = await buildTransaction({
				item: selectedToken,
				from: address,
				to: destinationAddress,
				amount,
			});

			if (txBuild.type === "eth") {
				const { provider, tx } = txBuild;
				const signedTx = await signEthereumTx(key.key.id, tx);

				if (!signedTx) {
					throw new Error("Failed to sign transaction");
				}

				const res = await provider.broadcastTransaction(
					signedTx.serialized,
				);

				const receipt = await provider.waitForTransaction(res.hash);

				if (!receipt) {
					throw new Error("Failed to get transaction receipt");
				}

				setReceipt(receipt);
			} else if (txBuild.type === "cosmos") {
				const { fee, msgs } = txBuild as TxBuild<"cosmos">;

				const chain = COSMOS_CHAINS.find(
					(item) => item.chainName === chainName,
				);
				/* const repo = walletManager.getWalletRepo(chainName);
				repo.activate();
				const rpc = await repo.getRpcEndpoint(); */

				const endpoint =
					chain?.rpc ??
					`https://rpc.cosmos.directory/${chainName}`; /* rpc
					? typeof rpc === "string"
						? rpc
						: rpc.url
					: `https://rpc.cosmos.directory/${chainName}`*/

				const client = await SigningStargateClient.connectWithSigner(
					endpoint,
					signer,
				);

				const res = await client.signAndBroadcast(address, msgs, fee);

				if (res.code) {
					console.error(res);
					throw new Error("tx failed");
				}

				// fixme types
				setReceipt(res as any);
			} else {
				throw new Error(`not implemented: ${txBuild.type}`);
			}
		} catch (err) {
			console.error(err);
		}

		await selectedToken.refetch();
		setPending(false);
	}

	function pasteFromClipboard(e: React.MouseEvent<HTMLButtonElement>) {
		e.preventDefault();

		if (!navigator?.clipboard) {
			console.error("Clipboard API not available");
			return;
		}

		navigator.clipboard
			.readText()
			.then((text) => {
				setDestinationAddress(text);
			})
			.catch((err) => {
				console.error("Failed to read clipboard contents: ", err);
			});
	}

	const maxAmount = selectedToken
		? bigintToFixed(selectedToken.balance, {
				decimals: selectedToken.decimals,
			})
		: "0";

	const Token = TokenIcons[selectedToken?.token ?? ""] ?? AssetPlaceholder;
	const Network =
		NetworkIcons[selectedToken?.chainName ?? ""] ?? AssetPlaceholder;

	return (
		<div className="max-w-[520px] w-[520px] text-center tracking-wide pb-5">
			<div className="font-bold text-5xl mb-12 leading-[56px]">Send</div>

			<form action="" onSubmit={(e) => e.preventDefault()}>
				<div>
					<KeySelector
						currentKey={key}
						token={selectedToken?.token}
						className="relative mb-8 z-50"
					/>

					{noAssets && (
						<div className="flex rounded-lg px-4 h-[56px] bg-negative-secondary mb-8 items-center gap-3">
							<Icons.redInfo />
							No available assets in this key. Select an another
							key
						</div>
					)}

					<div
						className={clsx(
							"relative z-40 mb-[2px] text-left rounded-lg rounded-bl-none rounded-br-none p-6 border-[1px] border-border-quaternary",
							noAssets && "pointer-events-none opacity-30	",
							amountWarning && "bg-negative-secondary",
							!amountWarning && "bg-secondary-bg",
						)}
					>
						<div className="text-label-secondary mb-3">
							You&apos;re sending
						</div>

						<div className="relative flex items-center justify-between">
							<input
								className={clsx(
									"block w-full h-10 bg-transparent outline-none foces:outline-none text-[32px] font-bold",
								)}
								id="address"
								onChange={(e) =>
									setAmount(numRestrict(e.target.value))
								}
								value={amount}
								placeholder="Amount"
							/>

							<div className="relative z-40">
								<div
									onClick={setModal.bind(null, {
										type: "select-asset",
										params: {
											keyResponse: key,
										},
									})}
									className="cursor-pointer h-[32px] bg-fill-quaternary rounded-[20px] p-1 pr-2 flex items-center gap-[6px]"
								>
									<div className="relative w-6 h-6 ">
										<Token className="w-6 h-6 object-contain" />
										<Network className="absolute bottom-[-1px] right-[-1px] w-[10px] h-[10px]" />
									</div>
									{selectedToken?.token}
									<Icons.chevronDown
										className={clsx(
											"ml-auto invert dark:invert-0",
										)}
									/>
								</div>
							</div>
						</div>

						<div className="flex mt-1 justify-between">
							<div className="text-label-secondary opacity-50 text-xs">
								{/* todo useFiatConversion hook */}
								{formatter.format(
									(amount ? parseFloat(amount) : 0) *
										(fiatConversion && selectedToken
											? bigintToFloat(
													(selectedToken.price *
														BigInt(10) **
															BigInt(
																fiatConversion.decimals,
															)) /
														fiatConversion.value,
													selectedToken.priceDecimals,
												)
											: 0),
								)}
							</div>
							<div
								onClick={() => {
									setAmount(maxAmount);
								}}
								className={clsx(
									"text-xs cursor-pointer",
									amountWarning && "text-negative",
									!amountWarning && "text-pixel-pink",
								)}
							>
								Max:{maxAmount}{" "}
							</div>
						</div>
					</div>

					<div
						className={clsx(
							"relative z-40 mb-8 text-left rounded-lg flex items-center justify-between rounded-tl-none rounded-tr-none p-6 border-[1px] border-border-quaternary",
							noAssets && "pointer-events-none opacity-30	",
							addressWarning && "bg-negative-secondary",
							!addressWarning && "bg-secondary-bg",
						)}
					>
						{destinationAddress && (
							<label
								className="text-label-secondary text-xs absolute top-[16px] left-5"
								htmlFor="destinationAddress"
							>
								To address
							</label>
						)}
						<input
							className={clsx(
								"block w-full bg-transparent outline-none foces:outline-none",
								destinationAddress && "translate-y-[8px]",
							)}
							id="destinationAddress"
							onChange={(e) =>
								setDestinationAddress(e.target.value)
							}
							value={destinationAddress}
							placeholder="To..."
						/>
						{destinationAddress ? (
							<>
								{addressWarning ? (
									<Icons.alert className="ml-4 mr-4" />
								) : null}
								<button
									className="text-label-secondary font-semibold"
									onClick={() => setDestinationAddress("")}
								>
									<img src="/images/x.svg" alt="" />
								</button>
							</>
						) : (
							<button
								onClick={pasteFromClipboard}
								className="text-label-secondary font-semibold"
							>
								Paste
							</button>
						)}

						{addressWarning && (
							<div className="absolute left-0 -bottom-1 translate-y-full text-negative text-xs">
								Add correct address
							</div>
						)}
					</div>
				</div>
			</form>

			<div className="mt-12 pt-6">
				<button
					onClick={submit}
					className={clsx(
						`bg-foreground h-14 flex items-center justify-center w-full font-semibold text-background hover:bg-accent transition-all duration-200`,
						(amount == "" ||
							destinationAddress == "" ||
							noAssets) &&
							"pointer-events-none opacity-50",
					)}
				>
					Send
				</button>
			</div>

			<SignRequestDialog
				state={state}
				reset={reset}
				error={error}
				pending={pending}
				step={{
					title: "Broadcast",
					description: `Transaction was broadcasted to ${selectedToken?.chainName} network`,
				}}
			/>
		</div>
	);
}
