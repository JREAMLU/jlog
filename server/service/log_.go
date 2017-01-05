package service

//
// import (
// 	"fmt"
// 	"log"
// 	"net"
// 	"os"
// 	"strconv"
// 	"time"
//
// 	"github.com/JREAMLU/core/com"
// )
//
// var topic = `jlog`
//
// // Request UDP request
// type Request struct {
// 	isCancel bool
// 	reqSeq   int
// 	reqPkg   []byte
// 	rspChan  chan<- []byte
// }
//
// // Server log :port & goroutine
// func Server(resolveNet, listenNet, port string) {
// 	udpAddr, err := net.ResolveUDPAddr(resolveNet, com.StringJoin(":", port))
// 	if err != nil {
// 		log.Printf("%v Fatal error %v", os.Stderr, err.Error())
// 		os.Exit(1)
// 	}
//
// 	conn, err := net.ListenUDP(listenNet, udpAddr)
// 	if err != nil {
// 		log.Printf("%v Fatal error %v", os.Stderr, err.Error())
// 		os.Exit(1)
// 	}
//
// 	defer conn.Close()
//
// 	reqChan := make(chan *Request, 100000000)
// 	go connHandler(reqChan)
//
// 	var seq int
// 	for {
// 		buf := make([]byte, 1024)
// 		rlen, remote, err := conn.ReadFromUDP(buf)
// 		if err != nil {
// 			fmt.Println("conn.ReadFromUDP fail.", err)
// 			continue
// 		}
// 		seq++
// 		go processHandler(conn, remote, buf[:rlen], reqChan, seq)
// 	}
// }
//
// func processHandler(conn *net.UDPConn, remote *net.UDPAddr, msg []byte, reqChan chan<- *Request, seq int) {
// 	rspChan := make(chan []byte, 1)
// 	reqChan <- &Request{false, seq, []byte(strconv.Itoa(seq)), rspChan}
// 	select {
// 	case rsp := <-rspChan:
// 		fmt.Printf("recv rsp. rsp=%v \n", string(rsp))
// 	case <-time.After(2 * time.Second):
// 		fmt.Printf("wait for rsp timeout.")
// 		reqChan <- &Request{isCancel: true, reqSeq: seq}
// 		conn.WriteToUDP([]byte("wait for rsp timeout."), remote)
// 		return
// 	}
//
// 	conn.WriteToUDP([]byte("all process succ."), remote)
// }
//
// func connHandler(reqChan <-chan *Request) {
// 	addr, err := net.ResolveUDPAddr("udp4", ":6001")
// 	if err != nil {
// 		fmt.Println("net.ResolveUDPAddr fail.", err)
// 		os.Exit(1)
// 	}
//
// 	conn, err := net.DialUDP("udp", nil, addr)
// 	if err != nil {
// 		fmt.Println("net.DialUDP fail.", err)
// 		os.Exit(1)
// 	}
// 	defer conn.Close()
//
// 	sendChan := make(chan []byte, 1000)
// 	go sendHandler(conn, sendChan)
//
// 	recvChan := make(chan []byte, 1000)
// 	go recvHandler(conn, recvChan)
//
// 	reqMap := make(map[int]*Request)
// 	for {
// 		select {
// 		case req := <-reqChan:
// 			if req.isCancel {
// 				delete(reqMap, req.reqSeq)
// 				fmt.Printf("CancelRequest recv. reqSeq=%v \n", req.reqSeq)
// 				continue
// 			}
// 			reqMap[req.reqSeq] = req
// 			sendChan <- req.reqPkg
// 			fmt.Printf("NormalRequest recv. reqSeq=%d reqPkg=%s \n", req.reqSeq, string(req.reqPkg))
// 		case rsp := <-recvChan:
// 			seq, err := strconv.Atoi(string(rsp))
// 			if err != nil {
// 				fmt.Printf("strconv.Atoi fail. err=%v \n", err)
// 				continue
// 			}
// 			req, ok := reqMap[seq]
// 			if !ok {
// 				fmt.Printf("seq not found. seq=%v \n", seq)
// 				continue
// 			}
// 			req.rspChan <- rsp
// 			fmt.Printf("send rsp to client. rsp=%v \n", string(rsp))
// 			delete(reqMap, req.reqSeq)
// 		}
// 	}
// }
//
// func sendHandler(conn *net.UDPConn, sendChan <-chan []byte) {
// 	for data := range sendChan {
// 		wlen, err := conn.Write(data)
// 		if err != nil || wlen != len(data) {
// 			fmt.Println("conn.Write fail.", err)
// 			continue
// 		}
// 		fmt.Printf("conn.Write succ. data=%v \n", string(data))
// 	}
// }
//
// func recvHandler(conn *net.UDPConn, recvChan chan<- []byte) {
// 	for {
// 		buf := make([]byte, 1024)
// 		rlen, err := conn.Read(buf)
// 		if err != nil || rlen <= 0 {
// 			fmt.Println("recv", err)
// 			continue
// 		}
// 		fmt.Printf("conn.Read succ. data=%v \n", string(buf))
// 		recvChan <- buf[:rlen]
// 	}
// }
