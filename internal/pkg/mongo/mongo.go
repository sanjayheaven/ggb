package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/sanjayheaven/ggb/configs"
	"github.com/sanjayheaven/ggb/internal/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(config *configs.Mongo) {

	logger := logger.LogrusLogger

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	address := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", config.User, config.Password, config.Host, config.Port, config.DataBase)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(address))
	if err != nil {
		panic(`😫: Connected failed, check your MongoDB with ` + address)
	}

	logger.Printf(`🍟: Successfully connected to MongoDB at ` + address)

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
