package service

// import (
// 	"log"
// 	"net"
// 	"os"
//
// 	"github.com/JREAMLU/core/com"
// 	"github.com/JREAMLU/core/mq"
// )
//
// var topic = `jlog`
//
// // Server log :port & goroutine
// func Server(resolveNet, listenNet, port string) {
// 	udpAddr, err := net.ResolveUDPAddr(resolveNet, com.StringJoin(":", port))
// 	if err != nil {
// 		log.Printf("%v Fatal error %v", os.Stderr, err.Error())
// 		os.Exit(1)
// 	}
// 	conn, err := net.ListenUDP(listenNet, udpAddr)
// 	if err != nil {
// 		log.Printf("%v Fatal error %v", os.Stderr, err.Error())
// 		os.Exit(1)
// 	}
// 	defer conn.Close()
// 	for {
// 		handleClient(conn)
// 	}
// }
//
// func handleClient(conn *net.UDPConn) {
// 	packet := make([]byte, 1024)
// 	n, addr, err := conn.ReadFromUDP(packet)
// 	packet = packet[:n-1]
// 	if err != nil {
// 		log.Println("failed to read UDP msg because of ", err.Error())
// 		return
// 	}
// 	go pushKafka(conn, addr, packet)
// }
//
// func pushKafka(conn *net.UDPConn, addr *net.UDPAddr, packet []byte) {
// 	err := mq.PushKafka(topic, string(packet))
// 	if err != nil {
// 		log.Println("log: ", string(packet), err)
// 	}
// 	// conn.WriteToUDP([]byte(time.Now().String()), addr)
// }
