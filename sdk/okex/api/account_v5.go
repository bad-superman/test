package api

import (
	"fmt"
	"strings"

	"github.com/bad-superman/test/logging"
)

const (
	_accountBalanceURL = "/api/v5/account/balance"
)

// 获取交易账户中资金余额信息。
// ccy	String	否	币种，如 BTC
// 支持多币种查询（不超过20个），币种之间半角逗号分隔
func (o *OkexClient) AccountBalance(coins []string) ([]Datum, error) {
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
