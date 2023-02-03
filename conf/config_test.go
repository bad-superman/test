package conf

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestConfig(t *testing.T) {
	convey.Convey("初始化", t, func() {
		path := "./config.toml"
		c := Init(path)
		convey.So(c, convey.ShouldNotBeNil)
	})
}
