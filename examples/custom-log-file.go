package examples

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func CustomLogFile() {
	router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	router.GET("/examples/custom-log-file", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})
	router.Run()
}
