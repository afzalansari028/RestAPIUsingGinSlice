package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

var persons = []Person{
	{Id: "10", Name: "Hardik", Age: 31, Address: "Hyderabad"},
	{Id: "30", Name: "Virat", Age: 38, Address: "Banglore"},
	{Id: "20", Name: "Rohit", Age: 42, Address: "Mumbai"},
}

func AddOnePerson(c *gin.Context) {

	jsonBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("Error while reading body")
	}

	var person Person
	err = json.Unmarshal(jsonBytes, &person)
	if err != nil {
		fmt.Println("Error while reunmarshaling")
	}
	persons = append(persons, person)
	c.JSON(200, "One person added")
}
