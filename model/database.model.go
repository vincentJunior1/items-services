package model

import (
	"fmt"
	cModels "items/controllers/models"
	"items/model/mapping"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	mysqlDatabase struct {
		DB *gorm.DB
	}

	MysqlDatabase interface {
		Register(ctx *gin.Context, data mapping.User) error
		CreateItems(ctx *gin.Context, data mapping.Items) error
		Login(ctx *gin.Context, email string) (mapping.User, error)
		GetItems(ctx *gin.Context, params cModels.ParamsGetItems) ([]mapping.Items, int64, error)
	}
)

// Inisiasi Database
func InitDatabase() MysqlDatabase {
	fmt.Println("<<< Initialize Database Connection >>>")
	return &mysqlDatabase{
		DB: ConnectionMysql(),
	}
}

// logmode ...

var logMode = map[string]logger.LogLevel{
	"silent": logger.Silent,
	"error":  logger.Error,
	"warn":   logger.Warn,
	"info":   logger.Info,
}

// configurasi mysql / database
func ConnectionMysql() *gorm.DB {
	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")
	debug := os.Getenv("DATABASE_DEBUG_MYSQL")
	mode := os.Getenv("LOG_MODE_MYSQL")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logMode[mode]),
	})

	if err != nil {
		fmt.Println("Error Connect database", err)
		panic("Error Connection Database")
	}

	fmt.Println("Mysql Connected Successfully")

	if debug == "true" {
		return db.Debug()
	}

	return db
}
