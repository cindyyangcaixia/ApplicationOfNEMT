package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/cindyyangcaixia/ApplicationOfNEMT/docs"
	"github.com/cindyyangcaixia/ApplicationOfNEMT/routers"

	"github.com/cindyyangcaixia/ApplicationOfNEMT/models"
	"github.com/cindyyangcaixia/ApplicationOfNEMT/pkg/setting"
)

func init() {
	setting.Setup()
	models.Setup()

}

//	@title			Application of National College Entrance Examination
//	@version		1.0
//	@description	This application is for college candidates to fill in their volunteers as a reference.

//	@contact.name	API Support
//	@contact.email	cindy.yang.caixia@gmail.com

//	@host		localhost:8000
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
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
