package handler

import "github.com/gin-gonic/gin"

func GetAllPersons(c *gin.Context) {
	c.JSON(200, persons)
}

func GetOnePerson(c *gin.Context) {

	id := c.Param("id")
	var person Person
	for i := 0; i < len(persons); i++ {
		if id == persons[i].Id {
			person = persons[i]
		}
	}
	c.JSON(200, person)
}
