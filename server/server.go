package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shubhamjagdhane/simple-load-balancer/entity"
	"github.com/shubhamjagdhane/simple-load-balancer/logger"
	"github.com/shubhamjagdhane/simple-load-balancer/server/middleware"
	"github.com/shubhamjagdhane/simple-load-balancer/tracer"
)

type Server struct {
	cfg                 *entity.Config
	logger              logger.ILogger
	serverPool          []*entity.ServerPool
	selectedServerIndex uint
}

func New(cfg *entity.Config, logger logger.ILogger) *Server {
	var SVR *Server

	if SVR != nil {
		return SVR
	}
	SVR = &Server{
		cfg:                 cfg,
		logger:              logger,
		serverPool:          make([]*entity.ServerPool, 0),
		selectedServerIndex: 0,
	}

	return SVR
}

func (s Server) Start() {
	tracer.New(s.cfg.Tracer.Enable, s.cfg.Tracer.ProjectID, s.cfg.Tracer.TracerName, s.logger)
	defer tracer.Shutdown()

	addr := fmt.Sprintf("%v:%v", s.cfg.HTTPAddress, s.cfg.Port)

	r := mux.NewRouter()
	r.Use(middleware.ValidateHeaderMiddleware)

	r.Handle("/urls/register", middleware.ErrHandler(s.RegisterUrlHandler)).Methods(http.MethodPost)
	r.Handle("/proxy", middleware.ErrHandler(s.ProxyHandler))
	r.Handle("/servers", middleware.ErrHandler(s.GetServers)).Methods(http.MethodGet)

	err := http.ListenAndServe(addr, r)
	if err != nil {
		s.logger.Fatalf("error listening to address %v, err=%v", addr, err)
	}
	s.logger.Debugf("HTTP server start %v", addr)
}
