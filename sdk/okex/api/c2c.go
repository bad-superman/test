package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/bad-superman/test/logging"
)

const (
	_userAgent        = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"
	_c2cOrderBooksURL = "/v3/c2c/tradingOrders/books?t=%d&quoteCurrency=%s&baseCurrency=%s&side=all&paymentMethod=all&userType=blockTrade&receivingAds=false&urlId=0"
)

// 获取大宗交易otc价格
// @@params
// quoteCurrency 发币
// baseCurrency 虚拟币
func (o *OkexClient) C2COrderBooks(quoteCurrency, baseCurrency string) (C2COrderBooksData, error) {
	var (
		t   = time.Now().UnixMilli()
		res = new(C2COrderBooksResp)
	)
	url := fmt.Sprintf(_c2cOrderBooksURL, t, quoteCurrency, baseCurrency)
	url = _okexRestApiHost + url
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("user-agent", _userAgent)
	resp, err := o.client.Do(req)
	if err != nil {
		logging.Errorf("OkexClient|C2COrderBooks error,err:%v", err)
		return res.Data, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		logging.Errorf("OkexClient|C2COrderBooks call error,code:%d,resp:%s",
			resp.StatusCode, string(body))
		return res.Data, fmt.Errorf("http code %d", resp.StatusCode)
	}
	err = json.Unmarshal(body, res)
	if err != nil {
		logging.Errorf("OkexClient|C2COrderBooks error,err:%v", err)
		return res.Data, err
	}
	return res.Data, nil
}
