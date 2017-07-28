package server

import (
	"net/http"
	"lmq/util"
	"lmq/db"
	"lmq/api/router"
	"encoding/json"
	"fmt"
	"lmq/lmq"
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
		router.NewGetRoute("/hello", SayHello),
	}
}

func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func AddMessage(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	message := new(lmq.Message)
	message.Platform = req.FormValue("platform")
	message.Module = req.FormValue("module")
	message.Tag = req.FormValue("tag")
	message.Url = req.FormValue("url")
	message.Params = req.FormValue("params")
	msgStr,_ := json.Marshal(message)
	fmt.Println(string(msgStr))
	retCode := http.StatusOK
	m := make(map[string]interface{})
	var errno int
	if len(message.Platform) == 0 || len(message.Module) == 0 || len(message.Url) == 0{
		fmt.Println("404 bad request")
		retCode = http.StatusBadRequest
		errno = util.HTTP_PARAM_ERROR
	}else{
		ok := lmq.ExistModule(message.Platform, message.Module)
		if ok {
			msgId := db.SaveMessage(message)
			m["id"] = msgId
			if msgId > 0 {
				errno = util.HTTP_SUCCESS
				lmq.AddQueue(message)
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
