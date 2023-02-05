package v2

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

func (m *Manager) SendMsg(ctx context.Context, msg interface{}) error {
	url := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", m.token)

	valueByte, _ := json.Marshal(msg)

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

	logging.Debugf("SendDingTalk, content=%s, accessToken=%s, resp=%v", valueByte, m.token, string(body))
	return nil
}
