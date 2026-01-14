package examples

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func BindAndValidation() {
	router := gin.Default()

	router.POST("/examples/login-json", func(ctx *gin.Context) {
		var json Login
		if err := ctx.ShouldBindJSON(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status": "unauthorized",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status": "you are logged in",
		})
	})
	// start server
	router.Run(":8080")
}
