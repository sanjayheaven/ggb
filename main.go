package main

import (
	"go-gin-boilerplate/config"
	"go-gin-boilerplate/routes"
	"net/http"

	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// load env config
	EnvConfig := config.LoadConfig()

	// init routes
	r, err := routes.Init()

	if err != nil {
		panic(err)
	}
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
