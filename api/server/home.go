package server


import (
	"net/http"

	"lmq/util"
	"lmq/api/router"
	"lmq/lmq"
	"encoding/json"
)

type homeRouter struct {
	routes  []router.Route
}

func NewHomeRouter() router.Router {
	r := &homeRouter{ }
	r.initRoutes()
	return r
}

// Routes returns the available routers to the checkpoint controller
func (r *homeRouter) Routes() []router.Route {
	return r.routes
}

func (r *homeRouter) initRoutes() {
	r.routes = []router.Route{
		router.NewGetRoute("/", Home),
	}
}

func Home(w http.ResponseWriter, req *http.Request) {
	moduleList := lmq.GetModuleList()
	jsonStr, ok := json.Marshal(moduleList)
	m := make(map[string]interface{})
	var errno int
	var retCode int
	if ok == nil {
		retCode = http.StatusOK
		errno = util.HTTP_SUCCESS
	}else{
		retCode = http.StatusInternalServerError
		errno = util.HTTP_FAILED
	}
	m["errno"] = errno
	m["errmsg"] = util.GetCodeString(errno)
	m["data"] = string(jsonStr)
	SendHttpResponse(w, m, retCode)
}

