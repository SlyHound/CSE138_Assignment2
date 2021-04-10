package utility

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	keyLimit int = 50 // maximum number of characters allowed for a key
)

func PrintHelloWorld(str string) {
	fmt.Println(str)
}

func PutRequest(r *gin.Engine) {
	r.PUT("/key-value-store/:key", func(c *gin.Context) {
		key := c.Param("key")
		body, _ := ioutil.ReadAll(c.Request.Body)
		strBody := string(body[:])
		if strBody == "{}" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Value is missing", "message": "Error in PUT"})
		} else if len(key) > keyLimit {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Key is too long", "message": "Error in PUT"})
		} else {
			fmt.Println("strBody", strBody)
		}
	})
}
