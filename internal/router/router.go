package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/leonardo-luz/project-builder-api/internal/config"
)

func SetupRouter(database *gorm.DB, cfg *config.Config) (*gin.Engine, error) {
	router := gin.Default()

	router.Use(config.CorsConfig())

	if err := router.SetTrustedProxies(cfg.TrustedProxies); err != nil {
		return nil, err
	}

	router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"status": "OK"})
	})
	router.GET("/favicon.ico", func(context *gin.Context) {
		context.Status(http.StatusNoContent)
	})

	api := router.Group("/api/v1")
	{
		UserRouter(api.Group("/users"), database)
		ProjectRouter(api.Group("/projects"), database)
	}

	return router, nil
}
