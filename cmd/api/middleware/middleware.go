package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tusfendi/festival-movie-be/config"
)

func JwtAuthMiddleware(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := config.TokenValid(key, c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"response": "gagal", "error": "Anda tidak ada Akses", "error_detail": err.Error()})
			c.Abort()
			return
		}
		// TBU
		c.Set("user_id", 1)
		c.Set("level", "ADMIN")

		c.Next()
	}
}

// cors middleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, x-user-agent, client-key")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS, DELETE, PACTH")

		if c.Request.Method == "OPTIONS" || c.Request.Method == "PUT" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
