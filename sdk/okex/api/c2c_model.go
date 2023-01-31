package api

type C2COrderBooksResp struct {
	Code         int64             `json:"code"`
	Data         C2COrderBooksData `json:"data"`
	DetailMsg    string            `json:"detailMsg"`
	ErrorCode    string            `json:"error_code"`
	ErrorMessage string            `json:"error_message"`
	Msg          string            `json:"msg"`
	RequestID    string            `json:"requestId"`
}

type C2COrderBooksData struct {
	Buy       []C2CBuyElement `json:"buy"`
	Recommend C2CRecommend    `json:"recommend"`
	Sell      []C2CBuyElement `json:"sell"`
}

type C2CBuyElement struct {
	AlreadyTraded             bool        `json:"alreadyTraded"`
	AvailableAmount           string      `json:"availableAmount"`
	BaseCurrency              string      `json:"baseCurrency"`
	Black                     bool        `json:"black"`
	CancelledOrderQuantity    int64       `json:"cancelledOrderQuantity"`
	CompletedOrderQuantity    int64       `json:"completedOrderQuantity"`
	CompletedRate             string      `json:"completedRate"`
	CreatorType               string      `json:"creatorType"`
	GuideUpgradeKyc           bool        `json:"guideUpgradeKyc"`
	ID                        string      `json:"id"`
	Intention                 bool        `json:"intention"`
	MaxCompletedOrderQuantity int64       `json:"maxCompletedOrderQuantity"`
	MaxUserCreatedDate        int64       `json:"maxUserCreatedDate"`
	MerchantID                string      `json:"merchantId"`
	MinCompletedOrderQuantity int64       `json:"minCompletedOrderQuantity"`
	MinCompletionRate         string      `json:"minCompletionRate"`
	MinKycLevel               int64       `json:"minKycLevel"`
	MinSellOrders             int64       `json:"minSellOrders"`
	Mine                      bool        `json:"mine"`
	NickName                  string      `json:"nickName"`
	PaymentMethods            []string    `json:"paymentMethods"`
	Price                     string      `json:"price"`
	PublicUserID              string      `json:"publicUserId"`
	QuoteCurrency             string      `json:"quoteCurrency"`
	QuoteMaxAmountPerOrder    string      `json:"quoteMaxAmountPerOrder"`
	QuoteMinAmountPerOrder    string      `json:"quoteMinAmountPerOrder"`
	QuoteScale                int64       `json:"quoteScale"`
	QuoteSymbol               string      `json:"quoteSymbol"`
	ReceivingAds              bool        `json:"receivingAds"`
	SafetyLimit               bool        `json:"safetyLimit"`
	Side                      string      `json:"side"`
	UserActiveStatusVo        interface{} `json:"userActiveStatusVo"`
	UserType                  string      `json:"userType"`
	VerificationType          int64       `json:"verificationType"`
}

type C2CRecommend struct {
	RecommendedAd interface{} `json:"recommendedAd"`
	UserGroup     int64       `json:"userGroup"`
}
