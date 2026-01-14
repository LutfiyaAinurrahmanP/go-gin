package examples

import "github.com/gin-gonic/gin"

func GroupingRoute() {
	r := gin.Default()

	v1 := r.Group("/examples")
	v1.POST("/login", func(ctx *gin.Context) {})
	v1.POST("/submit", func(ctx *gin.Context) {})
	v1.POST("/read", func(ctx *gin.Context) {})

	r.Run("8080")
}
