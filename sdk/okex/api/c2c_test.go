package api

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestC2C(t *testing.T) {
	o := NewOkexClient()

	convey.Convey("获取otc深度", t, func() {
		_, err := o.C2COrderBooks("CNY", "USDT")
		convey.So(err, convey.ShouldBeNil)
	})
}
