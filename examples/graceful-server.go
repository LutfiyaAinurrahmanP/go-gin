package examples

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func GracefulServer() {
	r := gin.Default()
	r.GET("/examples/graceful-server", func(ctx *gin.Context) {
		time.Sleep(5 * time.Second)
		ctx.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r.Handler(),
	}

	go func() {
		// Melayani koneksi
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Tunggu sinyal interrupt untuk mematikan server dengan anggun
	// dengan batas waktu 5 detik
	quit := make(chan os.Signal, 1)
	// Kill (tanpa parameter) secara bawaan mengirimkan syscall.SIGTERM
	// Kill -2 adalah syscall.SIGINT
	// Kill -9 adalah syscall.SIGKILL tetapi tidak dapat ditangkap. jadi tidak perlu menambahkannya
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
