package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"*"},
			AllowedHeaders:   []string{"Origin"},
			ExposedHeaders:   []string{"Content-Length", "Authorization"},
			AllowCredentials: true,
			MaxAge:           12,
		})
	}
}
