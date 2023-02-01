package api

type ExchangeRateResp struct {
	Code string             `json:"code"`
	Msg  string             `json:"msg"`
	Data []ExchangeRateData `json:"data"`
}

type ExchangeRateData struct {
	UsdCny string `json:"usdCny"`
}
