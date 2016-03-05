package goan

import (
	"crypto/md5"
	"errors"
	"fmt"
	"math/rand"
	"os"

	"gopkg.in/mgo.v2"
)

//Config holds basic configuration options for the application, including database settings and authentication
type Config struct {
	AuthenticationToken string
	Port                string
	DatabaseType        string // currently only Mongo, with MySQL and others coming soon*
	DatabaseName        string
	DatabaseHost        string
	DatabaseUser        string
	DatabasePassword    string
	DatabasePort        string
	DatabaseURL         string //overrides other entries
	DatabaseMongo       mgo.Session
}

//LoadConfig will load up a new configuration struct with sane defaults if none provided
func LoadConfig() (Config, error) {
	config := Config{}
	config.AuthenticationToken = os.Getenv("GOAN_AUTHTOKEN")
	config.Port = os.Getenv("GOAN_API_PORT")
	config.DatabaseType = os.Getenv("GOAN_DBTYPE")
	config.DatabaseURL = os.Getenv("GOAN_DBURL")
	config.DatabaseName = os.Getenv("GOAN_DBNAME")
	config.DatabaseHost = os.Getenv("GOAN_DBHOST")
	config.DatabasePort = os.Getenv("GOAN_DBPORT")
	config.DatabaseUser = os.Getenv("GOAN_DBUSER")
	config.DatabasePassword = os.Getenv("GOAN_DBPASSWORD")
	goanMode := os.Getenv("GOAN_MODE")

	if config.AuthenticationToken == "" {
		//randomize it with bcrypt on each server start up and prompt the user to specify one
		r1 := rand.Intn(100000000)
		r2 := rand.Intn(20000000)
		plain := fmt.Sprintf("%s-%d-%d", "go-goan-kv", r1, r2)
		h := md5.New()
		h.Write([]byte(plain))
		code := string(fmt.Sprintf("%x", h.Sum(nil)))
		config.AuthenticationToken = code
		fmt.Printf("\n\t----------------------------------------------")
		fmt.Printf("\n\t-   Auth: %s   -", code)
		fmt.Printf("\n\t----------------------------------------------\n\n")
	}

	if config.Port == "" {
		config.Port = "44889"
	}
	config.Port = fmt.Sprintf(":%s", config.Port)

	//handle the database
	if config.DatabaseType == "" {
		config.DatabaseType = "mongo"
	}
	if config.DatabaseType == "mongo" {
		//setup mongo connection
		if config.DatabaseURL == "" {
			//build it with the supplied params
			//mongodb://myuser:mypass@localhost:40001
			config.DatabaseURL = "mongodb://"
			if config.DatabaseHost == "" {
				config.DatabaseHost = "localhost"
			}
			if config.DatabasePort == "" {
				config.DatabasePort = "27017"
			}
			if config.DatabaseUser != "" && config.DatabasePassword != "" {
				config.DatabaseURL = fmt.Sprintf("%s%s:%s@%s:%s", config.DatabaseURL,
					config.DatabaseUser, config.DatabasePassword, config.DatabaseHost, config.DatabasePort)
			} else {
				config.DatabaseURL = fmt.Sprintf("%s%s:%s", config.DatabaseURL, config.DatabaseHost, config.DatabasePort)
			}
		}
		if config.DatabaseName == "" {
			config.DatabaseName = "goan"
		}
		if goanMode != "testing" {
			session, err := mgo.Dial(config.DatabaseURL)
			if err != nil {
				return config, err
			}
			session.SetMode(mgo.Monotonic, true)
			config.DatabaseMongo = *session
		} else {
			return config, errors.New("testing mode")
		}

	} else if config.DatabaseType == "mysql" {

	}

	return config, nil
}
