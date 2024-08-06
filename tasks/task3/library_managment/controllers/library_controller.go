package controller

import (
	"fmt"
	"os"
	"os/exec"

	m "example.com/library-managment/models"
	s "example.com/library-managment/services"
)


func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}


type LibraryController struct {
	library *s.Library
}

func NewLibraryController(library *s.Library) *LibraryController {
	return &LibraryController{library: library}
}

func (lc *LibraryController) Run() {
	for {
		fmt.Println("Library Management System")
		fmt.Println("1. Add a new book")
		fmt.Println("2. Remove an existing book")
		fmt.Println("3. Borrow a book")
		fmt.Println("4. Return a book")
		fmt.Println("5. List all available books")
		fmt.Println("6. List all borrowed books by a member")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			clearConsole()
			lc.addBook()
		case 2:
			clearConsole()
			lc.removeBook()
		case 3:
			clearConsole()
			lc.borrowBook()
		case 4:
			clearConsole()
			lc.returnBook()
		case 5:
			clearConsole()
			lc.listAvailableBooks()
		case 6:
			clearConsole()
			lc.listBorrowedBooks()
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func (lc *LibraryController) addBook() {
	var id int
	var title, author string
	fmt.Print("Enter book ID: ")
	fmt.Scan(&id)
	fmt.Print("Enter book title: ")
	fmt.Scan(&title)
	fmt.Print("Enter book author: ")
	fmt.Scan(&author)
	lc.library.AddBook(m.Book{ID: id, Title: title, Author: author, Status: "Available"})
	fmt.Println("Book added successfully.")
}

func (lc *LibraryController) removeBook() {
	var id int
	fmt.Print("Enter book ID to remove: ")
	fmt.Scan(&id)
	lc.library.RemoveBook(id)
	fmt.Println("Book removed successfully.")
}

func (lc *LibraryController) borrowBook() {
	var bookID, memberID int
	fmt.Print("Enter book ID to borrow: ")
	fmt.Scan(&bookID)
	fmt.Print("Enter member ID: ")
	fmt.Scan(&memberID)
	err := lc.library.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully.")
	}
}

func (lc *LibraryController) returnBook() {
	var bookID, memberID int
	fmt.Print("Enter book ID to return: ")
	fmt.Scan(&bookID)
	fmt.Print("Enter member ID: ")
	fmt.Scan(&memberID)
	err := lc.library.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully.")
	}
}

func (lc *LibraryController) listAvailableBooks() {
	books := lc.library.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No available books.")
	} else {
		fmt.Println("Available Books:")
		for _, book := range books {
			fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}

func (lc *LibraryController) listBorrowedBooks() {
	var memberID int
	fmt.Print("Enter member ID: ")
	fmt.Scan(&memberID)
	books := lc.library.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Println("No borrowed books.")
	} else {
		fmt.Println("Borrowed Books:")
		for _, book := range books {
			fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}
