package main

import (
	"log"
	"strings"

	"github.com/JREAMLU/core/mq"
	"github.com/JREAMLU/jlog/server/service"
)

var addrs = `172.16.9.4:9092,172.16.9.4:9092`

func main() {
	// init Kafka
	err := mq.InitKafka(strings.Split(addrs, ","))
	if err != nil {
		log.Printf("Init Kafka Error:%s", err.Error())
		return
	}
	defer mq.Close()
	log.Printf("Kafka Connected")

	service.Server("udp4", "udp", "1200")
}
