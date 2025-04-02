package main

/*
	Write the following structs and their corresponding methods:

	Book: a struct that has four fields: an int
	id, string title,
	string author, and int
	quantity.


	Library: a struct that one field named books which
	is a slice of pointer to the Book struct.
	Library should also contain the following two methods.



	CheckoutBook(id int): a method that returns a pointer to a
	Book whose id field matches the
	id parameter, as well as a bool indicating
	whether the book can be checked out. A book can only be checked out if it
	has at least one copy remaining and exists in the library. If the book
	exists and has at least one copy you should first decrement it's number of
	copies by one and then return the pointer to it along with the
	bool true. If the book does not exist or has
	insufficient copies the method should return nil and
	false.


	ReturnBook(id int): a method that returns a boolean
	indicating if the book corresponding with the id parameter
	can be returned. A book can only be returned if it exists in the library,
	if it does not your method should return false. If the book
	does exist, you should increment it's number of copies before returning
	true.



	For all test cases you may assume each book has a unique id.


	Please note that a book that exists in the library can be returned even if
	it hasn't been checked out.
*/

// Write your code here.
type Book struct {
    id int
    title string
    author string
    quantity int
}

type Library struct {
    books []*Book
}

func (l *Library) CheckoutBook(id int) (*Book, bool) {
    var book *Book = nil
    
    for _, b := range l.books {
        if b.id == id {
            book = b
        }
    }

    if book != nil {
        if book.quantity >= 1 {
            book.quantity -= 1
            return book, true
        } 
    }

    return nil, false
}


func (l *Library) ReturnBook(id int) bool {
    var book *Book = nil
    
    for _, b := range l.books {
        if b.id == id {
            book = b
        }
    }
    
    if book != nil {
        book.quantity += 1
        return true
    }

    return false
}