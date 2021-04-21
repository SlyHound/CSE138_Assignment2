package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type apiTest struct {
	route      string
	statusCode int
	method     string
	expected   bool
	value      string
	body       map[string]string
}

//Run after
var deleteTests = []apiTest{
	{"/key-value-store/test", 404, "DELETE", false, "", nil}, //Delete value that doesn't exist
	{"/key-value-store/test2", 200, "DELETE", true, "", nil}, //Delete value that does exist
	//Delete value that does not exist on main instance as forwarding instance
	//Delete value that exists on main instance as forwarding instance
	//Delete value when main instance is down and forwarding instance is up
}

var getTests = []apiTest{
	{"/key-value-store/test", 404, "GET", false, "", nil}, //Get Value that doesn't exist
	{"/key-value-store/test2", 200, "GET", true, "", nil}, //Get Value that does exist

}

var putTests = []apiTest{
	{"key-value-store/test2", 201, "PUT", false, "", map[string]string{"value": "distributed systems"}}, //Put a new value to server
	{"key-value-store/test2", 200, "PUT", true, "", map[string]string{"value": "single system"}},        //Put value that already exists
	//Put request with no value
	//Put request to non-existent key with >50 chars
	//Put request

}

var genTests = []apiTest{
	//malformed request
	//empty request
	//server down
}

// inspired by medium post here: https://medium.com/@craigchilds94/testing-gin-json-responses-1f258ce3b0b1
func performRequest(r http.Handler, method, path string, reqBody map[string]string) *httptest.ResponseRecorder {
	var req *http.Request
	if reqBody == nil {
		req, _ = http.NewRequest(method, path, nil)
	} else {
		body, err := json.Marshal(reqBody)
		println(string(body))
		if err != nil {
			log.Fatal(err)
		} else {
			req, _ = http.NewRequest(method, path, bytes.NewBuffer(body))
			println(req.Body)
		}

	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestPut(t *testing.T) {
	var kvStore = make(map[string]string)
	router := setupRouter(kvStore)
	for _, test := range putTests {
		w := performRequest(router, test.method, test.route, test.body)
		assert.Equal(t, test.statusCode, w.Code)
		var response map[string]interface{}
		err := json.Unmarshal([]byte(w.Body.String()), &response)
		value, replaced := response["replaced"]
		assert.Nil(t, err)
		if replaced {
			assert.Equal(t, test.expected, value)
		} else {
			err, _ := response["error"]
			t.Errorf("%s", err)
		}

	}
}

func TestGetFound(t *testing.T) {
	test := apiTest{"/key-value-store/test2", 200, "GET", true, "", nil}
	//our mock server setup here
	w := performRequest(router, test.method, test.route, nil)
	assert.Equal(t, test.statusCode, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	value, exists := response["doesExist"]
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, test.expected, value)
}

func TestDelete(t *testing.T) {
	//our mock server setup here
	var kvStore = make(map[string]string)
	router := setupRouter(kvStore)
	for _, test := range deleteTests {
		w := performRequest(router, test.method, test.route, nil)
		assert.Equal(t, test.statusCode, w.Code)
		var response map[string]interface{}
		err := json.Unmarshal([]byte(w.Body.String()), &response)
		value, exists := response["doesExist"]
		assert.Nil(t, err)
		assert.True(t, exists)
		assert.Equal(t, test.expected, value)
	}
}
