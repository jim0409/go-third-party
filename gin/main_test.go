package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String())

	var data map[string]interface{}
	wbBytes := w.Body.Bytes()
	err := json.Unmarshal(wbBytes, &data)
	assert.Nil(t, err)
	assert.Equal(t, "pong", data["message"])
}

// test CRUD router
// test POST method
func TestPostRoute(t *testing.T) {

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var data map[string]interface{}
	wbBytes := w.Body.Bytes()
	err := json.Unmarshal(wbBytes, &data)
	assert.Nil(t, err)
	assert.Equal(t, "test_Create_with_POST_method", data["message"])
}

// test POST method
func TestPOSTRoute(t *testing.T) {

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var data map[string]interface{}
	wbBytes := w.Body.Bytes()
	err := json.Unmarshal(wbBytes, &data)
	assert.Nil(t, err)
	assert.Equal(t, "test_Create_with_POST_method", data["message"])
}

// test READ method
func TestGETRoute(t *testing.T) {

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var data map[string]interface{}
	wbBytes := w.Body.Bytes()
	err := json.Unmarshal(wbBytes, &data)
	assert.Nil(t, err)
	assert.Equal(t, "test_Read_with_GET_method", data["message"])
}

// test PUT method
func TestPUTRoute(t *testing.T) {

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var data map[string]interface{}
	wbBytes := w.Body.Bytes()
	err := json.Unmarshal(wbBytes, &data)
	assert.Nil(t, err)
	assert.Equal(t, "test_Update_with_PUT_method", data["message"])
}

// test DELETE method
func TestDELETERoute(t *testing.T) {

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var data map[string]interface{}
	wbBytes := w.Body.Bytes()
	err := json.Unmarshal(wbBytes, &data)
	assert.Nil(t, err)
	assert.Equal(t, "test_Delete_with_DELETE_method", data["message"])
}
