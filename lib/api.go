//Package goan contains the actual API; the main.go file simply shells out the execution to LoadAPI()
package goan

import (
	"fmt"
    "strconv"
    "time"
	"github.com/gin-gonic/gin"
)

//Sort holds a structure of sort options
type Sort struct {
    Start int
    Count int
    Field string
    Direction string
}

//LoadAPI is the entry point for the application. It will start the GIN server
func LoadAPI() (*gin.Engine, *Config){
	fmt.Printf("\nLoading API...\n")
    
    //logging first
    SetupLogger()
    
    //Now DB stuff
	config, err := LoadConfig()
	if err != nil {
        LogError.Println(err)
		panic(err)
	}
	if config.DatabaseType == "mongo" {
		defer config.DatabaseMongo.Close()
	}

	//startup the API and setup the routes
	router := gin.Default()
    router.Use(CORSMiddleware())
	v1 := router.Group("/v1")
	{
        //Get basic status
        v1.GET("/", CheckAuthentication(&config), func(c *gin.Context) {
            c.JSON(501, gin.H{"status":"Not Implemented Yet"})
        })
        //Save an entry
		v1.POST("/", CheckAuthentication(&config), func(c *gin.Context) {
			authenticated := c.MustGet("Authenticated").(bool)
            entryType := c.DefaultPostForm("type", "General")
            reference := c.DefaultPostForm("reference", "")
            notes := c.DefaultPostForm("notes", "")
            code, json := SaveEntry(entryType, reference, notes, &config, authenticated)
            c.JSON(code, json)
		})
        
        //Get all the unique types in the db
        v1.GET("/types", CheckAuthentication(&config), func(c *gin.Context) {
            authenticated := c.MustGet("Authenticated").(bool)
            code, json := GetDistinctEntries(&config, authenticated)
            c.JSON(code, json)
        })

        //Get the entries based upon the type
		v1.GET("/types/:entryType", CheckAuthentication(&config), func(c *gin.Context) {
            authenticated := c.MustGet("Authenticated").(bool)
            from := c.DefaultQuery("from", "1969-01-01")
            to := c.DefaultQuery("to", "2020-12-31")
            fromStamp, _ := time.Parse("2006-01-02", from)
            toStamp, _ := time.Parse("2006-01-02", to)
            sort := ParseSort(c, "date")
            code, json := GetEntriesByType(c.Param("entryType"), fromStamp, toStamp, sort, &config, authenticated)
            c.JSON(code, json)
		})
	}
    
    return router, &config
}

//ParseSort parses the Gin context and takes in a default field to sort on.
//All other fields are set with sane defaults.
func ParseSort(c *gin.Context, field string) (Sort) {
    sort := Sort{}
    if c == nil {
        sort.Start = 0
        sort.Count = 100000
        sort.Field = field
        sort.Direction = "asc"
    }else{
        sort.Start, _ = strconv.Atoi(c.DefaultQuery("start", "0"))
        sort.Count, _ = strconv.Atoi(c.DefaultQuery("count", "100000"))
        sort.Field = c.DefaultQuery("sort", field)
        sort.Direction = c.DefaultQuery("sortDirection", "asc")
    }
    return sort
}

//ModifySortForMongo takes a sort object and modifes it for use in the Mongo API
func (sort *Sort) ModifySortForMongo() {
    if sort.Direction == "desc" || sort.Direction == "DESC" {
        sort.Direction = "-"
    } else {
        sort.Direction = ""
    }
    
    if sort.Field == "created" {
        sort.Field = "entryCreated"
    }
    
    if sort.Field == "id" || sort.Field == "MongoID" {
        sort.Field = "_id"
    }
}