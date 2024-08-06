package services

import (
	"errors"
	m "example.com/library-managment/models"
)

type LibraryManager interface {
	AddMember(member m.Member)
	AddBook(book m.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []m.Book
	ListBorrowedBooks(memberID int) []m.Book
}


type Library struct {
	books map[int]m.Book
	members map[int]m.Member
}

func NewLibrary() Library {
	return Library{
		books:   make(map[int]m.Book),
		members: make(map[int]m.Member),
	}
}

func (l *Library) AddMember(member m.Member){
	l.members[member.ID] = member
}

func (l *Library) AddBook(book m.Book){
	l.books[book.ID] = book
}

func (l *Library) RemoveBook(bookId int){
	delete(l.books,bookId)
}

func (l *Library) BorrowBook(bookId, memberId int)error {
	book,exists := l.books[bookId]

	if !exists{
		return errors.New("book not found")
	}

	if book.Status == "Borrowed" {
		return errors.New("book already borrowed")
	}

	member,exists := l.members[memberId]

	if !exists{
		return errors.New("member not found")
	}

	book.Status = "Borrowed"
	l.books[bookId] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.members[memberId] = member
	return nil
}


func (l *Library) ReturnBook(bookID int, memberID int) error {
	member, exists := l.members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	book, exists := l.books[bookID]
	if !exists {
		return errors.New("book not found")
	}

	if book.Status == "Available" {
		return errors.New("book is already available")
	}

	book.Status = "Available"
	l.books[bookID] = book

	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}

	l.members[memberID] = member
	return nil
}

func (l *Library) ListAvailableBooks() []m.Book {
	var availableBooks []m.Book
	for _, book := range l.books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []m.Book {
	member, exists := l.members[memberID]
	if !exists {
		return nil
	}
	return member.BorrowedBooks
}