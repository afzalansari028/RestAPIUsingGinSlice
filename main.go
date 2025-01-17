package main

import (
	"RestAPIUsingGinSlice/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()

	myAppRouter := router.Group("/my-app")

	secret := map[string]string{
		"username": "password",
	}
	// enable basic auth security
	myAppRouter.Use(gin.BasicAuth(secret))

	myAppRouter.POST("/addperson", handler.AddOnePerson)
	myAppRouter.GET("/getallpersons", handler.GetAllPersons)
	myAppRouter.GET("/getoneperson/:id", handler.GetOnePerson)
	myAppRouter.DELETE("/deleteperson/:id", handler.DeleteOnePerson)
	myAppRouter.PUT("/updateperson", handler.UpdatePerson)

	router.Run(":8080")

}
