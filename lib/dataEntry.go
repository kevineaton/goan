package goan

import (
	"github.com/gin-gonic/gin"
	"time"

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
	MongoID      bson.ObjectId `bson:"_id"`
	SQLID        int           `bson:"_sqlid"`
	EntryType    string        `bson:"entryType" json:"entryType"`
	Reference    string        `bson:"reference" json:"reference"`
	EntryCreated time.Time     `bson:"entryCreated" json:"entryCreated"`
	Notes        string        `bson:"notes" json:"notes"`
}

//EntryReturnHelper formats the return of a DataEntry object back to the client
func (dataEntry DataEntry) EntryReturnHelper() gin.H {
	ret := gin.H{
		"category":  dataEntry.EntryType,
		"reference": dataEntry.Reference,
		"created":   dataEntry.EntryCreated,
		"notes":     dataEntry.Notes,
	}

	return ret
}

//SaveEntry saves a new event entry to the database
func SaveEntry(entryType string, reference string, notes string, config *Config, authenticated bool) (int, gin.H) {
	//parse
	if !authenticated {
		return 401, gin.H{"status": "Unauthorized"}
	}
	data := DataEntry{}
	//the id will be set either by the database (SQL) or the function call (mongo)
	data.EntryType = entryType
	data.Reference = reference
	data.EntryCreated = time.Now()
	data.Notes = notes

	//how to save?
	if config.DatabaseType == "mongo" {
		err := SaveEntryMongo(&data, config)
		if err != nil {
			LogWarning.Println("Invalid post, data not saved")
			return 500, gin.H{"status": "Could not save that entry"}
		}
		dataRet := data.EntryReturnHelper()
		return 200, gin.H{"status": "inserted", "data": dataRet}
	}
	return 501, gin.H{"status": "Other databases not Implemented Yet"}
}

//GetEntriesByType gets a list of all entries filtered by a specific type
func GetEntriesByType(entryType string, from time.Time, to time.Time, sort Sort, config *Config, authenticated bool) (int, gin.H) {
	if !authenticated {
		return 401, gin.H{"status": "Unauthorized"}
	}
	if config.DatabaseType == "mongo" {
		matches, err := GetEntriesByTypeMongo(entryType, from, to, sort, config)
		if err != nil {
			return 500, gin.H{"status": "There was a problem"}
		}
		//loop and build
		ret := []gin.H{}
		for _, entry := range matches {
			ret = append(ret, entry.EntryReturnHelper())
		}
		count := len(matches)
		return 200, gin.H{"status": "OK", "count": count, "data": ret}
	}
	return 501, gin.H{"status": "Other databases not Implemented Yet"}
}

//GetDistinctEntries gets the distinct entries
func GetDistinctEntries(config *Config, authenticated bool) (int, gin.H) {
	if !authenticated {
		return 401, gin.H{"status": "Unauthorized"}
	}

	if config.DatabaseType == "mongo" {
		matches, err := GetDistinctEntriesMongo(config)
		if err != nil {
			return 500, gin.H{"status": "There was a problem"}
		}
		//loop and build
		count := len(matches)
		return 200, gin.H{"status": "OK", "count": count, "data": matches}
	}
	return 501, gin.H{"status": "Other databases not Implemented Yet"}
}
