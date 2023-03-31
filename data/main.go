package main

import (
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/bad-superman/test/conf"
	"github.com/bad-superman/test/dao"
	"github.com/bad-superman/test/data/process"
	"github.com/bad-superman/test/logging"
	okex_sdk "github.com/bad-superman/test/sdk/okex"
	okex_ws_sdk "github.com/bad-superman/test/sdk/okex/ws"
)

func init() {
	// 初始化配置
	conf.Init("./config.toml")
}

var btc_usdt_ask float64
var btc_usdt_bid float64

var btc_usd_ask float64
var btc_usd_bid float64

func DepthCallback(d interface{}) error {
	logging.Debug("GetDepth Msg: %s", d)
	data, ok := d.(*okex_ws_sdk.WSDepthTableV5Response)
	if !ok {
		return nil
	}
	ask, _ := strconv.ParseFloat(data.Data[0].Asks[0][0], 64)
	bid, _ := strconv.ParseFloat(data.Data[0].Bids[0][0], 64)
	if data.Arg.InstId == "BTC-USDT" {
		btc_usdt_ask = ask
		btc_usdt_bid = bid
	}

	if data.Arg.InstId == "BTC-USD-230630" {
		btc_usd_ask = ask
		btc_usd_bid = bid
	}
	return nil
}

func InterestRateUpload() {
	influxDB := dao.NewInfluxDB()
	c := time.Tick(15 * time.Second)
	for {
		<-c
		if btc_usdt_ask == 0 || btc_usdt_bid == 0 {
			continue
		}
		if btc_usd_ask == 0 || btc_usd_bid == 0 {
			continue
		}
		// 正向基差：spot买 feature卖
		gap_forward := (btc_usd_bid - btc_usdt_ask) / btc_usdt_ask * 100
		// 反向基差：feature买 spot卖
		gap_reverse := (btc_usdt_bid - btc_usd_ask) / btc_usd_ask * 100
		logging.Debugf("btc_usdt_ask:%.2f btc_usdt_bid:%.2f\n", btc_usdt_ask, btc_usdt_bid)
		logging.Debugf("btc_usd_ask:%.2f btc_usd_bid:%.2f\n", btc_usd_ask, btc_usd_bid)
		logging.Debugf("gap_z:%.2f gap_f:%.2f\n", gap_forward, gap_reverse)
		fields := map[string]interface{}{
			"forward": gap_forward,
			"reverse": gap_reverse,
		}
		tags := map[string]string{
			"path": "btc_usdt_spot-btc_usd_quater",
		}
		influxDB.WritePoint("interest_rate", fields, tags, time.Now())
	}
}

func main() {
	// otc 数据
	process.NewDataCron().Run()
	go InterestRateUpload()

	agent := &okex_ws_sdk.OKWSAgent{}
	config := &okex_sdk.Config{
		WSEndpoint:    okex_ws_sdk.WS_API_HOST,
		TimeoutSecond: 10,
		IsPrint:       false,
	}

	// 设置base url
	// agent.

	// Step1: Start agent.
START:
	agent.Start(config)

	agent.WithCallback("books5", DepthCallback)

	// Step2: Subscribe channel
	// Step2.0: Subscribe public channel swap/depths successfully.
	args := okex_ws_sdk.DepthArg{
		OpArgBase: okex_ws_sdk.OpArgBase{Channel: "books5"},
		InstId:    "BTC-USDT",
	}

	args1 := okex_ws_sdk.DepthArg{
		OpArgBase: okex_ws_sdk.OpArgBase{Channel: "books5"},
		InstId:    "BTC-USD-230630",
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
