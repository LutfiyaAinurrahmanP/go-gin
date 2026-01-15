package examples

import "github.com/gin-gonic/gin"

func MultipartUrlencodedForm() {
	r := gin.Default()
	r.POST("/examples/multipart-urlencoded-form/form_post", func(ctx *gin.Context) {
		m := ctx.PostForm("message")
		n := ctx.DefaultPostForm("nick", "anonymous")

		ctx.JSON(200, gin.H{
			"status":  "posted",
			"message": m,
			"nick":    n,
		})
	})

	r.Run(":8080")
}
