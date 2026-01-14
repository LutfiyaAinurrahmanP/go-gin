package examples

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func GoroutinesInsideMiddleware() {
	r := gin.Default()

	r.GET("/examples/long_async", func(ctx *gin.Context) {
		// buat salinan untuk digunakan dalam goroutine
		// Saat memulai Goroutine baru di dalam sebuah middleware atau handler, Anda TIDAK BOLEH menggunakan context asli di dalamnya, Anda harus menggunakan salinan hanya-baca.
		cCp := ctx.Copy()
		go func() {
			time.Sleep(5 * time.Second)

			// perhatikan bahwa anda menggunakan konteks yang disalin "cCp", INI PENTING
			log.Panicln("Done, in path " + cCp.Request.URL.Path)
		}()
	})

	r.GET("/examples/long_sync", func(ctx *gin.Context) {
		time.Sleep(5 * time.Second)

		log.Panicln("Done, in path " + ctx.Request.URL.Path)
	})

	r.Run(":8080")
}
