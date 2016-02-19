package goan

import (
    "fmt"
    "time"
    gin "github.com/gin-gonic/gin"
)

type DataEntry struct {
    Id string
    EntryType string
    Reference string
    EntryCreated time.Time
    Notes string
}


func SaveEntry(c *gin.Context, config *Config) {
    //parse
    fmt.Println("Here")
    if !c.MustGet("Authenticated").(bool) {
		c.JSON(401, gin.H{"status": "Unauthorized"})
	} else {
        data := DataEntry{}
		data.EntryType= c.DefaultPostForm("type", "General")
        data.Reference = c.DefaultPostForm("reference", "")
        data.EntryCreated = time.Now()
        data.Notes = c.DefaultPostForm("notes", "")
		
        //how to save?
        if config.DatabaseType == "mongo" {
            err := SaveEntryMongo(&data, config)
            if err != nil {
                c.JSON(500, gin.H{"status": "Could not save that entry"})
            } else {
                dataRet := gin.H{"type": data.EntryType, "reference": data.Reference, "created": data.EntryCreated, "notes": data.Notes}
                c.JSON(200, gin.H{"status": "Inserted", "data": dataRet})
            }
        }
    }
}