package manager

import (
	"fmt"
	"go-mongod/config"

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

func (i *infraManager) DbConn() *mongo.Database {
	credential := options.Credential{
		Username: i.cfg.User,
		Password: i.cfg.Password,
	}

	mongoUrl := fmt.Sprintf("mongo://%s:%s", i.cfg.ApiHost, i.cfg.ApiPort)
	clientOptions := options.Client()
	clientOptions.ApplyURI(mongoUrl).SetAuth(credential)
}
