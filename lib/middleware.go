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
