package logger

import (
	"fmt"
	"log"
)

const (
	infoColor    = "[INFO] \033[1;34m%s\033[0m\n"
	noticeColor  = "[NOTICE] \033[1;36m%s\033[0m\n"
	warningColor = "[WARNING] \033[1;33m%s\033[0m\n"
	errorColor   = "[ERROR] \033[1;31m%s\033[0m\n"
	debugColor   = "[DEBUG] \033[0;36m%s\033[0m\n"
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
	log.Printf(warningColor, value)
}

func Errorln(v ...any) {
	log.Printf(errorColor, v)
}

func Errorf(format string, v ...any) {
	value := fmt.Sprintf(format, v...)
	log.Printf(errorColor, value)
}
