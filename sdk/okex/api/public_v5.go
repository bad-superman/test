package api

import (
	"fmt"

	"github.com/bad-superman/test/logging"
	"github.com/bad-superman/test/sdk/okex"
)

const (
	_instrumentsURL = "/api/v5/public/instruments?instType=%s&uly=%s&instFamily=%s&instId=%s"
	_markPriceURL   = "/api/v5/public/mark-price?instType=%s&uly=%s&instFamily=%s&instId=%s"
)

func (o *OkexClient) Instruments(instType okex.InstrumentType, uly, instFamily, instId string) ([]InstrumentData, error) {
	res := new(InstrumentsResp)
	url := fmt.Sprintf(_instrumentsURL, instType, uly, instFamily, instId)
	err := o.get(url, res)
	if err != nil {
		logging.Errorf("OkexClient|Instruments error,err:%v", err)
		return res.Data, err
	}
	return res.Data, err
}

// 获取标记价格
func (o *OkexClient) MarkPrice(instType okex.InstrumentType, uly, instFamily, instId string) ([]MarkPriceData, error) {
	res := new(MarkPriceResp)
	url := fmt.Sprintf(_markPriceURL, instType, uly, instFamily, instId)
	err := o.get(url, res)
	if err != nil {
		logging.Errorf("OkexClient|MarkPrice error,err:%v", err)
		return res.Data, err
	}
	return res.Data, err
}
