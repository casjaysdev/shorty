// File: internal/lib/logger/logger.go
// Purpose: Provides a leveled and structured logger for the Shorty backend using the standard log package.

package logger

import (
	"log"
	"os"

	"shorty/internal/config"
)

type Logger struct {
	debug bool
	info  *log.Logger
	err   *log.Logger
}

func New(cfg *config.Config) *Logger {
	flags := log.LstdFlags | log.Lmsgprefix

	return &Logger{
		debug: cfg.Debug,
		info:  log.New(os.Stdout, "INFO  ", flags),
		err:   log.New(os.Stderr, "ERROR ", flags),
	}
}

func (l *Logger) Infof(format string, v ...any) {
	l.info.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...any) {
	l.err.Printf(format, v...)
}

func (l *Logger) Fatalf(format string, v ...any) {
	l.err.Fatalf(format, v...)
}

func (l *Logger) Debugf(format string, v ...any) {
	if l.debug {
		l.info.Printf("DEBUG  "+format, v...)
	}
}
