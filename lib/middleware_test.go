package goan

import (
	_ "net/http"
	_ "net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_Middleware_Default_Bad(t *testing.T) {
	config, _ := LoadConfig()
	router := gin.New()
	router.Use(CheckAuthentication(&config))
	router.Use(CORSMiddleware())
	router.GET("/?auth=reallybadauth", func(c *gin.Context) {
		assert.False(t, c.MustGet("Authenticated").(bool))
	})
	_ = performRequest(router, "GET", "/testing?auth=reallybadauth")
}

func Test_Middleware_Default(t *testing.T) {
	originalAuthenticationToken := os.Getenv("GOAN_AUTHTOKEN")
	_ = os.Setenv("GOAN_AUTHTOKEN", "reallybadauth")
	config, _ := LoadConfig()

	assert.Equal(t, config.AuthenticationToken, "reallybadauth")

	router := gin.New()
	router.Use(CheckAuthentication(&config))
	router.Use(CORSMiddleware())
	router.GET("/test?auth=reallybadauth", func(c *gin.Context) {
		assert.True(t, c.MustGet("Authenticated").(bool))
	})
	_ = performRequest(router, "GET", "/test?auth=reallybadauth")
	_ = os.Setenv("GOAN_AUTHTOKEN", originalAuthenticationToken)
}

func Test_CORS_Default(t *testing.T) {
	originalAuthenticationToken := os.Getenv("GOAN_AUTHTOKEN")
	_ = os.Setenv("GOAN_AUTHTOKEN", "reallybadauth")
	config, _ := LoadConfig()

	assert.Equal(t, config.AuthenticationToken, "reallybadauth")

	router := gin.New()
	router.Use(CheckAuthentication(&config))
	router.Use(CORSMiddleware())
	router.GET("/test?auth=reallybadauth", func(c *gin.Context) {
		assert.Equal(t, c.MustGet("Access-Control-Origin").(string), "*")
		assert.True(t, c.MustGet("Authenticated").(bool))
	})
	_ = performRequest(router, "GET", "/test?auth=reallybadauth")
	_ = os.Setenv("GOAN_AUTHTOKEN", originalAuthenticationToken)
}

func Test_CORS_Options(t *testing.T) {
	originalAuthenticationToken := os.Getenv("GOAN_AUTHTOKEN")
	_ = os.Setenv("GOAN_AUTHTOKEN", "reallybadauth")
	config, _ := LoadConfig()

	assert.Equal(t, config.AuthenticationToken, "reallybadauth")

	router := gin.New()
	router.Use(CheckAuthentication(&config))
	router.Use(CORSMiddleware())
	router.OPTIONS("/test?auth=reallybadauth", func(c *gin.Context) {
		assert.Equal(t, c.MustGet("Access-Control-Origin").(string), "*")
	})
	_ = performRequest(router, "OPTIONS", "/test?auth=reallybadauth")
	_ = os.Setenv("GOAN_AUTHTOKEN", originalAuthenticationToken)
}
