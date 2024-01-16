package cmd

import (
	config "go-gin-boilerplate/configs"
	"go-gin-boilerplate/internal/pkg/logger"
	"go-gin-boilerplate/internal/pkg/mongo"
	"go-gin-boilerplate/internal/pkg/mysql"
	"go-gin-boilerplate/internal/pkg/redis"
	routes "go-gin-boilerplate/internal/router"

	"net/http"

	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var (
	// host           string
	// port           uint
	ServerStartCmd = &cobra.Command{
		Use:   "server",
		Short: `Start the server`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			// log.Println("Hello World!")
			start()
		},
	}
)

func init() {

	// ServerStartCmd.Flags().StringVarP(&host, "host", "H", "127.0.0.1", "HTTP server host") // server host
	// ServerStartCmd.Flags().UintVarP(&port, "port", "p", 8080, "HTTP server port")          // server port
	rootCmd.AddCommand(ServerStartCmd) // add server start command
}

func start() {

	// init routes
	r, err := routes.Init()
	if err != nil {
		panic(err)
	}

	// init logger
	logger.Init()
	logger := logger.Logger

	// load env config
	config.LoadConfig()
	EnvConfig := config.EnvConfig

	// connect database
	mysql.Connect(&EnvConfig.Mysql)
	// connect mongo
	mongo.Connect(&EnvConfig.Mongo)
	// connect redis
	redis.Connect(&EnvConfig.Redis)

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
