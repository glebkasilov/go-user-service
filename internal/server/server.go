package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"user/internal/config"
	userRouter "user/internal/server/users"
	userService "user/internal/service/users"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	server *http.Server
}

func New(userService userService.UserService) *Server {
	cfg := config.Config().Server

	router := gin.Default()

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: router,
	}

	group := router.Group("/api")
	userRouter.Register(&userService, group) // dangerous place

	return &Server{
		router: router,
		server: httpServer,
	}
}

func (s *Server) Start() {
	if err := s.server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func (s *Server) GracefulStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		panic(err)
	}
}
