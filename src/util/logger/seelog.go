package logger

import (
	"github.com/cihub/seelog"
)

var Logger seelog.LoggerInterface
func init() {
	Logger = seelog.Disabled
	loadAppConfig()
}

func loadAppConfig() {
	logger, err := seelog.LoggerFromConfigAsFile("./conf/log.xml")
	if err != nil {
		panic(err)
	}

	Logger = logger
}

func DisableLog() {
	Logger = seelog.Disabled
}

func UseLogger(newLogger seelog.LoggerInterface) {
	Logger = newLogger
}
