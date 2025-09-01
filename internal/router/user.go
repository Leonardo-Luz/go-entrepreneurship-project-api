package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardo-luz/project-builder-api/internal/handler"
	"github.com/leonardo-luz/project-builder-api/internal/repository"
	"gorm.io/gorm"
)

func UserRouter(group *gin.RouterGroup, database *gorm.DB) {
	repository := repository.NewUserRepository(database)
	handler := handler.NewUserHandler(repository)
	{
		group.GET("/", handler.GetAllHandler)
		group.GET("/:id", handler.GetByIDHandler)
		group.POST("/", handler.CreateHandler)
		group.PUT("/:id", handler.UpdateHandler)
		group.DELETE("/:id", handler.DeleteHandler)
	}
}
