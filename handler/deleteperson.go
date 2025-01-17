package handler

import "github.com/gin-gonic/gin"

func DeleteOnePerson(c *gin.Context) {

	id := c.Param("id")

	for i := 0; i < len(persons); i++ {
		if id == persons[i].Id {
			persons = append(persons[:i], persons[i+1:]...)
		}
	}
	c.JSON(200, "One person deleted")
}
