package examples

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Redirect() {
	r := gin.Default()

	// mengarahkan ulang http baik lokasi internal maupun external
	r.GET("/examples/test", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "https://www.youtube.com/")
	})

	// mengarahkan ulang http dari post
	r.POST("/examples/test", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/foo")
	})

	// mengarahkan ulang router
	r.GET("/examples/test2", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/test2"
		r.HandleContext(ctx)
	})
	r.GET("/examples/test2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"hello": "world"})
	})
}
