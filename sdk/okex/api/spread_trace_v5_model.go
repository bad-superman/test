package api

import "github.com/bad-superman/test/sdk/okex"

type (
	SprdOrder struct {
		SprdId  okex.SprdID      `json:"sprdId"`
		ClOrdId string           `json:"clOrdId,omitempty"`
		Tag     string           `json:"tag,omitempty"`
		Side    okex.OrderSide   `json:"side"`
		OrdType okex.OrderType   `json:"ordType"`
		Sz      okex.JSONFloat64 `json:"sz"`
		Px      okex.JSONFloat64 `json:"px"`
	}

	SprdOrderResp struct {
		Code string           `json:"code"`
		Msg  string           `json:"msg"`
		Data []TradeOrderData `json:"data"`
	}

	SprdCancelOrderReq struct {
		OrdID   string `json:"ordId,omitempty"`
		ClOrdID string `json:"clOrdId,omitempty"`
	}

	SprdCancelOrderResp struct {
		Code string           `json:"code"`
		Msg  string           `json:"msg"`
		Data []TradeOrderData `json:"data"`
	}

	SprdMassCancelReq struct {
		SprdId okex.SprdID `json:"sprdId"`
	}

	SprdMassCancelData struct {
		Result bool `json:"result"`
	}

	SprdMassCancelResp struct {
		Code string               `json:"code"`
		Msg  string               `json:"msg"`
		Data []SprdMassCancelData `json:"data"`
	}

	SprdAmendOrderReq struct {
		OrdID   string           `json:"ordId,omitempty"`
		ClOrdId string           `json:"clOrdId,omitempty"`
		ReqId   string           `json:"reqId,omitempty"`
		NewSz   okex.JSONFloat64 `json:"newSz"`
		NewPx   okex.JSONFloat64 `json:"newPx"`
	}

	SprdAmendOrderResp struct {
		Code string           `json:"code"`
		Msg  string           `json:"msg"`
		Data []TradeOrderData `json:"data"`
	}

	SprdOrderInfoData struct {
		SprdID          okex.SprdID      `json:"sprdId"`
		OrdID           string           `json:"ordId"`
		ClOrdID         string           `json:"clOrdId"`
		Tag             string           `json:"tag"`
		Px              okex.JSONFloat64 `json:"px"`
		Sz              okex.JSONFloat64 `json:"sz"`
		OrdType         okex.OrderType   `json:"ordType"`
		Side            okex.OrderSide   `json:"side"`
		FillSz          okex.JSONFloat64 `json:"fillSz"`
		FillPx          okex.JSONFloat64 `json:"fillPx"`
		TradeID         string           `json:"tradeId"`
		AccFillSz       okex.JSONFloat64 `json:"accFillSz"`
		PendingFillSz   okex.JSONFloat64 `json:"pendingFillSz"`
		PendingSettleSz okex.JSONFloat64 `json:"pendingSettleSz"`
		CanceledSz      okex.JSONFloat64 `json:"canceledSz"`
		State           okex.SprdState   `json:"state"`
		AvgPx           okex.JSONFloat64 `json:"avgPx"`
		CancelSource    string           `json:"cancelSource"`
		UTime           okex.JSONInt64   `json:"uTime"`
		CTime           okex.JSONInt64   `json:"cTime"`
	}

	GetSprdOrderResp struct {
		Code string              `json:"code"`
		Msg  string              `json:"msg"`
		Data []SprdOrderInfoData `json:"data"`
	}
)

type (
	SprdPendingOrderReq struct {
		SprdId  okex.SprdID    `url:"sprdId"`
		OrdType okex.OrderType `url:"ordType"`
		State   okex.SprdState `url:"state"`
		BeginId string         `url:"beginId"` //请求的起始订单ID，请求此ID之后（更新的数据）的分页内容，不包括 beginId
		EndId   string         `url:"endId"`   //请求的结束订单ID，请求此ID之前（更旧的数据）的分页内容，不包括 endId
		Limit   okex.JSONInt64 `url:"limit"`
	}

	SprdOrderHistoryReq struct {
		SprdId  okex.SprdID    `url:"sprdId"`
		OrdType okex.OrderType `url:"ordType"`
		State   okex.SprdState `url:"state"`
		BeginId string         `url:"beginId"`
		EndId   string         `url:"endId"`
		Begin   okex.JSONInt64 `url:"Begin"` //筛选的开始时间戳，Unix 时间戳为毫秒数格式
		End     okex.JSONInt64 `url:"end"`   //筛选的结束时间戳，Unix 时间戳为毫秒数格式
		Limit   okex.JSONInt64 `url:"limit"`
	}
)

type SpreadLeg struct {
	InstID string `json:"instId"`
	Side   string `json:"side"`
}

type SpreadData struct {
	SprdID   okex.SprdID      `json:"sprdId"`
	SprdType okex.SprdType    `json:"sprdType"`
	State    okex.SprdState   `json:"state"`
	BaseCcy  string           `json:"baseCcy"`
	SzCcy    string           `json:"szCcy"`
	QuoteCcy string           `json:"quoteCcy"`
	TickSz   okex.JSONFloat64 `json:"tickSz"`
	MinSz    okex.JSONInt64   `json:"minSz"`
	LotSz    okex.JSONInt64   `json:"lotSz"`
	ListTime okex.JSONInt64   `json:"listTime"`
	Legs     []SpreadLeg      `json:"legs"`
	ExpTime  okex.JSONInt64   `json:"expTime"`
	UTime    okex.JSONInt64   `json:"uTime"`
}

type SprdSpreadsResp struct {
	Code string       `json:"code"`
	Msg  string       `json:"msg"`
	Data []SpreadData `json:"data"`
}
