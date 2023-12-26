package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/t01gyl0p/zolo-assignment-server/types"
	"gorm.io/gorm"
)

type BookController struct {
	DB *gorm.DB
}

func (c *BookController) AddBook(ctx *gin.Context) {
	var body types.Book
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var book = types.Book{
		Name: body.Name,
	}

	if err := c.DB.Create(&book).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to add the book"})
		return
	}

	ctx.JSON(201, book)
}

func (c *BookController) BrowseBooks(ctx *gin.Context) {
	var books []types.Book
	if err := c.DB.Find(&books).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to fetch books"})
		return
	}

	ctx.JSON(200, books)
}

func (c *BookController) ListBorrows(ctx *gin.Context) {
	var borrows []types.Borrow
	if err := c.DB.Find(&borrows).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to fetch borrows"})
		return
	}

	ctx.JSON(200, borrows)
}

func (c *BookController) BorrowBook(ctx *gin.Context) {
	var body types.BorrowBookRequestBody
	var bookID = ctx.Param("book_id")
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var book types.Book
	if err := c.DB.First(&book, bookID).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Book not found"})
		return
	}

	borrow := types.Borrow{
		BookID:          book.ID,
		BorrowStartTime: body.BorrowStartTime,
		BorrowEndTime:   body.BorrowEndTime,
	}

	if err := c.DB.Create(&borrow).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to borrow the book"})
		return
	}

	ctx.JSON(201, borrow)
}

func (c *BookController) DeleteBook(ctx *gin.Context) {
	bookId := ctx.Param("book_id")
	var book types.Book
	if err := c.DB.Delete(&book, bookId).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete the book"})
		return
	}
	ctx.Status(204)
}

func (c *BookController) ReturnBook(ctx *gin.Context) {
	bookID := ctx.Param("book_id")
	borrowID := ctx.Param("borrow_id")

	var book types.Book
	if err := c.DB.First(&book, bookID).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Book not found"})
		return
	}

	var borrow types.Borrow
	if err := c.DB.First(&borrow, borrowID).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Borrow record not found"})
		return
	}

	if borrow.Returned {
		ctx.JSON(400, gin.H{"error": "Book already returned"})
		return
	}

	borrow.Returned = true

	if err := c.DB.Save(&borrow).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to return the book"})
		return
	}

	ctx.JSON(200, gin.H{"message": fmt.Sprintf("Book %s returned successfully", bookID)})
}
