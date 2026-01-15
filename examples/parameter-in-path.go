package examples

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ParametersInPath() {
	r := gin.Default()

	// Handler ini akan cocok dengan /user/john tetapi tidak akan cocok dengan /user/ atau /user
	r.GET("examples/parameter/user/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "Hello %s", name)
	})

	// Namun, yang ini akan cocok dengan /user/john/ dan juga /user/john/send
	// Jika tidak ada router lain yang cocok dengan /user/john, maka akan diarahkan ke /user/john/
	r.GET("/user/:name/*action", func(ctx *gin.Context) {
		name := ctx.Param("name")
		action := ctx.Param("action")
		message := name + " is " + action
		ctx.String(http.StatusOK, message)
	})

	r.Run(":8080")
}
