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
)

func setUpRoutes(app *fiber.App) {
	fmt.Println("setUpRoutes()")
	app.Get("/api/v1/books", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.AddBook)
	app.Put("/api/v1/book/:id", book.UpdateBook)
	app.Delete("api/v1/book/:id", book.DeleteBook)
}

func connectToDb() {
	fmt.Println("connectToDb()")
	dbConfig := config.InitDbConfig()
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName)
	var err error
	db.GormInstance, err = gorm.Open("postgres", connectionString)
	utils.MustNotError(err)
	dataBase := db.GormInstance.DB()
	err = dataBase.Ping()
	utils.MustNotError(err)
	fmt.Println("Successfully connected to DB")
	db.GormInstance.AutoMigrate(&book.Books{})
	fmt.Println("Successfully migrated data to DB")
}

func main() {
	fmt.Println("main()")
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			fmt.Printf("Error: %+v", err.Error())
			return c.Status(404).SendString("hi, i'm an custom error")
		},
	})
	connectToDb()
	defer db.GormInstance.Close()
	setUpRoutes(app)
	app.Listen(":3000")
}
