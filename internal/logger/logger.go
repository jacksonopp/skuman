package logger

import (
	"fmt"
	"log"
)

const (
	infoColor    = "\033[1;34m%s\033[0m\n"
	noticeColor  = "\033[1;36m%s\033[0m\n"
	warningColor = "\033[1;33m%s\033[0m\n"
	errorColor   = "\033[1;31m%s\033[0m\n"
	debugColor   = "\033[0;36m%s\033[0m\n"
)

func Infoln(v ...any) {
	log.Printf(infoColor, v)
}

func Infof(format string, v ...any) {
	value := fmt.Sprintf(format, v...)
	log.Printf(infoColor, value)
}

func Warningln(v ...any) {
	log.Printf(warningColor, v)
}

func Warningf(format string, v ...any) {
	value := fmt.Sprintf(format, v...)
	log.Printf(infoColor, value)
}

func Errorln(v ...any) {
	log.Printf(errorColor, v)
}

func Errorf(format string, v ...any) {
	value := fmt.Sprintf(format, v...)
	log.Printf(errorColor, value)
}
