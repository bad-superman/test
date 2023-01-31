package api

import (
	"testing"

	"github.com/bad-superman/test/logging"
	"github.com/smartystreets/goconvey/convey"
)

func TestAccountBalance(t *testing.T) {
	o := NewOkexClient()

	convey.Convey("无参数", t, func() {
		balances, err := o.AccountBalance(nil)
		t.Log(balances)
		convey.So(len(balances), convey.ShouldBeGreaterThan, 0)
		convey.So(err, convey.ShouldBeNil)
	})

	convey.Convey("查询单币种", t, func() {
		balances, err := o.AccountBalance([]string{"BTC"})
		t.Log(balances)
		convey.So(len(balances), convey.ShouldEqual, 1)
		convey.So(err, convey.ShouldBeNil)
	})
}

func TestAccountPositions(t *testing.T) {
	o := NewOkexClient()

	convey.Convey("无参数", t, func() {
		positions, err := o.AccountPositions("", nil, nil)
		logging.Debug(positions)
		convey.So(len(positions), convey.ShouldBeGreaterThan, 0)
		convey.So(err, convey.ShouldBeNil)
	})
}
