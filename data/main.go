package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/bad-superman/test/conf"
	"github.com/bad-superman/test/dao"
	"github.com/bad-superman/test/data/process"
	"github.com/bad-superman/test/logging"
	"github.com/bad-superman/test/sdk/okex"
	okex_sdk "github.com/bad-superman/test/sdk/okex"
	okex_api "github.com/bad-superman/test/sdk/okex/api"
	okex_ws_sdk "github.com/bad-superman/test/sdk/okex/ws"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/muesli/cache2go"
)

var (
	_tokens        = []string{"BTC", "ETH"}
	_okexClient    *okex_api.OkexClient
	_instrumentMap = make(map[string]okex_api.InstrumentData)
	_cache         = cache2go.Cache("data")
)

func init() {
	// 初始化配置
	conf.Init("./config.toml")
	c := conf.GetConfig()
	_okexClient = okex_api.NewOkexClientByName(c, "test")
}

func DepthCallback(d interface{}) error {
	logging.Debug("GetDepth Msg: %s", d)
	data, ok := d.(*okex_ws_sdk.WSDepthTableV5Response)
	if !ok {
		return nil
	}
	// store in memory cache,expire 5s
	_cache.Add(data.Arg.InstId, 5*time.Second, data)
	return nil
}

func InterestRateUpload(config *conf.Config) {
	influxDB := dao.NewInfluxDBV2(config)
	c := time.Tick(15 * time.Second)
	for {
		<-c
		points := make([]*write.Point, 0)
		for instId, instrument := range _instrumentMap {
			// get book data in cache
			item, err := _cache.Value(instId)
			if err != nil {
				logging.Errorf("InterestRateUpload|get book data from cache error,instId:%d,err:%v", instId, err)
				continue
			}
			v := item.Data()
			depthData := v.(*okex_ws_sdk.WSDepthTableV5Response)
			ask, _ := strconv.ParseFloat(depthData.Data[0].Asks[0][0], 64)
			bid, _ := strconv.ParseFloat(depthData.Data[0].Bids[0][0], 64)
			if ask <= 0 || bid <= 0 {
				logging.Errorf("InterestRateUpload|book data zero,instId:%s,ask:%.2f,bid:%.2f", instId, ask, bid)
				continue
			}
			fields := map[string]interface{}{
				"instId": instId,
				"ask":    ask,
				"bid":    bid,
			}
			tags := map[string]string{
				"instrument_type": string(instrument.InstType),
				"alias":           string(instrument.Alias),
				"uly":             instrument.Uly,
				"inst_family":     instrument.InstFamily,
			}
			logging.Infof("InterestRateUpload|Point info,fields:%+v tags:%+v", fields, tags)
			point := write.NewPoint("book_data", tags, fields, time.Now())
			points = append(points, point)
		}
		if len(points) == 0 {
			continue
		}
		err := influxDB.WritePoints(points)
		if err != nil {
			logging.Errorf("InterestRateUpload|WritePoints fail,err:%v", err)
		}
	}
}

// OKWSAgent SubscribeV5 every 5min
func CronSubscribe(agent *okex_ws_sdk.OKWSAgent, args []interface{}) {
	c := time.Tick(5 * time.Minute)
	for {
		<-c
		err := agent.SubscribeV5(args)
		if err != nil {
			logging.Errorf("main|CronSubscribe error,err:%v", err)
		}
	}
}

// prepare instrument map
func prepareInstrumentMap() {
	// prepare instruments
	for _, token := range _tokens {
		// spot instrument of USDT
		InstId := fmt.Sprintf("%s-USDT", token)
		USDInstId := fmt.Sprintf("%s-USD", token)
		// swap base USD
		USDSwapInstId := fmt.Sprintf("%s-USD-SWAP", token)
		_instrumentMap[USDSwapInstId] = okex_api.InstrumentData{
			InstID:     USDSwapInstId,
			InstType:   okex.SwapInstrument,
			Alias:      "swap",
			Uly:        USDInstId,
			InstFamily: USDInstId,
		}
		// swap base USDT
		USDTSwapInstId := fmt.Sprintf("%s-USDT-SWAP", token)
		_instrumentMap[USDTSwapInstId] = okex_api.InstrumentData{
			InstID:     USDTSwapInstId,
			InstType:   okex.SwapInstrument,
			Alias:      "swap",
			Uly:        InstId,
			InstFamily: InstId,
		}
		// spot
		_instrumentMap[InstId] = okex_api.InstrumentData{
			InstID:     InstId,
			InstType:   okex.SpotInstrument,
			Alias:      "spot",
			Uly:        InstId,
			InstFamily: InstId,
		}
		// get all future InstId base USD
		instruments, err := _okexClient.Instruments(okex.FuturesInstrument, fmt.Sprintf("%s-USD", token), "", "")
		if err != nil {
			logging.Errorf("get Instruments error,token:%s,err:%v", token, err)
			logging.Sync()
			os.Exit(0)
		}
		for _, instrument := range instruments {
			_instrumentMap[instrument.InstID] = instrument
		}
		// get all future InstId base USDT
		instruments, err = _okexClient.Instruments(okex.FuturesInstrument, fmt.Sprintf("%s-USDT", token), "", "")
		if err != nil {
			logging.Errorf("get Instruments error,token:%s,err:%v", token, err)
			logging.Sync()
			os.Exit(0)
		}
		for _, instrument := range instruments {
			_instrumentMap[instrument.InstID] = instrument
		}
	}
}

func main() {
	// otc 数据
	process.NewDataCron().Run()
	go InterestRateUpload(conf.GetConfig())

	prepareInstrumentMap()
	subArgs := make([]interface{}, 0)
	for _, instrument := range _instrumentMap {
		subArgs = append(subArgs, okex_ws_sdk.DepthArg{
			OpArgBase: okex_ws_sdk.OpArgBase{Channel: "books5"},
			InstId:    instrument.InstID,
		})
	}
	logging.Infof("main|subArgs:%+v", subArgs)

	agent := &okex_ws_sdk.OKWSAgent{}
	config := &okex_sdk.Config{
		WSEndpoint:    okex_ws_sdk.WS_API_HOST,
		TimeoutSecond: 10,
		IsPrint:       false,
	}
	// 定时订阅
	go CronSubscribe(agent, subArgs)
	// 设置base url
	// agent.
	// Step1: Start agent.
START:
	agent.Start(config)

	agent.WithCallback("books5", DepthCallback)

	// Step2: Subscribe channel
	// Step2.0: Subscribe public channel swap/depths successfully.
	agent.SubscribeV5(subArgs)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	for {
		select {
		case <-signalCh:
			os.Exit(0)
		case <-agent.IsStop():
			logging.Errorf("stoped, goto restart")
			goto START
		}
	}
}
