package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

func TestGetAllPersons(t *testing.T) {

	// create moke server
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/getallpersons", GetAllPersons)

	//create http request
	req, err := http.NewRequest(http.MethodGet, "/getallpersons", nil)
	if err != nil {
		fmt.Println("Error while making http call:", err)
	}
	// req.Header.Set("Authorization", "Basic "+"dXNlcm5hbWU6cGFzc3dvcmQ=")

	// record response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}
func TestGetOnePerson(t *testing.T) {

	// create moke server
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/getoneperson/:id", GetOnePerson)

	// create http request
	req, err := http.NewRequest("GET", "/getoneperson/10", nil)
	if err != nil {
		fmt.Println("Error while making http call:", err)
	}

	// record response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
