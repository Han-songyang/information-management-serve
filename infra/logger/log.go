package logger

import (
	"bytes"
	"fmt"
	"github.com/han-songyang/information-management-serve/infra/conf"
	"github.com/sirupsen/logrus"
	"os"
)

type LogFormatter struct{}

var log = logrus.New()

// Format 格式化输出日志
func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog string

	newLog = fmt.Sprintf("[%s] [%s] %s\n", timestamp, entry.Level, entry.Message)

	b.WriteString(newLog)
	return b.Bytes(), nil
}

// InitLog 初始化日志
func InitLog() {
	log.SetFormatter(&LogFormatter{})

	path := conf.Config.Section("log").Key("path").String()

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("failed to log to file.")
	}

	log.SetLevel(logrus.InfoLevel)

	fmt.Println("log init success")
}

func InfoLog(module, index string, args ...interface{}) {
	log.WithFields(logrus.Fields{
		"module": module,
		"index":  index,
	}).Info(args)
}

func ErrorLog(module, index string, args ...interface{}) {
	log.WithFields(logrus.Fields{
		"module": module,
		"index":  index,
	}).Error(args)
}

func WarnLog(module, index string, args ...interface{}) {
	log.WithFields(logrus.Fields{
		"module": module,
		"index":  index,
	}).Warn(args)
}
