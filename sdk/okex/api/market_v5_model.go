package api

import "github.com/bad-superman/test/sdk/okex"

type ExchangeRateResp struct {
	Code string             `json:"code"`
	Msg  string             `json:"msg"`
	Data []ExchangeRateData `json:"data"`
}

type ExchangeRateData struct {
	UsdCny string `json:"usdCny"`
}

type (
	CandlesResp struct {
		Code string       `json:"code"`
		Msg  string       `json:"msg"`
		Data []CandleData `json:"data"`
	}

	CandleData []okex.JSONFloat64
)
