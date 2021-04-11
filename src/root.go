package main

import (
	"utility"

	"github.com/gin-gonic/gin"
)

const (
	port = ":8085"
)

func main() {

	router := gin.Default()

	var dict = make(map[string]string) // key-value store for PUT, GET, & DELETE requests (exported variable)

	utility.PutRequest(router, dict)
	utility.GetRequest(router, dict)
	utility.DeleteRequest(router, dict)

	router.Run(port)
}
