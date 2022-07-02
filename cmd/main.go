package main

import (
	"context"
	"github.com/Lameaux/auth/internal/config"
	"github.com/Lameaux/core/logger"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app := config.NewApp()
	logger.Infow("Starting", "app", config.AppName, "version", config.AppVersion)

	//srv := startAPIServer(app)
	//startWorkers(app)

	// Listen for the interrupt signal.
	<-ctx.Done()
	logger.Infow("the interrupt received, shutting down gracefully, press Ctrl+C again to force")
	stop()

	//shutdownAPIServer(srv)
	//shutdownWorkers()
	app.Config.Shutdown()

	logger.Infow("bye")
}
