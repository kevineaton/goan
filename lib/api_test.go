package goan

import (
	"github.com/gin-gonic/gin"
	_ "io/ioutil"
)

func (suite *TestSuite) Test_API_Sort() {
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		sort := ParseSort(c, "id")
		suite.Equal(sort.Field, "id")
		suite.Equal(sort.Direction, "asc")
		suite.Equal(sort.Start, 0)
		suite.Equal(sort.Count, 100000)

		//parse it for mongo
		sort.ModifySortForMongo()
		suite.Equal(sort.Field, "_id")
		suite.Equal(sort.Direction, "")
		suite.Equal(sort.Start, 0)
		suite.Equal(sort.Count, 100000)
	})
	_ = performRequest(router, "GET", "/")
}

func (suite *TestSuite) Test_API_Custom_Sort() {
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		sort := ParseSort(c, "id")
		suite.Equal(sort.Field, "created")
		suite.Equal(sort.Direction, "desc")
		suite.Equal(sort.Start, 3)
		suite.Equal(sort.Count, 10)

		//parse it for mongo
		sort.ModifySortForMongo()
		suite.Equal(sort.Field, "entryCreated")
		suite.Equal(sort.Direction, "-")
		suite.Equal(sort.Start, 3)
		suite.Equal(sort.Count, 10)
	})
	_ = performRequest(router, "GET", "/?sort=created&sortDirection=desc&start=3&count=10")
}

//Todo: Finish _API_ end points; session closing bug
func (suite *TestSuite) Test_API_Types() {
	router, config := LoadAPI()
	auth := config.AuthenticationToken
	suite.Equal(config.Port, ":44889")
	_ = performRequest(router, "GET", "/v1/types?auth="+auth)
}
func (suite *TestSuite) Test_API_Types_Bad() {
	router, _ := LoadAPI()
	_ = performRequest(router, "GET", "/v1/types")
}

func (suite *TestSuite) Test_API_Type() {
	router, config := LoadAPI()
	auth := config.AuthenticationToken
	_ = performRequest(router, "GET", "/v1/types/testing-data?auth="+auth)
}

func (suite *TestSuite) Test_API_Type_Bad() {
	router, _ := LoadAPI()
	_ = performRequest(router, "GET", "/v1/types/testing-data")
}

func (suite *TestSuite) Test_API_Status() {
	router, config := LoadAPI()
	auth := config.AuthenticationToken
	_ = performRequest(router, "GET", "/v1/?auth="+auth)
}

func (suite *TestSuite) Test_API_Status_Bad() {
	router, _ := LoadAPI()
	_ = performRequest(router, "GET", "/v1/")
}

func (suite *TestSuite) Test_API_Create() {
	router, config := LoadAPI()
	auth := config.AuthenticationToken
	_ = performRequest(router, "POST", "/v1/?auth="+auth)
}

func (suite *TestSuite) Test_API_Create_Bad() {
	router, _ := LoadAPI()
	_ = performRequest(router, "POST", "/v1/")
}
