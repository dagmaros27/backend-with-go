# Console-Based Library Management System

## Overview

The Console-Based Library Management System is a command-line application developed in Go that manages a library's book collection. This system allows users to perform various operations such as adding, removing, borrowing, and returning books, as well as listing available and borrowed books.

## Features

- Add Book: Add a new book to the library.
- Remove Book: Remove a book from the library.
- Borrow Book: Borrow a book from the library.
- Return Book: Return a borrowed book.
- List Available Books: Display a list of books currently available in the library.
- List Borrowed Books: Display a list of books currently borrowed by users.

## Installation

1. Install Go: Ensure Go is installed on your system. You can download and install it from Go's official website.
2. Clone Repository: Clone the repository containing the code.
   ```
   git clone https://github.com/dagmaros27/backend-with-go.git
   ```
3. Navigate to Project Directory:
   ```
   cd tasks/task3/library_management
   ```
4. Build and Run:
   ```
   go build ./library_management
   ```

## Usage

After running the application, you will be presented with a menu of options. Enter the corresponding number to select an operation. The available commands are:

- Add Book:

  - Input: Id, Title, Author
  - Example: Add Book -> Enter Id:1 -> Enter Title: The Great Gatsby -> Enter Author: F. Scott Fitzgerald

- Remove Book:

  - Input: Id
  - Example: Remove Book -> Enter Id: 1

- Borrow Book:

  - Input: Book Id, Member Id
  - Example: Borrow Book -> Enter Book ID:1 -> Enter Member ID:1

- Return Book:

  - Input: Book Id, Member Id
  - Example: Return Book -> Enter Book ID:1 -> Enter Member ID:1

- List Available Books:

  - Output: Displays a list of books currently available in the library.

- List Borrowed Books:
  - Output: Displays a list of books that are currently borrowed along with the borrower's name.

Code Structure

- main.go: The entry point of the application that handles user input and invokes corresponding functions.
- book.go: Defines the model for book.
- member.go: Defines the model for member.

- library_service.go: Implements the business logic of the library, including functions for adding, removing, borrowing, and returning books. This service layer interacts with the data structures defined in models.go.

- library_controller.go: Manages the interaction between the user interface (handled in main.go) and the library_service. It processes user commands and calls the appropriate functions in library_service.
