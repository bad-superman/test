package eastmoney

type (
	LSJZResp struct {
		Data       LSJZData    `json:"Data"`
		ErrCode    int64       `json:"ErrCode"`
		ErrMsg     interface{} `json:"ErrMsg"`
		TotalCount int64       `json:"TotalCount"`
		Expansion  interface{} `json:"Expansion"`
		PageSize   int64       `json:"PageSize"`
		PageIndex  int64       `json:"PageIndex"`
	}

	LSJZData struct {
		LSJZList  []LSJZList  `json:"LSJZList"`
		FundType  string      `json:"FundType"`
		SYType    interface{} `json:"SYType"`
		IsNewType bool        `json:"isNewType"`
		Feature   interface{} `json:"Feature"`
	}

	LSJZList struct {
		FSRQ      string      `json:"FSRQ"`  //净值日期
		DWJZ      string      `json:"DWJZ"`  //单位净值
		LJJZ      string      `json:"LJJZ"`  // 历史净值
		JZZZL     string      `json:"JZZZL"` // 日增长率
		SGZT      string      `json:"SGZT"`  //申购状态
		SHZT      string      `json:"SHZT"`  //赎回状态
		FHFCZ     string      `json:"FHFCZ"`
		FHFCBZ    string      `json:"FHFCBZ"`
		NAVTYPE   string      `json:"NAVTYPE"`
		SDATE     interface{} `json:"SDATE"`
		ACTUALSYI string      `json:"ACTUALSYI"`
		DTYPE     interface{} `json:"DTYPE"`
		FHSP      string      `json:"FHSP"`
	}
)
