package goan

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

//TestSuite holds the data for running all of the tests
type TestSuite struct {
	suite.Suite
	Config Config
}

var originalDatabaseName string

//SetupSuite sets up the test suite by creating a new database and establishing a connection
func (suite *TestSuite) SetupSuite() {
	originalDatabaseName = os.Getenv("GOAN_DBNAME")
	_ = os.Setenv("GOAN_DBNAME", "testing-data")

	con, err := LoadConfig()
	if err != nil {
		panic(err)
	}
	suite.Config = con
}

//TearDownSuite teears down the test suite and closes the database connection
func (suite *TestSuite) TearDownSuite() {
	_ = os.Setenv("GOAN_DBNAME", originalDatabaseName)
	if suite.Config.DatabaseType == "mongo" {
		defer suite.Config.DatabaseMongo.Close()
		DeleteAllTestingEntriesMongo("testing-data", &suite.Config)
	}
}

//Test_Suite_Run runs the test suite
func Test_Suite_Run(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
