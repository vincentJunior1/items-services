package routes

import (
	"fmt"
	ctrl "items/controllers"

	"items/helpers"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type (
	Router struct {
		Log         *logrus.Logger
		Controllers ctrl.Controllers
		Gin         *gin.Engine
	}
	RouterInterface interface {
		StartServer() error
	}
)

func InitRoutes(ctrl ctrl.Controllers, log *logrus.Logger) RouterInterface {
	return &Router{
		Controllers: ctrl,
		Gin:         gin.New(),
		Log:         log,
	}
}

func (r *Router) StartServer() error {
	fmt.Println("Initialize Router")

	// Bikin variable yang menyimpan engine dari si Router
	items := r.Gin.Group("/items")
	items.POST("/", r.Controllers.CreateItems)
	items.GET("/", r.Controllers.GetItems)

	if err := helpers.StartGinServer(r.Gin); err != nil {
		r.Log.Println("Start Server Err")
		return err
	}

	return nil
}
