package book

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/mfduar8766/GoRestAPI/db"
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
	fmt.Println("GetBooks()")
	var books []Books
	db.GormInstance.Find(&books)
	return c.JSON(books)
}

// GetBook - Used to GET book by ID
func GetBook(c *fiber.Ctx) error {
	fmt.Println("GetBook()")

	return c.Send([]byte("One Books"))
}

// AddBook - Used to add new book
func AddBook(c *fiber.Ctx) error {
	fmt.Println("AddBook()")
	var book Books
	err := json.Unmarshal(c.Body(), &book)
	utils.MustNotError(err)
	db.GormInstance.Create(&book)
	return c.JSON(book)
}

// UpdateBook - Used to update book by ID
func UpdateBook(c *fiber.Ctx) error {
	fmt.Println("UpdateBook()")
	id := c.Params("id")
	var book Books
	db.GormInstance.First(&book, id)
	if book.Title == "" {
		return c.Status(500).Send([]byte("Book does not exist"))
	}
	err := json.Unmarshal(c.Body(), &book)
	utils.MustNotError(err)
	var newBook = new(Books)
	newBook.Title = book.Title
	newBook.Author = book.Author
	db.GormInstance.Update(newBook)
	return c.JSON(newBook)
}

// DeleteBook - Used to DELETE book by ID
func DeleteBook(c *fiber.Ctx) error {
	fmt.Println("DeleteBook()")
	id := c.Params("id")
	var book Books
	db.GormInstance.First(&book, id)
	if book.Title == "" {
		return c.Status(500).Send([]byte("Book does not exist"))
	}
	db.GormInstance.Delete(&book)
	return c.Status(200).Send([]byte("Successfully deleted book"))
}
