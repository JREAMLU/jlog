package client

import (
	"fmt"
	"log"
	"net"
	"os"
)

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
	fmt.Println(string(buf[0:n]))
	os.Exit(0)

}
