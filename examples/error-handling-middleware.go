package examples

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		ctx.Next() // Langkah 1: proses permintaan terlebih dahulu

		// Langkah 2: Periksa apakah ada error yang ditabahkan kedalam context
		if len(ctx.Errors) > 0 {
			// Langkah 3: Gunakan error terakhir
			err := ctx.Errors.Last().Err

			// Langkah 4: Tanggapi dengan pesan error generik
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
		}

		// Langkah lain jika ada error yang ditemukan
	}
}

func ErrorHandlingMiddleware()  {
	r := gin.Default()
	// pasang middleware penanganan error
	r.Use(ErrorHandler())

	r.GET("/examples/ok", func(ctx *gin.Context) {
		somethingWentWrong := false

		if somethingWentWrong {
			ctx.Error(errors.New("something went wrong"))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Everything is fine!",
		})
	})

	r.GET("/examples/err", func(ctx *gin.Context) {
		somethingWentWrong := true

		if somethingWentWrong {
			ctx.Error(errors.New("something went wrong"))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Everything is fine!",
		})
	})

	r.Run(":8080")
}