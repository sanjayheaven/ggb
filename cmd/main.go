package main

import (
	"go-gin-boilerplate/api/swagger"
	"go-gin-boilerplate/configs"
	"go-gin-boilerplate/internal/models"
	Logger "go-gin-boilerplate/internal/pkg"
	routes "go-gin-boilerplate/internal/router"

	"net/http"

	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	// init routes
	r, err := routes.Init()
	if err != nil {
		panic(err)
	}
	r.Use(gin.Recovery())

	// init logger
	Logger.Init()
	logger := Logger.Logger

	// load env config
	EnvConfig := configs.LoadConfig()

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

	logger.Printf("👻 Server is now listening at port:  %s\n", EnvConfig.Server.Port)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("server listen error: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	i := <-quit
	logger.Println("server receive a signal: ", i.String())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("server shutdown error: %s\n", err)
	}
	logger.Println("Server exiting")

}
