import _m0 from "protobufjs/minimal";
import { Any } from "../../../google/protobuf/any";
import { Duration } from "../../../google/protobuf/duration";
import { Coin } from "../../base/v1beta1/coin";
export declare const protobufPackage = "cosmos.gov.v1";
/** Since: cosmos-sdk 0.46 */
/** VoteOption enumerates the valid vote options for a given governance proposal. */
export declare enum VoteOption {
    /** VOTE_OPTION_UNSPECIFIED - VOTE_OPTION_UNSPECIFIED defines a no-op vote option. */
    VOTE_OPTION_UNSPECIFIED = 0,
    /** VOTE_OPTION_YES - VOTE_OPTION_YES defines a yes vote option. */
    VOTE_OPTION_YES = 1,
    /** VOTE_OPTION_ABSTAIN - VOTE_OPTION_ABSTAIN defines an abstain vote option. */
    VOTE_OPTION_ABSTAIN = 2,
    /** VOTE_OPTION_NO - VOTE_OPTION_NO defines a no vote option. */
    VOTE_OPTION_NO = 3,
    /** VOTE_OPTION_NO_WITH_VETO - VOTE_OPTION_NO_WITH_VETO defines a no with veto vote option. */
    VOTE_OPTION_NO_WITH_VETO = 4,
    UNRECOGNIZED = -1
}
export declare function voteOptionFromJSON(object: any): VoteOption;
export declare function voteOptionToJSON(object: VoteOption): string;
/** ProposalStatus enumerates the valid statuses of a proposal. */
export declare enum ProposalStatus {
    /** PROPOSAL_STATUS_UNSPECIFIED - PROPOSAL_STATUS_UNSPECIFIED defines the default proposal status. */
    PROPOSAL_STATUS_UNSPECIFIED = 0,
    /**
     * PROPOSAL_STATUS_DEPOSIT_PERIOD - PROPOSAL_STATUS_DEPOSIT_PERIOD defines a proposal status during the deposit
     * period.
     */
    PROPOSAL_STATUS_DEPOSIT_PERIOD = 1,
    /**
     * PROPOSAL_STATUS_VOTING_PERIOD - PROPOSAL_STATUS_VOTING_PERIOD defines a proposal status during the voting
     * period.
     */
    PROPOSAL_STATUS_VOTING_PERIOD = 2,
    /**
     * PROPOSAL_STATUS_PASSED - PROPOSAL_STATUS_PASSED defines a proposal status of a proposal that has
     * passed.
     */
    PROPOSAL_STATUS_PASSED = 3,
    /**
     * PROPOSAL_STATUS_REJECTED - PROPOSAL_STATUS_REJECTED defines a proposal status of a proposal that has
     * been rejected.
     */
    PROPOSAL_STATUS_REJECTED = 4,
    /**
     * PROPOSAL_STATUS_FAILED - PROPOSAL_STATUS_FAILED defines a proposal status of a proposal that has
     * failed.
     */
    PROPOSAL_STATUS_FAILED = 5,
    UNRECOGNIZED = -1
}
export declare function proposalStatusFromJSON(object: any): ProposalStatus;
export declare function proposalStatusToJSON(object: ProposalStatus): string;
/** WeightedVoteOption defines a unit of vote for vote split. */
export interface WeightedVoteOption {
    /** option defines the valid vote options, it must not contain duplicate vote options. */
    option: VoteOption;
    /** weight is the vote weight associated with the vote option. */
    weight: string;
}
/**
 * Deposit defines an amount deposited by an account address to an active
 * proposal.
 */
