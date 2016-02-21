package goan

import (
	"testing"
    "os"
    
	_"net/http"
	_"net/http/httptest"

	"github.com/stretchr/testify/assert"
    gin "github.com/gin-gonic/gin"
)

func Test_Middleware_Default_Bad(t *testing.T) {
    config, _ := LoadConfig()
    router := gin.New()
    router.Use(CheckAuthentication(&config))
    router.GET("/?auth=reallybadauth", func(c *gin.Context){
        assert.False(t, c.MustGet("Authenticated").(bool))
    })
    _ = performRequest(router, "GET", "/?auth=reallybadauth")
}

func Test_Middleware_Default(t *testing.T) {
    originalAuthenticationToken := os.Getenv("GOAN_AUTHTOKEN")
	_ = os.Setenv("GOAN_AUTHTOKEN", "reallybadauth")
	config, _ := LoadConfig()
    
    assert.Equal(t, config.AuthenticationToken, "reallybadauth")
    
    router := gin.New()
    router.Use(CheckAuthentication(&config))
    router.GET("/test?auth=reallybadauth", func(c *gin.Context){
        assert.True(t, c.MustGet("Authenticated").(bool))
    })
    _ = performRequest(router, "GET", "/test?auth=reallybadauth")
    _ = os.Setenv("GOAN_AUTHTOKEN", originalAuthenticationToken)
}