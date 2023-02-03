package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/bad-superman/test/conf"
	"github.com/bad-superman/test/logging"
	"github.com/bad-superman/test/sdk/utils"
)

const (
	_okexRestApiHost = "https://aws.okx.com"
	_okexApiKey      = "5647dfa2-8bdd-46f4-a9dc-07c7bb27ea37"
	_okexSecretKey   = "B9BB7597D2DF626FEA7C3F48B3E5A85A"
	_okexPassphrase  = "Nuanguang@909"
)

type OkexClient struct {
	isSimulatedTrading bool // 模拟盘
	name               string
	apiKey             string
	secretKey          string
	passphrase         string
	client             http.Client
}

func NewOkexClient() *OkexClient {
	return &OkexClient{
		apiKey:     _okexApiKey,
		secretKey:  _okexSecretKey,
		passphrase: _okexPassphrase,
		client: http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func NewOkexClientByName(c *conf.Config, name string) *OkexClient {
	var okexConfig = &conf.OkexConfig{}
	for _, cnf := range c.OkexConfigs {
		if cnf.Name == name {
			okexConfig = cnf
			break
		}
	}
	return &OkexClient{
		name:       name,
		apiKey:     okexConfig.ApiKey,
		secretKey:  okexConfig.SecretKey,
		passphrase: okexConfig.Passphrase,
		client: http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// OK-ACCESS-KEY字符串类型的APIKey。
// OK-ACCESS-SIGN使用HMAC SHA256哈希函数获得哈希值，再使用Base-64编码（请参阅签名）。
// OK-ACCESS-TIMESTAMP发起请求的时间（UTC），如：2020-12-08T09:08:57.715Z
// OK-ACCESS-PASSPHRASE您在创建API密钥时指定的Passphrase。
func (o *OkexClient) post(url string, data interface{}, result interface{}) error {
	return o.request(http.MethodPost, url, data, result)
}

func (o *OkexClient) get(url string, result interface{}) error {
	return o.request(http.MethodGet, url, nil, result)
}
func (o *OkexClient) request(method string, url string, data interface{}, result interface{}) error {
	var (
		ts = utils.IsoTime()
	)
	bodyByte, err := json.Marshal(data)
	if err != nil {
		return err
	}
	body := bytes.NewReader(bodyByte)
	s := o.sign(ts, method, url, string(bodyByte))

	url = _okexRestApiHost + url
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		logging.Errorf("OkexClient|NewRequest error,err:%v", err)
	}
	req.Header.Add("OK-ACCESS-KEY", o.apiKey)
	req.Header.Add("OK-ACCESS-SIGN", s)
	req.Header.Add("OK-ACCESS-TIMESTAMP", ts)
	req.Header.Add("OK-ACCESS-PASSPHRASE", o.passphrase)
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	// 模拟盘
	if o.isSimulatedTrading {
		req.Header.Add("x-simulated-trading", "1")
	}

	resp, err := o.client.Do(req)
	if err != nil {
		return err
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		logging.Errorf("OkexClient|http code error,code:%d,resp:%s",
			resp.StatusCode, string(respBody))
		return fmt.Errorf("http code error,%d", resp.StatusCode)
	}
	err = json.Unmarshal(respBody, result)
	return err
}

// 请求签名
// timestamp + method + requestPath + body字符串（+表示字符串连接），以及SecretKey，使用HMAC SHA256方法加密，通过Base-64编码输出而得到的。
// 如：sign=CryptoJS.enc.Base64.stringify(CryptoJS.HmacSHA256(timestamp + 'GET' + '/api/v5/account/balance?ccy=BTC', SecretKey))
func (o *OkexClient) sign(ts, method, path, body string) string {
	// logging.Debugf("OkexClient|%s", ts+method+path+body)
	s, _ := utils.HmacSha256Base64Signer(ts+method+path+body, o.secretKey)
	return string(s)
}
