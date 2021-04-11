package main

import (
	"fmt"
	"utility"

	"github.com/gin-gonic/gin"
)

const (
	port = ":8085"
)

func main() {

	router := gin.Default()

	var dict = make(map[string]string) // key-value store for PUT, GET, & DELETE requests (exported variable)
	// fmt.Println(router, Dict)
	utility.PutRequest(router, dict)
	for key, value := range dict {
		fmt.Printf("key %s, value %s\n", key, value)
	}

	router.Run(port)
}
