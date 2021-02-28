package book

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/mfduar8766/GoRestAPI/db"
	"github.com/mfduar8766/GoRestAPI/logger"
	"github.com/mfduar8766/GoRestAPI/utils"
)

// Books - Books data structure
type Books struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
}

// GetBooks - Used to GET all books
func GetBooks(c *fiber.Ctx) error {
	logger.LogInstance.Info("GetBooks()")
	var books []Books
	db.GormInstance.Find(&books)
	return c.JSON(books)
}

// GetBook - Used to GET book by ID
func GetBook(c *fiber.Ctx) error {
	logger.LogInstance.Info("GetBook()")
	id := c.Params("id")
	var book Books
	db.GormInstance.First(&book, id)
	if book.Title == "" {
		return c.Status(500).Send(utils.CreateMessage("Book does not exist"))
	}
	db.GormInstance.Find(&book)
	return c.JSON(book)
}

// AddBook - Used to add new book
func AddBook(c *fiber.Ctx) error {
	logger.LogInstance.Info("AddBook()")
	var book Books
	if err := c.BodyParser(&book); err != nil {
		c.Status(500).Send(utils.CreateMessage("Error parsing incommig request body"))
	}
	db.GormInstance.Create(&book)
	return c.JSON(book)
}

// UpdateBook - Used to update book by ID
func UpdateBook(c *fiber.Ctx) error {
	logger.LogInstance.Info("UpdateBook()")
	id := c.Params("id")
	var book Books
	var newBook Books
	db.GormInstance.First(&book, id)
	if book.Title == "" || book.Author == "" {
		return c.Status(500).Send(utils.CreateMessage("Book does not exist"))
	}
	if err := json.Unmarshal(c.Body(), &newBook); err != nil {
		c.Status(500).Send(utils.CreateMessage("Error parsing incomming data"))
	}
	db.GormInstance.Model(&book).Updates(Books{Title: newBook.Title, Author: newBook.Author})
	return c.JSON(book)
}

// DeleteBook - Used to DELETE book by ID
func DeleteBook(c *fiber.Ctx) error {
	logger.LogInstance.Info("DeleteBook()")
	id := c.Params("id")
	var book Books
	db.GormInstance.First(&book, id)
	if book.Title == "" {
		return c.Status(500).Send(utils.CreateMessage("Book does not exist"))
	}
	db.GormInstance.Delete(&book)
	return c.Status(200).Send(utils.CreateMessage("Successfully deleted book"))
}
