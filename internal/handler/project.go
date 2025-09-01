package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/leonardo-luz/project-builder-api/internal/model"
	"github.com/leonardo-luz/project-builder-api/internal/repository"
)

type ProjectHandler struct {
	repository repository.ProjectRepository
}

func NewProjectHandler(repository repository.ProjectRepository) *ProjectHandler {
	return &ProjectHandler{repository}
}

func (handler *ProjectHandler) GetAllHandler(context *gin.Context) {
	projects, err := handler.repository.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Projects"})
		return
	}

	context.JSON(http.StatusOK, projects)
}

func (handler *ProjectHandler) GetByIDHandler(context *gin.Context) {
	id := context.Param("id")

	if parsedId, err := uuid.Parse(id); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	} else {
		project, err := handler.repository.GetByID(parsedId)

		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": "Failed to fetch Project"})
			return
		}

		context.JSON(http.StatusOK, project)
	}
}

func (handler *ProjectHandler) CreateHandler(context *gin.Context) {
	var project model.Project

	if err := context.ShouldBindJSON(&project); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := handler.repository.Create(&project); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Project"})
		return
	}

	context.JSON(http.StatusCreated, project)
}

func (handler *ProjectHandler) UpdateHandler(context *gin.Context) {
	id := context.Param("id")
	var project model.Project

	if err := context.ShouldBindJSON(&project); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if parsedId, err := uuid.Parse(id); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	} else {
		project.ID = parsedId
	}

	if err := handler.repository.Update(&project); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Project"})
		return
	}

	context.JSON(http.StatusOK, project)
}

func (handler *ProjectHandler) DeleteHandler(context *gin.Context) {
	id := context.Param("id")

	if parsedId, err := uuid.Parse(id); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	} else {
		if err := handler.repository.Delete(parsedId); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Project"})
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{"message": "Project deleted"})
}
