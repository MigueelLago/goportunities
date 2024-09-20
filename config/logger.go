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
		debug:   log.New(writer, "--DEBUG: ", logger.Flags()),
		info:    log.New(writer, "--INFO: ", logger.Flags()),
		warning: log.New(writer, "--WARNING: ", logger.Flags()),
		err:     log.New(writer, "--ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

// Criando Logs não formatados
func (l *Logger) Debug(values ...interface{}) {
	l.debug.Println(values...)
}

func (l *Logger) Info(values ...interface{}) {
	l.info.Println(values...)
}

func (l *Logger) Warn(values ...interface{}) {
	l.warning.Println(values...)
}

func (l *Logger) Error(values ...interface{}) {
	l.err.Println(values...)
}

// Criando Logs formatados
func (l *Logger) DebugF(format string, values ...interface{}) {
	l.debug.Printf(format, values...)
}

func (l *Logger) InfoF(format string, values ...interface{}) {
	l.info.Printf(format, values...)
}

func (l *Logger) WarnF(format string, values ...interface{}) {
	l.warning.Printf(format, values...)
}

func (l *Logger) ErrorF(format string, values ...interface{}) {
	l.err.Printf(format, values...)
}
