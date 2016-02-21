package goan

import (
	"testing"
    "os"
    "time"
    _"math/rand"
    _"fmt"
    
	_"github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "gopkg.in/mgo.v2/bson"
)

type DataEntryTestSuite struct {
    suite.Suite
    Config Config
}

var originalDatabaseName string

func (suite *DataEntryTestSuite) SetupSuite() {
    originalDatabaseName = os.Getenv("GOAN_DBNAME")
    _ = os.Setenv("GOAN_DBNAME", "testing-data")
    
    con, err := LoadConfig()
    if err != nil {
        panic(err)
    }
    suite.Config = con
}

func (suite *DataEntryTestSuite) TearDownSuite() {
    _ = os.Setenv("GOAN_DBNAME", originalDatabaseName)
    if suite.Config.DatabaseType == "mongo" {
        defer suite.Config.DatabaseMongo.Close()
        DeleteAllTestingEntriesMongo("testing-data", &suite.Config)
    }
}

func Test_DataEntrySuite(t *testing.T) {
    suite.Run(t, new(DataEntryTestSuite))
}

func (suite *DataEntryTestSuite) Test_DataEntry_ReturnHelper() {
    de := DataEntry{}
    de.SQLId = 0
    de.MongoId = bson.NewObjectId()
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
    _, exists = ret["SQLId"]
    suite.False(exists)
    _, exists = ret["MongoId"]
    suite.False(exists)
}