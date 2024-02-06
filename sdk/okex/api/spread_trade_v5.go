package api

import (
	"fmt"

	"github.com/bad-superman/test/logging"
	"github.com/google/go-querystring/query"
)

const (
	_sprdOrderURL                = "POST /api/v5/sprd/order"
	_sprdCancelOrderURL          = "POST /api/v5/sprd/cancel-order"
	_sprdMassCancelURL           = "POST /api/v5/sprd/mass-cancel"
	_sprdaMendOrderURL           = "POST /api/v5/sprd/amend-order"
	_sprdOrdersPendingURL        = "/api/v5/sprd/orders-pending"
	_sprdOrdersHistoryURL        = "/api/v5/sprd/orders-history"
	_sprdOrdersHistoryArchiveURL = "/api/v5/sprd/orders-history-archive"
	_sprdTradesURL               = "/api/v5/sprd/trades"
	_sprdSpreadsURL              = "/api/v5/sprd/spreads?baseCcy=%s&instId=%s&sprdId=%s&state=%s"
)

// 价差交易

// 下单
// https://aws.okx.com/docs-v5/zh/#spread-trading-rest-api-place-order
func (o *OkexClient) SprdOrder(req *SprdOrder) ([]TradeOrderData, error) {
	res := new(SprdOrderResp)
	err := o.post(_sprdOrderURL, req, res)
	if err != nil {
		logging.Errorf("OkexClient|SprdOrder error,err:%v", err)
		return res.Data, err
	}
	return res.Data, nil
}

// 撤单
// https://aws.okx.com/docs-v5/zh/#spread-trading-rest-api-cancel-order
// 撤单返回sCode等于0不能严格认为该订单已经被撤销，只表示您的撤单请求被系统服务器所接受，撤单结果以订单频道推送的状态或者查询订单状态为准
func (o *OkexClient) CancelSprdOrder(req *CancelOrderReq) error {
	res := new(SprdCancelOrderResp)
	err := o.post(_sprdCancelOrderURL, req, res)
	if err != nil {
		logging.Errorf("OkexClient|CancelSprdOrder error,err:%v", err)
		return err
	}
	if len(res.Data) == 0 {
		logging.Error("OkexClient|CancelSprdOrder error,no data")
		return err
	}
	code := res.Data[0].SCode
	msg := res.Data[0].SMsg
	if code == "0" {
		return nil
	}
	err = fmt.Errorf("code:%s msg:%s", code, msg)
	logging.Errorf("OkexClient|CancelSprdOrder error,err:%v", err)
	return err
}

// 全部撤单
// https://aws.okx.com/docs-v5/zh/#spread-trading-rest-api-cancel-all-orders
// 返回结果中result=true 代表您的请求已被成功接收，并将会被处理。撤单的实际结果会通过`sprd-orders`频道推送。
func (o *OkexClient) CancelAllSprdOrder(req *SprdMassCancelReq) error {
	res := new(SprdMassCancelResp)
	err := o.post(_sprdMassCancelURL, req, res)
	if err != nil {
		logging.Errorf("OkexClient|CancelAllSprdOrder error,err:%v", err)
		return err
	}
	if len(res.Data) == 0 {
		logging.Error("OkexClient|CancelAllSprdOrder error,no data")
		return err
	}
	result := res.Data[0].Result
	if result {
		return nil
	}
	err = fmt.Errorf("code:%s msg:%s", res.Code, res.Msg)
	logging.Errorf("OkexClient|CancelAllSprdOrder error,err:%v", err)
	return err
}

