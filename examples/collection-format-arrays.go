package examples

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Filters struct {
	Tags   []string `form:"tags" collection_format:"csv"`     // /search?tags=go,web,api
	Labels []string `form:"labels" collection_format:"multi"` // /search?labels=bug&labels=helpwanted
	IdsSSV []int    `form:"ids_ssv" collection_format:"ssv"`  // /search?ids_ssv=1 2 3
	IdsTSV []int    `form:"ids_tsv" collection_format:"tsv"`  // /search?ids_tsv=1\t2\t3
	Levels []int    `form:"levels" collection_format:"pipes"` // /search?levels=1|2|3
}

func CollectionFormatArrays() {
	router := gin.Default()
	router.GET("/examples/search", func(ctx *gin.Context) {
		var f Filters
		if err := ctx.ShouldBind(&f); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, f)
	})
	router.Run(":8080")
}

// http://localhost:8080/examples/search?tags=go,gin,api&labels=backend&labels=golang&labels=rest&ids_ssv=1%202%203%204&ids_isv=5%096%097&levels=10|20|30
