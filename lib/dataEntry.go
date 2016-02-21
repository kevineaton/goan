package goan

import (
    "fmt"
    "reflect"
    "time"
    gin "github.com/gin-gonic/gin"
    
    "gopkg.in/mgo.v2/bson"
)

//DataEntry is the primary data structure for the events that are logged
//
// The primary fields are:
// - EntryType: The type of event, such as "login" or "userSubscribed"
// - Reference: An optional field for referencing another record. For example, if the "entryType" is "user", this could hold a reference to the user Id
// - EntryCreated: The time the event occurred
// - Notes: Any optional notes to go along with the event
type DataEntry struct {
    MongoId bson.ObjectId `bson:"_id"`
    SQLId int `bson:"_sqlid"`
    EntryType string `bson:"entryType" json:"entryType"`
    Reference string `bson:"reference" json:"reference"`
    EntryCreated time.Time `bson:"entryCreated" json:"entryCreated"`
    Notes string `bson:"notes" json:"notes"`
}

//EntryReturnHelper formats the return of a DataEntry object back to the client
func (dataEntry DataEntry) EntryReturnHelper() gin.H {
    ret :=gin.H{
        "category": dataEntry.EntryType,
        "reference": dataEntry.Reference,
        "created": dataEntry.EntryCreated,
        "notes": dataEntry.Notes,
    }
    
    return ret
}

//SaveEntry saves a new event entry to the database
func SaveEntry(c *gin.Context, config *Config) {
    //parse
    if !c.MustGet("Authenticated").(bool) {
		c.JSON(401, gin.H{"status": "Unauthorized"})
	} else {
        data := DataEntry{}
        //the id will be set either by the database (SQL) or the function call (mongo)
		data.EntryType= c.DefaultPostForm("type", "General")
        data.Reference = c.DefaultPostForm("reference", "")
        data.EntryCreated = time.Now()
        data.Notes = c.DefaultPostForm("notes", "")
		
        //how to save?
        if config.DatabaseType == "mongo" {
            err := SaveEntryMongo(&data, config)
            if err != nil {
                c.JSON(500, gin.H{"status": "Could not save that entry"})
                panic(err)
            } else {
                dataRet := data.EntryReturnHelper()
                fmt.Println(reflect.TypeOf(dataRet))
                c.JSON(200, gin.H{"status": "inserted", "data": dataRet})
            }
        }
    }
}

//GetEntriesByType gets a list of all entries filtered by a specific type
func GetEntriesByType(entryType string, c *gin.Context, config *Config) {
    if !c.MustGet("Authenticated").(bool) {
		c.JSON(401, gin.H{"status": "Unauthorized"})
	} else {
        if config.DatabaseType == "mongo" {
            matches, err := GetEntriesByTypeMongo(entryType, config)
            if err != nil {
                c.JSON(500, gin.H{"status": "There was a problem"})
            } else {
                //loop and build
                ret := []gin.H{}
                for _, entry := range(matches) {
                    fmt.Printf("\n\tFound: %s", entry.EntryType)
                    ret = append(ret, entry.EntryReturnHelper())
                }
                count := len(matches)
                c.JSON(200, gin.H{"status": "OK", "count": count, "data": ret})
            }
        }
    }
}