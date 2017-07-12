/******************************************************************************
 * Copyright (c) 2016 Baidu.com, Inc. All Rights Reserved
 * author: liangqixuan <liangqixuan@baidu.com>
 * date : 2016/7/25
 * ***************************************************************************/

package logger

import (
	"github.com/cihub/seelog"
)

var Logger seelog.LoggerInterface
var LockLogger seelog.LoggerInterface
var DBLogger seelog.LoggerInterface

func init() {
	Logger = seelog.Disabled
	LockLogger = seelog.Disabled
	DBLogger = seelog.Disabled

	loadAppConfig()
}

func loadAppConfig() {
	logger, err := seelog.LoggerFromConfigAsFile("./conf/log.xml")
	if err != nil {
		panic(err)
	}

	lockLogger, err := seelog.LoggerFromConfigAsFile("./conf/lockLog.xml")
	if err != nil {
		panic(err)
	}

	dbLogger, err := seelog.LoggerFromConfigAsFile("./conf/dbLog.xml")
	if err != nil {
		panic(err)
	}

	Logger = logger
	LockLogger = lockLogger
	DBLogger = dbLogger
}

// DisableLog disables all library log output
func DisableLog() {
	Logger = seelog.Disabled
}

// UseLogger uses a specified seelog.LoggerInterface to output library log.
// Use this func if you are using Seelog logging system in your app.
func UseLogger(newLogger seelog.LoggerInterface) {
	Logger = newLogger
}
