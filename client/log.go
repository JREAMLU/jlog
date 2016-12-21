package client

import (
	"fmt"
	"log"
	"net"
	"os"
)

// Logger log it
// TODO 日志等级 格式化日至 日志内容 代码行数 服务开始时建立连接 整个服务结束 defer conn.Close()
func Logger(v interface{}) {
	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	if err != nil {
		log.Printf("%v Fatal error %v", os.Stderr, err.Error())
		os.Exit(1)
	}
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Printf("%v Fatal error %v", os.Stderr, err.Error())
		os.Exit(1)
	}
	_, err = conn.Write([]byte(v.(string)))
	if err != nil {
		log.Printf("%v Fatal error %v", os.Stderr, err.Error())
		os.Exit(1)
	}
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	if err != nil {
		log.Printf("%v Fatal error %v", os.Stderr, err.Error())
		os.Exit(1)
	}
	//日志格式
	fmt.Println(string(buf[0:n]))
	//连接
	os.Exit(0)
	// conn.Close()

}
