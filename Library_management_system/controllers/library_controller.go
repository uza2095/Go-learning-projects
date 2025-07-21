package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

// LibraryController handles console input and invokes appropriate service methods
type LibraryController struct {
	libraryService services.LibraryManager
	scanner        *bufio.Scanner
}

// NewLibraryController creates a new library controller
func NewLibraryController(libraryService services.LibraryManager) *LibraryController {
	return &LibraryController{
		libraryService: libraryService,
		scanner:        bufio.NewScanner(os.Stdin),
	}
}

// Start begins the console interface
func (lc *LibraryController) Start() {
	fmt.Println("=== Welcome to Library Management System ===")
	
	// Add some sample data
	lc.initializeSampleData()
	
	for {
		lc.displayMenu()
		choice := lc.getInput("Enter your choice: ")
		
		switch choice {
		case "1":
			lc.addBook()
		case "2":
			lc.removeBook()
		case "3":
			lc.addMember()
		case "4":
			lc.borrowBook()
		case "5":
			lc.returnBook()
		case "6":
			lc.listAvailableBooks()
		case "7":
			lc.listBorrowedBooks()
		case "8":
			lc.listAllBooks()
		case "9":
			lc.listAllMembers()
		case "0":
			fmt.Println("Thank you for using Library Management System!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
		
		fmt.Println("\nPress Enter to continue...")
		lc.scanner.Scan()
	}
}

// displayMenu shows the main menu options
func (lc *LibraryController) displayMenu() {
	fmt.Println("\n=== Library Management System ===")
	fmt.Println("1. Add Book")
	fmt.Println("2. Remove Book")
	fmt.Println("3. Add Member")
	fmt.Println("4. Borrow Book")
	fmt.Println("5. Return Book")
	fmt.Println("6. List Available Books")
	fmt.Println("7. List Borrowed Books by Member")
	fmt.Println("8. List All Books")
	fmt.Println("9. List All Members")
	fmt.Println("0. Exit")
}

// getInput prompts user for input and returns the trimmed string
func (lc *LibraryController) getInput(prompt string) string {
	fmt.Print(prompt)
	lc.scanner.Scan()
	return strings.TrimSpace(lc.scanner.Text())
}

// getIntInput prompts user for integer input
func (lc *LibraryController) getIntInput(prompt string) (int, error) {
	input := lc.getInput(prompt)
	return strconv.Atoi(input)
}

// addBook handles adding a new book
func (lc *LibraryController) addBook() {
	fmt.Println("\n=== Add New Book ===")
	
	id, err := lc.getIntInput("Enter Book ID: ")
	if err != nil {
		fmt.Println("Invalid ID. Please enter a valid number.")
		return
	}
	
	// Check if book already exists
	if _, err := lc.libraryService.GetBook(id); err == nil {
		fmt.Println("Book with this ID already exists.")
		return
	}
	
	title := lc.getInput("Enter Book Title: ")
	author := lc.getInput("Enter Book Author: ")
	
	if title == "" || author == "" {
		fmt.Println("Title and Author cannot be empty.")
		return
	}
	
	book := models.NewBook(id, title, author)
	lc.libraryService.AddBook(book)
	
	fmt.Printf("Book '%s' by %s has been added successfully!\n", title, author)
}

// removeBook handles removing a book
func (lc *LibraryController) removeBook() {
	fmt.Println("\n=== Remove Book ===")
	
	id, err := lc.getIntInput("Enter Book ID to remove: ")
	if err != nil {
		fmt.Println("Invalid ID. Please enter a valid number.")
		return
	}
	
	err = lc.libraryService.RemoveBook(id)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	
	fmt.Println("Book removed successfully!")
}

// addMember handles adding a new member
func (lc *LibraryController) addMember() {
	fmt.Println("\n=== Add New Member ===")
	
	id, err := lc.getIntInput("Enter Member ID: ")
	if err != nil {
		fmt.Println("Invalid ID. Please enter a valid number.")
		return
	}
	
	// Check if member already exists
	if _, err := lc.libraryService.GetMember(id); err == nil {
		fmt.Println("Member with this ID already exists.")
		return
	}
	
	name := lc.getInput("Enter Member Name: ")
	if name == "" {
		fmt.Println("Name cannot be empty.")
		return
	}
	
	member := models.NewMember(id, name)
	lc.libraryService.AddMember(member)
	
	fmt.Printf("Member '%s' has been added successfully!\n", name)
}

// borrowBook handles book borrowing
func (lc *LibraryController) borrowBook() {
	fmt.Println("\n=== Borrow Book ===")
	
	bookID, err := lc.getIntInput("Enter Book ID to borrow: ")
	if err != nil {
		fmt.Println("Invalid Book ID. Please enter a valid number.")
		return
	}
	
	memberID, err := lc.getIntInput("Enter Member ID: ")
	if err != nil {
		fmt.Println("Invalid Member ID. Please enter a valid number.")
		return
	}
	
	err = lc.libraryService.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	
	fmt.Println("Book borrowed successfully!")
}

// returnBook handles book returning
func (lc *LibraryController) returnBook() {
	fmt.Println("\n=== Return Book ===")
	
	bookID, err := lc.getIntInput("Enter Book ID to return: ")
	if err != nil {
		fmt.Println("Invalid Book ID. Please enter a valid number.")
		return
	}
	
	memberID, err := lc.getIntInput("Enter Member ID: ")
	if err != nil {
		fmt.Println("Invalid Member ID. Please enter a valid number.")
		return
	}
	
	err = lc.libraryService.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	
	fmt.Println("Book returned successfully!")
}

// listAvailableBooks displays all available books
func (lc *LibraryController) listAvailableBooks() {
	fmt.Println("\n=== Available Books ===")
	
	books := lc.libraryService.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No available books found.")
		return
	}
	
	fmt.Printf("%-5s %-30s %-20s %-10s\n", "ID", "Title", "Author", "Status")
	fmt.Println(strings.Repeat("-", 70))
	
	for _, book := range books {
		fmt.Printf("%-5d %-30s %-20s %-10s\n", book.ID, book.Title, book.Author, book.Status)
	}
}

// listBorrowedBooks displays books borrowed by a specific member
func (lc *LibraryController) listBorrowedBooks() {
	fmt.Println("\n=== Borrowed Books by Member ===")
	
	memberID, err := lc.getIntInput("Enter Member ID: ")
	if err != nil {
		fmt.Println("Invalid Member ID. Please enter a valid number.")
		return
	}
	
	// Check if member exists
	member, err := lc.libraryService.GetMember(memberID)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	
	books := lc.libraryService.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Printf("Member '%s' has not borrowed any books.\n", member.Name)
		return
	}
	
	fmt.Printf("Books borrowed by '%s':\n", member.Name)
	fmt.Printf("%-5s %-30s %-20s %-10s\n", "ID", "Title", "Author", "Status")
	fmt.Println(strings.Repeat("-", 70))
	
	for _, book := range books {
		fmt.Printf("%-5d %-30s %-20s %-10s\n", book.ID, book.Title, book.Author, book.Status)
	}
}

