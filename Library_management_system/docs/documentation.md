# Library Management System Documentation

## Overview
This is a console-based library management system implemented in Go.

## Architecture

### Project Structure
```
library_management/
├── main.go                     # Entry point of the application
├── controllers/
│   └── library_controller.go  # Handles console input and invokes service methods
├── models/
│   ├── book.go                # Defines the Book struct
│   └── member.go              # Defines the Member struct
├── services/
│   └── library_service.go     # Contains business logic and data manipulation
├── docs/
│   └── documentation.md       # System documentation
└── go.mod                     # Module definition and dependencies
```

## Core Components

### Models

#### Book Struct
```go
type Book struct {
    ID     int
    Title  string
    Author string
    Status string // "Available" or "Borrowed"
}
```

**Methods:**
- `NewBook(id int, title, author string) Book` - Creates a new book instance
- `IsAvailable() bool` - Checks if the book is available for borrowing
- `SetBorrowed()` - Marks the book as borrowed
- `SetAvailable()` - Marks the book as available

#### Member Struct
```go
type Member struct {
    ID            int
    Name          string
    BorrowedBooks []Book
}
```

**Methods:**
- `NewMember(id int, name string) Member` - Creates a new member instance
- `AddBorrowedBook(book Book)` - Adds a book to the member's borrowed books list
- `RemoveBorrowedBook(bookID int) bool` - Removes a book from the member's borrowed books list
- `HasBorrowedBook(bookID int) bool` - Checks if the member has borrowed a specific book
- `GetBorrowedBooksCount() int` - Returns the number of books borrowed by the member

### Interfaces

#### LibraryManager Interface
```go
type LibraryManager interface {
    AddBook(book Book)
    RemoveBook(bookID int) error
    BorrowBook(bookID int, memberID int) error
    ReturnBook(bookID int, memberID int) error
    ListAvailableBooks() []Book
    ListBorrowedBooks(memberID int) []Book
    AddMember(member Member)
    GetMember(memberID int) (*Member, error)
    GetBook(bookID int) (*Book, error)
    ListAllBooks() []Book
    ListAllMembers() []Member
}
```

### Services

#### Library Service
The `Library` struct implements the `LibraryManager` interface and contains:
- `books map[int]Book` - Stores all books with book ID as the key
- `members map[int]Member` - Stores all members with member ID as the key

**Key Methods:**
- `AddBook(book Book)` - Adds a new book to the library
- `RemoveBook(bookID int) error` - Removes a book from the library by its ID
- `BorrowBook(bookID int, memberID int) error` - Allows a member to borrow a book if available
- `ReturnBook(bookID int, memberID int) error` - Allows a member to return a borrowed book
- `ListAvailableBooks() []Book` - Lists all available books in the library
- `ListBorrowedBooks(memberID int) []Book` - Lists all books borrowed by a specific member

### Controllers

#### Library Controller
Handles console interaction and user input validation. Provides a menu-driven interface for:
1. Adding new books
2. Removing existing books
3. Adding new members
4. Borrowing books
5. Returning books
6. Listing available books
7. Listing borrowed books by member
8. Listing all books
9. Listing all members

## Features

### Core Functionality
- **Book Management**: Add, remove, and track books in the library
- **Member Management**: Register and manage library members
- **Borrowing System**: Allow members to borrow available books
- **Return System**: Process book returns and update availability
- **Listing Operations**: View available books, borrowed books, and member information

### Error Handling
The system includes comprehensive error handling for:
- Book not found scenarios
- Member not found scenarios
- Attempting to borrow unavailable books
- Attempting to remove borrowed books
- Invalid input validation
- Duplicate ID prevention

### Data Validation
- Input validation for numeric IDs
- Empty string validation for names and titles
- Existence checks before operations
- Status validation for book operations

## Usage

### Running the Application
```bash
go run main.go
```

### Sample Data
The application initializes with sample data including:
- 5 sample books (Go programming, Clean Code, Design Patterns, etc.)
- 3 sample members (Alice Johnson, Bob Smith, Charlie Brown)

### Menu Options
1. **Add Book**: Enter book ID, title, and author
2. **Remove Book**: Enter book ID to remove (only if not borrowed)
3. **Add Member**: Enter member ID and name
4. **Borrow Book**: Enter book ID and member ID
5. **Return Book**: Enter book ID and member ID
6. **List Available Books**: Display all available books
7. **List Borrowed Books by Member**: Enter member ID to see their borrowed books
8. **List All Books**: Display all books with their status
9. **List All Members**: Display all members with borrowed book count
0. **Exit**: Close the application

## Technical Implementation

### Go Features Demonstrated
- **Structs**: Book and Member structs with fields and methods
- **Interfaces**: LibraryManager interface defining contracts
- **Maps**: Used for efficient book and member storage and retrieval
- **Slices**: Used for storing borrowed books and listing operations
- **Error Handling**: Proper error handling with custom error messages
- **Methods**: Struct methods for encapsulation and behavior
- **Packages**: Organized code structure with separate packages
- **Pointers**: Used for efficient memory management and reference passing

### Design Patterns
- **Repository Pattern**: Library service acts as a repository for books and members
- **Controller Pattern**: Separation of concerns between UI and business logic
- **Interface Segregation**: Clean interface definition for library operations

## Future Enhancements
- Persistent data storage (file or database)
- Book reservation system
- Due date tracking and late fees
- Search functionality by title or author
- Book categories and genres
- Member borrowing limits
- Audit trail for all operations

## Testing
To test the application:
1. Run the application
2. Use the sample data or add your own books and members
3. Test borrowing and returning operations
4. Verify error handling with invalid inputs
5. Check listing operations for accuracy

## Dependencies
- Go 1.21 or higher
- No external dependencies required (uses only standard library)