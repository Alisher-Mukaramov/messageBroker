package server

import (
	"github.com/gin-gonic/gin"
	"messageBroker/internal/controller"
	"net/http"
	"os"
	"time"
)

const version = "v1"

type Server struct {
	controller *controller.Controller
	engine     *gin.Engine
}

func New(c *controller.Controller) *Server {
	return &Server{controller: c}
}

func (s *Server) Run(port string) {
	s.engine = gin.New()
	s.engine.Use(gin.Recovery())
	s.engine.Use(gin.Logger())

	host := os.Getenv("HOST")

	if host == "" {
		host = "127.0.0.1"
	}

	if port == "" {
		port = "8080"
	}

	ep := NewEndpoint(version, s)
	ep.Init()

	srv := &http.Server{
		Addr:           port,
		Handler:        s.engine,
		ReadTimeout:    300 * time.Second,
		WriteTimeout:   300 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	srv.ListenAndServe()
	s.engine.Run()
}
