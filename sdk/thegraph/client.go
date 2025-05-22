package thegraph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	_endpoint = "https://gateway.thegraph.com/api/subgraphs/id/"
)

// Client 是 thegraph 的 SDK 客户端
// Endpoint: GraphQL API 地址
// ApiKey: 认证用的 API Key（可选）
type Client struct {
	ApiKey string
}

// NewClient 创建一个新的 thegraph 客户端
func NewClient(apiKey string) *Client {
	return &Client{
		ApiKey: apiKey,
	}
}

// Query 方法，发送 GraphQL 请求并将 data 字段解析到 result 对象
// payload: GraphQLPayload结构体
// result: 需要解析的对象指针
func (c *Client) Query(url string, payload GraphQLPayload, result interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal payload failed: %w", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("create request failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if c.ApiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.ApiKey)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response failed: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error: %s, body: %s", resp.Status, string(respBody))
	}

	// 解析响应
	if err := json.Unmarshal(respBody, result); err != nil {
		return fmt.Errorf("unmarshal response failed: %w", err)
	}
	return nil
}

const (
	_queryIndexersQueryStr = `
	{
		indexers(first: %d, skip: %d) {
			id
			createdAt
			url
			geoHash
			defaultDisplayName
			stakedTokens
			allocatedTokens
			unstakedTokens
			lockedTokens
			tokensLockedUntil
			allocationCount
			totalAllocationCount
			queryFeesCollected
			queryFeeRebates
			rewardsEarned
			indexerIndexingRewards
			delegatorIndexingRewards
			indexerRewardsOwnGenerationRatio
			transferredToL2
			firstTransferredToL2At
			firstTransferredToL2AtBlockNumber
			firstTransferredToL2AtTx
			lastTransferredToL2At
			lastTransferredToL2AtBlockNumber
			lastTransferredToL2AtTx
			stakedTokensTransferredToL2
			idOnL2
			idOnL1
			delegatedCapacity
			tokenCapacity
			availableStake
			delegatedTokens
			ownStakeRatio
			delegatedStakeRatio
			delegatorShares
			delegationExchangeRate
			indexingRewardCut
			indexingRewardEffectiveCut
			overDelegationDilution
			delegatorQueryFees
			queryFeeCut
			queryFeeEffectiveCut
			delegatorParameterCooldown
			lastDelegationParameterUpdate
			forcedClosures
			totalReturn
			annualizedReturn
			stakingEfficiency
		}
	}`
)

func (c *Client) QueryIndexers(subgraphId string, first, offset int, result *QueryIndexersResponse) error {
	payload := GraphQLPayload{
		Query:         fmt.Sprintf(_queryIndexersQueryStr, first, offset),
		OperationName: "Subgraphs",
	}
	url := fmt.Sprintf("%s%s", _endpoint, subgraphId)
	return c.Query(url, payload, result)
}
