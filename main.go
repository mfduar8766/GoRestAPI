package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mfduar8766/GoRestAPI/book"
	"github.com/mfduar8766/GoRestAPI/config"
	"github.com/mfduar8766/GoRestAPI/db"
	"github.com/mfduar8766/GoRestAPI/utils"
	"github.com/mfduar8766/GoRestAPI/logger"
)

func setUpRoutes(app *fiber.App) {
	logger.LogInstance.Info("setUpRoutes()")
	app.Get("/api/v1/books", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.AddBook)
	app.Put("/api/v1/book/:id", book.UpdateBook)
	app.Delete("api/v1/book/:id", book.DeleteBook)
}

func connectToDb() {
	logger.LogInstance.Info("connectToDb()")
	dbConfig := config.InitDbConfig()
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName)
	var err error
	db.GormInstance, err = gorm.Open("postgres", connectionString)
	utils.MustNotError(err)
	dataBase := db.GormInstance.DB()
	err = dataBase.Ping()
	utils.MustNotError(err)
	logger.LogInstance.Info("Successfully connected to DB")
	db.GormInstance.AutoMigrate(&book.Books{})
	logger.LogInstance.Info("Successfully migrated data to DB")
}

func main() {
	logger.CreateLogger("logger", "logs.txt")
	logger.LogInstance.Info("Init App Running Main()")
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(404).SendString(err.Error())
		},
	})
	connectToDb()
	defer db.GormInstance.Close()
	setUpRoutes(app)
	app.Listen(":3000")
}
