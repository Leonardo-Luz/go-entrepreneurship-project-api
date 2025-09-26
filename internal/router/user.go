package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardo-luz/go-entrepreneurship-project-api/internal/config"
	"github.com/leonardo-luz/go-entrepreneurship-project-api/internal/handler"
	"github.com/leonardo-luz/go-entrepreneurship-project-api/internal/middleware"
	"github.com/leonardo-luz/go-entrepreneurship-project-api/internal/repository"
	"gorm.io/gorm"
)

func UserRouter(group *gin.RouterGroup, database *gorm.DB, cfg *config.Config) {
	repository := repository.NewUserRepository(database)
	handler := handler.NewUserHandler(repository, cfg)

	{
		group.GET("/", handler.GetAllHandler)
		group.GET("/:id", handler.GetByIDHandler)
	}

	protected := group.Group("/")

	protected.Use(middleware.AuthMiddleware(cfg.JWTSecret, database))
	{
		protected.POST("/", handler.CreateHandler)
		protected.PUT("/:id", handler.UpdateHandler)
		protected.DELETE("/:id", handler.DeleteHandler)
	}
}
