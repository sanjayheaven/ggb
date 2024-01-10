package main

import (
	"go-gin-boilerplate/api/swagger"
	"go-gin-boilerplate/configs"
	"go-gin-boilerplate/internal/models"
	"go-gin-boilerplate/internal/router"
	"net/http"

	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	// load env config
	EnvConfig := configs.LoadConfig()

	// init routes
	r, err := routes.Init()
	if err != nil {
		panic(err)
	}

	// init swagger
	swagger.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// connect database
	models.Connect(EnvConfig.Mysql.Dsn)

	// graceful shutdown
	server := &http.Server{
		Addr:    EnvConfig.Server.Port,
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server listen error: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	i := <-quit
	log.Println("server receive a signal: ", i.String())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown error: %s\n", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}

	log.Printf("👻 Server is now listening at port:  %s\n", EnvConfig.Server.Port)
}
