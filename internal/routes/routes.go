package routes

import (
	coreconfig "github.com/Lameaux/core/config"
	"github.com/gin-gonic/gin"

	"github.com/Lameaux/auth/internal/config"
	ih "github.com/Lameaux/auth/internal/index/handlers"
	coremiddlewares "github.com/Lameaux/core/middlewares"
)

func Gin(app *config.App) *gin.Engine { //nolint:funlen
	r := gin.Default()

	switch app.Config.Env {
	case coreconfig.EnvTest:
		gin.SetMode(gin.TestMode)
	case coreconfig.EnvProduction:
		gin.SetMode(gin.ReleaseMode)
	}

	r.Use(coremiddlewares.Timeout(app.Config.WaitTimeout))

	i := ih.NewHandler()

	r.GET("/", i.Index)
	r.GET("/health", i.Index)

	return r
}
