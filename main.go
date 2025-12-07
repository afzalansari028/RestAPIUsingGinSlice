package main

import (
	"RestAPIUsingGinSlice/handler"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	env := os.Getenv("ENV")

	var envFile string
	// get env file
	if env == "local" {
		envFile = fmt.Sprintf("./resources/.env.%s", env)
	} else {
		envFile = fmt.Sprintf("./resources/.env.%s", env)
	}
	// load env file based on env set : for local we set manually, for others that will be set on yml file if kubernetes.
	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Println("Env file not leaded..")
	}

	checkENV := os.Getenv("XYZ")
	fmt.Println("ENVIRONMENT::", checkENV)

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

// set env -->   $env:ENV="local"
// check env --> $env:ENV
