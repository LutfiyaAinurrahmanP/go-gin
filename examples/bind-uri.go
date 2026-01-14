package examples

import "github.com/gin-gonic/gin"

type Person3 struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func BindingUri() {
	r := gin.Default()
	r.GET("/examples/binding-uri/:name/:id", func(ctx *gin.Context) {
		var person Person3
		if err := ctx.ShouldBindUri(&person); err != nil {
			ctx.JSON(400, gin.H{
				"msg": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"name": person.Name,
			"uuid": person.ID,
		})
	})
	r.Run(":8080")
}
