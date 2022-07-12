package delivery

import (
	"fmt"
	"go-mongod/config"
	"go-mongod/delivery/controller"
	"go-mongod/manager"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	useCaseManager manager.UseCaseManager
	engine *gin.Engine
	host string
}

func (a *appServer) initHandlers() {
	controller.NewProductController(a.engine, a.useCaseManager.ProductRepo())
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}

func NewServer() *appServer {
	r := gin.Default()
	c := config.NewConfig()
	infraManager := manager.NewInfraManager(c)
	repoManager := manager.NewRepositoryManager(infraManager)
	usecaseManager := manager.NewUseCaseManager(repoManager)
	host := fmt.Sprintf("%s:%s", c.ApiHost, c.ApiPort)
	return &appServer{
		useCaseManager: usecaseManager,
		engine: r,
		host: host,
	}

}
