package logs

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

const levelLoggerImpl = -1

const defaultAsyncMsgLen = 1e3

// RFC5424 log message levels.
const (
	LevelEmergency = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInformational
	LevelDebug
)

const (
	// AdapterConsole console
	AdapterConsole = "console"
	// AdapterUDP udp
	AdapterUDP = "udp"
)

// JLogger looger
type JLogger struct {
	lock                sync.Mutex
	level               int
	init                bool
	enableFuncCallDepth bool
	loggerFuncCallDepth int
	asynchronous        bool
	msgChanLen          int64
	msgChan             chan *logMsg
	signalChan          chan string
	wg                  sync.WaitGroup
	outputs             []*nameLogger
}

type newLoggerFunc func() Logger

var jLogger = NewLogger()
var logMsgPool *sync.Pool
var adapters = make(map[string]newLoggerFunc)
var levelPrefix = [LevelDebug + 1]string{"[M] ", "[A] ", "[C] ", "[E] ", "[W] ", "[N] ", "[I] ", "[D] "}

type nameLogger struct {
	Logger
	name string
}

type logMsg struct {
	level int
	msg   string
	when  time.Time
}

// Logger defines the behavior of a log provider.
type Logger interface {
	Init(config string) error
	WriteMsg(when time.Time, msg string, level int) error
	Destroy()
	Flush()
}

// NewLogger new logger
func NewLogger(channelLens ...int64) *JLogger {
	jl := new(JLogger)
	jl.level = LevelDebug
	jl.loggerFuncCallDepth = 2
	jl.msgChanLen = append(channelLens, 0)[0]
	if jl.msgChanLen <= 0 {
		jl.msgChanLen = defaultAsyncMsgLen
	}
	jl.signalChan = make(chan string, 1)
	jl.setLogger(AdapterConsole)
	return jl
}

// Register makes a log provide available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, log newLoggerFunc) {
	if log == nil {
		panic("logs: Register provide is nil")
	}
	if _, dup := adapters[name]; dup {
		panic("logs: Register called twice for provider " + name)
	}
	adapters[name] = log
}

func (jl *JLogger) writeMsg(logLevel int, msg string, v ...interface{}) error {
	if !jl.init {
		jl.lock.Lock()
		jl.setLogger(AdapterConsole)
		jl.lock.Unlock()
	}

	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}
	when := time.Now()
	if jl.enableFuncCallDepth {
		_, file, line, ok := runtime.Caller(jl.loggerFuncCallDepth)
		if !ok {
			file = "???"
			line = 0
		}
		_, filename := path.Split(file)
		msg = "[" + filename + ":" + strconv.FormatInt(int64(line), 10) + "] " + msg
	}

	//set level info in front of filename info
	if logLevel == levelLoggerImpl {
		// set to emergency to ensure all log will be print out correctly
		logLevel = LevelEmergency
	} else {
		msg = levelPrefix[logLevel] + msg
	}

	if jl.asynchronous {
		lm := logMsgPool.Get().(*logMsg)
		lm.level = logLevel
		lm.msg = msg
		lm.when = when
		jl.msgChan <- lm
	} else {
		jl.writeToLoggers(when, msg, logLevel)
	}
	return nil
}

// SetLevel set level
func (jl *JLogger) SetLevel(l int) {
	jl.level = l
}

func (jl *JLogger) writeToLoggers(when time.Time, msg string, level int) {
	for _, l := range jl.outputs {
		err := l.WriteMsg(when, msg, level)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to WriteMsg to adapter:%v,error:%v\n", l.name, err)
		}
	}
}

// SetLogger sets a new logger.
func SetLogger(adapter string, config ...string) error {
	err := jLogger.SetLogger(adapter, config...)
	if err != nil {
		return err
	}
	return nil
}

// SetLogger provides a given logger adapter into BeeLogger with config string.
// config need to be correct JSON as string: {"interval":360}.
func (jl *JLogger) SetLogger(adapterName string, configs ...string) error {
	jl.lock.Lock()
	defer jl.lock.Unlock()
	if !jl.init {
		jl.outputs = []*nameLogger{}
		jl.init = true
	}
	return jl.setLogger(adapterName, configs...)
}

// SetLogger provides a given logger adapter into BeeLogger with config string.
// config need to be correct JSON as string: {"interval":360}.
func (jl *JLogger) setLogger(adapterName string, configs ...string) error {
	config := append(configs, "{}")[0]
	for _, l := range jl.outputs {
		if l.name == adapterName {
			return fmt.Errorf("logs: duplicate adaptername %q (you have set this logger before)", adapterName)
		}
	}

	log, ok := adapters[adapterName]
	if !ok {
		return fmt.Errorf("logs: unknown adaptername %q (forgotten Register?)", adapterName)
	}

	lg := log()
	err := lg.Init(config)
	if err != nil {
		fmt.Fprintln(os.Stderr, "logs.BeeLogger.SetLogger: "+err.Error())
		return err
	}
	jl.outputs = append(jl.outputs, &nameLogger{name: adapterName, Logger: lg})
	return nil
}

// SetLevel set level
func SetLevel(l int) {
	jLogger.SetLevel(l)
}

// DelLogger del logger console
func (jl *JLogger) DelLogger(adapterName string) error {
	jl.lock.Lock()
	defer jl.lock.Unlock()
	outputs := []*nameLogger{}
	for _, lg := range jl.outputs {
		if lg.name == adapterName {
			lg.Destroy()
		} else {
			outputs = append(outputs, lg)
		}
	}
	if len(outputs) == len(jl.outputs) {
		return fmt.Errorf("logs: unknown adaptername %q (forgotten Register?)", adapterName)
	}
	jl.outputs = outputs
	return nil
}

// DelLogger del logger console
func DelLogger(adapterName string) error {
	return jLogger.DelLogger(adapterName)
}

// Critical Log CRITICAL level message.
func (jl *JLogger) Critical(format string, v ...interface{}) {
	if LevelCritical > jl.level {
		return
	}
	jl.writeMsg(LevelCritical, format, v...)
}

// Critical logs a message at critical level.
func Critical(f interface{}, v ...interface{}) {
	jLogger.Critical(formatLog(f, v...))
}

func formatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			//format string
		} else {
			//do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}
