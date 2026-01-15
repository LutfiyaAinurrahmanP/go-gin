package examples

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Person4 struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func OnlyBindQueryString() {
	route := gin.Default()
	route.Any("/examples/only-bind-query-string/testing", startPage2)
	route.Run(":8080")
}

func startPage2(c *gin.Context) {
	var person Person4
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}
