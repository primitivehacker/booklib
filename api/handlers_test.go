package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllBooks(t *testing.T) {
	// Arrange
	expectedResult := []Book{
		{ID: "1", Title: "Title 1", Author: "Author 1", Quantity: 3, Isbn: "asdf3234"},
		{ID: "2", Title: "Title 2", Author: "Author 2", Quantity: 4, Isbn: "asdf23412"},
		{ID: "3", Title: "Title 3", Author: "Author 3", Quantity: 1, Isbn: "qerwog34"},
	}

	// Act
	actualResult := GetAllBooks()

	// Assert
	assert.Equal(t, expectedResult, actualResult)
}

func TestBookById(t *testing.T) {
	// Arrange
	expectedResult := &Book{ID: "1", Title: "Title 1", Author: "Author 1", Quantity: 3, Isbn: "asdf3234"}

	// Act
	actualResult, err := BookById("1")

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, actualResult)
}

func TestBookByIdNotFound(t *testing.T) {
	// Arrange

	// Act
	actualResult, err := BookById("10")

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, "book not found", err.Error())
	assert.Nil(t, actualResult)
}

func TestCheckOutBook(t *testing.T) {
	// Arrange
	expectedResult := &Book{ID: "1", Title: "Title 1", Author: "Author 1", Quantity: 3, Isbn: "asdf3234"},

	// Act
	actualResult, err := CheckOutBook("1")

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, actualResult)
}

func TestCheckOutBookNotFound(t *testing.T) {
	// Arrange

	// Act
	actualResult, err := CheckOutBook("10")

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, "Book not found.", err.Error())
	assert.Nil(t, actualResult)
}

func TestCheckOutBookNotAvailable(t *testing.T) {
	// Arrange
	CheckOutBook("3")  // Check out the last available book

	// Act
	actualResult, err := CheckOutBook("3")

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, "Book not available.", err.Error())
	assert.Nil(t, actualResult)
}

func TestReturnBook(t *testing.T) {
	// Arrange
	CheckOutBook("2")  // Check out a book
	expectedResult := &Book{ID: "2", Title: "Title 2", Author: "Author 2", Quantity: 4, Isbn: "asdf23412"},

	// Act
	actualResult, err := ReturnBook("2")

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, actualResult)
}

func TestReturnBookNotFound(t *testing.T) {
	// Arrange

	// Act
	actualResult, err := ReturnBook("10")

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, "Book not found.", err.Error())
	assert.Nil(t, actualResult)
}

func TestCreateBook(t *testing.T) {
	// Arrange
	newBook := Book{ID: "4", Title: "Title 4", Author: "Author 4", Quantity: 2, Isbn: "aslkj2342"}
	expectedResult := []Book{
		{ID: "1", Title: "Title 1", Author: "Author 1", Quantity: 3, Isbn: "asdf3234"},
		{ID: "2", Title: "Title 2", Author: "Author 2", Quantity: 4, Isbn: "asdf23412"},
		{ID: "3", Title: "Title 3", Author: "Author 3", Quantity: 1, Isbn: "qerwog34"},
		{ID: "4", Title: "Title 4", Author: "Author 4", Quantity: 2, Isbn: "aslkj2342"},
	}

	// Act
	CreateBook(newBook)
	actualResult := books

	// Assert
	assert.Equal(t, expectedResult, actualResult)
}