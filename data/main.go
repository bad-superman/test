package main

import (
	"time"

	okex_sdk "github.com/bad-superman/test/sdk/okex"
)

type WsBookArgs struct {
	Channel string `json:"channel"`
	InstId  string `json:"instId"`
}

func main() {
	agent := okex_sdk.OKWSAgent{}
	config := &okex_sdk.Config{
		WSEndpoint:    okex_sdk.WS_API_HOST,
		TimeoutSecond: 10,
		IsPrint:       true,
	}

	// 设置base url
	// agent.

	// Step1: Start agent.
	agent.Start(config)

	// Step2: Subscribe channel
	// Step2.0: Subscribe public channel swap/depths successfully.
	args := WsBookArgs{
		Channel: "books5",
		InstId:  "BTC-USDT",
	}
	agent.SubscribeV5([]interface{}{args})

	// Step3: Client receive depths from websocket server.
	// Step3.0: Receive partial depths
	// Step3.1: Receive update depths (It may take a very long time to see Update Event.)

	time.Sleep(60 * time.Second)

	// Step4. Stop all the go routine run in background.
	agent.Stop()
	time.Sleep(1 * time.Second)
}
