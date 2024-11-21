package application

import (
	"log/slog"
	"user/internal/server"
	"user/internal/service/users"
	"user/internal/storage"
)

type Application struct {
	server *server.Server
}

func New(logger *slog.Logger) *Application {
	storage := storage.New()
	userSevice := sevice.New(storage, logger)
	server := server.New(*userSevice)
	return &Application{
		server: server,
	}
}

func (a *Application) Start() {
	a.server.Start()
}

func (a *Application) GracefulStop() {
	a.server.GracefulStop()
}
