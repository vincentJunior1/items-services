package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	cModels "items/controllers/models"
	hModels "items/helpers/models"
	"items/model/mapping"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func (c *controllers) Register(ctx *gin.Context) {
	fmt.Println("<<< Login >>>")
	res := hModels.Response{}
	payload := cModels.Register{}

	if err := ctx.BindJSON(&payload); err != nil {
		fmt.Println("Error:", err)
		res.Meta.Code = http.StatusBadRequest
		res.Meta.Message = "Bad Request"

		ctx.JSON(res.Meta.Code, res)
		return
	}

	newPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error:", err)
		res.Meta.Code = http.StatusInternalServerError
		res.Meta.Message = "Server Error"

		ctx.JSON(res.Meta.Code, res)
		return
	}

	payload.Password = string(newPassword)

	data := mapping.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: string(newPassword),
	}

	if err := c.repository.Register(ctx, data); err != nil {
		fmt.Println("Error:", err)
		res.Meta.Code = http.StatusUnprocessableEntity
		res.Meta.Message = "Unprocessable Entity"

		ctx.JSON(res.Meta.Code, res)
		return
	}

	res.Meta.Code = http.StatusOK
	res.Meta.Message = "Success Register"

	ctx.JSON(res.Meta.Code, res)
}

func (c *controllers) Login(ctx *gin.Context) {
	fmt.Println("<<< Login >>>")
	res := hModels.Response{}
	payload := cModels.Login{}

	if err := ctx.BindJSON(&payload); err != nil {
		fmt.Println("Error:", err)
		res.Meta.Code = http.StatusBadRequest
		res.Meta.Message = "Bad Request"

		ctx.JSON(res.Meta.Code, res)
		return
	}

	user, err := c.repository.Login(ctx, payload.Email)

	if err != nil {
		fmt.Println("Error:", err)
		res.Meta.Code = http.StatusUnauthorized
		res.Meta.Message = "Email Not Found"

		ctx.JSON(res.Meta.Code, res)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		fmt.Println("Error:", err)
		res.Meta.Code = http.StatusUnauthorized
		res.Meta.Message = "Wrong Password KONTOL!!!"

		ctx.JSON(res.Meta.Code, res)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": cModels.LoginRes{
			Id:   user.Id,
			Name: user.Name,
		},
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_CODE")))

	if err != nil {
		fmt.Println("Error:", err)
		res.Meta.Code = http.StatusInternalServerError
		res.Meta.Message = "Server Error"

		ctx.JSON(res.Meta.Code, res)
		return
	}

	res.Meta.Code = http.StatusOK
	res.Meta.Message = "Success"
	res.Data = tokenString

	ctx.JSON(res.Meta.Code, res)
}

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

	minPrice := 0
	maxPrice := 0
	resItems := []cModels.Items{}
	for i, val := range data {
		if i == 0 {
			minPrice = val.Price
			maxPrice = val.Price
		}

		if val.Price > maxPrice {
			maxPrice = val.Price
		}

		if val.Price < minPrice {
			minPrice = val.Price
		}

		fmt.Println(val)
		resItems = append(resItems, cModels.Items{
			Name:     val.Name,
			Price:    val.Price,
			Quantity: val.Quantity,
		})
	}

	res.Meta.Code = http.StatusOK
	res.Meta.Message = "Success"
	res.Data = cModels.ResponseItems{
		MinPrice: minPrice,
		MaxPrice: maxPrice,
		Items:    resItems,
	}
	res.Page = &hModels.Pagination{
		Limit:     params.Limit,
		Page:      params.Page,
		TotalData: total,
	}
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
