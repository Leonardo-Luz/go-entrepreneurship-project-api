package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/leonardo-luz/project-builder-api/internal/config"
	"github.com/leonardo-luz/project-builder-api/internal/handler"
	"github.com/leonardo-luz/project-builder-api/internal/repository"
)

func AuthRouter(group *gin.RouterGroup, database *gorm.DB, cfg *config.Config) {
	repository := repository.NewUserRepository(database)
	handler := handler.NewUserHandler(repository, cfg)

	{
		group.POST("/register", handler.RegisterHandler)
		group.POST("/login", handler.LoginHandler)
		group.POST("/logout", handler.LogoutHandler)
	}
}
