package api

type AccountBalanceResp struct {
	Code string               `json:"code"`
	Msg  string               `json:"msg"`
	Data []AccountBalanceData `json:"data"`
}

type AccountBalanceData struct {
	AdjEq       string              `json:"adjEq"`
	Details     []map[string]string `json:"details"`
	Imr         string              `json:"imr"`
	ISOEq       string              `json:"isoEq"`
	MgnRatio    string              `json:"mgnRatio"`
	Mmr         string              `json:"mmr"`
	NotionalUsd string              `json:"notionalUsd"`
	OrdFroz     string              `json:"ordFroz"`
	TotalEq     string              `json:"totalEq"`
	UTime       string              `json:"uTime"`
}

type AccountPositionsResp struct {
	Code string                 `json:"code"`
	Msg  string                 `json:"msg"`
	Data []AccountPositionsData `json:"data"`
}

type AccountPositionsData struct {
	Adl            string           `json:"adl"`
	AvailPos       string           `json:"availPos"`
	AvgPx          string           `json:"avgPx"`
	CTime          string           `json:"cTime"`
	Ccy            string           `json:"ccy"`
	DeltaBS        string           `json:"deltaBS"`
	DeltaPA        string           `json:"deltaPA"`
	GammaBS        string           `json:"gammaBS"`
	GammaPA        string           `json:"gammaPA"`
	Imr            string           `json:"imr"`
	InstID         string           `json:"instId"`
	InstType       string           `json:"instType"`
	Interest       string           `json:"interest"`
	Last           string           `json:"last"`
	UsdPx          string           `json:"usdPx"`
	Lever          string           `json:"lever"`
	Liab           string           `json:"liab"`
	LiabCcy        string           `json:"liabCcy"`
	LiqPx          string           `json:"liqPx"`
	MarkPx         string           `json:"markPx"`
	Margin         string           `json:"margin"`
	MgnMode        string           `json:"mgnMode"`
	MgnRatio       string           `json:"mgnRatio"`
	Mmr            string           `json:"mmr"`
	NotionalUsd    string           `json:"notionalUsd"`
	OptVal         string           `json:"optVal"`
	PTime          string           `json:"pTime"`
	Pos            string           `json:"pos"`
	PosCcy         string           `json:"posCcy"`
	PosID          string           `json:"posId"`
	PosSide        string           `json:"posSide"`
	SpotInUseAmt   string           `json:"spotInUseAmt"`
	SpotInUseCcy   string           `json:"spotInUseCcy"`
	ThetaBS        string           `json:"thetaBS"`
	ThetaPA        string           `json:"thetaPA"`
	TradeID        string           `json:"tradeId"`
	BizRefID       string           `json:"bizRefId"`
	BizRefType     string           `json:"bizRefType"`
	QuoteBAL       string           `json:"quoteBal"`
	BaseBAL        string           `json:"baseBal"`
	BaseBorrowed   string           `json:"baseBorrowed"`
	BaseInterest   string           `json:"baseInterest"`
	QuoteBorrowed  string           `json:"quoteBorrowed"`
	QuoteInterest  string           `json:"quoteInterest"`
	UTime          string           `json:"uTime"`
	Upl            string           `json:"upl"`
	UplRatio       string           `json:"uplRatio"`
	VegaBS         string           `json:"vegaBS"`
	VegaPA         string           `json:"vegaPA"`
	CloseOrderAlgo []CloseOrderAlgo `json:"closeOrderAlgo"`
}

type CloseOrderAlgo struct {
	AlgoID          string `json:"algoId"`
	SlTriggerPx     string `json:"slTriggerPx"`
	SlTriggerPxType string `json:"slTriggerPxType"`
	TpTriggerPx     string `json:"tpTriggerPx"`
	TpTriggerPxType string `json:"tpTriggerPxType"`
	CloseFraction   string `json:"closeFraction"`
}

type PositionRiskResp struct {
	Code string             `json:"code"`
	Data []PositionRiskData `json:"data"`
	Msg  string             `json:"msg"`
}

type PositionRiskData struct {
	AdjEq   string    `json:"adjEq"`   //美金层面有效保证金,适用于跨币种保证金模式 和组合保证金模式
	BALData []BALData `json:"balData"` //币种资产信息
	PosData []PosData `json:"posData"` //持仓详细信息
	Ts      string    `json:"ts"`
}

type BALData struct {
	Ccy   string `json:"ccy"`
	DisEq string `json:"disEq"`
	Eq    string `json:"eq"`
}

type PosData struct {
	BaseBAL     string `json:"baseBal"`
	Ccy         string `json:"ccy"`
	InstID      string `json:"instId"`
	InstType    string `json:"instType"`
	MgnMode     string `json:"mgnMode"` // 保证金模式 cross：全仓 isolated：逐仓
	NotionalCcy string `json:"notionalCcy"`
	NotionalUsd string `json:"notionalUsd"`
	Pos         string `json:"pos"`    //以张为单位的持仓数量，逐仓自主划转模式下，转入保证金后会产生pos为0的仓位
	PosCcy      string `json:"posCcy"` //仓位资产币种，仅适用于币币杠杆仓位
	PosID       string `json:"posId"`
	PosSide     string `json:"posSide"`  //持仓方向 long：双向持仓多头 short：双向持仓空头 net：单向持仓（交割/永续/期权：pos为正代表多头，pos为负代表空头。币币杠杆：posCcy为交易货币时，代表多头；posCcy为计价货币时，代表空头。）
	QuoteBAL    string `json:"quoteBal"` //计价币余额 ，适用于 币币杠杆（逐仓自主划转模式和逐仓一键借币模式）
}
