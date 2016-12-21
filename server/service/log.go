package service

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/JREAMLU/core/com"
)

// Server log :port & goroutine
func Server(port string) {
	udpAddr, err := net.ResolveUDPAddr("udp4", com.StringJoin(":", port))
	if err != nil {
		log.Printf("%v Fatal error %v", os.Stderr, err.Error())
		os.Exit(1)
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Printf("%v Fatal error %v", os.Stderr, err.Error())
		os.Exit(1)
	}
	for {
		handleClient(conn)
	}
}

// TODO thorw kakfa
func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	fmt.Println("content: ", string(buf[0:]))
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}
