package main

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

//Get Tests
func TestGetWithValue(t *testing.T) {
	kvStore := make(map[string]string)
	kvStore["course1"] = "Data Structures"
	apitest.New().
		Handler(setupRouter(kvStore)).Get("/key-value-store/course1").Expect(t).Body(`{"doesExist":true, "message":"Retrieved successfully", "value":"Data Structures"}`).Status(http.StatusOK).End()
}

func TestGetWithoutValue(t *testing.T) {
	kvStore := make(map[string]string)
	apitest.New().
		Handler(setupRouter(kvStore)).Get("/key-value-store/course1").Expect(t).Body(`{"doesExist":false, "error":"Key does not exist", "message":"Error in GET"}`).Status(http.StatusNotFound).End()
}

//Put Tests
func TestPutWithNoKV(t *testing.T) {
	kvStore := make(map[string]string)
	apitest.New().
		Handler(setupRouter(kvStore)).Put("/key-value-store/course1").Body(`{"value": "Distributed Systems"}`).Expect(t).Body(`{"message":"Added successfully","replaced":false}`).Status(http.StatusCreated).End()
}

func TestPutWithKV(t *testing.T) {
	kvStore := make(map[string]string)
	kvStore["course1"] = "Single System"
	apitest.New().
		Handler(setupRouter(kvStore)).Put("/key-value-store/course1").Body(`{"value": "Distributed Systems"}`).Expect(t).Body(`{"message":"Updated successfully","replaced":true}`).Status(http.StatusOK).End()
}

func TestPutWithNoKey(t *testing.T) {
	kvStore := make(map[string]string)
	apitest.New().
		Handler(setupRouter(kvStore)).Put("/key-value-store/course1").Body(`{}`).Expect(t).Body(`{"error":"Value is missing","message":"Error in PUT"}`).Status(http.StatusBadRequest).End()
}

func TestPutWithKeyGT50Chars(t *testing.T) {
	kvStore := make(map[string]string)
	kvStore["course1"] = "Single System"
	apitest.New().
		Handler(setupRouter(kvStore)).Put("/key-value-store/Dasldiufhalsidhuflaisdhgflasdhfgalsdkjfhgasljdhfglasdhfgkasjdhfgaskjdfhgasljdhfgasldjhfgsadljfhgasldjfhg").Body(`{"value": "haha"}`).Expect(t).Body(`{"error":"Key is too long","message":"Error in PUT"}`).Status(http.StatusBadRequest).End()
}

//DELETE Requests
func TestDeleteKeyExist(t *testing.T) {
	kvStore := make(map[string]string)
	kvStore["course1"] = "Distributed Systems"
	apitest.New().
		Handler(setupRouter(kvStore)).Delete("/key-value-store/course1").Expect(t).Body(`{"doesExist":true,"message":"Deleted successfully"}`).Status(http.StatusOK).End()
}

func TestDeleteKeyNoExist(t *testing.T) {
	kvStore := make(map[string]string)
	apitest.New().
		Handler(setupRouter(kvStore)).Delete("/key-value-store/course1").Expect(t).Body(`{"doesExist":false,"error":"Key does not exist","message":"Error in DELETE"}`).Status(http.StatusNotFound).End()
}
