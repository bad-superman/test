package thegraph

import "time"

// GraphQLPayload 定义了GraphQL请求体的结构
// Query: GraphQL查询语句
// Variables: 查询变量（可选）
// OperationName: 操作名（可选）
type GraphQLPayload struct {
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables,omitempty"`
	OperationName string                 `json:"operationName,omitempty"`
}

type Indexer struct {
	ID                                string    `json:"id" gorm:"column:id;primaryKey"`
	CreatedAt                         int64     `json:"createdAt" gorm:"column:created_at"`
	URL                               string    `json:"url" gorm:"column:url"`
	GeoHash                           string    `json:"geoHash" gorm:"column:geo_hash"`
	DefaultDisplayName                string    `json:"defaultDisplayName" gorm:"column:default_display_name"`
	StakedTokens                      string    `json:"stakedTokens" gorm:"column:staked_tokens"`
	AllocatedTokens                   string    `json:"allocatedTokens" gorm:"column:allocated_tokens"`
	UnstakedTokens                    string    `json:"unstakedTokens" gorm:"column:unstaked_tokens"`
	LockedTokens                      string    `json:"lockedTokens" gorm:"column:locked_tokens"`
	TokensLockedUntil                 int       `json:"tokensLockedUntil" gorm:"column:tokens_locked_until"`
	AllocationCount                   int64     `json:"allocationCount" gorm:"column:allocation_count"`
	TotalAllocationCount              string    `json:"totalAllocationCount" gorm:"column:total_allocation_count"`
	QueryFeesCollected                string    `json:"queryFeesCollected" gorm:"column:query_fees_collected"`
	QueryFeeRebates                   string    `json:"queryFeeRebates" gorm:"column:query_fee_rebates"`
	RewardsEarned                     string    `json:"rewardsEarned" gorm:"column:rewards_earned"`
	IndexerIndexingRewards            string    `json:"indexerIndexingRewards" gorm:"column:indexer_indexing_rewards"`
	DelegatorIndexingRewards          string    `json:"delegatorIndexingRewards" gorm:"column:delegator_indexing_rewards"`
	IndexerRewardsOwnGenerationRatio  string    `json:"indexerRewardsOwnGenerationRatio" gorm:"column:indexer_rewards_own_generation_ratio"`
	TransferredToL2                   bool      `json:"transferredToL2" gorm:"column:transferred_to_l2"`
	FirstTransferredToL2At            string    `json:"firstTransferredToL2At" gorm:"column:first_transferred_to_l2_at"`
	FirstTransferredToL2AtBlockNumber string    `json:"firstTransferredToL2AtBlockNumber" gorm:"column:first_transferred_to_l2_at_block_number"`
	FirstTransferredToL2AtTx          string    `json:"firstTransferredToL2AtTx" gorm:"column:first_transferred_to_l2_at_tx"`
	LastTransferredToL2At             string    `json:"lastTransferredToL2At" gorm:"column:last_transferred_to_l2_at"`
	LastTransferredToL2AtBlockNumber  string    `json:"lastTransferredToL2AtBlockNumber" gorm:"column:last_transferred_to_l2_at_block_number"`
	LastTransferredToL2AtTx           string    `json:"lastTransferredToL2AtTx" gorm:"column:last_transferred_to_l2_at_tx"`
	StakedTokensTransferredToL2       string    `json:"stakedTokensTransferredToL2" gorm:"column:staked_tokens_transferred_to_l2"`
	IDOnL2                            string    `json:"idOnL2" gorm:"column:id_on_l2"`
	IDOnL1                            string    `json:"idOnL1" gorm:"column:id_on_l1"`
	DelegatedCapacity                 string    `json:"delegatedCapacity" gorm:"column:delegated_capacity"`
	TokenCapacity                     string    `json:"tokenCapacity" gorm:"column:token_capacity"`
	AvailableStake                    string    `json:"availableStake" gorm:"column:available_stake"`
	DelegatedTokens                   string    `json:"delegatedTokens" gorm:"column:delegated_tokens"`
	OwnStakeRatio                     string    `json:"ownStakeRatio" gorm:"column:own_stake_ratio"`
	DelegatedStakeRatio               string    `json:"delegatedStakeRatio" gorm:"column:delegated_stake_ratio"`
	DelegatorShares                   string    `json:"delegatorShares" gorm:"column:delegator_shares"`
	DelegationExchangeRate            string    `json:"delegationExchangeRate" gorm:"column:delegation_exchange_rate"`
	IndexingRewardCut                 int       `json:"indexingRewardCut" gorm:"column:indexing_reward_cut"`
	IndexingRewardEffectiveCut        string    `json:"indexingRewardEffectiveCut" gorm:"column:indexing_reward_effective_cut"`
	OverDelegationDilution            string    `json:"overDelegationDilution" gorm:"column:over_delegation_dilution"`
	DelegatorQueryFees                string    `json:"delegatorQueryFees" gorm:"column:delegator_query_fees"`
	QueryFeeCut                       int       `json:"queryFeeCut" gorm:"column:query_fee_cut"`
	QueryFeeEffectiveCut              string    `json:"queryFeeEffectiveCut" gorm:"column:query_fee_effective_cut"`
	DelegatorParameterCooldown        int       `json:"delegatorParameterCooldown" gorm:"column:delegator_parameter_cooldown"`
	LastDelegationParameterUpdate     int       `json:"lastDelegationParameterUpdate" gorm:"column:last_delegation_parameter_update"`
	ForcedClosures                    int       `json:"forcedClosures" gorm:"column:forced_closures"`
	TotalReturn                       string    `json:"totalReturn" gorm:"column:total_return"`
	AnnualizedReturn                  string    `json:"annualizedReturn" gorm:"column:annualized_return"`
	StakingEfficiency                 string    `json:"stakingEfficiency" gorm:"column:staking_efficiency"`
	UpdatedAt                         time.Time `json:"-" gorm:"column:updated_at"`
}

