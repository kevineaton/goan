package goan

import (
    "fmt"
    
    _"gopkg.in/mgo.v2"
    _"gopkg.in/mgo.v2/bson"
)
func SaveEntryMongo(data *DataEntry, config *Config) error {
    fmt.Println("Saving")
    //we create the id ourselves since that is what mongo wants
    
    collection := config.DatabaseMongo.DB(config.DatabaseName).C("entries")
    err := collection.Insert(&data)
    return err
}