package server


import (
	"net/http"

	"../../util"
	"../router"
	"../../lmq"
)

type homeRouter struct {
	routes  []router.Route
}

func NewHomeRouter() router.Router {
	r := &homeRouter{ }
	r.initRoutes()
	return r
}

func (r *homeRouter) Routes() []router.Route {
	return r.routes
}

func (r *homeRouter) initRoutes() {
	r.routes = []router.Route{
		router.NewGetRoute("/", Home),
	}
}

func Home(w http.ResponseWriter, req *http.Request) {
	m := make(map[string]interface{})
	retCode := http.StatusOK
	errno := util.HTTP_SUCCESS
	m["errno"] = errno
	m["errmsg"] = util.GetCodeString(errno)
	m["data"] = lmq.GetModuleList()
	SendHttpResponse(w, m, retCode)
}