// listAllBooks displays all books in the library
func (lc *LibraryController) listAllBooks() {
	fmt.Println("\n=== All Books ===")
	
	books := lc.libraryService.ListAllBooks()
	if len(books) == 0 {
		fmt.Println("No books found in the library.")
		return
	}
	
	fmt.Printf("%-5s %-30s %-20s %-10s\n", "ID", "Title", "Author", "Status")
	fmt.Println(strings.Repeat("-", 70))
	
	for _, book := range books {
		fmt.Printf("%-5d %-30s %-20s %-10s\n", book.ID, book.Title, book.Author, book.Status)
	}
}

// listAllMembers displays all members in the library
func (lc *LibraryController) listAllMembers() {
	fmt.Println("\n=== All Members ===")
	
	members := lc.libraryService.ListAllMembers()
	if len(members) == 0 {
		fmt.Println("No members found in the library.")
		return
	}
	
	fmt.Printf("%-5s %-20s %-15s\n", "ID", "Name", "Borrowed Books")
	fmt.Println(strings.Repeat("-", 45))
	
	for _, member := range members {
		fmt.Printf("%-5d %-20s %-15d\n", member.ID, member.Name, member.GetBorrowedBooksCount())
	}
}

// initializeSampleData adds some sample books and members for testing
func (lc *LibraryController) initializeSampleData() {
	// Add sample books
	books := []models.Book{
		models.NewBook(1, "The Go Programming Language", "Alan Donovan"),
		models.NewBook(2, "Clean Code", "Robert Martin"),
		models.NewBook(3, "Design Patterns", "Gang of Four"),
		models.NewBook(4, "Effective Go", "The Go Team"),
		models.NewBook(5, "Concurrency in Go", "Katherine Cox-Buday"),
	}
	
	for _, book := range books {
		lc.libraryService.AddBook(book)
	}
	
	// Add sample members
	members := []models.Member{
		models.NewMember(1, "Alice Johnson"),
		models.NewMember(2, "Bob Smith"),
		models.NewMember(3, "Charlie Brown"),
	}
	
	for _, member := range members {
		lc.libraryService.AddMember(member)
	}
	
	fmt.Println("Sample data initialized successfully!")
}