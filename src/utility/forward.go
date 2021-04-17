package utility

import (
	"fmt"
	"net/http"

	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type JSON struct {
	Error    string `json:"error"`
	Response string `json:"message"`
}

func ForwardRequest(r *gin.Engine) {
	fmt.Println("In ForwardRequest call")
	r.GET("/key-value-store/:key", handleRequests())
	r.PUT("/key-value-store/:key", handleRequests())
	r.DELETE("/key-value-store/:key", handleRequests())
}

func handleRequests() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.URL.Host = "10.10.0.2:8085"
		c.Request.URL.Scheme = "http"

		fwdRequest, err := http.NewRequest(c.Request.Method, c.Request.URL.String(), c.Request.Body)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}

		fwdRequest.Header = c.Request.Header

		httpForwarder := &http.Client{}
		fwdResponse, err := httpForwarder.Do(fwdRequest)

		if err != nil {
			msg := "Error in " + fwdRequest.Method
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Main instance is down", "message": msg})
		}
		if fwdResponse != nil {
			body, _ := ioutil.ReadAll(fwdResponse.Body)
			rawJSON := json.RawMessage(body)
			c.JSON(fwdResponse.StatusCode, rawJSON)
			defer fwdResponse.Body.Close()
		}
	}
}
