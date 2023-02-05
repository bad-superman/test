package v2

// 参照文档 https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq
type Text struct {
	Link Link `json:"text"`
	Base
}

type MarkDown struct {
	Markdown Link `json:"markdown"`
	Base
}
type Links struct {
	Link Link `json:"link"`
	Base
}

type ActionCardAll struct {
	AllActionCard AllActionCard `json:"actionCard"`
	Base
}

type ActionCardSingel struct {
	SingelActionCard SingelActionCard `json:"actionCard"`
	Base
}

type FeedCard struct {
	FeedCard FeedCardAll `json:"feedCard"`
	Base
}

type Btns struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

type FeedCardAll struct {
	Links Link `json:"links"`
}

type At struct {
	AtMobiles []string `json:"atMobiles"` //手机号码
	IsAtAll   bool     `json:"isAtAll"`   // @所有人
}

type Base struct {
	Msgtype string `json:"msgtype"`
	At      At     `json:"at"`
}
type Link struct {
	Content    string `json:"content"`
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicURL     string `json:"picUrl"`
	MessageURL string `json:"messageUrl"`
}

type SingelActionCard struct {
	Title          string `json:"title"`
	Text           string `json:"text"`
	BtnOrientation string `json:"btnOrientation"`
	Btns           []Btns `json:"btns"`
}

type AllActionCard struct {
	Title          string `json:"title"`
	Text           string `json:"text"`
	BtnOrientation string `json:"btnOrientation"`
	SingleTitle    string `json:"singleTitle"`
	SingleURL      string `json:"singleURL"`
}

const (
	Text_type = iota
	Markdown_type
	ActionCardAll_type
	ActionCardSingel_type
	FeedCard_type
	Links_type
)

func NewText() *Text {
	msg := &Text{}
	msg.Msgtype = "text"
	return msg
}

func NewMarkDown() *MarkDown {
	msg := &MarkDown{}
	msg.Msgtype = "markdown"
	return msg
}
func NewActionCardAll() *ActionCardAll {
	msg := &ActionCardAll{}
	msg.Msgtype = "actionCard"
	return msg
}
func NewActionCardSingel() *ActionCardSingel {
	msg := &ActionCardSingel{}
	msg.Msgtype = "actionCard"
	return msg
}

func NewFeedCard() *FeedCard {
	msg := &FeedCard{}
	msg.Msgtype = "feedCard"
	return msg
}

func NewLinks() *Links {
	msg := &Links{}
	msg.Msgtype = "link"
	return msg
}
