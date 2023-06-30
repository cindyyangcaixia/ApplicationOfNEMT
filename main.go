package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cindyyangcaixia/gin-example/models"
	"github.com/cindyyangcaixia/gin-example/pkg/setting"
	"github.com/cindyyangcaixia/gin-example/routers"
)

func init() {
	setting.Setup()
	models.Setup()

}

func main() {
	routersInit := routers.InitRouter()

	var endPoint = fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := &http.Server{
		Addr:    endPoint,
		Handler: routersInit,
		// ReadTimeout:    setting.ServerSetting.ReadTimeout,
		// WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	log.Printf("[info] start http server listening %s", endPoint)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shuting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown:", err)
	}
	log.Println("Server exiting")
}
