package goan

import (
	"testing"
    "os"
    "time"
    "math/rand"
    "fmt"
    
	_"github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "gopkg.in/mgo.v2/bson"
)

type DataEntryTestMongoSuite struct {
    suite.Suite
    Config Config
}

var originalDatabaseNameMongo string

func (suite *DataEntryTestMongoSuite) SetupSuite() {
    originalDatabaseNameMongo = os.Getenv("GOAN_DBNAME")
    _ = os.Setenv("GOAN_DBNAME", "testing")
    
    con, err := LoadConfig()
    if err != nil {
        panic(err)
    }
    suite.Config = con
}

func (suite *DataEntryTestMongoSuite) TearDownSuite() {
    _ = os.Setenv("GOAN_DBNAME", originalDatabaseNameMongo)
    if suite.Config.DatabaseType == "mongo" {
        defer suite.Config.DatabaseMongo.Close()
        DeleteAllTestingEntriesMongo("testing-mongo", &suite.Config)
    }
}

func Test_DataEntryMongoSuite(t *testing.T) {
    suite.Run(t, new(DataEntryTestMongoSuite))
}

func (suite *DataEntryTestMongoSuite) Test_DataEntry_MongoSave() {
    de := DataEntry{}
    r1 := rand.Intn(100000000)
    current := time.Now()
    de.SQLId = 0
    de.MongoId = bson.NewObjectId()
    de.EntryType = "testing-mongo"
    de.Reference = fmt.Sprintf("%d", r1)
    de.EntryCreated = current
    de.Notes = "Sample Notes"
    
    err := SaveEntryMongo(&de, &suite.Config)
    if err != nil {
        suite.False(true)
    }
    matches, err := GetEntriesByTypeMongo("testing-mongo", &suite.Config)
    if err != nil {
        suite.False(true)
    }
    found := false
    r1Check := fmt.Sprintf("%d", r1)
    for _, entry := range(matches) {
        if entry.Reference == r1Check {
            found = true
        }
    }
    
    suite.True(found)
}