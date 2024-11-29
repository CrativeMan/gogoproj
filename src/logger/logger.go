package logger

import (
	"fmt"
	"os"
	"time"
)

type Logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type logger struct{}

func NewLogger() Logger {
	return &logger{}
}

func (l *logger) Info(msg string) {
	time := time.Now().Format("15:04:05")
	fmt.Fprintf(os.Stdout, "[%s] INFO: %s\n", time, msg)
}

func (l *logger) Warn(msg string) {
	time := time.Now().Format("15:04:05")
	fmt.Fprintf(os.Stdout, "[%s] WARN: %s\n", time, msg)
}

func (l *logger) Error(msg string) {
	time := time.Now().Format("15:04:05")
	fmt.Fprintf(os.Stderr, "[%s] ERROR: %s\n", time, msg)
}
