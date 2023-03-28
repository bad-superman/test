package api

import (
	"github.com/bad-superman/test/logging"
)

const (
	_exchangeRate           = "/api/v5/market/exchange-rate"
	_historyIndexCandlesURL = "/api/v5/market/history-index-candles"
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

// func (o *OkexClient) HistoryIndexCandles(instId string, after, before time.Time, bar okex.BarSize, limit int) ([]CandleData, error) {
// 	res := new(CandlesResp)
// 	afterStr := after.Unix()
// 	err := o.get(_historyIndexCandlesURL, res)
// 	if err != nil {
// 		logging.Errorf("OkexClient|ExchangeRate error,err:%v", err)
// 		return res.Data, err
// 	}
// 	return res.Data, err
// }
