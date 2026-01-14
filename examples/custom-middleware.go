package examples

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()

		// set example variable
		ctx.Set("example", "12345")

		// before request
		fmt.Println("Sebelum handler")
		ctx.Next()

		// after request
		fmt.Println("Sesudah handler")
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := ctx.Writer.Status()
		log.Println(status)
	}
}

func Middleware() {
	r := gin.New()
	r.Use(Logger()) //middleware digunakan disini

	r.GET("/examples/custom-middleware", func(ctx *gin.Context) {
		example := ctx.MustGet("example").(string)

		// it would print : "12345"
		log.Println(example)
	})
	r.Run(":8080")
}
