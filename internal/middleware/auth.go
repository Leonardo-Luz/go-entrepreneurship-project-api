package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardo-luz/project-builder-api/internal/auth"
	"github.com/leonardo-luz/project-builder-api/internal/model"
	"gorm.io/gorm"
)

func AuthMiddleware(secret string, db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		cookie, err := context.Request.Cookie("jwt")
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			context.Abort()
			return
		}

		claims, err := auth.ValidateJWT(cookie.Value, secret)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			context.Abort()
			return
		}

		var user model.User
		if err := db.First(&user, "id = ?", claims.UserID).Error; err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			context.Abort()
			return
		}

		context.Set("user", &user)
		context.Next()
	}
}
