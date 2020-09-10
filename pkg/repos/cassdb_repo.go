package repos

import (
	"github.com/berksArrUs/pkg/domain"
	"errors"
)

// stubs for book storage in CQL

type BookCassDb struct {
}

func NewBookCassDb() BookCassDb {
	return BookCassDb{}
}

func(c *BookCassDb) Add(bookData domain.Book) (string, error) {
	return "", errors.New("boom")
}

func(c *BookCassDb) Retrieve(id string) ([]domain.Book, error) {
	return []domain.Book{}, errors.New("No books found")
}

func(c BookCassDb) RetrieveAll() map[string]domain.Book {
	return map[string]domain.Book{}
}

func(c *BookCassDb) Update(id string, bookData domain.Book) error {
	return nil
}

func(c *BookCassDb) Delete(id string) error {
	return errors.New("could not find book to delete")
}
