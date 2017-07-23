package server

import (
	"net/http"

	"lmq/util"
	"lmq/api/router"
	"lmq/lmq"
)

type moduleRouter struct {
	routes  []router.Route
}

func NewModuleRouter() router.Router {
	r := &moduleRouter{ }
	r.initRoutes()
	return r
}

// Routes returns the available routers to the checkpoint controller
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
