package logging

import (
	"fmt"
	"io"

	"github.com/fatih/color"
)

const (
	defaultPrefixInfo  = "✓"
	defaultPrefixError = "✗"
)

type Logger struct {
	writer io.Writer

	prefixInfo  string
	prefixError string
}

func NewLogger(w io.Writer) *Logger {
	l := &Logger{
		writer:      w,
		prefixInfo:  defaultPrefixInfo,
		prefixError: defaultPrefixError,
	}

	return l
}

type Option func(*Logger)

func WithPrefixInfo(p string) Option {
	return func(l *Logger) {
		l.prefixInfo = p
	}
}

func WithPrefixError(p string) Option {
	return func(l *Logger) {
		l.prefixError = p
	}
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
		color.New(color.FgGreen).Sprint(l.prefixInfo),
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
		color.New(color.FgRed).Sprint(l.prefixError),
	}
	a = append(a, args...)

	l.printf(format, a...)
}
