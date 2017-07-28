package main

import (
	"lmq/api"
	"net/http"
	"lmq/lmq"
	"lmq/db"
	"lmq/util/logger"
	"lmq/event"
)

func main() {
	event.TestEvent()
	lmq.InitModule()
	db.InitDB();
	logger.Logger.Trace("server start...")
	server := api.NewServer()
	server.InitRouter()
	http.ListenAndServe(":8001", server.Router)
	logger.Logger.Trace("server start OK")
}
