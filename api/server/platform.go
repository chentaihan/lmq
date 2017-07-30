package server

import (
	"net/http"

	"lmq/util"
	"lmq/api/router"
	"lmq/lmq"
	"lmq/util/logger"
)

type platformRouter struct {
	routes  []router.Route
}

func NewPlatformRouter() router.Router {
	r := &platformRouter{ }
	r.initRoutes()
	return r
}

// Routes returns the available routers to the checkpoint controller
func (r *platformRouter) Routes() []router.Route {
	return r.routes
}

func (r *platformRouter) initRoutes() {
	r.routes = []router.Route{
		router.NewGetRoute("/addplatform", AddPlatform),
		router.NewGetRoute("/deleteplatform", DeletePlatform),
	}
}

func AddPlatform(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	platform := req.FormValue("platform")
	isOK := lmq.AddPlatform(platform)
	logger.Logger.Tracef("AddPlatform ret=%d",isOK)
	m := make(map[string]interface{})
	var errno int
	var retCode int
	if isOK == true {
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

func DeletePlatform(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	platform := req.FormValue("platform")
	ok := lmq.DeletePlatform(platform)
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