// 修改订单
// https://aws.okx.com/docs-v5/zh/#spread-trading-rest-api-amend-order
// newSz 若修改订单时，订单修改后的新数量小于或等于 (accFillSz + canceledSz + pendingSettleSz)，在 pendingSettleSz 结算后，订单状态会根据 canceledSz 的不同而不同。当 canceledSz = 0，订单状态将被改为 filled；当 canceledSz > 0，订单状态将被改为 canceled。
// 修改订单返回sCode等于0不能严格认为该订单已经被修改，只表示您的修改订单请求被系统服务器所接受，改单结果以订单频道推送的状态或者查询订单状态为准
func (o *OkexClient) ModifySprdOrder(req *SprdAmendOrderReq) error {
	res := new(SprdAmendOrderResp)
	err := o.post(_sprdaMendOrderURL, req, res)
	if err != nil {
		logging.Errorf("OkexClient|ModifySprdOrder error,err:%v", err)
		return err
	}
	if len(res.Data) == 0 {
		logging.Error("OkexClient|ModifySprdOrder error,no data")
		return err
	}
	code := res.Data[0].SCode
	msg := res.Data[0].SMsg
	if code == "0" {
		return nil
	}
	err = fmt.Errorf("code:%s msg:%s", code, msg)
	logging.Errorf("OkexClient|ModifySprdOrder error,err:%v", err)
	return err
}

// 获取订单信息
// https://aws.okx.com/docs-v5/zh/#spread-trading-rest-api-get-order-details
// 订单数量等式: pendingFillSz + canceledSz + accFillSz = sz
func (o *OkexClient) GetSprdOrderInfo(ordId, clOrdId string) (SprdOrderInfoData, error) {
	res := new(GetSprdOrderResp)
	url := fmt.Sprintf("%s?ordId=%s&clOrdId=%s", _sprdOrderURL, ordId, clOrdId)
	err := o.get(url, res)
	if err != nil {
		logging.Errorf("OkexClient|GetSprdOrderInfo error,err:%v", err)
		return SprdOrderInfoData{}, err
	}
	if len(res.Data) == 0 {
		return SprdOrderInfoData{}, fmt.Errorf("no data")
	}
	return res.Data[0], nil
}

// 获取未成交订单列表
// https://aws.okx.com/docs-v5/zh/#spread-trading-rest-api-get-active-orders
func (o *OkexClient) SprdPendingOrders(req *SprdPendingOrderReq) ([]SprdOrderInfoData, error) {
	res := new(GetSprdOrderResp)
	v, _ := query.Values(req)
	url := _sprdOrdersPendingURL + "?" + v.Encode()
	err := o.get(url, res)
	if err != nil {
		logging.Errorf("OkexClient|SprdPendingOrders error,err:%v", err)
	}
	return res.Data, nil
}

// 获取历史订单记录（近21天)
//https://aws.okx.com/docs-v5/zh/#spread-trading-rest-api-get-active-orders
func (o *OkexClient) GetSprdOrdersHistory(req *SprdOrderHistoryReq) ([]SprdOrderInfoData, error) {
	res := new(GetSprdOrderResp)
	v, _ := query.Values(req)
	url := _sprdOrdersHistoryURL + "?" + v.Encode()
	err := o.get(url, res)
	if err != nil {
		logging.Errorf("OkexClient|GetSprdOrdersHistory error,err:%v", err)
	}
	return res.Data, nil
}

// 获取历史订单记录（近三月)
//https://aws.okx.com/docs-v5/zh/#spread-trading-rest-api-get-orders-history-last-3-months
func (o *OkexClient) GetSprdOrdersArchiveHistory(req *SprdOrderHistoryReq) ([]SprdOrderInfoData, error) {
	res := new(GetSprdOrderResp)
	v, _ := query.Values(req)
	url := _sprdOrdersHistoryArchiveURL + "?" + v.Encode()
	err := o.get(url, res)
	if err != nil {
		logging.Errorf("OkexClient|GetSprdOrdersArchiveHistory error,err:%v", err)
	}
	return res.Data, nil
}

func (o *OkexClient) SprdSpreads(baseCcy, instId, sprdId, state string) ([]SpreadData, error) {
	res := new(SprdSpreadsResp)
	url := fmt.Sprintf(_sprdSpreadsURL, baseCcy, instId, sprdId, state)
	err := o.get(url, res)
	if err != nil {
		logging.Errorf("OkexClient|SprdSpreads error,err:%v", err)
		return res.Data, err
	}
	return res.Data, err
}
