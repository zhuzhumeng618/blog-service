package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

// Level 自定义日志级别类型
type Level int8

// Fields 建立日志消息类型与日志信息映射关系
type Fields map[string]interface{}

// 自定义日志级别常量
const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

// String 获取日志级别对应的字符串日志等级描述
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	// 如果都不满足则返回空字符串
	return ""
}

// Logger 日志模型
type Logger struct {
	newLogger *log.Logger     // 标准库中的 Logger
	ctx       context.Context // 标准库中的 Context，上下文
	level     Level
	fields    Fields
	callers   []string
}

// NewLogger 创建日志实例
// io.Writer 日志信息写入的目的地
// prefix 日志信息的前缀
// flag 日志信息的标记 log.Lshortfile | log.Ldate | log.Ltime
func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	// 创建标准库 logger 日志库实例
	l := log.New(w, prefix, flag)
	return &Logger{
		newLogger: l,
	}
}

// clone 拷贝日志实例
func (l *Logger) clone() *Logger {
	// 将 Logger 实例的值赋值给 nl，不是地址！！！
	nl := *l
	return &nl
}

// WithLevel 带有日志级别的 Logger
func (l *Logger) WithLevel(level Level) *Logger {
	ll := l.clone()
	ll.level = level
	return ll
}

// WithFields 带有字段的 Logger
func (l *Logger) WithFields(fields Fields) *Logger {
	ll := l.clone()
	// 如果 fields 没有值，fields 是 map，需要初始化之后使用
	if ll.fields == nil {
		ll.fields = make(Fields)
	}
	// 将 fields 遍历存储到本 Logger 实例的 fields 中
	for k, v := range fields {
		ll.fields[k] = v
	}
	return ll
}

// WithContext 带有上下文的 Logger
func (l *Logger) WithContext(ctx context.Context) *Logger {
	ll := l.clone()
	ll.ctx = ctx
	return ll
}

// WithCaller 设置当前某一层调用栈的信息（程序计数器、文件信息、行号）
func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}
	return ll
}

// WithCallersFrames 设置当前整个调用栈的信息
func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	var callers []string
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		callers = append(callers, fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function))
		if !more {
			break
		}
	}
	ll := l.clone()
	ll.callers = callers
	return ll
}

// JSONFormat 日志内容的格式化
func (l *Logger) JSONFormat(message string) map[string]interface{} {
	// 最少预留 4 个格子，用于 level time message callers 信息的存储
	data := make(Fields, len(l.fields)+4)
	data["level"] = l.level.String()
	data["time"] = time.Now().Local().UnixNano() // 当前时间的纳秒级
	data["message"] = message
	data["callers"] = l.callers
	// 如果 l 实例中本来就含有 fields 信息，则将 l 中的 fields 不重复的 key-value 存储到 data 中
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}

	return data
}

// Output 日志输出动作
func (l *Logger) Output(message string) {
	body, _ := json.Marshal(l.JSONFormat(message))
	content := string(body)
	switch l.level {
	case LevelDebug:
		l.newLogger.Print(content)
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelWarn:
		l.newLogger.Print(content)
	case LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panic(content)
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.WithLevel(LevelDebug).Output(fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.WithLevel(LevelDebug).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Info(v ...interface{}) {
	l.WithLevel(LevelInfo).Output(fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.WithLevel(LevelInfo).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(v ...interface{}) {
	l.WithLevel(LevelWarn).Output(fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.WithLevel(LevelWarn).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.WithLevel(LevelError).Output(fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.WithLevel(LevelError).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(v ...interface{}) {
	l.WithLevel(LevelFatal).Output(fmt.Sprint(v...))
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.WithLevel(LevelFatal).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(v ...interface{}) {
	l.WithLevel(LevelPanic).Output(fmt.Sprint(v...))
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	l.WithLevel(LevelPanic).Output(fmt.Sprintf(format, v...))
}
