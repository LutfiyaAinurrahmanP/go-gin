package examples

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HttpExample1() {
	router := gin.Default()
	http.ListenAndServe(":8080", router)
}

func HttpExample2(){
	router := gin.Default()

  s := &http.Server{
    Addr:           ":8080",
    Handler:        router,
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }
  s.ListenAndServe()
}