package log

import (
	"blog/config"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	Logger  = logrus.New()
	logFile *os.File
	logName string
)

type Format struct{}

func (l *Format) Format(entry *logrus.Entry) ([]byte, error) {
	timeStamp := time.Now().Local().Format("2006-01-02 15:04:05")
	var fileName string
	var rowNumber int

	if entry.Caller != nil {
		fileName = filepath.Base(entry.Caller.File)
		rowNumber = entry.Caller.Line
	}

	msg := fmt.Sprintf("%s [%s:%d] [%s] %s\n", timeStamp, fileName, rowNumber, strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}

func StartLog() {
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetFormatter(new(Format))
	Logger.SetReportCaller(true)
	go logMonitor()
}

func logMonitor() {

	for {
		logDir := config.Settings.Log
		// MkdirAll -> mkdir -p xx
		Logger.Info(logDir)
		if err := os.MkdirAll(logDir, 0644); err != nil {
			if logFile != nil {
				_ = logFile.Close()
			}
			Logger.SetOutput(os.Stdout)
			Logger.WithField("Mkdir Failed", err.Error()).Error("Mkdir Failed")
			time.Sleep(time.Minute)
			continue
		}

		// format
		logID := time.Now().Local().Format("2006-01-02")
		newLogName := filepath.Join(config.Settings.Log, logID+".log")
		if logName == newLogName {
			return
		}
		Logger.Info(newLogName)
		newLogFile, err := os.OpenFile(newLogName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			if logFile != nil {
				_ = logFile.Close()
			}
			Logger.SetOutput(os.Stdout)
			Logger.Errorf("Failed to open log file for output: %s ", err.Error())
			time.Sleep(time.Minute)
			continue
		}
		Logger.Info("Open New File Success")
		Logger.SetOutput(io.MultiWriter(os.Stdout, newLogFile))

		// 关闭旧文件
		if logFile != nil {
			_ = logFile.Close()
		}
		logFile = newLogFile
		logName = newLogName

		time.Sleep(time.Minute)
	}

}
