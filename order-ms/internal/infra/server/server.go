package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/paulozy/btg-challenge/order-ms/internal/infra/server/middlewares"
)

type Handler struct {
	Path   string
	Method string
	Auth   bool
	Func   gin.HandlerFunc
}

type Server struct {
	Host     string
	Port     string
	Env      string
	Router   *gin.Engine
	Handlers []Handler
}

func NewServer(host, port, env string) *Server {
	server := &Server{
		Host:   host,
		Port:   port,
		Env:    env,
		Router: gin.Default(),
	}

	return server
}

func (s *Server) AddHandlers() {
	s.Handlers = append(s.Handlers, Routes...)
}

func (s *Server) Start() {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}

	s.Router.Use(cors.New(config))
	s.Router.Use(middlewares.EnsureAllowedToRequestMiddleware())

	for _, h := range s.Handlers {
		s.Router.Handle(h.Method, h.Path, h.Func)
	}

	var address string

	if s.Host == "" && s.Env == "production" {
		address = "0.0.0.0:" + s.Port
	} else {
		address = s.Host + ":" + s.Port
	}

	s.Router.Run(address)
}
