package service

import (
	"fmt"
	"net"
	"os"

	"github.com/JREAMLU/core/com"
	"github.com/JREAMLU/core/mq"
	"github.com/astaxie/beego"
)

// Server log :port & goroutine
func Server(resolveNet, listenNet, port, topic string) {
	udpAddr, err := net.ResolveUDPAddr(resolveNet, com.StringJoin(":", port))
	if err != nil {
		beego.Info(fmt.Sprintf("%v Fatal error %v", os.Stderr, err.Error()))
		os.Exit(1)
	}
	conn, err := net.ListenUDP(listenNet, udpAddr)
	if err != nil {
		beego.Error(fmt.Sprintf("%v Fatal error %v", os.Stderr, err.Error()))
		os.Exit(1)
	}
	defer conn.Close()
	for {
		handleClient(conn, topic)
	}
}

func handleClient(conn *net.UDPConn, topic string) {
	packet := make([]byte, 1024*10)
	n, _, err := conn.ReadFromUDP(packet)
	packet = packet[:n-1]
	if err != nil {
		beego.Error("failed to read UDP msg because of ", err.Error())
		return
	}
	go pushKafka(packet, topic)
}

func pushKafka(packet []byte, topic string) {
	err := mq.PushKafka(topic, string(packet))
	if err != nil {
		beego.Error("failed to push kafka: ", string(packet), err)
	}
}
