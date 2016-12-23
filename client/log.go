package client

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/JREAMLU/core/com"
	"github.com/astaxie/beego/logs"
)

// LoggerClient logger client
type LoggerClient struct {
	Conn    *net.UDPConn
	Console bool
	Level   int
}

const (
	// UDP4 udp4
	UDP4 = "udp4"
)

const (
	// LevelEmergency emergency
	LevelEmergency = iota
	// LevelAlert alert
	LevelAlert
	// LevelCritical critical
	LevelCritical
	// LevelError error
	LevelError
	// LevelWarning warning
	LevelWarning
	// LevelNotice notice
	LevelNotice
	// LevelInformational information
	LevelInformational
	// LevelDebug debug2
	LevelDebug
)

// LoggerConn UDPConn
var LoggerConn *net.UDPConn

// InitLogger init logger client
// TODO 日志等级 格式化日至 日志内容 代码行数 服务开始时建立连接 整个服务结束 defer conn.Close()
func InitLogger(addr string) (LoggerClient, error) {
	var loggerClient LoggerClient
	udpAddr, err := net.ResolveUDPAddr(UDP4, com.StringJoin(":", addr))
	if err != nil {
		return loggerClient, fmt.Errorf("%v Fatal error %v", os.Stderr, err.Error())
	}
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return loggerClient, fmt.Errorf("%v Fatal error %v", os.Stderr, err.Error())
	}
	loggerClient.Conn = conn
	return loggerClient, nil
}

// Write udp to server
func Write(v interface{}) error {
	_, err := LoggerConn.Write([]byte(v.(string)))
	if err != nil {
		return fmt.Errorf("%v Fatal error %v", os.Stderr, err.Error())
	}
	return nil
}

// SetLevel sets the global log level used by the simple logger.
func SetLevel(l int) {
	logs.SetLevel(l)
}

// Critical logs a message at critical level.
func Critical(v ...interface{}) {
	logs.Critical(generateFmtStr(len(v)), v...)
}

func generateFmtStr(n int) string {
	return strings.Repeat("%v ", n)
}
