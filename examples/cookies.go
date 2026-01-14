package examples

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Cookies() {
	router := gin.Default()

	router.GET("/examples/cookies", func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			ctx.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}

		fmt.Printf("Cookie value: %s \n", cookie)
	})

	router.Run()
}

func CookiesHttp()  {
	r := gin.Default()
  r.GET("/set-cookie", func(c *gin.Context) {
    c.SetCookieData(&http.Cookie{
      Name:   "session_id",
      Value:  "abc123",
      Path:   "/",
      Domain:   "localhost",
      Expires:  time.Now().Add(24 * time.Hour),
      MaxAge:   86400,
      Secure:   true,
      HttpOnly: true,
      SameSite: http.SameSiteLaxMode,
      // Partitioned: true, // Go 1.22+
    })
    c.String(http.StatusOK, "ok")
  })
  r.Run(":8080")
}