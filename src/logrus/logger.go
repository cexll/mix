package logrus

import (
    "fmt"
    l "github.com/sirupsen/logrus"
    "path/filepath"
    "runtime"
)

type Logger struct {
    *l.Logger
}

func (t *Logger) ErrorStack(err interface{}, stack []byte) {
    if len(stack) > 0 {
        t.Logger.Errorf(fmt.Sprintf("%s\n%s", err, string(stack)))
    } else {
        t.Logger.Errorf(fmt.Sprintf("%s", err))
    }
}

func NewLogger() *Logger {
    logger := l.New()
    logger.ReportCaller = true // 显示调用信息

    formatter := new(l.TextFormatter)
    formatter.FullTimestamp = true
    formatter.TimestampFormat = "2006-01-02 15:04:05"
    formatter.DisableQuote = true // 不转义换行符，为了保存错误堆栈到日志文件
    formatter.CallerPrettyfier = func(frame *runtime.Frame) (function string, file string) {
        return "", fmt.Sprintf("%s:%d", filepath.Base(frame.File), frame.Line)
    }
    logger.Formatter = formatter

    return &Logger{logger}
}
