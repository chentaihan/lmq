package main

import (
	"fmt"

	"lmq/api"
	"net/http"
	"lmq/lmq"
	"lmq/db"
	"lmq/util/logger"
	"lmq/event"
	"lmq/container"
	"lmq/util"
)

func main() {
	util.LoadConfig()
	db.InitDB();
	lmq.InitModule()
	logger.Logger.Trace("init success ...")
	startWorker()
	server := api.NewServer()
	server.InitRouter()
	httpPort := util.LmqConfig.Serve.HttpPort
	if httpPort <= 1024 {
		httpPort = 9527
	}
	http.ListenAndServe(fmt.Sprintf(":%d", httpPort), server.Router)
	logger.Logger.Trace("server start OK")
}

func startWorker(){
	moduleList := lmq.GetModuleList()
	moduleNameList := make([]string, len(moduleList))
	esQueueList := make([](*container.CQueue), len(moduleList))
	for _, value := range moduleList {
		for _, module := range value {
			name := fmt.Sprintf("%s_%s", module.Platform, module.Name)
			moduleNameList = append(moduleNameList, name)
			esQueueList = append(esQueueList, module.Queue)
		}
	}
	event.InitEventCenter(moduleNameList, esQueueList)
	event.StartEventCenter()
	logger.Logger.Trace("startWorker OK")
}