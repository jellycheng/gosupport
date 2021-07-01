package xlog

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type LogLevel uint32

// 7种日志级别：低->高
const (
	TraceLevel LogLevel = iota
	// 调试模式
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	// 满足该等级，打印日志则退出应用，不建议用
	FatalLevel
	// 满足该等级，会发起恐慌，不建议用
	PanicLevel
)

func (level LogLevel) ToString() (string, error) {
	switch level {
	case TraceLevel:
		return "trace", nil
	case DebugLevel:
		return "debug", nil
	case InfoLevel:
		return "info", nil
	case WarnLevel:
		return "warn", nil
	case ErrorLevel:
		return "error", nil
	case FatalLevel:
		return "fatal", nil
	case PanicLevel:
		return "panic", nil
	default:
		return "", fmt.Errorf("无效的level值 %d", level)
	}

}


// 自定义日志结构体，继承原生go日志结构体，在此基础上增加日志级别封装
type Logger struct {
	*log.Logger
	level LogLevel  //日志级别，一个级别对应一个方法，后期考虑可对应多个方法（如带f、ln等结尾的）
}

func (l *Logger) output(level LogLevel, v ...interface{}) {
	if l.level <= level { // 打印满足级别的日志
		l.Logger.Println(v...)
	}
}

// 是否满足日志等级
func (l Logger) IsLevelEnabled(level LogLevel) bool {
	if l.level <= level {
		return true
	}
	return false
}

func (l *Logger) Trace(v ...interface{}) {
	l.output(TraceLevel, v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.output(DebugLevel, v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.output(InfoLevel, v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.output(WarnLevel, v...)
}

func (l *Logger) Warning(v ...interface{})  {
	l.output(WarnLevel, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.output(ErrorLevel, v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.output(FatalLevel, v...)
	if l.IsLevelEnabled(FatalLevel) == true {
		os.Exit(1)
	}
}

func (l *Logger) Panic(v ...interface{}) {
	l.output(PanicLevel, v...)
	if l.IsLevelEnabled(PanicLevel) == true {
		s := fmt.Sprintln(v...)
		panic(s)
	}
}

func (l *Logger) GetGoLogger() *log.Logger {
	return l.Logger
}

func Debug(v ...interface{}) {
	debug.Debug(v...)
}

func Info(v ...interface{}) {
	info.Info(v...)
}

func Warn(v ...interface{}) {
	warn.Warn(v...)
}

func Error(v ...interface{}) {
	err01.Error(v...)
}

// 返回xlog.Logger结构体指针
func New(out io.Writer, prefix string, flag int, level LogLevel) *Logger {
	return &Logger{
		Logger: log.New(out, prefix, flag),
		level:  level,
	}
}

// 等级字符串 转 日志等级类型
func ParseLevel(levelStr string) (LogLevel, error) {
	switch strings.ToLower(levelStr) {
	case "panic":
		return PanicLevel, nil
	case "fatal":
		return FatalLevel, nil
	case "error":
		return ErrorLevel, nil
	case "warn", "warning":
		return WarnLevel, nil
	case "info":
		return InfoLevel, nil
	case "debug":
		return DebugLevel, nil
	case "trace":
		return TraceLevel, nil
	default:
		var l LogLevel
		return l, fmt.Errorf("无效的日志Level: %q", levelStr)
	}
}

// 日志级别前缀
const (
	InfoPrefix  = "[info] "
	DebugPrefix = "[debug] "
	WarnPrefix  = "[warn] "
	ErrorPrefix  = "[error] "
)

var (
	debug = New(os.Stdout, DebugPrefix, log.LstdFlags | log.Llongfile, DebugLevel)
	info = New(os.Stdout, InfoPrefix, log.LstdFlags | log.Llongfile, InfoLevel)
	warn = New(os.Stdout, WarnPrefix, log.LstdFlags | log.Llongfile, WarnLevel)
	err01 = New(os.Stdout, ErrorPrefix, log.LstdFlags | log.Llongfile, WarnLevel)
)

