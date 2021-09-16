package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %d %s %s",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC822),
			param.Method,
			param.StatusCode,
			param.Latency,
			param.Path,
		)
	})
}

/*
CORSMiddleware - it is a middleware function
*/
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		email := session.Get("email")
		log.Println("email", email)
		if email == nil {
			log.Println("please login to see the user details")
			http.Redirect(c.Writer, c.Request, "/api/login", http.StatusUnauthorized)
		}
		c.Next()
	}
}
