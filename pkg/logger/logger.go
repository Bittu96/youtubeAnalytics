package logger

import log "github.com/sirupsen/logrus"

func Warn(msg string) {
	log.Warn(msg)
}

func Error(msg string) {
	log.Error(msg)
}
