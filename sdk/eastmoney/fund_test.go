package eastmoney

import (
	"testing"
	"time"

	"github.com/bad-superman/test/dao"
	"github.com/bad-superman/test/sdk/utils"
	"github.com/smartystreets/goconvey/convey"
)

func TestLSJZ(t *testing.T) {
	e := NewEastmoneyClient()

	var (
		fundCode = "000828"
		page     = 1
		size     = 30
	)
	influxdb := dao.NewInfluxDB()
	convey.Convey("历史净值", t, func() {
		data, err := e.FundHistoricalNetValue(fundCode, page, size)
		for _, l := range data.Data.LSJZList {
			t, _ := time.ParseInLocation("2006-01-02", l.FSRQ, time.Local)
			lsjz := utils.StringToFloat64(l.LJJZ)
			swjz := utils.StringToFloat64(l.DWJZ)
			fields := map[string]interface{}{
				"fund_code": fundCode,
				"lsjz":      lsjz,
				"swjz":      swjz,
			}
			tags := map[string]string{
				"fund_code": fundCode,
			}
			influxdb.WritePoint("fund_net_value", fields, tags, t)
		}
		convey.So(err, convey.ShouldBeNil)
		convey.So(data, convey.ShouldNotBeNil)
	})
}
