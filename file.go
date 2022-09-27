package fileLogger

import (
	"fmt"
	"github.com/tarasov-krk/logger"
	"log"
	"os"
	"time"
)

var goLog *log.Logger
var logFile *os.File

type fileLogger struct {
	logger *log.Logger
	level  logger.Level
}

func New(filePath string) *fileLogger {
	goLog = log.New(os.Stderr, "", log.LstdFlags)

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("error opening file: %v" + err.Error())
	}
	logFile = f

	goLog.SetOutput(f)

	return &fileLogger{
		logger: goLog,
		level:  logger.DefaultLogLevel,
	}
}

func (l *fileLogger) SetLevel(level logger.Level) {
	l.level = level
}

func (l *fileLogger) GetLevel() logger.Level {
	return l.level
}

func (l *fileLogger) Write(level string, msg ...any) {
	l.logger.Println(level, fmt.Sprint(msg...))
}

func GetLogFile() *os.File {
	return logFile
}

func GetDefaultLogFileName() string {
	now := time.Now()
	logFileName := fmt.Sprintf("%d.%d.%d", now.Day(), now.Month(), now.Year())
	return "./logs/" + logFileName + ".log"
}
