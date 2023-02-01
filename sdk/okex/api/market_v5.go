package api

import "github.com/bad-superman/test/logging"

const (
	_exchangeRate = "/api/v5/market/exchange-rate"
)

func (o *OkexClient) ExchangeRate() ([]ExchangeRateData, error) {
	res := new(ExchangeRateResp)
	err := o.get(_exchangeRate, res)
	if err != nil {
		logging.Errorf("OkexClient|ExchangeRate error,err:%v", err)
		return res.Data, err
	}
	return res.Data, err
}
