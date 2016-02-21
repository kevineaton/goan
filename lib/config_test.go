package goan

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Config_Default(t *testing.T) {
	config, _ := LoadConfig()
	//test defaults
	assert.Equal(t, config.Port, ":44889")
	assert.NotEqual(t, config.AuthenticationToken, "")
}

func Test_Config_Env(t *testing.T) {
	originalAuthenticationToken := os.Getenv("GOAN_AUTHTOKEN")
	originalPort := os.Getenv("GOAN_API_PORT")
	originalDatabaseType := os.Getenv("GOAN_DBTYPE")
	originalDatabaseURL := os.Getenv("GOAN_DBURL")
	originalDatabaseName := os.Getenv("GOAN_DBNAME")
	originalDatabaseHost := os.Getenv("GOAN_DBHOST")
	originalDatabasePort := os.Getenv("GOAN_DBPORT")
	originalDatabaseUser := os.Getenv("GOAN_DBUSER")
	originalDatabasePassword := os.Getenv("GOAN_DBPASSWORD")
	originalEnv := os.Getenv("GOAN_MODE")

	//set new data to test the environment
	_ = os.Setenv("GOAN_AUTHTOKEN", "reallybadauth")
	_ = os.Setenv("GOAN_API_PORT", "9999")
	_ = os.Setenv("GOAN_DBTYPE", "mongo")
	_ = os.Setenv("GOAN_DBURL", "")
	_ = os.Setenv("GOAN_DBNAME", "testing")
	_ = os.Setenv("GOAN_DBHOST", "serverdoesntexist")
	_ = os.Setenv("GOAN_DBPORT", "888999")
	_ = os.Setenv("GOAN_DBUSER", "dbuser")
	_ = os.Setenv("GOAN_DBPASSWORD", "plainpassword")
	_ = os.Setenv("GOAN_MODE", "testing")

	config, _ := LoadConfig()
	//test defaults
	assert.Equal(t, config.Port, ":9999")
	assert.Equal(t, config.AuthenticationToken, "reallybadauth")
	assert.Equal(t, config.DatabaseType, "mongo")
	assert.Equal(t, config.DatabaseName, "testing")
	assert.Equal(t, config.DatabaseHost, "serverdoesntexist")
	assert.Equal(t, config.DatabasePort, "888999")
	assert.Equal(t, config.DatabaseUser, "dbuser")
	assert.Equal(t, config.DatabasePassword, "plainpassword")

	//now go back
	_ = os.Setenv("GOAN_AUTHTOKEN", originalAuthenticationToken)
	_ = os.Setenv("GOAN_API_PORT", originalPort)
	_ = os.Setenv("GOAN_DBTYPE", originalDatabaseType)
	_ = os.Setenv("GOAN_DBURL", originalDatabaseURL)
	_ = os.Setenv("GOAN_DBNAME", originalDatabaseName)
	_ = os.Setenv("GOAN_DBHOST", originalDatabaseHost)
	_ = os.Setenv("GOAN_DBPORT", originalDatabasePort)
	_ = os.Setenv("GOAN_DBUSER", originalDatabaseUser)
	_ = os.Setenv("GOAN_DBPASSWORD", originalDatabasePassword)
	_ = os.Setenv("GOAN_MODE", originalEnv)
}

func Test_Config_DBURL_Env(t *testing.T) {
	//this one takes the longest due to looking up a bad db
	originalDatabaseURL := os.Getenv("GOAN_DBURL")
	originalDatabaseType := os.Getenv("GOAN_DBTYPE")
	originalEnv := os.Getenv("GOAN_MODE")

	//set new data to test the environment
	_ = os.Setenv("GOAN_DBURL", "mongodb://localhost:27017/testing")
	_ = os.Setenv("GOAN_DBTYPE", "mongo")
	_ = os.Setenv("GOAN_MODE", "testing")

	config, _ := LoadConfig()
	assert.Equal(t, config.DatabaseURL, "mongodb://localhost:27017/testing")

	//now go back
	_ = os.Setenv("GOAN_DBURL", originalDatabaseURL)
	_ = os.Setenv("GOAN_DBTYPE", originalDatabaseType)
	_ = os.Setenv("GOAN_MODE", originalEnv)
}
