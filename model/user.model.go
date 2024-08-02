package model

import (
	"items/model/mapping"

	"github.com/gin-gonic/gin"
)

func (db *mysqlDatabase) Register(ctx *gin.Context, data mapping.User) error {
	query := db.DB.Model(&data)
	query.Create(&data)

	return query.Error
}

func (db *mysqlDatabase) Login(ctx *gin.Context, email string) (mapping.User, error) {
	var data mapping.User
	query := db.DB.Model(&data)
	query = query.Where("email = ?", email)
	query.First(&data)

	return data, query.Error
}
