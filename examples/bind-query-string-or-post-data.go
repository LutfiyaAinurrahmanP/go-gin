package examples

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Person2 struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func BindingQueryStringOrPost() {
	r := gin.Default()
	r.GET("/examples/bind-query-string-or-post-data", startPage)
	r.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person2

	err := c.ShouldBind(&person)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"name":     person.Name,
		"address":  person.Address,
		"birthday": person.Birthday.Format("2006-01-02"),
	})

	log.Println(person.Name)
	log.Println(person.Address)
	log.Println(person.Birthday)

	c.String(200, "Success")
}
