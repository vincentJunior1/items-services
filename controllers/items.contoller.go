package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	cModels "items/controllers/models"
	hModels "items/helpers/models"
	"items/model/mapping"

	"github.com/gin-gonic/gin"
)

func (c *controllers) GetItems(ctx *gin.Context) {
	fmt.Println("<<< Get Items Controllers >>>")
	res := hModels.Response{}
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))

	params := cModels.ParamsGetItems{
		Search: ctx.Query("search"),
		Page:   page,
		Limit:  limit,
	}

	data, total, err := c.repository.GetItems(ctx, params)

	if err != nil {
		fmt.Println("Error get data ", err)
		res.Meta.Code = http.StatusInternalServerError
		res.Meta.Message = "Server Error"

		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res.Meta.Code = http.StatusOK
	res.Meta.Message = "Success"
	res.Data = data
	res.Page = hModels.Pagination{
		Limit:     params.Limit,
		Page:      params.Page,
		TotalData: total,
	}
	fmt.Println("KONTOL")
	ctx.JSON(http.StatusOK, res)
}

func (c *controllers) CreateItems(ctx *gin.Context) {
	fmt.Println("<<<CreateItemsController>>>")
	res := hModels.Response{}
	payload := cModels.ReqGetItems{}

	if err := ctx.BindJSON(&payload); err != nil {
		fmt.Println("Error bind json:", err)
		res.Meta.Code = http.StatusBadRequest
		res.Meta.Message = "Bad Request"
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	item := mapping.Items{
		Name:     payload.Name,
		Price:    payload.Price,
		Quantity: payload.Quantiy,
	}

	if err := c.repository.CreateItems(ctx, item); err != nil {
		fmt.Println("Error Create Items:", err)
		res.Meta.Code = http.StatusUnprocessableEntity
		res.Meta.Message = "Unprocessable Entity"

		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	res.Meta.Code = http.StatusOK
	res.Meta.Message = "Success"

	ctx.JSON(http.StatusOK, res)
}
