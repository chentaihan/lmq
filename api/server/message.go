package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"lmq/util"
	"lmq/db"
	"lmq/api/router"
	"lmq/lmq"
	"lmq/event"
	"lmq/util/logger"
)

type messageRouter struct {
	routes  []router.Route
}

func NewMessageRouter() router.Router {
	r := &messageRouter{ }
	r.initRoutes()
	return r
}

func (r *messageRouter) Routes() []router.Route {
	return r.routes
}

func (r *messageRouter) initRoutes() {
	r.routes = []router.Route{
		router.NewGetRoute("/addmessage", AddMessage),
	}
}

func AddMessage(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	message := lmq.NewMessage()
	message.Platform = req.FormValue("platform")
	message.Module = req.FormValue("module")
	message.Tag = req.FormValue("tag")
	message.Url = req.FormValue("url")
	message.Params = req.FormValue("params")
	msgStr,_ := json.Marshal(message)
	logger.Logger.Tracef("AddMessage content=%s", string(msgStr))
	retCode := http.StatusOK
	m := make(map[string]interface{})
	var errno int
	if len(message.Platform) == 0 || len(message.Module) == 0 || len(message.Url) == 0{
		retCode = http.StatusBadRequest
		errno = util.HTTP_PARAM_ERROR
	}else{
		ok := lmq.ExistModule(message.Platform, message.Module)
		if ok {
			msgId := db.SaveMessage(message)
			m["id"] = msgId
			if msgId > 0 {
				logger.Logger.Tracef("AddMessage SaveMessage success msgId=%d", msgId)
				errno = util.HTTP_SUCCESS
				if ok = lmq.AddMessageQueue(message); ok{
					queueName := fmt.Sprintf("%s_%s", message.Platform, message.Module)
					event.SendSignal(queueName, event.EVENT_TYPE_ADD_MESSAGE)
				}else{
					logger.Logger.Errorf("AddMessage AddQueue failed msgId=%d", msgId)
				}
			}else{
				retCode = http.StatusInternalServerError
				errno = util.HTTP_SAVEMESSAGE_FAILED
			}
		}else{
			errno = util.HTTP_PARAM_MODULE_NOT_EXIST
		}
	}
	m["errmsg"] = util.GetCodeString(errno)
	SendHttpResponse(w, m, retCode)
}
