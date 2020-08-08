package logger

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/terryli1643/apidemo/libs/configure"
	"github.com/terryli1643/apidemo/libs/daemon"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type RotateFileHook struct {
	Config    configure.TRotateFileConfig
	Formatter logrus.Formatter
	logWriter io.Writer
}

func NewRotateFileHook(config configure.TRotateFileConfig, formatter logrus.Formatter) logrus.Hook {

	hook := RotateFileHook{
		Config:    config,
		Formatter: formatter,
	}
	hook.logWriter = &lumberjack.Logger{
		Filename:   config.Filename,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
	}

	return &hook
}

func (hook *RotateFileHook) Levels() []logrus.Level {

	var levels []logrus.Level
	levels = append(levels, logrus.AllLevels...)

	return levels
}

func (hook *RotateFileHook) Fire(entry *logrus.Entry) (err error) {
	b, err := hook.Formatter.Format(entry)
	if err != nil {
		return err
	}
	hook.logWriter.Write(b)
	if daemon.IsDaemonMode() {
		entry.Logger.Out = ioutil.Discard
	} else {
		entry.Logger.Out = os.Stdout
	}

	return nil
}
