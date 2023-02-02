package api

import (
	"testing"

	"github.com/bad-superman/test/sdk/okex"
	"github.com/smartystreets/goconvey/convey"
)

func TestTrade(t *testing.T) {
	o := NewOkexClient()

	var (
		instId  = "BTC-USD-230331"
		tdMode  = okex.TradeCrossMode
		side    = okex.OrderBuy
		posSide okex.PositionSide
		sz      okex.JSONFloat64 = 1
		px      okex.JSONFloat64 = 1000
		ordType                  = okex.OrderLimit
	)

	convey.Convey("合约限价单", t, func() {
		order, err := o.TradeFuturesOrder(instId, tdMode, side, posSide, sz, px, ordType)
		convey.So(err, convey.ShouldBeNil)
		convey.So(order, convey.ShouldNotBeNil)

		err = o.CancelOrder(instId, order.OrdID, "")
		convey.So(err, convey.ShouldBeNil)
	})

}
