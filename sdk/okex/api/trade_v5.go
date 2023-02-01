package api

import (
	"fmt"

	"github.com/bad-superman/test/logging"
)

const (
	_tradeOrderURL       = "/api/v5/trade/order"
	_tradeCancelOrderURL = "/api/v5/trade/cancel-order"
	_tradeAmendOrder     = "/api/v5/trade/amend-order"
)

// 下单
// 只有当您的账户有足够的资金才能下单。
// 该接口支持带单合约的下单，但不支持为带单合约平仓
// https://aws.okx.com/docs-v5/zh/#rest-api-trade-place-order
func (o *OkexClient) TradeOrder(req *Order) ([]TradeOrderData, error) {
	res := new(TradeOrderResp)
	err := o.post(_tradeOrderURL, req, res)
	if err != nil {
		logging.Errorf("OkexClient|TradeOrder error,err:%v", err)
		return res.Data, err
	}
	return res.Data, nil
}

// 撤单
// https://aws.okx.com/docs-v5/zh/#rest-api-trade-cancel-order
// 撤单返回sCode等于0不能严格认为该订单已经被撤销，只表示您的撤单请求被系统服务器所接受，撤单结果以订单频道推送的状态或者查询订单状态为准
func (o *OkexClient) CancelOrder(instId, ordId, clOrdId string) error {
	req := CancelOrderReq{
		InstID:  instId,
		OrdID:   ordId,
		ClOrdID: clOrdId,
	}
	res := new(CancelOrderResp)
	err := o.post(_tradeCancelOrderURL, req, res)
	if err != nil {
		logging.Errorf("OkexClient|TradeOrder error,err:%v", err)
		return err
	}
	if len(res.Data) == 0 {
		logging.Error("OkexClient|TradeOrder error,no data")
		return err
	}
	code := res.Data[0].SCode
	msg := res.Data[0].SMsg
	if code == "0" {
		return nil
	}
	err = fmt.Errorf("code:%s msg:%s", code, msg)
	logging.Errorf("OkexClient|ModifyOrder error,err:%v", err)
	return err
}

// 获取订单信息
// https://aws.okx.com/docs-v5/zh/#rest-api-trade-get-order-details
func (o *OkexClient) GetOrderInfo(instId, ordId, clOrdId string) ([]Order, error) {
	req := GetOrderInfoReq{
		InstID:  instId,
		OrdID:   ordId,
		ClOrdID: clOrdId,
	}
	res := new(GetOrderInfoResp)
	err := o.post(_tradeOrderURL, req, res)
	if err != nil {
		logging.Errorf("OkexClient|GetOrderInfo error,err:%v", err)
	}
	return res.Data, nil
}

// 修改订单
// 修改当前未成交的挂单
// https://aws.okx.com/docs-v5/zh/#rest-api-trade-amend-order
func (o *OkexClient) ModifyOrder(req *ModifyOrderReq) error {
	res := new(ModifyOrderResp)
	err := o.post(_tradeAmendOrder, req, res)
	if err != nil {
		logging.Errorf("OkexClient|ModifyOrder error,err:%v", err)
		return err
	}
	if len(res.Data) == 0 {
		logging.Error("OkexClient|ModifyOrder error,no data")
		return err
	}
	code := res.Data[0].SCode
	msg := res.Data[0].SMsg
	if code == "0" {
		return nil
	}
	err = fmt.Errorf("code:%s msg:%s", code, msg)
	logging.Errorf("OkexClient|ModifyOrder error,err:%v", err)
	return err
}
