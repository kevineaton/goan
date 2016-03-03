package goan

import (
    "time"
    "fmt"
	"gopkg.in/mgo.v2/bson"
)

//SaveEntryMongo saves a DataEntry struct to Mongo
func SaveEntryMongo(data *DataEntry, config *Config) error {
	//we create the id ourselves since that is what mongo wants
	data.MongoID = bson.NewObjectId()
	collection := config.DatabaseMongo.DB(config.DatabaseName).C("entries")
	err := collection.Insert(&data)
	return err
}

//GetEntriesByTypeMongo returns a slice of DataEntry objects by the specified filters
func GetEntriesByTypeMongo(entryType string, from time.Time, to time.Time, sort Sort, config *Config) ([]DataEntry, error) {
	collection := config.DatabaseMongo.DB(config.DatabaseName).C("entries")
	matches := []DataEntry{}
    query := bson.M{
        "entryType": entryType,
        "entryCreated": bson.M{
            "$gte": from,
            "$lt": to,
        },
    }
    //modify the sort
    sort.ModifySortForMongo()
    sortString := fmt.Sprintf("%s%s", sort.Direction, sort.Field)
    
	err := collection.Find(query).Limit(sort.Count).Skip(sort.Start).Sort(sortString).All(&matches)
	if err != nil {
		panic("yikes")
	}
	return matches, err
}

//GetDistinctEntriesMongo gets the distinct entry types that have been input
func GetDistinctEntriesMongo(config *Config) ([]string, error) {
    collection := config.DatabaseMongo.DB(config.DatabaseName).C("entries")
    var result []string
    err := collection.Find(bson.M{}).Distinct("entryType", &result)
    if err != nil {
		panic("yikes")
	}
	return result, err
}



//DeleteAllTestingEntriesMongo removes all entries of a specific type. While it can be used for all removals,
//including non-testing, you probably don't really want to delete all of your entries
func DeleteAllTestingEntriesMongo(entryType string, config *Config) error {
	collection := config.DatabaseMongo.DB(config.DatabaseName).C("entries")
	_, err := collection.RemoveAll(bson.M{"entryType": entryType})
	return err
}
