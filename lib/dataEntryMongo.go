package goan

import (
    "gopkg.in/mgo.v2/bson"
)

//SaveEntryMongo saves a DataEntry struct to Mongo
func SaveEntryMongo(data *DataEntry, config *Config) error {
    //we create the id ourselves since that is what mongo wants
    data.MongoId = bson.NewObjectId()
    collection := config.DatabaseMongo.DB(config.DatabaseName).C("entries")
    err := collection.Insert(&data)
    return err
}

//GetEntriesByTypeMongo returns a slice of DataEntry objects by the specified filters
func GetEntriesByTypeMongo(entryType string, config *Config) ([]DataEntry, error) {
    collection := config.DatabaseMongo.DB(config.DatabaseName).C("entries")
    matches := []DataEntry{}
    err := collection.Find(bson.M{"entryType": entryType}).All(&matches)
    if err != nil {
        panic("yikes")
    }
    return matches, err
}

func DeleteAllTestingEntriesMongo(entryType string, config *Config) error {
    collection := config.DatabaseMongo.DB(config.DatabaseName).C("entries")
    _, err := collection.RemoveAll(bson.M{"entryType": entryType})
    return err
}