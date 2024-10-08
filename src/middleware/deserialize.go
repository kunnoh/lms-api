package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	userrepository "github.com/kunnoh/lms-api/src/repository/user.repository"
	"github.com/kunnoh/lms-api/src/utils"
)

func DeserializeUser(usersRepo userrepository.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		authHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": http.StatusForbidden, "status": "fail", "error": "you're not logged in"})
			return
		}

		verifiedTok, err := utils.ValidateToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": http.StatusForbidden, "status": "fail", "error": err.Error()})
			return
		}

		// Access the claims
		var claims jwt.MapClaims
		var ok bool

		if claims, ok = verifiedTok.Claims.(jwt.MapClaims); !ok || !verifiedTok.Valid {
			fmt.Println("Invalid token")
		}

		// Get the user ID from the claims and cast it to a string
		userId, ok := claims["sub"].(string)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": http.StatusForbidden, "status": "fail", "error": "invalid token"})
			return
		}

		// Get the user from the repository
		user, err := usersRepo.FindById(userId)
		if err != nil {
			if err.Error() == "user not found" {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": http.StatusUnauthorized, "status": "fail", "error": "you are not logged in"})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "status": "fail", "error": "error in login in"})
			return
		}

		ctx.Set("currentUser", user.Email)
		ctx.Set("currentUserID", user.UserId)
		ctx.Next()

	}
}
