//The goan package contains the actual API; the main.go file simply shells out the execution to LoadAPI()
package goan

import (
	"fmt"
	gin "github.com/gin-gonic/gin"
)

//Main is the entry point for the application. It will start the GIN server
func LoadAPI() (*gin.Engine, *Config){
	fmt.Printf("\nLoading API...\n")
	config, err := LoadConfig()
	if err != nil {
		panic(err)
	}
	if config.DatabaseType == "mongo" {
		defer config.DatabaseMongo.Close()
	}

	//startup the API and setup the routes
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/", CheckAuthentication(&config), func(c *gin.Context) {
			SaveEntry(c, &config)
		})

		v1.GET("/:entryType", CheckAuthentication(&config), func(c *gin.Context) {
			GetEntriesByType(c.Param("entryType"), c, &config)
		})
	}

	fmt.Printf("\nListening on port %s\n", config.Port)
    router.Run(config.Port)
    
    return router, &config
}