// TableName 返回表名
func (Indexer) TableName() string {
	return "thegraph_indexer"
}

// DDL 返回创建表的SQL语句
func (Indexer) DDL() string {
	return `CREATE TABLE IF NOT EXISTS indexers (
  id CHAR(40) PRIMARY KEY,
  created_at BIGINT,
  url VARCHAR(255),
  geo_hash VARCHAR(255),
  default_display_name VARCHAR(255),
  staked_tokens VARCHAR(255),
  allocated_tokens VARCHAR(255),
  unstaked_tokens VARCHAR(255),
  locked_tokens VARCHAR(255),
  tokens_locked_until INT,
  allocation_count BIGINT,
  total_allocation_count BIGINT,
  query_fees_collected VARCHAR(255),
  query_fee_rebates VARCHAR(255),
  rewards_earned VARCHAR(255),
  indexer_indexing_rewards VARCHAR(255),
  delegator_indexing_rewards VARCHAR(255),
  indexer_rewards_own_generation_ratio VARCHAR(255),
  transferred_to_l2 VARCHAR(255),
  first_transferred_to_l2_at VARCHAR(255),
  first_transferred_to_l2_at_block_number VARCHAR(255),
  first_transferred_to_l2_at_tx VARCHAR(255),
  last_transferred_to_l2_at VARCHAR(255),
  last_transferred_to_l2_at_block_number VARCHAR(255),
  last_transferred_to_l2_at_tx VARCHAR(255),
  staked_tokens_transferred_to_l2 VARCHAR(255),
  id_on_l2 VARCHAR(255),
  id_on_l1 VARCHAR(255),
  delegated_capacity VARCHAR(255),
  token_capacity VARCHAR(255),
  available_stake VARCHAR(255),
  delegated_tokens VARCHAR(255),
  own_stake_ratio VARCHAR(255),
  delegated_stake_ratio VARCHAR(255),
  delegator_shares VARCHAR(255),
  delegation_exchange_rate VARCHAR(255),
  indexing_reward_cut INT,
  indexing_reward_effective_cut VARCHAR(255),
  over_delegation_dilution VARCHAR(255),
  delegator_query_fees VARCHAR(255),
  query_fee_cut INT,
  query_fee_effective_cut VARCHAR(255),
  delegator_parameter_cooldown INT,
  last_delegation_parameter_update INT,
  forced_closures INT,
  total_return VARCHAR(255),
  annualized_return VARCHAR(255),
  staking_efficiency VARCHAR(255),
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
}

type QueryIndexersData struct {
	Indexers []Indexer `json:"indexers"`
}

type QueryIndexersResponse struct {
	Data QueryIndexersData `json:"data"`
}
