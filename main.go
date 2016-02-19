//This application runs a very basic Gin server that will take in data and send it back
package main

import (
	_"crypto/md5"
	"fmt"
	_"math/rand"
	_"os"
	_"strings"
    _"log"
    
    _"gopkg.in/mgo.v2"
    _"gopkg.in/mgo.v2/bson"
	gin "github.com/gin-gonic/gin"
    goan "github.com/kevineaton/goan/lib"
)

func main() {
    fmt.Printf("\nLoading...\n")
    config := goan.LoadConfig()
    if config.DatabaseType == "mongo" {
        defer config.DatabaseMongo.Close()
    }
    
    //startup the API and setup the routes
    router := gin.Default()
    v1 := router.Group("/v1") 
    {
        v1.POST("/", goan.CheckAuthentication(&config), func(c *gin.Context) {
            goan.SaveEntry(c, &config)
        })
    }
	
    fmt.Printf("\nListening on port %s\n", config.Port)
	router.Run(config.Port)
}

