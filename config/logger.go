package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

func (l *Logger) Debug(message string, v ...interface{}) {
	l.debug.Printf(message, v...)
}

func (l *Logger) Info(message string, v ...interface{}) {
	l.info.Printf(message, v...)
}

func (l *Logger) Warning(message string, v ...interface{}) {
	l.warning.Printf(message, v...)
}

func (l *Logger) Error(message string, v ...interface{}) {
	l.err.Printf(message, v...)
}
