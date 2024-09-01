package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kunnoh/lms-api/src/repository"
	"github.com/kunnoh/lms-api/src/utils"
)

func DeserializeUser(usersRepo *repository.UserRepository) gin.HandlerFunc {
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
		// fmt.Println(claims)
		// fmt.Println(verifiedTok)

		if claims, ok := verifiedTok.Claims.(jwt.MapClaims); ok && verifiedTok.Valid {
			// Access the claims
			fmt.Println("Audience:", claims["aud"])
			fmt.Println("Expiration:", claims["exp"])
			fmt.Println("Issuer:", claims["iss"])
			fmt.Println("Role:", claims["role"])
			fmt.Println("Subject:", claims["sub"])
		} else {
			fmt.Println("Invalid token")
		}

		// id, err_id := strconv.Atoi(fmt.Sprint(sub))
		// if err_id != nil{
		// 	fmt.Println(err_id)
		// }
		// res, err := usersRepo.FindById(sub)

	}
}
