package service

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLogs(t *testing.T) {
	Convey("func config()", t, func() {
		Convey("correct", func() {
			conf, err := GetConfig("../conf/server.toml")
			So(err, ShouldBeNil)
			So(conf.Mode, ShouldNotBeNil)
			So(conf.Servers[conf.Mode], ShouldNotBeNil)
		})
	})
}
