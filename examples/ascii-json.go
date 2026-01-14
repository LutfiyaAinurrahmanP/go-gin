package examples

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AsciiJSON() {
	r := gin.Default()

	r.GET("/examples/ascii-json", func(ctx *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		ctx.AsciiJSON(http.StatusOK, data)
	})

	r.Run()
}
