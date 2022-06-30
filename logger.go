package logger

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

type Logger struct {
	logPath       string
	logFilePrefix string
	logFileName   string
	log           *os.File
}

func NewDefaultLogger() *Logger {
	return NewLogger("logs", "log")
}

func NewLogger(filePath string, filePrefix string) *Logger {
	l := &Logger{
		logPath:       filePath,
		logFilePrefix: filePrefix,
	}

	var err error
	l.logFileName = fmt.Sprintf("%s_%s.txt", l.logFilePrefix, strings.Replace(time.Now().Format("20060102150405.000000"), ".", "_", -1))
	if _, statErr := os.Stat(l.logPath); os.IsNotExist(statErr) {
		if err = os.Mkdir(l.logPath, 0755); err != nil {
			panic(err)
		}
	}

	if l.log, err = os.OpenFile(l.GetLogFilePath(), os.O_RDWR|os.O_CREATE, 0600); err != nil {
		panic(err)
	}

	return l
}

func (l *Logger) Close() {
	if err := l.log.Close(); err != nil {
		panic(err)
	}
}

func (l *Logger) Logf(toConsole bool, format string, a ...any) {
	l.Log(fmt.Sprintf(format, a...), toConsole)
}

func (l *Logger) Log(message string, toConsole bool) {
	_, err := l.log.WriteString(message + "\n")
	if err != nil {
		fmt.Println("Log entry failed to write")
	}
	if toConsole {
		fmt.Println(message)
	}
}

func (l *Logger) GetLogPath() string {
	return l.logPath
}

func (l *Logger) GetLogFileName() string {
	return l.logFileName
}

func (l *Logger) GetLogFilePath() string {
	return path.Join(l.logPath, l.logFileName)
}
