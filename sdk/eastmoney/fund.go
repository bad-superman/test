package eastmoney

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/bad-superman/test/logging"
)

const (
	_host     = "http://api.fund.eastmoney.com"
	_fundLSJZ = "/f10/lsjz?callback=&fundCode=%s&pageIndex=%d&pageSize=%d&startDate=&endDate=&_=%d"
)

// 查询基金历史净值
func (e *EastmoneyClient) FundHistoricalNetValue(fundCode string, page, size int) (*LSJZResp, error) {
	url := fmt.Sprintf(_fundLSJZ, fundCode, page, size, time.Now().UnixMilli())
	url = _host + url
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Referer", "http://fundf10.eastmoney.com/")
	resp, err := e.client.Do(req)
	if err != nil {
		logging.Errorf("EastmoneyClient|FundHistoricalNetValue error,err:%v", err)
		return nil, err
	}
	result := new(LSJZResp)
	respBody, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		logging.Errorf("OkexClient|http code error,code:%d,resp:%s",
			resp.StatusCode, string(respBody))
		return nil, fmt.Errorf("http code error,%d", resp.StatusCode)
	}
	err = json.Unmarshal(respBody, result)
	return result, err
}
