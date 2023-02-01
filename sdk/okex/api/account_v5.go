package api

import (
	"fmt"
	"strings"

	"github.com/bad-superman/test/logging"
)

const (
	_accountBalanceURL      = "/api/v5/account/balance"
	_accountPositionsURL    = "/api/v5/account/positions"
	_accountPositionRiskURL = "/api/v5/account/account-position-risk"
)

// 获取交易账户中资金余额信息。
// ccy	String	否	币种，如 BTC
// 支持多币种查询（不超过20个），币种之间半角逗号分隔
// https://aws.okx.com/docs-v5/zh/#rest-api-account-get-balance
func (o *OkexClient) AccountBalance(coins []string) ([]AccountBalanceData, error) {
	url := _accountBalanceURL
	if len(coins) > 0 {
		coinStr := strings.Join(coins, ",")
		url = fmt.Sprintf("%s?ccy=%s", url, coinStr)
	}
	res := new(AccountBalanceResp)
	err := o.get(url, res)
	if err != nil {
		logging.Errorf("OkexClient|AccountBalance error,err:%v", err)
		return nil, err
	}
	return res.Data, nil
}

// 查看持仓信息
// 获取该账户下拥有实际持仓的信息。账户为单向持仓模式会显示净持仓（net），账户为双向持仓模式下会分别返回多头（long）或空头（short）的仓位。按照仓位创建时间倒序排列。
// @@params
// instType 产品类型
// MARGIN：币币杠杆
// SWAP：永续合约
// FUTURES：交割合约
// OPTION：期权
// instType和instId同时传入的时候会校验instId与instType是否一致。
//
// instId 交易产品ID，如：BTC-USD-190927-5000-C
// 支持多个instId查询（不超过10个），半角逗号分隔
//
// posId 持仓ID
// 支持多个posId查询（不超过20个），半角逗号分割
// https://aws.okx.com/docs-v5/zh/#rest-api-account-get-positions
func (o *OkexClient) AccountPositions(instType string, instId, posId []string) ([]AccountPositionsData, error) {
	url := _accountPositionsURL

	params := make([]string, 0)
	if len(instType) > 0 {
		param := fmt.Sprintf("instType=%s", instType)
		params = append(params, param)
	}
	if len(instId) > 0 {
		param := fmt.Sprintf("instType=%s", strings.Join(instId, ","))
		params = append(params, param)
	}
	if len(posId) > 0 {
		param := fmt.Sprintf("posId=%s", strings.Join(posId, ","))
		params = append(params, param)
	}
	if len(params) > 0 {
		url = fmt.Sprintf("%s?%s", url, strings.Join(params, "&"))
	}
	res := new(AccountPositionsResp)
	err := o.get(url, res)
	if err != nil {
		logging.Errorf("OkexClient|AccountBalance error,err:%v", err)
		return nil, err
	}
	return res.Data, nil
}

// 查看账户整体风险。
// @@params
// instType 产品类型
// MARGIN：币币杠杆
// SWAP：永续合约
// FUTURES：交割合约
// OPTION：期权
// https://aws.okx.com/docs-v5/zh/#rest-api-account-get-account-and-position-risk
func (o *OkexClient) PositionRisk(instType string) ([]PositionRiskData, error) {
	url := _accountPositionRiskURL
	if len(instType) > 0 {
		url = fmt.Sprintf("%s?instType=%s", url, instType)
	}
	res := new(PositionRiskResp)
	err := o.get(url, res)
	if err != nil {
		logging.Errorf("OkexClient|AccountBalance error,err:%v", err)
		return nil, err
	}
	return res.Data, nil
}
