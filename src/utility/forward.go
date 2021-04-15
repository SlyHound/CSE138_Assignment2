package utility

import (
	"fmt"
	"net/http"

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
	// fwdRequest, err := http.NewRequest(r.Method, r.URL.String(), r.Body)
	// fmt.Println("fwdRequest method(0):", fwdRequest.Method)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// fwdRequest.Header.Set("X-Forwarding-For", r.RemoteAddr)
	// fmt.Println("fwdRequest method:", fwdRequest.Method)

	// httpForwarder := &http.Client{}
	// fwdResponse, err := httpForwarder.Do(fwdRequest)
	// fmt.Println("fwdResponse method:", fwdResponse.Request.Method)
	// if err != nil {
	// 	w.WriteHeader(http.StatusServiceUnavailable) // sends a 503 status code
	// 	response := JSON{
	// 		Error:    "Main instance is Down",
	// 		Response: "Error in " + fwdResponse.Request.Method,
	// 	}
	// 	json.NewEncoder(w).Encode(response)
	// }

	// fwdResponse.Body.Close()
}

func handleRequests() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.URL.Host = "10.10.0.2:8085"
		c.Request.URL.Scheme = "http"

		fmt.Println("method & url:", c.Request.Method, c.Request.URL.String())
		fwdRequest, err := http.NewRequest(c.Request.Method, c.Request.URL.String(), c.Request.Body)
		fmt.Println("finished getting fwdRqst")
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}

		fwdRequest.Header = c.Request.Header
		// fwdRequest.Header.Set("Host", "127.0.0.1:8086")
		// for debugging purposes //
		fmt.Println(fwdRequest.Header)

		httpForwarder := &http.Client{}
		fwdResponse, err := httpForwarder.Do(fwdRequest)
		fmt.Println("finished getting fwdResponse")
		if err != nil {
			msg := "Error in " + fwdRequest.Method
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Main instance is down", "message": msg})
		}
		if fwdResponse != nil {
			c.JSON(fwdResponse.StatusCode, fwdResponse.Body)
			defer fwdResponse.Body.Close()
		}
	}
}
