package api

type AccountBalanceResp struct {
	Code string  `json:"code"`
	Msg  string  `json:"msg"`
	Data []Datum `json:"data"`
}

type Datum struct {
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
