package model

import (
	cModels "items/controllers/models"
	"items/model/mapping"

	"github.com/gin-gonic/gin"
)

func (db *mysqlDatabase) GetItems(ctx *gin.Context, params cModels.ParamsGetItems) ([]mapping.Items, int64, error) {
	var data []mapping.Items
	var total int64
	query := db.DB.Model(&data)

	if params.Search != "" {
		query = query.Where("name like ?", "%"+params.Search+"%")
	}
	query.Count(&total)
	offset := (params.Page - 1) * params.Limit
	query = query.Limit(params.Limit).Offset(offset)
	query.Find(&data)

	return data, total, query.Error
}

func (db *mysqlDatabase) CreateItems(ctx *gin.Context, data mapping.Items) error {
	query := db.DB.Model(&data)
	query.Create(&data)

	return query.Error
}

func (db *mysqlDatabase) Delete() {}
