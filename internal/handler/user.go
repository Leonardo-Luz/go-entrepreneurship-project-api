package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/leonardo-luz/project-builder-api/internal/model"
	"github.com/leonardo-luz/project-builder-api/internal/repository"
)

type UserHandler struct {
	repository repository.UserRepository
}

func NewUserHandler(repository repository.UserRepository) *UserHandler {
	return &UserHandler{repository}
}

func (handler *UserHandler) GetAllHandler(context *gin.Context) {
	users, err := handler.repository.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Users"})
		return
	}

	context.JSON(http.StatusOK, users)
}

func (handler *UserHandler) GetByIDHandler(context *gin.Context) {
	id := context.Param("id")

	if parsedId, err := uuid.Parse(id); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	} else {
		user, err := handler.repository.GetByID(parsedId)

		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": "Failed to fetch User"})
			return
		}

		context.JSON(http.StatusOK, user)
	}
}

func (handler *UserHandler) CreateHandler(context *gin.Context) {
	var user model.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := handler.repository.Create(&user); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create User"})
		return
	}

	context.JSON(http.StatusCreated, user)
}

func (handler *UserHandler) UpdateHandler(context *gin.Context) {
	id := context.Param("id")
	var user model.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if parsedId, err := uuid.Parse(id); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	} else {
		user.ID = parsedId
	}

	if err := handler.repository.Update(&user); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update User"})
		return
	}

	context.JSON(http.StatusOK, user)
}

func (handler *UserHandler) DeleteHandler(context *gin.Context) {
	id := context.Param("id")

	if parsedId, err := uuid.Parse(id); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	} else {
		if err := handler.repository.Delete(parsedId); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete User"})
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
