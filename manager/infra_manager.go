package manager

import (
	"context"
	"fmt"
	"go-mongod/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InfraManager interface {
	DbConn() *mongo.Database
}

type infraManager struct {
	db  *mongo.Database
	cfg config.Config
}

func (i *infraManager) initDb() {
	credential := options.Credential{
		Username: i.cfg.User,
		Password: i.cfg.Password,
	}

	mongoUrl := fmt.Sprintf("mongodb://%s:%s", i.cfg.Host, i.cfg.Port)
	clientOptions := options.Client()
	clientOptions.ApplyURI(mongoUrl).SetAuth(credential)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	i.db = client.Database(i.cfg.DbName)
	// defer func() {
	// 	if err := connect.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()
}

func (i *infraManager) DbConn() *mongo.Database {
	return i.db
}



func NewInfraManager(cfg config.Config) InfraManager {
	infra := infraManager{cfg: cfg}
	infra.initDb()
	return &infra
}