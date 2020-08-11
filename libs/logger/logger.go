package logger

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/terryli1643/apidemo/libs/configure"
	"github.com/terryli1643/apidemo/libs/daemon"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var myLog *Logger

type Logger struct {
	*logrus.Logger
}

func New() *Logger {
	if myLog == nil {

		config := configure.New().LogRotate

		myLog = &Logger{}
		myLog.Logger = logrus.StandardLogger()

		fullPath, _ := exec.LookPath(os.Args[0])
		fname := "log"
		if strings.TrimSpace(fname) != "" {
			fname = filepath.Base(fullPath)
		}

		formatter := &logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05.000",
			// DisableColors:   true,
		}
		myLog.SetFormatter(formatter)
		myLog.SetLogLevel(config.Level)
		if daemon.IsDaemonMode() {
			writer := &lumberjack.Logger{
				Filename:   "./log/" + fname + ".log",
				MaxSize:    config.MaxSize,
				MaxBackups: config.MaxBackups,
				MaxAge:     config.MaxAge,
			}
			myLog.SetOutput(writer)
			log.SetOutput(myLog.Writer())
			gin.DefaultWriter = myLog.Writer()
		}
		myLog.Info("logger init")
	}

	return myLog

}

type TLogConfig struct {
	// FileName   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Level      string
}

func (log *Logger) getLineNumer(skip int) string {
	if pc, file, line, ok := runtime.Caller(skip); ok {
		funcName := runtime.FuncForPC(pc).Name()
		return fmt.Sprintf(" (%v:%v:%v)", path.Base(funcName), path.Base(file), line)
	}
	return " (no line number)"
}

type FLogPrintf func(format string, args ...interface{})
type FLogPrint func(args ...interface{})

func (log *Logger) logPrintf(fn FLogPrintf, format string, args ...interface{}) {
	lineNum := log.getLineNumer(3)
	var arr []interface{}
	arr = append(arr, args...)
	arr = append(arr, lineNum)

	fn(format+"%v", arr...)
}

func (log *Logger) logErrorPrintf(fn FLogPrintf, format string, args ...interface{}) {
	lineNum := log.getLineNumer(3)
	var arr []interface{}
	arr = append(arr, args...)
	arr = append(arr, lineNum)
	arr = append(arr, "\n"+string(debug.Stack()))
	fn(format+"%v"+"%v", arr...)
}

func (log *Logger) logPrint(fn FLogPrint, args ...interface{}) {
	lineNum := log.getLineNumer(3)
	var arr []interface{}
	arr = append(arr, args...)
	arr = append(arr, lineNum)

	fn(arr...)
}

func (log *Logger) logErrorPrint(fn FLogPrint, args ...interface{}) {
	lineNum := log.getLineNumer(3)
	var arr []interface{}
	arr = append(arr, args...)
	arr = append(arr, lineNum)
	arr = append(arr, "\n"+string(debug.Stack()))
	fn(arr...)
}

func Debugf(format string, args ...interface{}) {
	New().logPrintf(myLog.Logger.Debugf, format, args...)
}

func Infof(format string, args ...interface{}) {
	New().logPrintf(myLog.Logger.Infof, format, args...)
}

func Warnf(format string, args ...interface{}) {
	New().logPrintf(myLog.Logger.Warnf, format, args...)
}

func Warningf(format string, args ...interface{}) {
	New().logPrintf(myLog.Logger.Warningf, format, args...)
}

func Errorf(format string, args ...interface{}) {
	New().logErrorPrintf(myLog.Logger.Errorf, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	New().logErrorPrintf(myLog.Logger.Fatalf, format, args...)
}

func Panicf(format string, args ...interface{}) {
	New().logErrorPrintf(myLog.Logger.Panicf, format, args...)
}

func Debug(args ...interface{}) {
	New().logPrint(myLog.Logger.Debug, args...)
}

func Info(args ...interface{}) {
	New().logPrint(myLog.Logger.Info, args...)
}

func Print(args ...interface{}) {
	New().logPrint(myLog.Logger.Print, args...)
}

func Warn(args ...interface{}) {
	New().logPrint(myLog.Logger.Warn, args...)
}

func Warning(args ...interface{}) {
	New().logPrint(myLog.Logger.Warning, args...)
}

func Error(args ...interface{}) {
	New().logErrorPrint(myLog.Logger.Error, args...)
}

func Fatal(args ...interface{}) {
	New().logErrorPrint(myLog.Logger.Fatal, args...)
}

func Panic(args ...interface{}) {
	New().logErrorPrint(myLog.Logger.Panic, args...)
}

func Debugln(args ...interface{}) {
	New().logPrint(myLog.Logger.Debugln, args...)
}

func Infoln(args ...interface{}) {
	New().logPrint(myLog.Logger.Infoln, args...)
}

func Println(args ...interface{}) {
	New().logPrint(myLog.Logger.Println, args...)
}

func Warnln(args ...interface{}) {
	New().logPrint(myLog.Logger.Warnln, args...)
}

func Warningln(args ...interface{}) {
	New().logPrint(myLog.Logger.Warningln, args...)
}

func Errorln(args ...interface{}) {
	New().logErrorPrint(myLog.Logger.Errorln, args...)
}

func Fatalln(args ...interface{}) {
	New().logErrorPrint(myLog.Logger.Fatalln, args...)
}

func Panicln(args ...interface{}) {
	New().logErrorPrint(myLog.Logger.Panicln, args...)
}

func WithField(key string, value interface{}) *logrus.Entry {

	lineNum := New().getLineNumer(2)

	fields := logrus.Fields{
		key:        value,
		"~LineNum": lineNum,
	}

	return myLog.Logger.WithFields(fields)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	lineNum := New().getLineNumer(2)
	fields["~LineNum"] = lineNum

	return myLog.Logger.WithFields(fields)
}

func (log *Logger) SetLogLevel(level string) {
	switch level {
	case "INFO":
		log.Logger.SetLevel(logrus.InfoLevel)
	case "WARN":
		log.Logger.SetLevel(logrus.WarnLevel)
	case "ERROR":
		log.Logger.SetLevel(logrus.ErrorLevel)
	case "DEBUG":
		log.Logger.SetLevel(logrus.DebugLevel)
	default:
		log.Logger.SetLevel(logrus.InfoLevel)
	}
}
