package api

import "github.com/bad-superman/test/sdk/okex"

type InstrumentData struct {
	Alias        okex.AliasType      `json:"alias"`
	BaseCcy      string              `json:"baseCcy"`
	Category     string              `json:"category"`
	CtMult       string              `json:"ctMult"`
	CtType       string              `json:"ctType"`
	CtVal        string              `json:"ctVal"`
	CtValCcy     string              `json:"ctValCcy"`
	ExpTime      okex.JSONInt64      `json:"expTime"`
	InstFamily   string              `json:"instFamily"`
	InstID       string              `json:"instId"`
	InstType     okex.InstrumentType `json:"instType"`
	Lever        string              `json:"lever"`
	ListTime     okex.JSONInt64      `json:"listTime"`
	LotSz        string              `json:"lotSz"`
	MaxIcebergSz string              `json:"maxIcebergSz"`
	MaxLmtSz     string              `json:"maxLmtSz"`
	MaxMktSz     string              `json:"maxMktSz"`
	MaxStopSz    string              `json:"maxStopSz"`
	MaxTriggerSz string              `json:"maxTriggerSz"`
	MaxTwapSz    string              `json:"maxTwapSz"`
	MinSz        string              `json:"minSz"`
	OptType      okex.OptionType     `json:"optType"`
	QuoteCcy     string              `json:"quoteCcy"`
	SettleCcy    string              `json:"settleCcy"`
	State        string              `json:"state"`
	Stk          string              `json:"stk"`
	TickSz       string              `json:"tickSz"`
	Uly          string              `json:"uly"`
}

type InstrumentsResp struct {
	Code string           `json:"code"`
	Msg  string           `json:"msg"`
	Data []InstrumentData `json:"data"`
}
