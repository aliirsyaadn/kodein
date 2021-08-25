package log

import (
	"fmt"
	"os"
)

func setFormat(format string) string {
	return fmt.Sprint(timeline(), traceline(), tagline, format, newline)
}

func Print(format string) {
	print(format)
}

func Printf(tag, format string, args ...interface{}) {
	args = append([]interface{}{tag}, args...)
	printf(setFormat(format), args...)
}

func Info(format string) {
	info(format)
}

func InfoDetail(tag, format string, args ...interface{}) {
	args = append([]interface{}{tag}, args...)
	infof(setFormat(format), args...)
}

func Debug(format string) {
	debug(format)
}

func DebugDetail(tag, format string, args ...interface{}) {
	args = append([]interface{}{tag}, args...)
	debugf(setFormat(format), args...)
}

func Warn(format string) {
	warn(format)
}

func WarnDetail(tag string, format string, args ...interface{}) {
	args = append([]interface{}{tag}, args...)
	warnf(setFormat(format), args...)
}

func Error(format string) {
	error(format)
}

func ErrorDetail(tag string, format string, args ...interface{}) {
	args = append([]interface{}{tag}, args...)
	errorf(setFormat(format), args...)
}

func Fatal(format string) {
	fatal(format)
	os.Exit(1)
}

func FatalDetail(tag string, format string, args ...interface{}) {
	args = append([]interface{}{tag}, args...)
	fatalf(setFormat(format), args...)
	os.Exit(1)
}
