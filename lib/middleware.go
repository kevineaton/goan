package goan

import (
	gin "github.com/gin-gonic/gin"
)

//CheckAuthentication will check the user's authentication token in either the
//post body or the query string (in that priority order)
func CheckAuthentication(config *Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		passedAuth := c.DefaultPostForm("auth", "")
		if passedAuth == "" {
			passedAuth = c.DefaultQuery("auth", "")
		}

		if passedAuth == config.AuthenticationToken {
			c.Set("Authenticated", true)
		} else {
			c.Set("Authenticated", false)
		}
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
