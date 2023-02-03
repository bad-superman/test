package api

import "github.com/bad-superman/test/sdk/okex"

type (
	Order struct {
		InstID       string              `json:"instId"`
		Ccy          string              `json:"ccy"`
		OrdID        string              `json:"ordId"`
		ClOrdID      string              `json:"clOrdId"`
		TradeID      string              `json:"tradeId"`
		Tag          string              `json:"tag"`
		Category     string              `json:"category"`
		FeeCcy       string              `json:"feeCcy"`
		RebateCcy    string              `json:"rebateCcy"`
		QuickMgnType string              `json:"quickMgnType"`
		ReduceOnly   string              `json:"reduceOnly"`
		Px           okex.JSONFloat64    `json:"px"`
		Sz           okex.JSONFloat64    `json:"sz"`
		Pnl          okex.JSONFloat64    `json:"pnl"`
		AccFillSz    okex.JSONFloat64    `json:"accFillSz"`
		FillPx       okex.JSONFloat64    `json:"fillPx"`
		FillSz       okex.JSONFloat64    `json:"fillSz"`
		FillTime     okex.JSONFloat64    `json:"fillTime"`
		AvgPx        okex.JSONFloat64    `json:"avgPx"`
		Lever        okex.JSONFloat64    `json:"lever"`
		Fee          okex.JSONFloat64    `json:"fee"`
		Rebate       okex.JSONFloat64    `json:"rebate"`
		State        okex.OrderState     `json:"state"`
		TdMode       okex.TradeMode      `json:"tdMode"`
		PosSide      okex.PositionSide   `json:"posSide,omitempty"`
		Side         okex.OrderSide      `json:"side,omitempty"`
		OrdType      okex.OrderType      `json:"ordType"`
		InstType     okex.InstrumentType `json:"instType"`
		TgtCcy       okex.QuantityType   `json:"tgtCcy"`
		UTime        okex.JSONTime       `json:"uTime"`
		CTime        okex.JSONTime       `json:"cTime"`
	}

	TradeOrderResp struct {
		Code string           `json:"code"`
		Msg  string           `json:"msg"`
		Data []TradeOrderData `json:"data"`
	}

	TradeOrderData struct {
		ClOrdID string `json:"clOrdId"`
		OrdID   string `json:"ordId"`
		Tag     string `json:"tag"`
		SCode   string `json:"sCode"`
		SMsg    string `json:"sMsg"`
	}
)

type (
	CancelOrderReq struct {
		InstID  string `json:"instId"`
		OrdID   string `json:"ordId,omitempty"`
		ClOrdID string `json:"clOrdId,omitempty"`
	}

	CancelOrderResp struct {
		Code string           `json:"code"`
		Msg  string           `json:"msg"`
		Data []TradeOrderData `json:"data"`
	}

	CancelOrderData struct {
		ClOrdID string `json:"clOrdId"`
		OrdID   string `json:"ordId"`
		SCode   string `json:"sCode"` // 事件执行结果的code，0代表成功
		SMsg    string `json:"sMsg"`
	}
)

type (
	GetOrderInfoReq struct {
		InstID  string `json:"instId"`
		OrdID   string `json:"ordId,omitempty"`
		ClOrdID string `json:"clOrdId,omitempty"`
	}

	GetOrderInfoResp struct {
		Code string  `json:"code"`
		Msg  string  `json:"msg"`
		Data []Order `json:"data"`
	}
)

type (
	ModifyOrderReq struct {
		InstID    string           `json:"instId"`
		CxlOnFail bool             `json:"cxlOnFail,omitempty"`
		OrdID     string           `json:"ordId,omitempty"`
		ClOrdID   string           `json:"clOrdId,omitempty"`
		ReqID     string           `json:"reqId,omitempty"`
		NewSz     okex.JSONFloat64 `json:"newSz,omitempty"`
		NewPx     okex.JSONFloat64 `json:"newPx,omitempty"`
	}

	ModifyOrderResp struct {
		Code string            `json:"code"`
		Msg  string            `json:"msg"`
		Data []ModifyOrderData `json:"data"`
	}

	ModifyOrderData struct {
		ClOrdID string `json:"clOrdId"`
		OrdID   string `json:"ordId"`
		ReqID   string `json:"reqId,omitempty"`
		SCode   string `json:"sCode"` // 事件执行结果的code，0代表成功
		SMsg    string `json:"sMsg"`
	}
)

type (
	PendingOrderReq struct {
		InstType okex.InstrumentType `url:"instType"`
		Uly      string              `url:"uly"`
		InstId   string              `url:"instId"`
		OrdType  okex.OrderType      `url:"ordType"`
		State    okex.OrderState     `url:"state"`
		After    string              `url:"after"`
		Before   string              `url:"before"`
		Limit    string              `url:"limit"`
	}
)
