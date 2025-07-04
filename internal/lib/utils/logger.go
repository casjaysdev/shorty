// File: internal/lib/utils/logger.go
// Purpose: Provides a centralized, leveled logger with formatting and file support.

package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	logFile     *os.File
	logToFile   = false
	quietOutput = false
)

func InitLogger(filename string, toFile bool, quiet bool) error {
	quietOutput = quiet
	if toFile {
		var err error
		logFile, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("cannot open log file: %w", err)
		}
		log.SetOutput(logFile)
		logToFile = true
	}
	return nil
}

func logPrefix(level string) string {
	ts := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("[%s] [%s]", ts, level)
}

func Info(msg string, args ...any) {
	write("INFO", msg, args...)
}

func Warn(msg string, args ...any) {
	write("WARN", msg, args...)
}

func Error(msg string, args ...any) {
	write("ERROR", msg, args...)
}

func Debug(msg string, args ...any) {
	write("DEBUG", msg, args...)
}

func Fatal(msg string, args ...any) {
	write("FATAL", msg, args...)
	os.Exit(1)
}

func write(level string, msg string, args ...any) {
	fullMsg := fmt.Sprintf("%s %s", logPrefix(level), fmt.Sprintf(msg, args...))
	if !quietOutput {
		fmt.Println(fullMsg)
	}
	if logToFile && logFile != nil {
		log.Println(fullMsg)
	}
}
