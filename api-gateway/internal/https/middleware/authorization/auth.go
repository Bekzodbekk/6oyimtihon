package authorization

import (
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

        if strings.Contains(url, "swagger") || url == "/api/v1/users" || url == "/api/v1/users/login" {
            ctx.Next()
            return
        }

        if token == "" {
            ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
                "error": "Authorization header is missing",
            })
            return
        }

        // Bearer prefix borligini tekshirish
        if !strings.HasPrefix(token, "Bearer ") {
            ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
                "error": "Authorization token is missing Bearer prefix",
            })
            return
        }

        // Bearer prefix ni ochirish
        token = strings.TrimPrefix(token, "Bearer ")

        // Tokenni extract qilish
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
