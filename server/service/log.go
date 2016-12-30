package service

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/JREAMLU/core/com"
	"github.com/JREAMLU/core/mq"
)

var topic = `jlog`

// Server log :port & goroutine
func Server(resolveNet, listenNet, port string) {
	udpAddr, err := net.ResolveUDPAddr(resolveNet, com.StringJoin(":", port))
	if err != nil {
		log.Printf("%v Fatal error %v", os.Stderr, err.Error())
		os.Exit(1)
	}
	conn, err := net.ListenUDP(listenNet, udpAddr)
	if err != nil {
		log.Printf("%v Fatal error %v", os.Stderr, err.Error())
		os.Exit(1)
	}
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	data := make([]byte, 1024)
	_, addr, err := conn.ReadFromUDP(data)
	if err != nil {
		log.Println("failed to read UDP msg because of ", err.Error())
		return
	}
	err = mq.PushKafka(topic, string(data))
	fmt.Println("log: ", string(data), err)
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}
