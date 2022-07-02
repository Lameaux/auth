package config

import (
	"os"
	"time"

	coreconfig "github.com/Lameaux/core/config"
	coretesting "github.com/Lameaux/core/testing"
)

var tables = []string{ //nolint:gochecknoglobals
	"users",
}

const (
	AppName    = "auth"
	AppVersion = "0.1"

	defaultWorkerSleep = 5
)

type App struct {
	Config coreconfig.AppConfig

	WorkerSleep time.Duration
}

func NewApp() *App {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = coreconfig.EnvDevelopment
	}

	app := defaultApp(env)
	return app
}

func NewTestApp() *App {
	coretesting.SetWorkingDir()

	app := defaultApp(coreconfig.EnvTest)

	coretesting.CleanupDatabase(app.Config.DBPool, tables)

	return app
}

func defaultApp(env string) *App {
	return &App{
		Config:      *coreconfig.NewAppConfig(env),
		WorkerSleep: time.Duration(defaultWorkerSleep) * time.Second,
	}
}
