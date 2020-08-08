package logger

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/terryli1643/apidemo/libs/configure"
)

var log *Logger

type Logger struct {
	*logrus.Logger
}

func InitLog() *Logger {
	if log == nil {
		config := configure.ServerConfig.LogRotate

		log = &Logger{}
		log.Logger = logrus.New()
		formatter := &logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05.000",
		}

		fullPath, _ := exec.LookPath(os.Args[0])
		fname := filepath.Base(fullPath)

		hook := NewRotateFileHook(configure.TRotateFileConfig{
			Filename:   "./log/" + fname + ".log",
			MaxSize:    config.MaxSize,
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge,
		}, formatter)

		log.AddHook(hook)
		log.Formatter = formatter
		log.SetLogLevel(config.Level)
		log.Info("logger init")
	}

	return log

}

type TLogConfig struct {
	// FileName   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Level      string
}

func (log *Logger) Output(calldepth int, s string) error {
	// line := log.getLineNumer(calldepth)
	// log.Logger.Debug(s, line)
	log.Debug(s)
	return nil
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

func (log *Logger) Debugf(format string, args ...interface{}) {
	log.logPrintf(log.Logger.Debugf, format, args...)
}

func (log *Logger) Infof(format string, args ...interface{}) {
	log.logPrintf(log.Logger.Infof, format, args...)
}

func (log *Logger) Warnf(format string, args ...interface{}) {
	log.logPrintf(log.Logger.Warnf, format, args...)
}

func (log *Logger) Warningf(format string, args ...interface{}) {
	log.logPrintf(log.Logger.Warningf, format, args...)
}

func (log *Logger) Errorf(format string, args ...interface{}) {
	log.logErrorPrintf(log.Logger.Errorf, format, args...)
}

func (log *Logger) Fatalf(format string, args ...interface{}) {
	log.logErrorPrintf(log.Logger.Fatalf, format, args...)
}

func (log *Logger) Panicf(format string, args ...interface{}) {
	log.logErrorPrintf(log.Logger.Panicf, format, args...)
}

func (log *Logger) Debug(args ...interface{}) {
	log.logPrint(log.Logger.Debug, args...)
}

func (log *Logger) Info(args ...interface{}) {
	log.logPrint(log.Logger.Info, args...)
}

func (log *Logger) Print(args ...interface{}) {
	log.logPrint(log.Logger.Print, args...)
}

func (log *Logger) Warn(args ...interface{}) {
	log.logPrint(log.Logger.Warn, args...)
}

func (log *Logger) Warning(args ...interface{}) {
	log.logPrint(log.Logger.Warning, args...)
}

func (log *Logger) Error(args ...interface{}) {
	log.logErrorPrint(log.Logger.Error, args...)
}

func (log *Logger) Fatal(args ...interface{}) {
	log.logErrorPrint(log.Logger.Fatal, args...)
}

func (log *Logger) Panic(args ...interface{}) {
	log.logErrorPrint(log.Logger.Panic, args...)
}

func (log *Logger) Debugln(args ...interface{}) {
	log.logPrint(log.Logger.Debugln, args...)
}

func (log *Logger) Infoln(args ...interface{}) {
	log.logPrint(log.Logger.Infoln, args...)
}

func (log *Logger) Println(args ...interface{}) {
	log.logPrint(log.Logger.Println, args...)
}

func (log *Logger) Warnln(args ...interface{}) {
	log.logPrint(log.Logger.Warnln, args...)
}

func (log *Logger) Warningln(args ...interface{}) {
	log.logPrint(log.Logger.Warningln, args...)
}

func (log *Logger) Errorln(args ...interface{}) {
	log.logErrorPrint(log.Logger.Errorln, args...)
}

func (log *Logger) Fatalln(args ...interface{}) {
	log.logErrorPrint(log.Logger.Fatalln, args...)
}

func (log *Logger) Panicln(args ...interface{}) {
	log.logErrorPrint(log.Logger.Panicln, args...)
}

// func (log *Logger) sqlLogPrint(fn FLogPrint, args ...interface{}) {
// 	lineNum := log.getLineNumer(3)
// 	var arr []interface{}
// 	arr = append(arr, args...)
// 	arr = append(arr, lineNum)

// 	fn(arr...)
// }

func (log *Logger) SqlDebug(args ...interface{}) {
	// log.sqlLogPrint(log.Logger.Debug, args...)
	argstr := fmt.Sprintf("%+v", args)

	ignorestr := `/*no print*/`
	if strings.Contains(argstr, ignorestr) {
		return
	}

	if strings.Contains(argstr, "Error") {
		log.Error(args...)
	} else {
		log.Debug(args...)
	}
}

func (log *Logger) WithField(key string, value interface{}) *logrus.Entry {

	lineNum := log.getLineNumer(2)

	fields := logrus.Fields{
		key:        value,
		"~LineNum": lineNum,
	}

	return log.Logger.WithFields(fields)
}

func (log *Logger) WithFields(fields logrus.Fields) *logrus.Entry {
	lineNum := log.getLineNumer(2)
	fields["~LineNum"] = lineNum

	return log.Logger.WithFields(fields)
}

func (log *Logger) SetLogLevel(level string) {
	switch level {
	case "INFO":
		log.SetLevel(logrus.InfoLevel)
	case "WARN":
		log.SetLevel(logrus.WarnLevel)
	case "ERROR":
		log.SetLevel(logrus.ErrorLevel)
	case "DEBUG":
		log.SetLevel(logrus.DebugLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}
}
