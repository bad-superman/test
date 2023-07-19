package api

import (
	"testing"

	"github.com/bad-superman/test/sdk/okex"
	"github.com/smartystreets/goconvey/convey"
)

func TestInstruments(t *testing.T) {
	o := NewOkexClient()

	convey.Convey("获取交易产品基础信息", t, func() {
		data, err := o.Instruments(okex.FuturesInstrument, "BTC-USD", "", "")
		t.Log(data)
		convey.So(err, convey.ShouldBeNil)
	})
}
