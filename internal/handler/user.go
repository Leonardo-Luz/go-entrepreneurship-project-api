package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/leonardo-luz/go-entrepreneurship-project-api/internal/auth"
	"github.com/leonardo-luz/go-entrepreneurship-project-api/internal/config"
	"github.com/leonardo-luz/go-entrepreneurship-project-api/internal/model"
	"github.com/leonardo-luz/go-entrepreneurship-project-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	repository repository.UserRepository
	cfg        *config.Config
}

func NewUserHandler(repository repository.UserRepository, cfg *config.Config) *UserHandler {
	return &UserHandler{repository, cfg}
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

func (handler *UserHandler) RegisterHandler(context *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := handler.repository.Create(&user); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User registered"})
}

func (handler *UserHandler) LoginHandler(context *gin.Context) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := context.ShouldBindJSON(&creds); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	user, err := handler.repository.GetByEmail(creds.Email)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := auth.GenerateJWT(user, handler.cfg.JWTSecret)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	context.SetCookie("jwt", token, int((24 * time.Hour).Seconds()), "/", "", true, true)
	context.JSON(http.StatusOK, gin.H{"message": "Logged in"})
}

func (handler *UserHandler) LogoutHandler(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "/", "", true, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}
