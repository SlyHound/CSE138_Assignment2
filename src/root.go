package main

import (
	"fmt"
	"os"
	"src/utility"

	"github.com/gin-gonic/gin"
)

const (
	port = ":8085"
)

func main() {

	router := gin.Default()

	var dict = make(map[string]string) // key-value store for PUT, GET, & DELETE requests (exported variable)

	// if there is a forwarding address, then we the forwarding instance is running & //
	// therefore we can forward requests to the main instance //
	if os.Getenv("FORWARDING_ADDRESS") != "" {
		utility.ForwardRequest(router)
	} else { // otherwise we can directly handle requests as we're the main instance //
		utility.PutRequest(router, dict)
		utility.GetRequest(router, dict)
		utility.DeleteRequest(router, dict)
	}
	err := router.Run(port)
	if err != nil {
		fmt.Println("There was an issue attempting to start the server", err, "was returned.")
	}
}
