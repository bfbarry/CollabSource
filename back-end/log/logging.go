package log

import (
	native "log"
	"fmt"
	"os"
	"time"
)
// useage...
// logger = logging.NewLogEngine("logs.log", WARNING)
// logger.log(INFO, "action occurred")
type LogLevel int
const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	CRITICAL
	MAX_LOG_LEVEL
)
type LogEngine struct {
	uri string // e.g., file address
	protocol string // e.g., local or server TODO make an enum
	filterLevel LogLevel
}

var L *LogEngine

type Msg struct {
	level LogLevel
	dateTime string
	message string
}

var messageFormat string
// UTILS

func init() {
	messageFormat = "[%s][%s]: %s\n" //level, dateTime, message
}

func levelAsString(l LogLevel) string {
	switch l {
	case DEBUG:
		return "DEBUG" 
	case INFO:
		return "INFO" 
	case WARN:
		return "WARN" 
	case ERROR:
		return "ERROR" 
	case CRITICAL:
		return "CRITICAL"
	default:
		return fmt.Sprintf("INVALID_LOG_LEVEL_%d", l)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func getCurrentTime() string {
	formatStr := "2006-01-02 15:04:05" //don't change
	return time.Now().Format(formatStr)
}

func InitLogEngine(uri string, protocol string, filterLevel LogLevel) {
	L = &LogEngine{
		uri: uri,
		protocol: protocol,
		filterLevel: filterLevel,
	}
	switch protocol {//TODO
	case "file":
		;
	case "stdout":
		;
	}
}

func newLog(level LogLevel, message string) *Msg {
	return &Msg{
		level: level,
		dateTime: getCurrentTime(),
		message: message,
	}
}

func (this *LogEngine) Log(level LogLevel, message string) {
	// TODO log levels
	l := newLog(level, message)
	if level < this.filterLevel {
		return
	}
	fmtLog := fmt.Sprintf(messageFormat, levelAsString(l.level), l.dateTime, l.message)
	switch this.protocol {
	case "stdout":
		native.Println(fmtLog)
	case "file":
		f, err := os.OpenFile(this.uri, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		checkErr(err)
		defer f.Close()

		_, err = f.WriteString(fmtLog)
		checkErr(err)

	}

}