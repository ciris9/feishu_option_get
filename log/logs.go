package log

import (
	"log"
	"option_get/config"
)

func Debugf(format string, v ...interface{}) {
	if config.IsDebug {
		log.Printf(format, v...)
	}
}

func Debug(msg ...any) {
	if config.IsDebug {
		log.Print(msg...)
	}
}
