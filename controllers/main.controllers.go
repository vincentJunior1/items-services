package controllers

import (
	"fmt"
	"items/model"

	"github.com/gin-gonic/gin"
)

type (
	controllers struct {
		repository model.MysqlDatabase
	}

	Controllers interface {
		GetItems(ctx *gin.Context)
		CreateItems(ctx *gin.Context)
	}
)

func InitControllers(db model.MysqlDatabase) Controllers {
	fmt.Println("<<< Init Controller >>>")
	return &controllers{
		repository: db,
	}
}
