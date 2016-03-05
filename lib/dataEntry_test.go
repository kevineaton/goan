package goan

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

func (suite *TestSuite) Test_DataEntry_ReturnHelper() {
	de := DataEntry{}
	de.SQLID = 0
	de.MongoID = bson.NewObjectId()
	de.EntryType = "testing-data"
	de.Reference = "TestingId"
	de.EntryCreated = time.Now()
	de.Notes = "Sample Notes"

	ret := de.EntryReturnHelper()
	suite.Equal(ret["notes"], "Sample Notes")
	suite.Equal(ret["reference"], "TestingId")
	suite.Equal(ret["category"], "testing-data")
	_, exists := ret["created"]
	suite.True(exists)
	_, exists = ret["SQLID"]
	suite.False(exists)
	_, exists = ret["MongoID"]
	suite.False(exists)
}

func (suite *TestSuite) Test_DataEntry_CRUD() {
	de := DataEntry{}
	de.SQLID = 0
	de.MongoID = bson.NewObjectId()
	de.EntryType = "testing-data"
	de.Reference = "TestingId"
	de.EntryCreated = time.Now()
	de.Notes = "Sample Notes"

	code, ret := SaveEntry(de.EntryType, de.Reference, de.Notes, &suite.Config, true)
	suite.Equal(code, 200)
	_, exists := ret["data"]
	suite.True(exists)

	//get a whole bunch and see if we have it
	//todo
	fromStamp, _ := time.Parse("2006-01-02", "2015-11-06")
	toStamp, _ := time.Parse("2006-01-02", "2020-01-08")
	sort := ParseSort(nil, "created")
	fetchCode, _ := GetEntriesByType("testing-data", fromStamp, toStamp, sort, &suite.Config, true)
	suite.Equal(fetchCode, 200)
}

func (suite *TestSuite) Test_DataEntry_Distinct_Types() {
	de := DataEntry{}
	de.SQLID = 0
	de.MongoID = bson.NewObjectId()
	de.EntryType = "testing-data"
	de.Reference = "TestingId"
	de.EntryCreated = time.Now()
	de.Notes = "Sample Notes"

	code, _ := SaveEntry(de.EntryType, de.Reference, de.Notes, &suite.Config, true)
	suite.Equal(code, 200)

	disCode, _ := GetDistinctEntries(&suite.Config, true)
	suite.Equal(disCode, 200)

}

func (suite *TestSuite) Test_DataEntry_Bad_Auth() {
	fromStamp, _ := time.Parse("2006-01-02", "2015-11-06")
	toStamp, _ := time.Parse("2006-01-02", "2020-01-08")
	sort := ParseSort(nil, "created")

	code, _ := SaveEntry("testing-data", "", "", &suite.Config, false)
	suite.Equal(code, 401)

	code, _ = GetDistinctEntries(&suite.Config, false)
	suite.Equal(code, 401)

	code, _ = GetEntriesByType("testing-data", fromStamp, toStamp, sort, &suite.Config, false)
	suite.Equal(code, 401)

}
