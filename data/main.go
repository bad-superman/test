package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/bad-superman/test/logging"
	okex_sdk "github.com/bad-superman/test/sdk/okex"
)

var btc_usdt_ask float64
var btc_usdt_bid float64

var btc_usd_ask float64
var btc_usd_bid float64

func DepthCallback(d interface{}) error {
	logging.Infof("GetDepth Msg: %s", d)
	data, ok := d.(*okex_sdk.WSDepthTableV5Response)
	if !ok {
		return nil
	}
	ask, _ := strconv.ParseFloat(data.Data[0].Asks[0][0], 64)
	bid, _ := strconv.ParseFloat(data.Data[0].Bids[0][0], 64)
	if data.Arg.InstId == "BTC-USDT" {
		btc_usdt_ask = ask
		btc_usdt_bid = bid
	}

	if data.Arg.InstId == "BTC-USD-230331" {
		btc_usd_ask = ask
		btc_usd_bid = bid
	}
	if btc_usdt_ask == 0 || btc_usdt_bid == 0 {
		return nil
	}
	if btc_usd_ask == 0 || btc_usd_bid == 0 {
		return nil
	}
	// 正向基差：spot买 feature卖
	gap_z := (btc_usd_bid - btc_usdt_ask) / btc_usdt_ask * 100
	// 反向基差：feature买 spot卖
	gap_f := (btc_usdt_bid - btc_usd_ask) / btc_usd_ask * 100
	log.Printf("btc_usdt_ask:%.2f btc_usdt_bid:%.2f\n", btc_usdt_ask, btc_usdt_bid)
	log.Printf("btc_usd_ask:%.2f btc_usd_bid:%.2f\n", btc_usd_ask, btc_usd_bid)
	log.Printf("gap_z:%.2f gap_f:%.2f\n", gap_z, gap_f)
	return nil
}

func main() {
	agent := &okex_sdk.OKWSAgent{}
	config := &okex_sdk.Config{
		WSEndpoint:    okex_sdk.WS_API_HOST,
		TimeoutSecond: 10,
		IsPrint:       true,
	}

	// 设置base url
	// agent.

	// Step1: Start agent.
START:
	agent.Start(config)

	agent.WithCallback("books5", DepthCallback)

	// Step2: Subscribe channel
	// Step2.0: Subscribe public channel swap/depths successfully.
	args := okex_sdk.DepthArg{
		OpArgBase: okex_sdk.OpArgBase{Channel: "books5"},
		InstId:    "BTC-USDT",
	}

	args1 := okex_sdk.DepthArg{
		OpArgBase: okex_sdk.OpArgBase{Channel: "books5"},
		InstId:    "BTC-USD-230331",
	}
	agent.SubscribeV5([]interface{}{args, args1})

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	for {
		select {
		case <-signalCh:
			goto STOP
		case <-agent.IsStop():
			logging.Errorf("stoped, goto restart")
			goto START
		}
	}
STOP:
	agent.Stop()
}
