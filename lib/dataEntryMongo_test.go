package goan

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	_ "github.com/stretchr/testify/assert"
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
	de.SQLID = 0
	de.MongoID = bson.NewObjectId()
	de.EntryType = "testing-mongo"
	de.Reference = fmt.Sprintf("%d", r1)
	de.EntryCreated = current
	de.Notes = "Sample Notes"

	err := SaveEntryMongo(&de, &suite.Config)
	if err != nil {
		suite.False(true)
	}
    from, _ := time.Parse("2006-01-02", "2016-01-01")
    to, _ := time.Parse("2006-01-02", "2020-01-01")
    sort := Sort{
        Start: 0,
        Count: 1000,
        Field: "date",
        Direction: "asc",
    }
	matches, err := GetEntriesByTypeMongo("testing-mongo", from, to, sort, &suite.Config)
	if err != nil {
		suite.False(true)
	}
	found := false
	r1Check := fmt.Sprintf("%d", r1)
	for _, entry := range matches {
		if entry.Reference == r1Check {
			found = true
		}
	}

	suite.True(found)
}
