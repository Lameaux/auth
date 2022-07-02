package main

import (
	"context"
	"github.com/Lameaux/auth/internal/config"
	"github.com/Lameaux/auth/internal/routes"
	"github.com/Lameaux/core/httpserver"
	"github.com/Lameaux/core/logger"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app := config.NewApp()
	logger.Infow("Starting", "app", config.AppName, "version", config.AppVersion)

	srv := httpserver.Start(&app.Config, routes.Gin(app))
	//startWorkers(app)

	// Listen for the interrupt signal.
	<-ctx.Done()
	logger.Infow("the interrupt received, shutting down gracefully, press Ctrl+C again to force")
	stop()

	httpserver.Shutdown(srv, httpserver.ShutdownTimeout)
	//shutdownWorkers()
	app.Config.Shutdown()

	logger.Infow("bye")
}
