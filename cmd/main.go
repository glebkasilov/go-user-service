package main

import (
	"log/slog"
	"os"
	"os/signal"
	"user/internal/application"
)

func main() {
	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	mainlogger := slog.New(h)

	app := application.New(mainlogger)

	app.Start()

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt)

	<-ch

	app.GracefulStop()
}
