package repos

import (
	"berksArrUs/pkg/domain"
	"errors"
	"fmt"
	"sync"
)

// simple in-memory cache for books

type BookMemCache struct {
	books map[string]domain.Book
	lock sync.RWMutex
}

func NewInMemCache() *BookMemCache {
	return &BookMemCache{
		books: make(map[string]domain.Book),
		lock:  sync.RWMutex{},
	}
}

func(c BookMemCache) Add(bookData domain.Book) (string, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	newKey := fmt.Sprintf("book_%d", len(c.books))
	c.books[newKey] = bookData
	return newKey, nil
}

func(c BookMemCache) Retrieve(id string) (domain.Book, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if book, exists := c.books[id]; exists {
		return book, nil
	}
	return domain.Book{}, errors.New("no book found")
}

func(c BookMemCache) RetrieveAll() map[string]domain.Book {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.books
}

func(c BookMemCache) Update(id string, bookData domain.Book) error {
	if _, err := c.Retrieve(id); err != nil {
		return errors.New("no book found to update")
	}
	c.lock.Lock()
	defer c.lock.Unlock()
	c.books[id] = bookData
	return nil
}

func(c BookMemCache) Delete(id string) error {
	if _, err := c.Retrieve(id); err != nil {
		return errors.New("no book found to delete")
	}
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.books, id)
	return nil
}