export interface Deposit {
    /** proposal_id defines the unique id of the proposal. */
    proposalId: number;
    /** depositor defines the deposit addresses from the proposals. */
    depositor: string;
    /** amount to be deposited by depositor. */
    amount: Coin[];
}
/** Proposal defines the core field members of a governance proposal. */
export interface Proposal {
    /** id defines the unique id of the proposal. */
    id: number;
    /** messages are the arbitrary messages to be executed if the proposal passes. */
    messages: Any[];
    /** status defines the proposal status. */
    status: ProposalStatus;
    /**
     * final_tally_result is the final tally result of the proposal. When
     * querying a proposal via gRPC, this field is not populated until the
     * proposal's voting period has ended.
     */
    finalTallyResult: TallyResult | undefined;
    /** submit_time is the time of proposal submission. */
    submitTime: Date | undefined;
    /** deposit_end_time is the end time for deposition. */
    depositEndTime: Date | undefined;
    /** total_deposit is the total deposit on the proposal. */
    totalDeposit: Coin[];
    /** voting_start_time is the starting time to vote on a proposal. */
    votingStartTime: Date | undefined;
    /** voting_end_time is the end time of voting on a proposal. */
    votingEndTime: Date | undefined;
    /**
     * metadata is any arbitrary metadata attached to the proposal.
     * the recommended format of the metadata is to be found here:
     * https://docs.cosmos.network/v0.47/modules/gov#proposal-3
     */
    metadata: string;
    /**
     * title is the title of the proposal
     *
     * Since: cosmos-sdk 0.47
     */
    title: string;
    /**
     * summary is a short summary of the proposal
     *
     * Since: cosmos-sdk 0.47
     */
    summary: string;
    /**
     * proposer is the address of the proposal sumbitter
     *
     * Since: cosmos-sdk 0.47
     */
    proposer: string;
    /**
     * expedited defines if the proposal is expedited
     *
     * Since: cosmos-sdk 0.50
     */
    expedited: boolean;
    /**
     * failed_reason defines the reason why the proposal failed
     *
     * Since: cosmos-sdk 0.50
     */
    failedReason: string;
}
/** TallyResult defines a standard tally for a governance proposal. */
export interface TallyResult {
    /** yes_count is the number of yes votes on a proposal. */
    yesCount: string;
    /** abstain_count is the number of abstain votes on a proposal. */
    abstainCount: string;
    /** no_count is the number of no votes on a proposal. */
    noCount: string;
    /** no_with_veto_count is the number of no with veto votes on a proposal. */
    noWithVetoCount: string;
}
/**
 * Vote defines a vote on a governance proposal.
 * A Vote consists of a proposal ID, the voter, and the vote option.
 */
export interface Vote {
    /** proposal_id defines the unique id of the proposal. */
    proposalId: number;
    /** voter is the voter address of the proposal. */
    voter: string;
    /** options is the weighted vote options. */
    options: WeightedVoteOption[];
    /**
     * metadata is any arbitrary metadata attached to the vote.
     * the recommended format of the metadata is to be found here: https://docs.cosmos.network/v0.47/modules/gov#vote-5
     */
    metadata: string;
}
/**
 * DepositParams defines the params for deposits on governance proposals.
 *
 * @deprecated
 */
export interface DepositParams {
    /** Minimum deposit for a proposal to enter voting period. */
    minDeposit: Coin[];
    /**
     * Maximum period for Atom holders to deposit on a proposal. Initial value: 2
     * months.
     */
    maxDepositPeriod: Duration | undefined;
}
/**
 * VotingParams defines the params for voting on governance proposals.
 *
 * @deprecated
 */
export interface VotingParams {
    /** Duration of the voting period. */
    votingPeriod: Duration | undefined;
}
/**
 * TallyParams defines the params for tallying votes on governance proposals.
 *
 * @deprecated
 */
export interface TallyParams {
    /**
     * Minimum percentage of total stake needed to vote for a result to be
     * considered valid.
     */
    quorum: string;
    /** Minimum proportion of Yes votes for proposal to pass. Default value: 0.5. */
    threshold: string;
    /**
     * Minimum value of Veto votes to Total votes ratio for proposal to be
     * vetoed. Default value: 1/3.
     */
    vetoThreshold: string;
}
/**
 * Params defines the parameters for the x/gov module.
 *
 * Since: cosmos-sdk 0.47
 */
