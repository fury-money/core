package config

const (
	// AccountAddressPrefix is the prefix of bech32 encoded address
	AccountAddressPrefix = "furya"

	// AppName is the application name
	AppName = "furya"

	// CoinType is the LUNA coin type as defined in SLIP44 (https://github.com/satoshilabs/slips/blob/master/slip-0044.md)
	CoinType = 330

	// BondDenom staking denom
	BondDenom = "uluna"

	// More denoms
	// Luna      = "luna"    // 1 (base denom unit)
	// MilliLuna = "mluna"   // 10^-3 (milli)
	MicroLuna = BondDenom // 10^-6 (micro)
	// NanoLuna  = "nluna"   // 10^-9 (nano)

	// Modules Msgs
	AuthzMsgExec                        = "/cosmos.authz.v1beta1.MsgExec"
	AuthzMsgGrant                       = "/cosmos.authz.v1beta1.MsgGrant"
	AuthzMsgRevoke                      = "/cosmos.authz.v1beta1.MsgRevoke"
	BankMsgSend                         = "/cosmos.bank.v1beta1.MsgSend"
	BankMsgMultiSend                    = "/cosmos.bank.v1beta1.MsgMultiSend"
	DistrMsgSetWithdrawAddr             = "/cosmos.distribution.v1beta1.MsgSetWithdrawAddress"
	DistrMsgWithdrawValidatorCommission = "/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission"
	DistrMsgFundCommunityPool           = "/cosmos.distribution.v1beta1.MsgFundCommunityPool"
	DistrMsgWithdrawDelegatorReward     = "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward"
	FeegrantMsgGrantAllowance           = "/cosmos.feegrant.v1beta1.MsgGrantAllowance"
	FeegrantMsgRevokeAllowance          = "/cosmos.feegrant.v1beta1.MsgRevokeAllowance"
	GovMsgVoteWeighted                  = "/cosmos.gov.v1beta1.MsgVoteWeighted"
	GovMsgSubmitProposal                = "/cosmos.gov.v1beta1.MsgSubmitProposal"
	GovMsgDeposit                       = "/cosmos.gov.v1beta1.MsgDeposit"
	GovMsgVote                          = "/cosmos.gov.v1beta1.MsgVote"
	StakingMsgEditValidator             = "/cosmos.staking.v1beta1.MsgEditValidator"
	StakingMsgDelegate                  = "/cosmos.staking.v1beta1.MsgDelegate"
	StakingMsgUndelegate                = "/cosmos.staking.v1beta1.MsgUndelegate"
	StakingMsgBeginRedelegate           = "/cosmos.staking.v1beta1.MsgBeginRedelegate"
	StakingMsgCreateValidator           = "/cosmos.staking.v1beta1.MsgCreateValidator"
	VestingMsgCreateVestingAccount      = "/cosmos.vesting.v1beta1.MsgCreateVestingAccount"
	TransferMsgTransfer                 = "/ibc.applications.transfer.v1.MsgTransfer"
	WasmMsgStoreCode                    = "/cosmwasm.wasm.v1.MsgStoreCode"
	WasmMsgExecuteContract              = "/cosmwasm.wasm.v1.MsgExecuteContract"
	WasmMsgInstantiateContract          = "/cosmwasm.wasm.v1.MsgInstantiateContract"
	WasmMsgMigrateContract              = "/cosmwasm.wasm.v1.MsgMigrateContract"

	// Module query
	QueryAllAllianceValidators                 = "/alliance.alliance.Query/AllAllianceValidators"
	QueryAllAlliancesDelegations               = "/alliance.alliance.Query/AllAlliancesDelegations"
	QueryAlliance                              = "/alliance.alliance.Query/Alliance"
	QueryAllianceDelegation                    = "/alliance.alliance.Query/AllianceDelegation"
	QueryAllianceDelegationRewards             = "/alliance.alliance.Query/AllianceDelegationRewards"
	QueryAllianceRedelegations                 = "/alliance.alliance.Query/AllianceRedelegations"
	QueryAllianceUnbondings                    = "/alliance.alliance.Query/AllianceUnbondings"
	QueryAllianceUnbondingsByDenomAndDelegator = "/alliance.alliance.Query/AllianceUnbondingsByDenomAndDelegator"
	QueryAllianceValidator                     = "/alliance.alliance.Query/AllianceValidator"
	QueryAlliances                             = "/alliance.alliance.Query/Alliances"
	QueryAlliancesDelegation                   = "/alliance.alliance.Query/AlliancesDelegation"
	QueryAlliancesDelegationByValidator        = "/alliance.alliance.Query/AlliancesDelegationByValidator"
	QueryIBCAlliance                           = "/alliance.alliance.Query/IBCAlliance"
	QueryIBCAllianceDelegation                 = "/alliance.alliance.Query/IBCAllianceDelegation"
	QueryIBCAllianceDelegationRewards          = "/alliance.alliance.Query/IBCAllianceDelegationRewards"
	QueryAllianceParams                        = "/alliance.alliance.Query/Params"
	QueryAccount                               = "/cosmos.auth.v1beta1.Query/Account"
	QueryAccountAddressByID                    = "/cosmos.auth.v1beta1.Query/AccountAddressByID"
	QueryAccountInfo                           = "/cosmos.auth.v1beta1.Query/AccountInfo"
	QueryAccounts                              = "/cosmos.auth.v1beta1.Query/Accounts"
	QueryAddressBytesToString                  = "/cosmos.auth.v1beta1.Query/AddressBytesToString"
	QueryAddressStringToBytes                  = "/cosmos.auth.v1beta1.Query/AddressStringToBytes"
	QueryBech32Prefix                          = "/cosmos.auth.v1beta1.Query/Bech32Prefix"
	QueryModuleAccountByName                   = "/cosmos.auth.v1beta1.Query/ModuleAccountByName"
	QueryModuleAccounts                        = "/cosmos.auth.v1beta1.Query/ModuleAccounts"
	QueryAuthParams                            = "/cosmos.auth.v1beta1.Query/Params"
	QueryGranteeGrants                         = "/cosmos.authz.v1beta1.Query/GranteeGrants"
	QueryGranterGrants                         = "/cosmos.authz.v1beta1.Query/GranterGrants"
	QueryGrants                                = "/cosmos.authz.v1beta1.Query/Grants"
	QueryAllBalances                           = "/cosmos.bank.v1beta1.Query/AllBalances"
	QueryBalance                               = "/cosmos.bank.v1beta1.Query/Balance"
	QueryDenomMetadata                         = "/cosmos.bank.v1beta1.Query/DenomMetadata"
	QueryDenomOwners                           = "/cosmos.bank.v1beta1.Query/DenomOwners"
	QueryDenomsMetadata                        = "/cosmos.bank.v1beta1.Query/DenomsMetadata"
	QueryBankParams                            = "/cosmos.bank.v1beta1.Query/Params"
	QuerySendEnabled                           = "/cosmos.bank.v1beta1.Query/SendEnabled"
	QuerySpendableBalanceByDenom               = "/cosmos.bank.v1beta1.Query/SpendableBalanceByDenom"
	QuerySpendableBalances                     = "/cosmos.bank.v1beta1.Query/SpendableBalances"
	QuerySupplyOf                              = "/cosmos.bank.v1beta1.Query/SupplyOf"
	QueryTotalSupply                           = "/cosmos.bank.v1beta1.Query/TotalSupply"
	QueryConsensusParams                       = "/cosmos.consensus.v1.Query/Params"
	QueryCommunityPool                         = "/cosmos.distribution.v1beta1.Query/CommunityPool"
	QueryDelegationRewards                     = "/cosmos.distribution.v1beta1.Query/DelegationRewards"
	QueryDelegationTotalRewards                = "/cosmos.distribution.v1beta1.Query/DelegationTotalRewards"
	QueryDistributionDelegatorValidators       = "/cosmos.distribution.v1beta1.Query/DelegatorValidators"
	QueryDelegatorWithdrawAddress              = "/cosmos.distribution.v1beta1.Query/DelegatorWithdrawAddress"
	QueryDistributionParams                    = "/cosmos.distribution.v1beta1.Query/Params"
	QueryValidatorCommission                   = "/cosmos.distribution.v1beta1.Query/ValidatorCommission"
	QueryValidatorDistributionInfo             = "/cosmos.distribution.v1beta1.Query/ValidatorDistributionInfo"
	QueryValidatorOutstandingRewards           = "/cosmos.distribution.v1beta1.Query/ValidatorOutstandingRewards"
	QueryValidatorSlashes                      = "/cosmos.distribution.v1beta1.Query/ValidatorSlashes"
	QueryAllEvidence                           = "/cosmos.evidence.v1beta1.Query/AllEvidence"
	QueryEvidence                              = "/cosmos.evidence.v1beta1.Query/Evidence"
	QueryAllowance                             = "/cosmos.feegrant.v1beta1.Query/Allowance"
	QueryAllowances                            = "/cosmos.feegrant.v1beta1.Query/Allowances"
	QueryAllowancesByGranter                   = "/cosmos.feegrant.v1beta1.Query/AllowancesByGranter"
	QueryDeposit                               = "/cosmos.gov.v1.Query/Deposit"
	QueryDeposits                              = "/cosmos.gov.v1.Query/Deposits"
	QueryGovParams                             = "/cosmos.gov.v1.Query/Params"
	QueryProposal                              = "/cosmos.gov.v1.Query/Proposal"
	QueryProposals                             = "/cosmos.gov.v1.Query/Proposals"
	QueryTallyResult                           = "/cosmos.gov.v1.Query/TallyResult"
	QueryVote                                  = "/cosmos.gov.v1.Query/Vote"
	QueryVotes                                 = "/cosmos.gov.v1.Query/Votes"
	QueryAnnualProvisions                      = "/cosmos.mint.v1beta1.Query/AnnualProvisions"
	QueryInflation                             = "/cosmos.mint.v1beta1.Query/Inflation"
	QueryMintParams                            = "/cosmos.mint.v1beta1.Query/Params"
	QueryParamsModuleParams                    = "/cosmos.params.v1beta1.Query/Params"
	QuerySubspaces                             = "/cosmos.params.v1beta1.Query/Subspaces"
	QuerySlashingParams                        = "/cosmos.slashing.v1beta1.Query/Params"
	QuerySigningInfo                           = "/cosmos.slashing.v1beta1.Query/SigningInfo"
	QuerySigningInfos                          = "/cosmos.slashing.v1beta1.Query/SigningInfos"
	QueryDelegation                            = "/cosmos.staking.v1beta1.Query/Delegation"
	QueryDelegatorDelegations                  = "/cosmos.staking.v1beta1.Query/DelegatorDelegations"
	QueryDelegatorUnbondingDelegations         = "/cosmos.staking.v1beta1.Query/DelegatorUnbondingDelegations"
	QueryDelegatorValidator                    = "/cosmos.staking.v1beta1.Query/DelegatorValidator"
	QueryDelegatorValidators                   = "/cosmos.staking.v1beta1.Query/DelegatorValidators"
	QueryHistoricalInfo                        = "/cosmos.staking.v1beta1.Query/HistoricalInfo"
	QueryStakingParams                         = "/cosmos.staking.v1beta1.Query/Params"
	QueryPool                                  = "/cosmos.staking.v1beta1.Query/Pool"
	QueryRedelegations                         = "/cosmos.staking.v1beta1.Query/Redelegations"
	QueryUnbondingDelegation                   = "/cosmos.staking.v1beta1.Query/UnbondingDelegation"
	QueryValidator                             = "/cosmos.staking.v1beta1.Query/Validator"
	QueryValidatorDelegations                  = "/cosmos.staking.v1beta1.Query/ValidatorDelegations"
	QueryValidatorUnbondingDelegations         = "/cosmos.staking.v1beta1.Query/ValidatorUnbondingDelegations"
	QueryValidators                            = "/cosmos.staking.v1beta1.Query/Validators"
	QueryAppliedPlan                           = "/cosmos.upgrade.v1beta1.Query/AppliedPlan"
	QueryAuthority                             = "/cosmos.upgrade.v1beta1.Query/Authority"
	QueryCurrentPlan                           = "/cosmos.upgrade.v1beta1.Query/CurrentPlan"
	QueryModuleVersions                        = "/cosmos.upgrade.v1beta1.Query/ModuleVersions"
	QueryUpgradedConsensusState                = "/cosmos.upgrade.v1beta1.Query/UpgradedConsensusState"
	QueryAllContractState                      = "/cosmwasm.wasm.v1.Query/AllContractState"
	QueryCode                                  = "/cosmwasm.wasm.v1.Query/Code"
	QueryCodes                                 = "/cosmwasm.wasm.v1.Query/Codes"
	QueryContractHistory                       = "/cosmwasm.wasm.v1.Query/ContractHistory"
	QueryContractInfo                          = "/cosmwasm.wasm.v1.Query/ContractInfo"
	QueryContractsByCode                       = "/cosmwasm.wasm.v1.Query/ContractsByCode"
	QueryContractsByCreator                    = "/cosmwasm.wasm.v1.Query/ContractsByCreator"
	QueryWasmParams                            = "/cosmwasm.wasm.v1.Query/Params"
	QueryPinnedCodes                           = "/cosmwasm.wasm.v1.Query/PinnedCodes"
	QueryRawContractState                      = "/cosmwasm.wasm.v1.Query/RawContractState"
	QuerySmartContractState                    = "/cosmwasm.wasm.v1.Query/SmartContractState"
	QueryCounterpartyPayee                     = "/ibc.applications.fee.v1.Query/CounterpartyPayee"
	QueryFeeEnabledChannel                     = "/ibc.applications.fee.v1.Query/FeeEnabledChannel"
	QueryFeeEnabledChannels                    = "/ibc.applications.fee.v1.Query/FeeEnabledChannels"
	QueryIncentivizedPacket                    = "/ibc.applications.fee.v1.Query/IncentivizedPacket"
	QueryIncentivizedPackets                   = "/ibc.applications.fee.v1.Query/IncentivizedPackets"
	QueryIncentivizedPacketsForChannel         = "/ibc.applications.fee.v1.Query/IncentivizedPacketsForChannel"
	QueryPayee                                 = "/ibc.applications.fee.v1.Query/Payee"
	QueryTotalAckFees                          = "/ibc.applications.fee.v1.Query/TotalAckFees"
	QueryTotalRecvFees                         = "/ibc.applications.fee.v1.Query/TotalRecvFees"
	QueryTotalTimeoutFees                      = "/ibc.applications.fee.v1.Query/TotalTimeoutFees"
	QueryInterchainAccount                     = "/ibc.applications.interchain_accounts.controller.v1.Query/InterchainAccount"
	QueryInterchainAccControllerParams         = "/ibc.applications.interchain_accounts.controller.v1.Query/Params"
	QueryInterchainAccHostParams               = "/ibc.applications.interchain_accounts.host.v1.Query/Params"
	QueryDenomHash                             = "/ibc.applications.transfer.v1.Query/DenomHash"
	QueryDenomTrace                            = "/ibc.applications.transfer.v1.Query/DenomTrace"
	QueryDenomTraces                           = "/ibc.applications.transfer.v1.Query/DenomTraces"
	QueryEscrowAddress                         = "/ibc.applications.transfer.v1.Query/EscrowAddress"
	QueryTransferParams                        = "/ibc.applications.transfer.v1.Query/Params"
	QueryTotalEscrowForDenom                   = "/ibc.applications.transfer.v1.Query/TotalEscrowForDenom"
	QueryChannel                               = "/ibc.core.channel.v1.Query/Channel"
	QueryChannelClientState                    = "/ibc.core.channel.v1.Query/ChannelClientState"
	QueryChannelConsensusState                 = "/ibc.core.channel.v1.Query/ChannelConsensusState"
	QueryChannels                              = "/ibc.core.channel.v1.Query/Channels"
	QueryConnectionChannels                    = "/ibc.core.channel.v1.Query/ConnectionChannels"
	QueryNextSequenceReceive                   = "/ibc.core.channel.v1.Query/NextSequenceReceive"
	QueryPacketAcknowledgement                 = "/ibc.core.channel.v1.Query/PacketAcknowledgement"
	QueryPacketAcknowledgements                = "/ibc.core.channel.v1.Query/PacketAcknowledgements"
	QueryPacketCommitment                      = "/ibc.core.channel.v1.Query/PacketCommitment"
	QueryPacketCommitments                     = "/ibc.core.channel.v1.Query/PacketCommitments"
	QueryPacketReceipt                         = "/ibc.core.channel.v1.Query/PacketReceipt"
	QueryUnreceivedAcks                        = "/ibc.core.channel.v1.Query/UnreceivedAcks"
	QueryUnreceivedPackets                     = "/ibc.core.channel.v1.Query/UnreceivedPackets"
	QueryClientParams                          = "/ibc.core.client.v1.Query/ClientParams"
	QueryClientState                           = "/ibc.core.client.v1.Query/ClientState"
	QueryClientStates                          = "/ibc.core.client.v1.Query/ClientStates"
	QueryClientStatus                          = "/ibc.core.client.v1.Query/ClientStatus"
	QueryConsensusState                        = "/ibc.core.client.v1.Query/ConsensusState"
	QueryConsensusStateHeights                 = "/ibc.core.client.v1.Query/ConsensusStateHeights"
	QueryConsensusStates                       = "/ibc.core.client.v1.Query/ConsensusStates"
	QueryUpgradedClientState                   = "/ibc.core.client.v1.Query/UpgradedClientState"
	QueryUpgradedConsensusStateV1              = "/ibc.core.client.v1.Query/UpgradedConsensusState"
	QueryClientConnections                     = "/ibc.core.connection.v1.Query/ClientConnections"
	QueryConnection                            = "/ibc.core.connection.v1.Query/Connection"
	QueryConnectionClientState                 = "/ibc.core.connection.v1.Query/ConnectionClientState"
	QueryConnectionConsensusState              = "/ibc.core.connection.v1.Query/ConnectionConsensusState"
	QueryConnectionParams                      = "/ibc.core.connection.v1.Query/ConnectionParams"
	QueryConnections                           = "/ibc.core.connection.v1.Query/Connections"
	QueryICQParams                             = "/icq.v1.Query/Params"
	QueryDeployerFeeShares                     = "/juno.feeshare.v1.Query/DeployerFeeShares"
	QueryFeeShare                              = "/juno.feeshare.v1.Query/FeeShare"
	QueryFeeShares                             = "/juno.feeshare.v1.Query/FeeShares"
	QueryFeeshareParams                        = "/juno.feeshare.v1.Query/Params"
	QueryWithdrawerFeeShares                   = "/juno.feeshare.v1.Query/WithdrawerFeeShares"
	QueryBeforeSendHookAddress                 = "/osmosis.tokenfactory.v1beta1.Query/BeforeSendHookAddress"
	QueryDenomAuthorityMetadata                = "/osmosis.tokenfactory.v1beta1.Query/DenomAuthorityMetadata"
	QueryDenomsFromCreator                     = "/osmosis.tokenfactory.v1beta1.Query/DenomsFromCreator"
	QueryTokeFactoryParams                     = "/osmosis.tokenfactory.v1beta1.Query/Params"
	QueryPobParams                             = "/pob.builder.v1.Query/Params"
	QueryPFMParams                             = "/router.v1.Query/Params"

	// UpgradeName gov proposal name
	Upgrade2_2_0  = "2.2.0"
	Upgrade2_3_0  = "2.3.0"
	Upgrade2_4_rc = "2.4.0-rc4" // This is pisco only since an incorrect plan name was used for the upgrade
	Upgrade2_4    = "v2.4"
	Upgrade2_5    = "v2.5"
	Upgrade2_6    = "v2.6"
	Upgrade2_7    = "v2.7"
	Upgrade2_8    = "v2.8"
	Upgrade2_9    = "v2.9"
)
