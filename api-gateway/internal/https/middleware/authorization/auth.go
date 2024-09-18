package authorization

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	t "api-gateway/internal/https/token"
	"api-gateway/logger"

	"github.com/gin-gonic/gin"
)

func MiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		logger.Info("Authorization header:", token)
		url := ctx.Request.URL.Path
		logger.Info("Request URL:", url)

		if strings.Contains(url, "swagger") || url == "/api/v1/users" || url == "/api/v1/users/login" || url == "/api/v1/users/verify" {
			ctx.Next()
			return
		}

		fmt.Println(token + "+")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Authorization header is missing",
			})
			return
		}

		if !strings.HasPrefix(token, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Authorization token is missing Bearer prefix",
			})
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		claims, err := t.ExtractClaim(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": err.Error(),
			})
			return
		}
		log.Println(claims)

		ctx.Next()
	}
}
