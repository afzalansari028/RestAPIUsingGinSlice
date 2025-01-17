package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func UpdatePerson(c *gin.Context) {
	var newPerson Person
	jsonBytes, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(jsonBytes, &newPerson)

	for i := 0; i < len(persons); i++ {
		if persons[i].Id == newPerson.Id {
			fmt.Println("id matched")
			persons[i].Name = newPerson.Name
			persons[i].Address = newPerson.Address
		}
	}
	c.JSON(200, "Updated successfully")
}
