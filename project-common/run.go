package common

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(r *gin.Engine, srvName string, addr string, stop func()) {
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	// 优雅启停
	go func() {
		log.Printf("%s running in %s \n", srvName, srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("Shutting Down project %s...\n", srvName)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if stop != nil {
		stop()
	}

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("%s Shutdown,cause by : %v", srvName, err)
	}
	select {
	case <-ctx.Done():
		log.Println("wait timeout...")

	}
	log.Printf("%s stop success...", srvName)
}
