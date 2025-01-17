package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

func TestAddOnePerson(t *testing.T) { // cmd to run this test func:- go test -run ^TestAddOnePerson$ ./...

	// create mock request
	requestBody := strings.NewReader(`{
	 "id":"10","name":"Hardik","age":31,"address":"Hyderabad"
	}`)

	// create moke server
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/addperson", AddOnePerson)

	//create http request
	req, err := http.NewRequest("POST", "/addperson", requestBody)
	if err != nil {
		fmt.Println("Error while making http call:", err)
	}
	// req.Header.Set("Authorization", "Basic "+"dXNlcm5hbWU6cGFzc3dvcmQ=")

	// record response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	jsonBytes, _ := io.ReadAll(w.Body)
	var result string
	json.Unmarshal(jsonBytes, &result)
	assert.Equal(t, ("One person added"), result)

}
