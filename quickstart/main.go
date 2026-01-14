package quickstart

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Example() {
	router := gin.Default()
	router.GET("/quickstart/example", func(ctx *gin.Context) {
		// H = untuk shourtcut maping data
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}

func ExampleWithHTTP() {
	router := gin.Default()

	router.GET("/quickstart/example-with-http", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
