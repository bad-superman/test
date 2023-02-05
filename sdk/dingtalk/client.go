package dingtalk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/bad-superman/test/logging"
)

type Manager struct {
	token string
}

func New(token string) *Manager {
	return &Manager{
		token: token,
	}
}

func (m *Manager) SendMarkDownMsg(ctx context.Context, title, content string) error {
	url := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", m.token)

	data := map[string]interface{}{
		"markdown": map[string]interface{}{
			"title": title,
			"text":  content,
		},
		"at": map[string]interface{}{
			"isAtAll": false,
		},
		"msgtype": "markdown",
	}

	valueByte, _ := json.Marshal(data)

	httpClient := &http.Client{
		Timeout: 3000 * time.Millisecond,
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(valueByte))
	if err != nil {
		logging.Errorf("failed to SendDingTalk, url=%s, value=%s, err=%v", url, string(valueByte), err)
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := httpClient.Do(req)
	if err != nil {
		logging.Errorf("failed to SendDingTalk, url=%s, value=%s, err=%v", url, string(valueByte), err)
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logging.Errorf("failed to SendDingTalk, url=%s, value=%s, err=%v", url, string(valueByte), err)
		return err
	}
	defer func() { _ = res.Body.Close() }()

	logging.Debugf("SendDingTalk, content=%s, accessToken=%s, resp=%v", content, m.token, string(body))
	return nil
}
