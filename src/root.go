package main

import (
	// "fmt"
	"utility"

	"github.com/gin-gonic/gin"
)

const (
	port = ":8085"
)

func main() {

	router := gin.Default()

	// var Dict map[string]string // key-value store for PUT, GET, & DELETE requests (exported variable)
	// fmt.Println(router, Dict)

	utility.PutRequest(router)

	router.Run(port)
}
