package models

// Book represents a book in the library
type Book struct {
	ID     int
	Title  string
	Author string
	Status string // "Available" or "Borrowed"
}

// NewBook creates a new book instance
func NewBook(id int, title, author string) Book {
	return Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: "Available",
	}
}

// IsAvailable checks if the book is available for borrowing
func (b *Book) IsAvailable() bool {
	return b.Status == "Available"
}

// SetBorrowed marks the book as borrowed
func (b *Book) SetBorrowed() {
	b.Status = "Borrowed"
}

// SetAvailable marks the book as available
func (b *Book) SetAvailable() {
	b.Status = "Available"
}