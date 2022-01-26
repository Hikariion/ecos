package logger

import log "github.com/sirupsen/logrus"

type Logger struct {
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}
func Errorf(format string, args ...interface{}) {
	log.Errorf("ERROR! "+format, args...)
}
func Warningf(format string, args ...interface{}) {
	log.Warningf(format, args...)
}
