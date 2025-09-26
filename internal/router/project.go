package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardo-luz/go-entrepreneurship-project-api/internal/handler"
	"github.com/leonardo-luz/go-entrepreneurship-project-api/internal/repository"
	"gorm.io/gorm"
)

func ProjectRouter(group *gin.RouterGroup, database *gorm.DB) {
	repository := repository.NewProjectRepository(database)
	handler := handler.NewProjectHandler(repository)
	{
		group.GET("/", handler.GetAllHandler)
		group.GET("/:id", handler.GetByIDHandler)
		group.POST("/", handler.CreateHandler)
		group.PUT("/:id", handler.UpdateHandler)
		group.DELETE("/:id", handler.DeleteHandler)
	}
}
