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

func setupRouter(kvStore map[string]string) *gin.Engine {
	router := gin.Default()
	// if there is a forwarding address, then we the forwarding instance is running & //
	// therefore we can forward requests to the main instance //
	if os.Getenv("FORWARDING_ADDRESS") != "" {
		utility.ForwardRequest(router)
	} else { // otherwise we can directly handle requests as we're the main instance //
		utility.PutRequest(router, kvStore)
		utility.GetRequest(router, kvStore)
		utility.DeleteRequest(router, kvStore)
	}
	return router
}

func main() {

	var kvStore = make(map[string]string) // key-value store for PUT, GET, & DELETE requests (exported variable)

	router := setupRouter(kvStore)
	err := router.Run(port)
	if err != nil {
		fmt.Println("There was an issue attempting to start the server", err, "was returned.")
	}
}
