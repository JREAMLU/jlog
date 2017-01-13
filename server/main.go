package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/JREAMLU/core/mq"
	"github.com/JREAMLU/jlog/server/service"
	"github.com/astaxie/beego"
)

var addrs = `172.16.9.4:9092,172.16.9.4:9092`
var global service.Config

func init() {
	var err error
	global, err = service.GetConfig("conf/server.toml")
	if err != nil {
		beego.Error("toml parasr error: ", err)
	}
	fmt.Println(global.Mode)
	fmt.Println(global)
}

func main() {
	// init Kafka
	err := mq.InitKafka(strings.Split(addrs, ","))
	if err != nil {
		log.Printf("Init Kafka Error:%s", err.Error())
		return
	}
	defer mq.Close()
	beego.Info("Kafka Connected")

	// service.Server("udp4", "udp", "1200")
	service.Server("udp4", "udp", "1200")
}
