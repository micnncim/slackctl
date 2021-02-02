package logging

import (
	"fmt"
	"io"

	"github.com/fatih/color"
)

const (
	symbolSuccess = "✓"
	symbolFailure = "✗"
)

type Logger struct {
	writer io.Writer
}

func NewLogger(w io.Writer) *Logger {
	return &Logger{writer: w}
}

func (l *Logger) printf(format string, args ...interface{}) {
	fmt.Fprintf(l.writer, format, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.Infof("%s\n", args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	format = fmt.Sprintf("%%s %s", format)

	a := []interface{}{
		color.New(color.FgGreen).Sprint(symbolSuccess),
	}
	a = append(a, args...)

	l.printf(format, a...)
}

func (l *Logger) Error(args ...interface{}) {
	l.Errorf("%s\n", args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	format = fmt.Sprintf("%%s %s", format)

	a := []interface{}{
		color.New(color.FgRed).Sprint(symbolFailure),
	}
	a = append(a, args...)

	l.printf(format, a...)
}
