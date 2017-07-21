package server

import (
	"net/http"
	"lmq/util"
	"lmq/db"
	"lmq/api/router"
	"encoding/json"
	"fmt"
)

type messageRouter struct {
	routes  []router.Route
}

func NewMessageRouter() router.Router {
	r := &messageRouter{ }
	r.initRoutes()
	return r
}

// Routes returns the available routers to the checkpoint controller
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
	message := new(db.Message)
	message.Platform = req.FormValue("platform")
	message.Module = req.FormValue("module")
	message.Tag = req.FormValue("tag")
	message.Url = req.FormValue("url")
	message.Params = req.FormValue("params")
	msgStr,_ := json.Marshal(message)
	fmt.Println(string(msgStr))
	retCode := http.StatusOK
	m := make(map[string]interface{})
	if len(message.Platform) == 0 || len(message.Module) == 0 || len(message.Url) == 0{
		fmt.Println("404 bad request")
		retCode = http.StatusBadRequest
		m["errno"] = util.HTTP_PARAM_ERROR
		m["errmsg"] = util.GetCodeString(util.HTTP_PARAM_ERROR)
	}else{
		fmt.Println("start add message")
		msgId := db.SaveMessage(message)
		m["id"] = msgId
		if msgId > 0 {
			m["errno"] = util.HTTP_SUCCESS
		}else{
			m["errno"] = util.HTTP_FAILED
		}
	}
	SendHttpResponse(w, m, retCode)
}
