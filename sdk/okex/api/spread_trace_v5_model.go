package api

import "github.com/bad-superman/test/sdk/okex"

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
