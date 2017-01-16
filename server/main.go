package main

import (
	"log"
	"strings"

	"github.com/JREAMLU/core/mq"
	"github.com/JREAMLU/jlog/server/service"
	"github.com/astaxie/beego"
)

var global service.Config

func init() {
	var err error
	global, err = service.GetConfig("conf/server.toml")
	if err != nil {
		beego.Error("toml parasr error: ", err)
	}
}

func main() {
	// init Kafka
	err := mq.InitKafka(strings.Split(global.Servers[global.Mode].Addr, ","))
	if err != nil {
		log.Printf("Init Kafka Error:%s", err.Error())
		return
	}
	defer mq.Close()
	beego.Info("Kafka Connected")

	service.Server(
		global.Servers[global.Mode].ResolveNet,
		global.Servers[global.Mode].ListenNet,
		global.Servers[global.Mode].Port,
		global.Servers[global.Mode].Topic,
	)
}
