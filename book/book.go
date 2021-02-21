package book

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

// Books - Books data structure
type Books struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// GetBooks - Used to GET all books
func GetBooks(c *fiber.Ctx) error {
	fmt.Println("GetBooks()")
	return c.Send([]byte("All Books"))
}

// GetBook - Used to GET book by ID
func GetBook(c *fiber.Ctx) error {
	fmt.Println("GetBook()")
	return c.Send([]byte("One Books"))
}

// AddBook - Used to add new book
func AddBook(c *fiber.Ctx) error {
	fmt.Println("AddBook()")
	return c.Send([]byte("Add Books"))
}

// UpdateBook - Used to update book by ID
func UpdateBook(c *fiber.Ctx) error {
	fmt.Println("UpdateBook()")
	return c.Send([]byte("Update Books"))
}

// DeleteBook - Used to DELETE book by ID
func DeleteBook(c *fiber.Ctx) error {
	fmt.Println("DeleteBook()")
	return c.Send([]byte("Delete Books"))
}
