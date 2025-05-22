package process

import (
	"time"

	"github.com/bad-superman/test/conf"
	"github.com/bad-superman/test/dao"
	"github.com/bad-superman/test/logging"
	okex_api "github.com/bad-superman/test/sdk/okex/api"
	"github.com/bad-superman/test/sdk/thegraph"
	"github.com/bad-superman/test/sdk/utils"
	"github.com/robfig/cron"
)

type DataCron struct {
	cron           *cron.Cron
	dao            *dao.Dao
	influxDb       *dao.InfluxDBV2
	okexClient     *okex_api.OkexClient
	thegraphClient *thegraph.Client
}

func NewDataCron() *DataCron {
	c := conf.GetConfig()
	return &DataCron{
		cron:           cron.New(),
		influxDb:       dao.NewInfluxDBV2(c),
		okexClient:     okex_api.NewOkexClientByName(c, "test"),
		thegraphClient: thegraph.NewClient(c.Thegraph.ApiKey),
		dao:            dao.New(c),
	}
}

func (d *DataCron) Run() {
	d.cron.AddFunc("0 */10 * * * *", d.OkexOTCCron)
	d.cron.AddFunc("0 */10 * * * *", d.OkexExchangeRate)
	d.cron.AddFunc("0 0 * * * *", d.SyncAllTheGraphIndexer)
	d.cron.Start()
}

func (d *DataCron) OkexOTCCron() {
	books, err := d.okexClient.C2COrderBooks("CNY", "USDT")
	if err != nil {
		logging.Errorf("DataCron|C2COrderBooks error,err:%v", err)
		return
	}

	var (
		ask float64
		bid float64
	)

	if len(books.Buy) > 0 {
		ask = utils.StringToFloat64(books.Buy[0].Price)
	}
	if len(books.Sell) > 0 {
		bid = utils.StringToFloat64(books.Sell[len(books.Sell)-1].Price)
	}

	fields := map[string]interface{}{
		"ask": ask,
		"bid": bid,
	}

	tags := map[string]string{
		"coin_quote": "USDT_CNY",
	}

	err = d.influxDb.WritePoint("otc_price", fields, tags, time.Now())
	if err != nil {
		logging.Errorf("DataCron|C2COrderBooks WritePoint error,err:%v", err)
		return
	}
	logging.Infof("DataCron|C2COrderBooks WritePoint ok,ask:%f,bid:%f", ask, bid)
}

func (d *DataCron) OkexExchangeRate() {
	data, err := d.okexClient.ExchangeRate()
	if err != nil {
		logging.Errorf("DataCron|ExchangeRate error,err:%v", err)
		return
	}

	var (
		usdCnyPrice float64
	)

	if len(data) > 0 {
		usdCnyPrice = utils.StringToFloat64(data[0].UsdCny)
	}

	fields := map[string]interface{}{
		"price": usdCnyPrice,
	}

	tags := map[string]string{
		"coin_quote": "USD_CNY",
	}

	err = d.influxDb.WritePoint("exchange_rate", fields, tags, time.Now())
	if err != nil {
		logging.Errorf("DataCron|ExchangeRate WritePoint error,err:%v", err)
		return
	}
	logging.Infof("DataCron|ExchangeRate WritePoint ok,usdCnyPrice:%f", usdCnyPrice)
}
