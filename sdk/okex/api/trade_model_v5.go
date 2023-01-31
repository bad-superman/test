package api

import "github.com/bad-superman/test/sdk/okex"

type (
	TradeOrderReq struct {
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
		Sz           okex.JSONInt64      `json:"sz"`
		Pnl          okex.JSONFloat64    `json:"pnl"`
		AccFillSz    okex.JSONInt64      `json:"accFillSz"`
		FillPx       okex.JSONFloat64    `json:"fillPx"`
		FillSz       okex.JSONInt64      `json:"fillSz"`
		FillTime     okex.JSONFloat64    `json:"fillTime"`
		AvgPx        okex.JSONFloat64    `json:"avgPx"`
		Lever        okex.JSONFloat64    `json:"lever"`
		Fee          okex.JSONFloat64    `json:"fee"`
		Rebate       okex.JSONFloat64    `json:"rebate"`
		State        okex.OrderState     `json:"state"`
		TdMode       okex.TradeMode      `json:"tdMode"`
		PosSide      okex.PositionSide   `json:"posSide"`
		Side         okex.OrderSide      `json:"side"`
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
