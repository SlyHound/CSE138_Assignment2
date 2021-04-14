package utility

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JSON struct {
	Error    string `json:"error"`
	Response string `json:"message"`
}

func ForwardRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("In ForwardRequest call")
		http.HandleFunc("/key-value-store/{key}", func(w http.ResponseWriter, r *http.Request) {
			fwdRequest, err := http.NewRequest(r.Method, r.URL.String(), r.Body)
			fmt.Println("fwdRequest method(0):", fwdRequest.Method)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			fwdRequest.Header.Set("X-Forwarding-For", r.RemoteAddr)
			fmt.Println("fwdRequest method:", fwdRequest.Method)

			httpForwarder := &http.Client{}
			fwdResponse, err := httpForwarder.Do(fwdRequest)
			fmt.Println("fwdResponse method:", fwdResponse.Request.Method)
			if err != nil {
				w.WriteHeader(http.StatusServiceUnavailable) // sends a 503 status code
				response := JSON{
					Error:    "Main instance is Down",
					Response: "Error in " + fwdResponse.Request.Method,
				}
				json.NewEncoder(w).Encode(response)
			}

			fwdResponse.Body.Close()
		})
	}
}
