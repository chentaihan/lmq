package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"lmq/api/server"
	"lmq/api/router"
)

const versionMatcher = "/v{version:[0-9.]+}"


type Message struct{
	ID int64
	Platform string
	Module string
	Tag string
	Url string
	Params string
}

type Server struct {
	routers []router.Router
	Router  *mux.Router
}

func NewServer() *Server{
	return &Server{}
}

func (s *Server) InitRouter() {
	routers := []router.Router{
		server.NewMessageRouter(),
	}
	s.routers = append(s.routers, routers...)
	s.Router = s.createMux()
}

// createMux initializes the main router the server uses.
func (s *Server) createMux() *mux.Router {
	m := mux.NewRouter()
	for _, apiRouter := range s.routers {
		for _, r := range apiRouter.Routes() {
			m.Path(versionMatcher + r.Path()).Methods(r.Method()).Handler(r.Handler())
			m.Path(r.Path()).Methods(r.Method()).Handler(r.Handler())
		}
	}

	err := server.NewRequestNotFoundError(fmt.Errorf("page not found"))
	notFoundHandler := server.MakeErrorHandler(err)
	m.HandleFunc(versionMatcher+"/{path:.*}", notFoundHandler)
	m.NotFoundHandler = notFoundHandler

	return m
}