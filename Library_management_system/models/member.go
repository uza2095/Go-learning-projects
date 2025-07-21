package models

// Member represents a library member
type Member struct {
	ID            int
	Name          string
	BorrowedBooks []Book
}

// NewMember creates a new member instance
func NewMember(id int, name string) Member {
	return Member{
		ID:            id,
		Name:          name,
		BorrowedBooks: make([]Book, 0),
	}
}

// AddBorrowedBook adds a book to the member's borrowed books list
func (m *Member) AddBorrowedBook(book Book) {
	m.BorrowedBooks = append(m.BorrowedBooks, book)
}

// RemoveBorrowedBook removes a book from the member's borrowed books list
func (m *Member) RemoveBorrowedBook(bookID int) bool {
	for i, book := range m.BorrowedBooks {
		if book.ID == bookID {
			// Remove the book from the slice
			m.BorrowedBooks = append(m.BorrowedBooks[:i], m.BorrowedBooks[i+1:]...)
			return true
		}
	}
	return false
}

// HasBorrowedBook checks if the member has borrowed a specific book
func (m *Member) HasBorrowedBook(bookID int) bool {
	for _, book := range m.BorrowedBooks {
		if book.ID == bookID {
			return true
		}
	}
	return false
}

// GetBorrowedBooksCount returns the number of books borrowed by the member
func (m *Member) GetBorrowedBooksCount() int {
	return len(m.BorrowedBooks)
}