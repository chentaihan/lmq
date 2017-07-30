package server

import (
	"net/http"

	"lmq/util"
	"lmq/api/router"
	"lmq/lmq"
	"fmt"
	"lmq/event"
)

type moduleRouter struct {
	routes  []router.Route
}

func NewModuleRouter() router.Router {
	r := &moduleRouter{ }
	r.initRoutes()
	return r
}

func (r *moduleRouter) Routes() []router.Route {
	return r.routes
}

func (r *moduleRouter) initRoutes() {
	r.routes = []router.Route{
		router.NewGetRoute("/addmodule", AddModule),
		router.NewGetRoute("/deletemodule", DeleteModule),
	}
}

func AddModule(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	platform := req.FormValue("platform")
	moduleName := req.FormValue("module")
	ok := lmq.AddModule(platform, moduleName)
	m := make(map[string]interface{})
	var errno int
	var retCode int
	if ok {
		if module := lmq.GetModule(platform, moduleName); module != nil {
			queueName := fmt.Sprintf("%s_%s", platform, moduleName)
			event.AddQueue(queueName, module.Queue)
		}
		retCode = http.StatusOK
		errno = util.HTTP_SUCCESS
	}else{
		retCode = http.StatusInternalServerError
		errno = util.HTTP_FAILED
	}
	m["errno"] = errno
	m["errmsg"] = util.GetCodeString(errno)
	SendHttpResponse(w, m, retCode)
}

func DeleteModule(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	platform := req.FormValue("platform")
	moduleName := req.FormValue("module")
	ok := lmq.DeleteModule(platform, moduleName)
	m := make(map[string]interface{})
	var errno int
	var retCode int
	if ok {
		queueName := fmt.Sprintf("%s_%s", platform, moduleName)
		event.SendSignal(queueName, event.EVENT_TYPE_DELETE_QUEUE)
		retCode = http.StatusOK
		errno = util.HTTP_SUCCESS
	}else{
		retCode = http.StatusInternalServerError
		errno = util.HTTP_FAILED
	}
	m["errno"] = errno
	m["errmsg"] = util.GetCodeString(errno)
	SendHttpResponse(w, m, retCode)
}
