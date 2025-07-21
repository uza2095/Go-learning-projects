package services

import (
	"errors"
	"library_management/models"
)

// LibraryManager interface defines the contract for library operations
type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
	AddMember(member models.Member)
	GetMember(memberID int) (*models.Member, error)
	GetBook(bookID int) (*models.Book, error)
	ListAllBooks() []models.Book
	ListAllMembers() []models.Member
}

// Library implements the LibraryManager interface
type Library struct {
	books   map[int]models.Book
	members map[int]models.Member
}

// NewLibrary creates a new library instance
func NewLibrary() *Library {
	return &Library{
		books:   make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
}

// AddBook adds a new book to the library
func (l *Library) AddBook(book models.Book) {
	l.books[book.ID] = book
}

// RemoveBook removes a book from the library by its ID
func (l *Library) RemoveBook(bookID int) error {
	book, exists := l.books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	
	if book.Status == "Borrowed" {
		return errors.New("cannot remove a borrowed book")
	}
	
	delete(l.books, bookID)
	return nil
}

// BorrowBook allows a member to borrow a book if it is available
func (l *Library) BorrowBook(bookID int, memberID int) error {
	// Check if book exists
	book, exists := l.books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	
	// Check if book is available
	if !book.IsAvailable() {
		return errors.New("book is already borrowed")
	}
	
	// Check if member exists
	member, exists := l.members[memberID]
	if !exists {
		return errors.New("member not found")
	}
	
	// Update book status
	book.SetBorrowed()
	l.books[bookID] = book
	
	// Add book to member's borrowed books
	member.AddBorrowedBook(book)
	l.members[memberID] = member
	
	return nil
}

// ReturnBook allows a member to return a borrowed book
func (l *Library) ReturnBook(bookID int, memberID int) error {
	// Check if book exists
	book, exists := l.books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	
	// Check if member exists
	member, exists := l.members[memberID]
	if !exists {
		return errors.New("member not found")
	}
	
	// Check if member has borrowed this book
	if !member.HasBorrowedBook(bookID) {
		return errors.New("member has not borrowed this book")
	}
	
	// Update book status
	book.SetAvailable()
	l.books[bookID] = book
	
	// Remove book from member's borrowed books
	member.RemoveBorrowedBook(bookID)
	l.members[memberID] = member
	
	return nil
}

// ListAvailableBooks lists all available books in the library
func (l *Library) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range l.books {
		if book.IsAvailable() {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

// ListBorrowedBooks lists all books borrowed by a specific member
func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := l.members[memberID]
	if !exists {
		return []models.Book{}
	}
	return member.BorrowedBooks
}

// AddMember adds a new member to the library
func (l *Library) AddMember(member models.Member) {
	l.members[member.ID] = member
}

// GetMember retrieves a member by ID
func (l *Library) GetMember(memberID int) (*models.Member, error) {
	member, exists := l.members[memberID]
	if !exists {
		return nil, errors.New("member not found")
	}
	return &member, nil
}

// GetBook retrieves a book by ID
func (l *Library) GetBook(bookID int) (*models.Book, error) {
	book, exists := l.books[bookID]
	if !exists {
		return nil, errors.New("book not found")
	}
	return &book, nil
}

// ListAllBooks returns all books in the library
func (l *Library) ListAllBooks() []models.Book {
	var allBooks []models.Book
	for _, book := range l.books {
		allBooks = append(allBooks, book)
	}
	return allBooks
}

// ListAllMembers returns all members in the library
func (l *Library) ListAllMembers() []models.Member {
	var allMembers []models.Member
	for _, member := range l.members {
		allMembers = append(allMembers, member)
	}
	return allMembers
}