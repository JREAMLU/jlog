package client

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLogger(t *testing.T) {
	Convey("func Logger()", t, func() {
		// defer func() {
		// 	if x := recover(); x != nil {
		// 		buf := make([]byte, 1<<20)
		// 		runtime.Stack(buf, false)
		//
		// 		spew.Dump(x)
		// 		// Logger(x)
		// 	}
		// }()
		// var s []string
		// s = append(s, "a")
		// fmt.Println(s[10])
		// for i := 0; i < 10; i++ {
		// log.Println("i: ", i)
		// Logger("abc")
		// }
		// InitLogger("1200")
		// Write("cde")
		Critical("cde", "abc")
	})
}
