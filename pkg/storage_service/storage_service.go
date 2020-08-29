package storage_service

import (
	"berksArrUs/pkg/domain"
	"berksArrUs/pkg/repos"
	"errors"
)

type BookRepo interface {
	Add(bookData domain.Book) (string, error)
	Retrieve(id string) (domain.Book, error)
	RetrieveAll() map[string]domain.Book
	Update(id string, bookData domain.Book) error
	Delete(id string) error
}

type BookStorageService struct {
	storage BookRepo
}

func NewBookStorageService(cassDbHost string) (*BookStorageService, error) {
	var repository BookRepo
	if cassDbHost != "" {
		return nil, errors.New("not yet implemented")
	}
	repository = repos.NewInMemCache()
	return &BookStorageService{storage: repository}, nil
}

func(s *BookStorageService) Add(bookData domain.Book) (string, error) {
	return s.storage.Add(bookData)
}

func(s *BookStorageService) Retrieve(id string) (domain.Book, error) {
	return s.storage.Retrieve(id)
}

func(s *BookStorageService) RetrieveAll() map[string]domain.Book {
	return s.storage.RetrieveAll()
}

func(s *BookStorageService) Update(id string, bookData domain.Book) error {
	return s.storage.Update(id, bookData)
}

func(s *BookStorageService) Delete(id string) error {
	return s.storage.Delete(id)
}