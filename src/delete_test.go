package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type deleteTest struct {
	route      string
	statusCode int
	expected   string
}

var deleteTests = []deleteTest{
	deleteTest{"/key-value-store/test", 200, "true"},
}

// inspired by medium post here: https://medium.com/@craigchilds94/testing-gin-json-responses-1f258ce3b0b1
func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestDelete(t *testing.T) {
	router := gin.Default()
	for _, test := range deleteTests {
		w := performRequest(router, "DELETE", test.route)
		assert.Equal(t, test.statusCode, w.Code)
		var response map[string]string
		err := json.Unmarshal([]byte(w.Body.String()), &response)
		println(err)
		// value, exists := response["doesExist"]
		// assert.Nil(t, err)
		// assert.True(t, exists)
		// assert.Equal(t, test.expected, value)
	}
}