export interface Params {
    /** Minimum deposit for a proposal to enter voting period. */
    minDeposit: Coin[];
    /**
     * Maximum period for Atom holders to deposit on a proposal. Initial value: 2
     * months.
     */
    maxDepositPeriod: Duration | undefined;
    /** Duration of the voting period. */
    votingPeriod: Duration | undefined;
    /**
     * Minimum percentage of total stake needed to vote for a result to be
     *  considered valid.
     */
    quorum: string;
    /** Minimum proportion of Yes votes for proposal to pass. Default value: 0.5. */
    threshold: string;
    /**
     * Minimum value of Veto votes to Total votes ratio for proposal to be
     *  vetoed. Default value: 1/3.
     */
    vetoThreshold: string;
    /** The ratio representing the proportion of the deposit value that must be paid at proposal submission. */
    minInitialDepositRatio: string;
    /**
     * The cancel ratio which will not be returned back to the depositors when a proposal is cancelled.
     *
     * Since: cosmos-sdk 0.50
     */
    proposalCancelRatio: string;
    /**
     * The address which will receive (proposal_cancel_ratio * deposit) proposal deposits.
     * If empty, the (proposal_cancel_ratio * deposit) proposal deposits will be burned.
     *
     * Since: cosmos-sdk 0.50
     */
    proposalCancelDest: string;
    /**
     * Duration of the voting period of an expedited proposal.
     *
     * Since: cosmos-sdk 0.50
     */
    expeditedVotingPeriod: Duration | undefined;
    /**
     * Minimum proportion of Yes votes for proposal to pass. Default value: 0.67.
     *
     * Since: cosmos-sdk 0.50
     */
    expeditedThreshold: string;
    /** Minimum expedited deposit for a proposal to enter voting period. */
    expeditedMinDeposit: Coin[];
    /** burn deposits if a proposal does not meet quorum */
    burnVoteQuorum: boolean;
    /** burn deposits if the proposal does not enter voting period */
    burnProposalDepositPrevote: boolean;
    /** burn deposits if quorum with vote type no_veto is met */
    burnVoteVeto: boolean;
    /**
     * The ratio representing the proportion of the deposit value minimum that must be met when making a deposit.
     * Default value: 0.01. Meaning that for a chain with a min_deposit of 100stake, a deposit of 1stake would be
     * required.
     *
     * Since: cosmos-sdk 0.50
     */
    minDepositRatio: string;
}
export declare const WeightedVoteOption: {
    encode(message: WeightedVoteOption, writer?: _m0.Writer): _m0.Writer;
    decode(input: _m0.Reader | Uint8Array, length?: number): WeightedVoteOption;
    fromJSON(object: any): WeightedVoteOption;
    toJSON(message: WeightedVoteOption): unknown;
    create<I extends {
        option?: VoteOption;
        weight?: string;
    } & {
        option?: VoteOption;
        weight?: string;
    } & { [K in Exclude<keyof I, keyof WeightedVoteOption>]: never; }>(base?: I): WeightedVoteOption;
    fromPartial<I_1 extends {
        option?: VoteOption;
        weight?: string;
    } & {
        option?: VoteOption;
        weight?: string;
    } & { [K_1 in Exclude<keyof I_1, keyof WeightedVoteOption>]: never; }>(object: I_1): WeightedVoteOption;
};
export declare const Deposit: {
    encode(message: Deposit, writer?: _m0.Writer): _m0.Writer;
    decode(input: _m0.Reader | Uint8Array, length?: number): Deposit;
    fromJSON(object: any): Deposit;
    toJSON(message: Deposit): unknown;
    create<I extends {
        proposalId?: number;
        depositor?: string;
        amount?: {
            denom?: string;
            amount?: string;
        }[];
    } & {
        proposalId?: number;
        depositor?: string;
        amount?: {
            denom?: string;
            amount?: string;
        }[] & ({
            denom?: string;
            amount?: string;
        } & {
            denom?: string;
            amount?: string;
        } & { [K in Exclude<keyof I["amount"][number], keyof Coin>]: never; })[] & { [K_1 in Exclude<keyof I["amount"], keyof {
            denom?: string;
            amount?: string;
        }[]>]: never; };
    } & { [K_2 in Exclude<keyof I, keyof Deposit>]: never; }>(base?: I): Deposit;
    fromPartial<I_1 extends {
        proposalId?: number;
        depositor?: string;
        amount?: {
            denom?: string;
            amount?: string;
        }[];
    } & {
        proposalId?: number;
        depositor?: string;
        amount?: {
            denom?: string;
            amount?: string;
        }[] & ({
            denom?: string;
            amount?: string;
        } & {
            denom?: string;
            amount?: string;
        } & { [K_3 in Exclude<keyof I_1["amount"][number], keyof Coin>]: never; })[] & { [K_4 in Exclude<keyof I_1["amount"], keyof {
            denom?: string;
            amount?: string;
        }[]>]: never; };
    } & { [K_5 in Exclude<keyof I_1, keyof Deposit>]: never; }>(object: I_1): Deposit;
};
export declare const Proposal: {
    encode(message: Proposal, writer?: _m0.Writer): _m0.Writer;
    decode(input: _m0.Reader | Uint8Array, length?: number): Proposal;
    fromJSON(object: any): Proposal;
    toJSON(message: Proposal): unknown;
    create<I extends {
        id?: number;
        messages?: {
            typeUrl?: string;
            value?: Uint8Array;
        }[];
        status?: ProposalStatus;
        finalTallyResult?: {
            yesCount?: string;
            abstainCount?: string;
            noCount?: string;
            noWithVetoCount?: string;
        };
        submitTime?: Date | undefined;
        depositEndTime?: Date | undefined;
        totalDeposit?: {
            denom?: string;
            amount?: string;
        }[];
        votingStartTime?: Date | undefined;
        votingEndTime?: Date | undefined;
        metadata?: string;
        title?: string;
        summary?: string;
        proposer?: string;
        expedited?: boolean;
        failedReason?: string;
    } & {
        id?: number;
        messages?: {
            typeUrl?: string;
            value?: Uint8Array;
        }[] & ({
            typeUrl?: string;
            value?: Uint8Array;
        } & {
            typeUrl?: string;
            value?: Uint8Array;
        } & { [K in Exclude<keyof I["messages"][number], keyof Any>]: never; })[] & { [K_1 in Exclude<keyof I["messages"], keyof {
            typeUrl?: string;
            value?: Uint8Array;
        }[]>]: never; };
        status?: ProposalStatus;
        finalTallyResult?: {
            yesCount?: string;
            abstainCount?: string;
            noCount?: string;
            noWithVetoCount?: string;
        } & {
            yesCount?: string;
            abstainCount?: string;
            noCount?: string;
            noWithVetoCount?: string;
        } & { [K_2 in Exclude<keyof I["finalTallyResult"], keyof TallyResult>]: never; };
        submitTime?: Date | undefined;
        depositEndTime?: Date | undefined;
        totalDeposit?: {
            denom?: string;
            amount?: string;
        }[] & ({
            denom?: string;
            amount?: string;
        } & {
            denom?: string;
            amount?: string;
        } & { [K_3 in Exclude<keyof I["totalDeposit"][number], keyof Coin>]: never; })[] & { [K_4 in Exclude<keyof I["totalDeposit"], keyof {
            denom?: string;
            amount?: string;
        }[]>]: never; };
        votingStartTime?: Date | undefined;
        votingEndTime?: Date | undefined;
        metadata?: string;
        title?: string;
        summary?: string;
        proposer?: string;
        expedited?: boolean;
        failedReason?: string;
    } & { [K_5 in Exclude<keyof I, keyof Proposal>]: never; }>(base?: I): Proposal;
    fromPartial<I_1 extends {
        id?: number;
        messages?: {
            typeUrl?: string;
            value?: Uint8Array;
        }[];
        status?: ProposalStatus;
        finalTallyResult?: {
            yesCount?: string;
            abstainCount?: string;
            noCount?: string;
            noWithVetoCount?: string;
        };
        submitTime?: Date | undefined;
        depositEndTime?: Date | undefined;
        totalDeposit?: {
            denom?: string;
            amount?: string;
        }[];
        votingStartTime?: Date | undefined;
        votingEndTime?: Date | undefined;
        metadata?: string;
        title?: string;
        summary?: string;
        proposer?: string;
        expedited?: boolean;
        failedReason?: string;
    } & {
        id?: number;
        messages?: {
            typeUrl?: string;
            value?: Uint8Array;
        }[] & ({
            typeUrl?: string;
            value?: Uint8Array;
        } & {
            typeUrl?: string;
            value?: Uint8Array;
        } & { [K_6 in Exclude<keyof I_1["messages"][number], keyof Any>]: never; })[] & { [K_7 in Exclude<keyof I_1["messages"], keyof {
            typeUrl?: string;
            value?: Uint8Array;
        }[]>]: never; };
        status?: ProposalStatus;
        finalTallyResult?: {
            yesCount?: string;
            abstainCount?: string;
            noCount?: string;
            noWithVetoCount?: string;
        } & {
            yesCount?: string;
            abstainCount?: string;
            noCount?: string;
            noWithVetoCount?: string;
        } & { [K_8 in Exclude<keyof I_1["finalTallyResult"], keyof TallyResult>]: never; };
        submitTime?: Date | undefined;
        depositEndTime?: Date | undefined;
        totalDeposit?: {
            denom?: string;
            amount?: string;
        }[] & ({
            denom?: string;
            amount?: string;
        } & {
            denom?: string;
            amount?: string;
        } & { [K_9 in Exclude<keyof I_1["totalDeposit"][number], keyof Coin>]: never; })[] & { [K_10 in Exclude<keyof I_1["totalDeposit"], keyof {
            denom?: string;
            amount?: string;
        }[]>]: never; };
        votingStartTime?: Date | undefined;
        votingEndTime?: Date | undefined;
        metadata?: string;
        title?: string;
        summary?: string;
        proposer?: string;
        expedited?: boolean;
        failedReason?: string;
    } & { [K_11 in Exclude<keyof I_1, keyof Proposal>]: never; }>(object: I_1): Proposal;
};
export declare const TallyResult: {
    encode(message: TallyResult, writer?: _m0.Writer): _m0.Writer;
    decode(input: _m0.Reader | Uint8Array, length?: number): TallyResult;
    fromJSON(object: any): TallyResult;
    toJSON(message: TallyResult): unknown;
    create<I extends {
        yesCount?: string;
        abstainCount?: string;
        noCount?: string;
        noWithVetoCount?: string;
    } & {
        yesCount?: string;
        abstainCount?: string;
        noCount?: string;
        noWithVetoCount?: string;
    } & { [K in Exclude<keyof I, keyof TallyResult>]: never; }>(base?: I): TallyResult;
    fromPartial<I_1 extends {
        yesCount?: string;
        abstainCount?: string;
        noCount?: string;
        noWithVetoCount?: string;
    } & {
        yesCount?: string;
        abstainCount?: string;
        noCount?: string;
        noWithVetoCount?: string;
    } & { [K_1 in Exclude<keyof I_1, keyof TallyResult>]: never; }>(object: I_1): TallyResult;
};
export declare const Vote: {
    encode(message: Vote, writer?: _m0.Writer): _m0.Writer;
    decode(input: _m0.Reader | Uint8Array, length?: number): Vote;
    fromJSON(object: any): Vote;
    toJSON(message: Vote): unknown;
    create<I extends {
        proposalId?: number;
        voter?: string;
        options?: {
            option?: VoteOption;
            weight?: string;
        }[];
        metadata?: string;
    } & {
        proposalId?: number;
        voter?: string;
        options?: {
            option?: VoteOption;
            weight?: string;
        }[] & ({
            option?: VoteOption;
            weight?: string;
        } & {
            option?: VoteOption;
            weight?: string;
        } & { [K in Exclude<keyof I["options"][number], keyof WeightedVoteOption>]: never; })[] & { [K_1 in Exclude<keyof I["options"], keyof {
            option?: VoteOption;
            weight?: string;
        }[]>]: never; };
        metadata?: string;
    } & { [K_2 in Exclude<keyof I, keyof Vote>]: never; }>(base?: I): Vote;
    fromPartial<I_1 extends {
        proposalId?: number;
        voter?: string;
        options?: {
            option?: VoteOption;
            weight?: string;
        }[];
        metadata?: string;
    } & {
        proposalId?: number;
        voter?: string;
        options?: {
            option?: VoteOption;
            weight?: string;
        }[] & ({
            option?: VoteOption;
            weight?: string;
        } & {
            option?: VoteOption;
            weight?: string;
        } & { [K_3 in Exclude<keyof I_1["options"][number], keyof WeightedVoteOption>]: never; })[] & { [K_4 in Exclude<keyof I_1["options"], keyof {
            option?: VoteOption;
            weight?: string;
        }[]>]: never; };
        metadata?: string;
    } & { [K_5 in Exclude<keyof I_1, keyof Vote>]: never; }>(object: I_1): Vote;
};
export declare const DepositParams: {
    encode(message: DepositParams, writer?: _m0.Writer): _m0.Writer;
    decode(input: _m0.Reader | Uint8Array, length?: number): DepositParams;
    fromJSON(object: any): DepositParams;
    toJSON(message: DepositParams): unknown;
    create<I extends {
        minDeposit?: {
            denom?: string;
            amount?: string;
        }[];
        maxDepositPeriod?: {
            seconds?: number;
            nanos?: number;
        };
    } & {
        minDeposit?: {
            denom?: string;
            amount?: string;
        }[] & ({
            denom?: string;
            amount?: string;
        } & {
            denom?: string;
            amount?: string;
        } & { [K in Exclude<keyof I["minDeposit"][number], keyof Coin>]: never; })[] & { [K_1 in Exclude<keyof I["minDeposit"], keyof {
            denom?: string;
            amount?: string;
        }[]>]: never; };
        maxDepositPeriod?: {
            seconds?: number;
            nanos?: number;
        } & {
            seconds?: number;
            nanos?: number;
        } & { [K_2 in Exclude<keyof I["maxDepositPeriod"], keyof Duration>]: never; };
    } & { [K_3 in Exclude<keyof I, keyof DepositParams>]: never; }>(base?: I): DepositParams;
    fromPartial<I_1 extends {
        minDeposit?: {
            denom?: string;
            amount?: string;
        }[];
        maxDepositPeriod?: {
            seconds?: number;
            nanos?: number;
        };
    } & {
        minDeposit?: {
            denom?: string;
            amount?: string;
        }[] & ({
            denom?: string;
            amount?: string;
        } & {
            denom?: string;
            amount?: string;
        } & { [K_4 in Exclude<keyof I_1["minDeposit"][number], keyof Coin>]: never; })[] & { [K_5 in Exclude<keyof I_1["minDeposit"], keyof {
            denom?: string;
            amount?: string;
        }[]>]: never; };
        maxDepositPeriod?: {
            seconds?: number;
            nanos?: number;
        } & {
            seconds?: number;
            nanos?: number;
        } & { [K_6 in Exclude<keyof I_1["maxDepositPeriod"], keyof Duration>]: never; };
    } & { [K_7 in Exclude<keyof I_1, keyof DepositParams>]: never; }>(object: I_1): DepositParams;
};
export declare const VotingParams: {
    encode(message: VotingParams, writer?: _m0.Writer): _m0.Writer;
    decode(input: _m0.Reader | Uint8Array, length?: number): VotingParams;
    fromJSON(object: any): VotingParams;
    toJSON(message: VotingParams): unknown;
    create<I extends {
        votingPeriod?: {
            seconds?: number;
            nanos?: number;
        };
    } & {
        votingPeriod?: {
            seconds?: number;
            nanos?: number;
        } & {
            seconds?: number;
            nanos?: number;
        } & { [K in Exclude<keyof I["votingPeriod"], keyof Duration>]: never; };
    } & { [K_1 in Exclude<keyof I, "votingPeriod">]: never; }>(base?: I): VotingParams;
    fromPartial<I_1 extends {
        votingPeriod?: {
            seconds?: number;
            nanos?: number;
        };
    } & {
        votingPeriod?: {
            seconds?: number;
            nanos?: number;
        } & {
            seconds?: number;
            nanos?: number;
        } & { [K_2 in Exclude<keyof I_1["votingPeriod"], keyof Duration>]: never; };
    } & { [K_3 in Exclude<keyof I_1, "votingPeriod">]: never; }>(object: I_1): VotingParams;
};
export declare const TallyParams: {
    encode(message: TallyParams, writer?: _m0.Writer): _m0.Writer;
    decode(input: _m0.Reader | Uint8Array, length?: number): TallyParams;
    fromJSON(object: any): TallyParams;
    toJSON(message: TallyParams): unknown;
    create<I extends {
        quorum?: string;
        threshold?: string;
        vetoThreshold?: string;
    } & {
        quorum?: string;
        threshold?: string;
        vetoThreshold?: string;
    } & { [K in Exclude<keyof I, keyof TallyParams>]: never; }>(base?: I): TallyParams;
    fromPartial<I_1 extends {
        quorum?: string;
        threshold?: string;
        vetoThreshold?: string;
    } & {
        quorum?: string;
        threshold?: string;
        vetoThreshold?: string;
    } & { [K_1 in Exclude<keyof I_1, keyof TallyParams>]: never; }>(object: I_1): TallyParams;
};
export declare const Params: {
    encode(message: Params, writer?: _m0.Writer): _m0.Writer;
    decode(input: _m0.Reader | Uint8Array, length?: number): Params;
    fromJSON(object: any): Params;
    toJSON(message: Params): unknown;
    create<I extends {
        minDeposit?: {
            denom?: string;
            amount?: string;
        }[];
        maxDepositPeriod?: {
            seconds?: number;
            nanos?: number;
        };
        votingPeriod?: {
            seconds?: number;
            nanos?: number;
        };
        quorum?: string;
        threshold?: string;
        vetoThreshold?: string;
        minInitialDepositRatio?: string;
        proposalCancelRatio?: string;
        proposalCancelDest?: string;
        expeditedVotingPeriod?: {
            seconds?: number;
            nanos?: number;
        };
        expeditedThreshold?: string;
        expeditedMinDeposit?: {
            denom?: string;
            amount?: string;
        }[];
        burnVoteQuorum?: boolean;
        burnProposalDepositPrevote?: boolean;
        burnVoteVeto?: boolean;
        minDepositRatio?: string;
    } & {
        minDeposit?: {
            denom?: string;
            amount?: string;
        }[] & ({
            denom?: string;
            amount?: string;
        } & {
            denom?: string;
            amount?: string;
        } & { [K in Exclude<keyof I["minDeposit"][number], keyof Coin>]: never; })[] & { [K_1 in Exclude<keyof I["minDeposit"], keyof {
            denom?: string;
            amount?: string;
        }[]>]: never; };
        maxDepositPeriod?: {
            seconds?: number;
            nanos?: number;
        } & {
            seconds?: number;
            nanos?: number;
        } & { [K_2 in Exclude<keyof I["maxDepositPeriod"], keyof Duration>]: never; };
        votingPeriod?: {
            seconds?: number;
            nanos?: number;
        } & {
            seconds?: number;
            nanos?: number;
        } & { [K_3 in Exclude<keyof I["votingPeriod"], keyof Duration>]: never; };
        quorum?: string;
        threshold?: string;
        vetoThreshold?: string;
        minInitialDepositRatio?: string;
        proposalCancelRatio?: string;
        proposalCancelDest?: string;
        expeditedVotingPeriod?: {
            seconds?: number;
            nanos?: number;
        } & {
            seconds?: number;
            nanos?: number;
        } & { [K_4 in Exclude<keyof I["expeditedVotingPeriod"], keyof Duration>]: never; };
        expeditedThreshold?: string;
        expeditedMinDeposit?: {
            denom?: string;
            amount?: string;
        }[] & ({
            denom?: string;
            amount?: string;
        } & {
            denom?: string;
            amount?: string;
        } & { [K_5 in Exclude<keyof I["expeditedMinDeposit"][number], keyof Coin>]: never; })[] & { [K_6 in Exclude<keyof I["expeditedMinDeposit"], keyof {
            denom?: string;
            amount?: string;
        }[]>]: never; };
        burnVoteQuorum?: boolean;
        burnProposalDepositPrevote?: boolean;
        burnVoteVeto?: boolean;
        minDepositRatio?: string;
    } & { [K_7 in Exclude<keyof I, keyof Params>]: never; }>(base?: I): Params;
    fromPartial<I_1 extends {
        minDeposit?: {
            denom?: string;
            amount?: string;
        }[];
        maxDepositPeriod?: {
            seconds?: number;
            nanos?: number;
        };
        votingPeriod?: {
            seconds?: number;
            nanos?: number;
        };
        quorum?: string;
        threshold?: string;
        vetoThreshold?: string;
        minInitialDepositRatio?: string;
        proposalCancelRatio?: string;
        proposalCancelDest?: string;
        expeditedVotingPeriod?: {
            seconds?: number;
            nanos?: number;
        };
        expeditedThreshold?: string;
        expeditedMinDeposit?: {
            denom?: string;
            amount?: string;
        }[];
        burnVoteQuorum?: boolean;
        burnProposalDepositPrevote?: boolean;
        burnVoteVeto?: boolean;
        minDepositRatio?: string;
    } & {
        minDeposit?: {
            denom?: string;
            amount?: string;
        }[] & ({
            denom?: string;
            amount?: string;
        } & {
            denom?: string;
            amount?: string;
        } & { [K_8 in Exclude<keyof I_1["minDeposit"][number], keyof Coin>]: never; })[] & { [K_9 in Exclude<keyof I_1["minDeposit"], keyof {
            denom?: string;
            amount?: string;
        }[]>]: never; };
        maxDepositPeriod?: {
            seconds?: number;
            nanos?: number;
        } & {
            seconds?: number;
            nanos?: number;
        } & { [K_10 in Exclude<keyof I_1["maxDepositPeriod"], keyof Duration>]: never; };
        votingPeriod?: {
            seconds?: number;
            nanos?: number;
        } & {
            seconds?: number;
            nanos?: number;
        } & { [K_11 in Exclude<keyof I_1["votingPeriod"], keyof Duration>]: never; };
        quorum?: string;
        threshold?: string;
        vetoThreshold?: string;
        minInitialDepositRatio?: string;
        proposalCancelRatio?: string;
        proposalCancelDest?: string;
        expeditedVotingPeriod?: {
            seconds?: number;
            nanos?: number;
        } & {
            seconds?: number;
            nanos?: number;
        } & { [K_12 in Exclude<keyof I_1["expeditedVotingPeriod"], keyof Duration>]: never; };
        expeditedThreshold?: string;
        expeditedMinDeposit?: {
            denom?: string;
            amount?: string;
        }[] & ({
            denom?: string;
            amount?: string;
        } & {
            denom?: string;
            amount?: string;
        } & { [K_13 in Exclude<keyof I_1["expeditedMinDeposit"][number], keyof Coin>]: never; })[] & { [K_14 in Exclude<keyof I_1["expeditedMinDeposit"], keyof {
            denom?: string;
            amount?: string;
        }[]>]: never; };
        burnVoteQuorum?: boolean;
        burnProposalDepositPrevote?: boolean;
        burnVoteVeto?: boolean;
        minDepositRatio?: string;
    } & { [K_15 in Exclude<keyof I_1, keyof Params>]: never; }>(object: I_1): Params;
};
type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;
export type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P : P & {
    [K in keyof P]: Exact<P[K], I[K]>;
} & {
    [K in Exclude<keyof I, KeysOfUnion<P>>]: never;
};
export {};