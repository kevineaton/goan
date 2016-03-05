//This application runs a very basic Gin server that will take in data and send it back
package main

import (
	"fmt"
	goan "github.com/kevineaton/goan/lib"
)

//Main is the entry point for the application. It will start the GIN server
func main() {
	router, config := goan.LoadAPI()
	fmt.Printf("\nListening on port %s\n", config.Port)
	router.Run(config.Port)
}
