package examples

import "github.com/gin-gonic/gin"

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func MultipartUrlencodedBinding() {
	r := gin.Default()
	r.POST("/examples/multipart-urlencoded-binding/login", func(ctx *gin.Context) {
		// Anda dapat melakukan bind multipart form dengan deklarasi binding eksplisit:
		// c.ShouldBindWith(&form, binding.Form)
		// atau Anda bisa langsung menggunakan autobinding dengan metode ShouldBind
		var form LoginForm
		// dalam kasus ini, binding yang sesuai akan dipilih secara otomatis
		if ctx.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				ctx.JSON(200, gin.H{
					"status": "you are logged in",
				})
			} else {
				ctx.JSON(401, gin.H{
					"status": "unauthorized",
				})
			}
		}
	})
	r.Run(":8080")
}
