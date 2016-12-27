package logs

import (
	"fmt"
	"time"
)

type udpWrite struct {
	Name string
}

func newUDP() Logger {
	u := &udpWrite{Name: "Jream"}
	return u
}

// Init log
func (r *udpWrite) Init(jsonConfig string) error {
	return nil
}

// WriteMsg write message
// TODO write udp to server
func (r *udpWrite) WriteMsg(when time.Time, msg string, level int) error {
	fmt.Println("<<<<<<<<<<")
	return nil
}

func (r *udpWrite) Destroy() {

}

func (r *udpWrite) Flush() {

}

func init() {
	Register(AdapterUDP, newUDP)
}
