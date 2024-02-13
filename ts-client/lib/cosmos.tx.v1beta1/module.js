// Generated by Ignite ignite.com/cli
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { Api } from "./rest";
import { BroadcastTxResponse } from "./types/cosmos/tx/v1beta1/service";
import { TxDecodeAminoRequest } from "./types/cosmos/tx/v1beta1/service";
import { SignDocDirectAux } from "./types/cosmos/tx/v1beta1/tx";
import { SignerInfo } from "./types/cosmos/tx/v1beta1/tx";
import { ModeInfo } from "./types/cosmos/tx/v1beta1/tx";
import { GetBlockWithTxsRequest } from "./types/cosmos/tx/v1beta1/service";
import { TxDecodeRequest } from "./types/cosmos/tx/v1beta1/service";
import { TxEncodeRequest } from "./types/cosmos/tx/v1beta1/service";
import { ModeInfo_Multi } from "./types/cosmos/tx/v1beta1/tx";
import { GetTxResponse } from "./types/cosmos/tx/v1beta1/service";
import { TxEncodeAminoResponse } from "./types/cosmos/tx/v1beta1/service";
import { Tip } from "./types/cosmos/tx/v1beta1/tx";
import { GetBlockWithTxsResponse } from "./types/cosmos/tx/v1beta1/service";
import { AuthInfo } from "./types/cosmos/tx/v1beta1/tx";
import { GetTxsEventRequest } from "./types/cosmos/tx/v1beta1/service";
import { GetTxsEventResponse } from "./types/cosmos/tx/v1beta1/service";
import { SimulateResponse } from "./types/cosmos/tx/v1beta1/service";
import { TxRaw } from "./types/cosmos/tx/v1beta1/tx";
import { SimulateRequest } from "./types/cosmos/tx/v1beta1/service";
import { GetTxRequest } from "./types/cosmos/tx/v1beta1/service";
import { Tx } from "./types/cosmos/tx/v1beta1/tx";
import { AuxSignerData } from "./types/cosmos/tx/v1beta1/tx";
import { BroadcastTxRequest } from "./types/cosmos/tx/v1beta1/service";
import { TxDecodeResponse } from "./types/cosmos/tx/v1beta1/service";
import { TxDecodeAminoResponse } from "./types/cosmos/tx/v1beta1/service";
import { TxBody } from "./types/cosmos/tx/v1beta1/tx";
import { ModeInfo_Single } from "./types/cosmos/tx/v1beta1/tx";
import { Fee } from "./types/cosmos/tx/v1beta1/tx";
import { TxEncodeResponse } from "./types/cosmos/tx/v1beta1/service";
import { TxEncodeAminoRequest } from "./types/cosmos/tx/v1beta1/service";
import { SignDoc } from "./types/cosmos/tx/v1beta1/tx";
export { BroadcastTxResponse, TxDecodeAminoRequest, SignDocDirectAux, SignerInfo, ModeInfo, GetBlockWithTxsRequest, TxDecodeRequest, TxEncodeRequest, ModeInfo_Multi, GetTxResponse, TxEncodeAminoResponse, Tip, GetBlockWithTxsResponse, AuthInfo, GetTxsEventRequest, GetTxsEventResponse, SimulateResponse, TxRaw, SimulateRequest, GetTxRequest, Tx, AuxSignerData, BroadcastTxRequest, TxDecodeResponse, TxDecodeAminoResponse, TxBody, ModeInfo_Single, Fee, TxEncodeResponse, TxEncodeAminoRequest, SignDoc };
export const registry = new Registry(msgTypes);
function getStructure(template) {
    const structure = { fields: [] };
    for (let [key, value] of Object.entries(template)) {
        let field = { name: key, type: typeof value };
        structure.fields.push(field);
    }
    return structure;
}
const defaultFee = {
    amount: [],
    gas: "200000",
};
export const txClient = ({ signer, prefix, addr } = { addr: "http://localhost:26657", prefix: "cosmos" }) => {
    return {
        async sendBroadcastTxResponse({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendBroadcastTxResponse: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.broadcastTxResponse({ value: BroadcastTxResponse.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendBroadcastTxResponse: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendTxDecodeAminoRequest({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendTxDecodeAminoRequest: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.txDecodeAminoRequest({ value: TxDecodeAminoRequest.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendTxDecodeAminoRequest: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendSignDocDirectAux({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendSignDocDirectAux: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.signDocDirectAux({ value: SignDocDirectAux.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendSignDocDirectAux: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendSignerInfo({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendSignerInfo: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.signerInfo({ value: SignerInfo.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendSignerInfo: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendModeInfo({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendModeInfo: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.modeInfo({ value: ModeInfo.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendModeInfo: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendGetBlockWithTxsRequest({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendGetBlockWithTxsRequest: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.getBlockWithTxsRequest({ value: GetBlockWithTxsRequest.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendGetBlockWithTxsRequest: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendTxDecodeRequest({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendTxDecodeRequest: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.txDecodeRequest({ value: TxDecodeRequest.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendTxDecodeRequest: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendTxEncodeRequest({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendTxEncodeRequest: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.txEncodeRequest({ value: TxEncodeRequest.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendTxEncodeRequest: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendModeInfo_Multi({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendModeInfo_Multi: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.modeInfoMulti({ value: ModeInfo_Multi.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendModeInfo_Multi: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendGetTxResponse({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendGetTxResponse: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.getTxResponse({ value: GetTxResponse.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendGetTxResponse: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendTxEncodeAminoResponse({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendTxEncodeAminoResponse: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.txEncodeAminoResponse({ value: TxEncodeAminoResponse.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendTxEncodeAminoResponse: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendTip({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendTip: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.tip({ value: Tip.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendTip: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendGetBlockWithTxsResponse({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendGetBlockWithTxsResponse: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.getBlockWithTxsResponse({ value: GetBlockWithTxsResponse.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendGetBlockWithTxsResponse: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendAuthInfo({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendAuthInfo: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.authInfo({ value: AuthInfo.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendAuthInfo: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendGetTxsEventRequest({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendGetTxsEventRequest: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.getTxsEventRequest({ value: GetTxsEventRequest.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendGetTxsEventRequest: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendGetTxsEventResponse({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendGetTxsEventResponse: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.getTxsEventResponse({ value: GetTxsEventResponse.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendGetTxsEventResponse: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendSimulateResponse({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendSimulateResponse: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.simulateResponse({ value: SimulateResponse.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendSimulateResponse: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendTxRaw({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendTxRaw: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.txRaw({ value: TxRaw.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendTxRaw: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendSimulateRequest({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendSimulateRequest: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.simulateRequest({ value: SimulateRequest.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendSimulateRequest: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendGetTxRequest({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendGetTxRequest: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.getTxRequest({ value: GetTxRequest.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendGetTxRequest: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendTx({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendTx: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.tx({ value: Tx.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendTx: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendAuxSignerData({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendAuxSignerData: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.auxSignerData({ value: AuxSignerData.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendAuxSignerData: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendBroadcastTxRequest({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendBroadcastTxRequest: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.broadcastTxRequest({ value: BroadcastTxRequest.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendBroadcastTxRequest: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendTxDecodeResponse({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendTxDecodeResponse: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.txDecodeResponse({ value: TxDecodeResponse.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendTxDecodeResponse: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendTxDecodeAminoResponse({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendTxDecodeAminoResponse: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.txDecodeAminoResponse({ value: TxDecodeAminoResponse.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendTxDecodeAminoResponse: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendTxBody({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendTxBody: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.txBody({ value: TxBody.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendTxBody: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendModeInfo_Single({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendModeInfo_Single: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.modeInfoSingle({ value: ModeInfo_Single.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendModeInfo_Single: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendFee({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendFee: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.fee({ value: Fee.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendFee: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendTxEncodeResponse({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendTxEncodeResponse: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.txEncodeResponse({ value: TxEncodeResponse.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendTxEncodeResponse: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendTxEncodeAminoRequest({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendTxEncodeAminoRequest: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.txEncodeAminoRequest({ value: TxEncodeAminoRequest.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendTxEncodeAminoRequest: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendSignDoc({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendSignDoc: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.signDoc({ value: SignDoc.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendSignDoc: Could not broadcast Tx: ' + e.message);
            }
        },
        broadcastTxResponse({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.BroadcastTxResponse", value: BroadcastTxResponse.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:BroadcastTxResponse: Could not create message: ' + e.message);
            }
        },
        txDecodeAminoRequest({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.TxDecodeAminoRequest", value: TxDecodeAminoRequest.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:TxDecodeAminoRequest: Could not create message: ' + e.message);
            }
        },
        signDocDirectAux({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.SignDocDirectAux", value: SignDocDirectAux.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:SignDocDirectAux: Could not create message: ' + e.message);
            }
        },
        signerInfo({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.SignerInfo", value: SignerInfo.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:SignerInfo: Could not create message: ' + e.message);
            }
        },
        modeInfo({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.ModeInfo", value: ModeInfo.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:ModeInfo: Could not create message: ' + e.message);
            }
        },
        getBlockWithTxsRequest({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.GetBlockWithTxsRequest", value: GetBlockWithTxsRequest.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:GetBlockWithTxsRequest: Could not create message: ' + e.message);
            }
        },
        txDecodeRequest({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.TxDecodeRequest", value: TxDecodeRequest.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:TxDecodeRequest: Could not create message: ' + e.message);
            }
        },
        txEncodeRequest({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.TxEncodeRequest", value: TxEncodeRequest.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:TxEncodeRequest: Could not create message: ' + e.message);
            }
        },
        modeInfoMulti({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.ModeInfo_Multi", value: ModeInfo_Multi.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:ModeInfo_Multi: Could not create message: ' + e.message);
            }
        },
        getTxResponse({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.GetTxResponse", value: GetTxResponse.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:GetTxResponse: Could not create message: ' + e.message);
            }
        },
        txEncodeAminoResponse({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.TxEncodeAminoResponse", value: TxEncodeAminoResponse.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:TxEncodeAminoResponse: Could not create message: ' + e.message);
            }
        },
        tip({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.Tip", value: Tip.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:Tip: Could not create message: ' + e.message);
            }
        },
        getBlockWithTxsResponse({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.GetBlockWithTxsResponse", value: GetBlockWithTxsResponse.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:GetBlockWithTxsResponse: Could not create message: ' + e.message);
            }
        },
        authInfo({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.AuthInfo", value: AuthInfo.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:AuthInfo: Could not create message: ' + e.message);
            }
        },
        getTxsEventRequest({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.GetTxsEventRequest", value: GetTxsEventRequest.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:GetTxsEventRequest: Could not create message: ' + e.message);
            }
        },
        getTxsEventResponse({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.GetTxsEventResponse", value: GetTxsEventResponse.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:GetTxsEventResponse: Could not create message: ' + e.message);
            }
        },
        simulateResponse({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.SimulateResponse", value: SimulateResponse.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:SimulateResponse: Could not create message: ' + e.message);
            }
        },
        txRaw({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.TxRaw", value: TxRaw.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:TxRaw: Could not create message: ' + e.message);
            }
        },
        simulateRequest({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.SimulateRequest", value: SimulateRequest.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:SimulateRequest: Could not create message: ' + e.message);
            }
        },
        getTxRequest({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.GetTxRequest", value: GetTxRequest.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:GetTxRequest: Could not create message: ' + e.message);
            }
        },
        tx({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.Tx", value: Tx.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:Tx: Could not create message: ' + e.message);
            }
        },
        auxSignerData({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.AuxSignerData", value: AuxSignerData.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:AuxSignerData: Could not create message: ' + e.message);
            }
        },
        broadcastTxRequest({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.BroadcastTxRequest", value: BroadcastTxRequest.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:BroadcastTxRequest: Could not create message: ' + e.message);
            }
        },
        txDecodeResponse({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.TxDecodeResponse", value: TxDecodeResponse.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:TxDecodeResponse: Could not create message: ' + e.message);
            }
        },
        txDecodeAminoResponse({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.TxDecodeAminoResponse", value: TxDecodeAminoResponse.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:TxDecodeAminoResponse: Could not create message: ' + e.message);
            }
        },
        txBody({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.TxBody", value: TxBody.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:TxBody: Could not create message: ' + e.message);
            }
        },
        modeInfoSingle({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.ModeInfo_Single", value: ModeInfo_Single.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:ModeInfo_Single: Could not create message: ' + e.message);
            }
        },
        fee({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.Fee", value: Fee.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:Fee: Could not create message: ' + e.message);
            }
        },
        txEncodeResponse({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.TxEncodeResponse", value: TxEncodeResponse.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:TxEncodeResponse: Could not create message: ' + e.message);
            }
        },
        txEncodeAminoRequest({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.TxEncodeAminoRequest", value: TxEncodeAminoRequest.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:TxEncodeAminoRequest: Could not create message: ' + e.message);
            }
        },
        signDoc({ value }) {
            try {
                return { typeUrl: "/cosmos.tx.v1beta1.SignDoc", value: SignDoc.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:SignDoc: Could not create message: ' + e.message);
            }
        },
    };
};
export const queryClient = ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseURL: addr });
};
class SDKModule {
    constructor(client) {
        this.registry = [];
        this.query = queryClient({ addr: client.env.apiURL });
        this.updateTX(client);
        this.structure = {};
        client.on('signer-changed', (signer) => {
            this.updateTX(client);
        });
    }
    updateTX(client) {
        const methods = txClient({
            signer: client.signer,
            addr: client.env.rpcURL,
            prefix: client.env.prefix ?? "cosmos",
        });
        this.tx = methods;
        for (let m in methods) {
            this.tx[m] = methods[m].bind(this.tx);
        }
    }
}
;
const IgntModule = (test) => {
    return {
        module: {
            CosmosTxV1Beta1: new SDKModule(test)
        },
        registry: msgTypes
    };
};
export default IgntModule;
