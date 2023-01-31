package process

import (
	"github.com/bad-superman/test/dao"
	"github.com/bad-superman/test/logging"
	okex_api "github.com/bad-superman/test/sdk/okex/api"
	"github.com/bad-superman/test/sdk/utils"
	"github.com/robfig/cron"
)

type DataCron struct {
	cron       *cron.Cron
	influxDb   *dao.InfluxDB
	okexClient *okex_api.OkexClient
}

func NewDataCron() *DataCron {
	return &DataCron{
		cron:       cron.New(),
		influxDb:   dao.NewInfluxDB(),
		okexClient: okex_api.NewOkexClient(),
	}
}

func (d *DataCron) Run() {
	d.cron.AddFunc("0 * * * * *", d.OkexOTCCron)
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

	err = d.influxDb.WritePoint("otc_price", fields, tags)
	if err != nil {
		logging.Errorf("DataCron|C2COrderBooks WritePoint error,err:%v", err)
		return
	}
	logging.Infof("DataCron|C2COrderBooks WritePoint ok,ask:%f,bid:%f", ask, bid)
}
