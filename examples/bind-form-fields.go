package examples

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name      string    `form:"name,default=William"`
	Age       int       `form:"age,default=10"`
	Friends   []string  `form:"friends,default=Will;Bill"` // multi/csv: gunakan ; dalam nilai default
	Addresses [2]string `form:"addresses,default=foo bar" collection_format:"ssv"`
	LapTimes  []int     `form:"lap_times,default=1;2;3" collection_format:"csv"`
}

func BindFromFields() {
	r := gin.Default()
	r.POST("/examples/bind-from-fields", func(ctx *gin.Context) {
		var req Person
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, req)
	})
	_ = r.Run(":8000")
}
