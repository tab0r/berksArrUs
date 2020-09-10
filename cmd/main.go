package main

import (
	"github.com/berksArrUs/pkg/domain"
	"github.com/berksArrUs/pkg/storage_service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type BookStorage interface {
	Add(bookData domain.Book) (string, error)
	//BulkAdd(books []domain.Book) ([]string, error)
	Retrieve(id string) (domain.Book, error)
	RetrieveAll() map[string]domain.Book
	//BulkRetrieve(ids []string) ([]domain.Book, error)
	Update(id string, bookData domain.Book) error
	Delete(id string) error
	//BulkDelete(ids []string) error
}

func routeListing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "routes for berk service")
	log.Println("Root endpoint hit")
}

// I prefer gorilla/mux router as it is a bit more flexible.
// But, I wanted to avoid importing anything outside base libraries
// I think this router builders are a nice alternative, and very flexible if I change storage implementations
func buildBooksRouter(storage BookStorage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Serve the resource.
			log.Println("retrieving book...")
			// get the id out of the path
			pathStrings := strings.Split(r.URL.Path, "/")
			if len(pathStrings) != 3 {
				errorString := fmt.Sprintf("bad path query: %s", r.URL.Path)
				log.Println(errorString)
				http.Error(w, errorString, http.StatusBadRequest)
				return
			}

			id := pathStrings[2]
			log.Println("path param is: " + id)
			book, err := storage.Retrieve(id)
			if err != nil {
				log.Printf("could not retrieve book")
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			_, _ = fmt.Fprint(w, fmt.Sprintf("bookData: %+v\n", book))
		case http.MethodPost:
			// Create a new record.
			log.Println("adding book...")
			var bookData domain.Book
			err := json.NewDecoder(r.Body).Decode(&bookData)
			if err != nil {
				log.Printf("could not add book; unmarshalling error")
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			} else {
				log.Printf("Storing book: %v", bookData)
				id, err := storage.Add(bookData)
				if err != nil {
					log.Printf("could not add book; storage error")
				} else {
					log.Printf("Stored book with id: %s \n", id)
					_, _ = fmt.Fprintf(w, id)
				}
			}
		case http.MethodPut:
			// Update an existing record.
			log.Println("updating book...")

			// get the id out of the path
			pathStrings := strings.Split(r.URL.Path, "/")
			if len(pathStrings) != 3 {
				errorString := fmt.Sprintf("bad path query: %s", r.URL.Path)
				log.Println(errorString)
				http.Error(w, errorString, http.StatusInternalServerError)
				return
			}

			id := pathStrings[2]
			log.Println("path param is: " + id)

			var bookData domain.Book
			err := json.NewDecoder(r.Body).Decode(&bookData)
			if err != nil {
				log.Printf("could not updating book; unmarshalling error")
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			err = storage.Update(id, bookData)
			if err != nil {
				log.Printf("could not update book")
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		case http.MethodDelete:
			// Remove the record.
			log.Println("deleting book...")

			// get the id out of the path
			pathStrings := strings.Split(r.URL.Path, "/")
			if len(pathStrings) != 3 {
				errorString := fmt.Sprintf("bad path query: %s", r.URL.Path)
				log.Println(errorString)
				http.Error(w, errorString, http.StatusInternalServerError)
				return
			}

			id := pathStrings[2]
			log.Println("path param is: " + id)

			err := storage.Delete(id)

			if err != nil {
				log.Printf("could not delete book")
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			//_, _ = fmt.Fprint(w, http.StatusOK)
		default:
			// Give an error message.
			_, _ = fmt.Fprintf(w, "invalid request")
		}
		return
	}
}

func buildBooksDumpRouter(storage BookStorage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Serve the resource.
			log.Println("retrieving all books...")
			books := storage.RetrieveAll()
			booksString := ""
			for k, b := range(books) {
				booksString = booksString + fmt.Sprintf("%s: %+v \n", k, b)
			}
			_, _ = fmt.Fprint(w, booksString)
		case http.MethodPost:
			// Create a new record.
			log.Println("bulk add books endpoint hit...")
			_, _ = fmt.Fprintf(w, "Bulk add is not yet implemented")
			//var bookData []domain.Book
			//err := json.NewDecoder(r.Body).Decode(&bookData)
			//if err != nil {
			//	log.Printf("could not add book; unmarshalling error")
			//	http.Error(w, err.Error(), http.StatusBadRequest)
			//	return
			//} else {
			//	log.Printf("Storing book: %v", bookData)
			//	id, err := storage.BulkAdd(bookData)
			//	if err != nil {
			//		log.Printf("could not add book; storage error")
			//	} else {
			//		log.Printf("Stored book with id: %s \n", id)
			//		_, _ = fmt.Fprintf(w, id)
			//	}
			//}
		default:
			// Give an error message.
			_, _ = fmt.Fprintf(w, "invalid request")
		}
	}
}

func searchRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "search not yet implemented")
}

func handleRequests(books BookStorage) {
	bookRouter := buildBooksRouter(books)
	bulkBookRouter := buildBooksDumpRouter(books)
	http.HandleFunc("/books", bookRouter)
	http.HandleFunc("/books/", bookRouter)
	http.HandleFunc("/bulk_books", bulkBookRouter)
	http.HandleFunc("/search", searchRoute)
	http.HandleFunc("/", routeListing)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("starting berks webserver")
	// not yet implemented
	//storageService, err := storage_service.NewBookStorageService("mycassdb.host")
	storageService, err := storage_service.NewBookStorageService("")
	if err != nil {
		log.Fatalln("Could not start webserver, error:", err)
	}
	handleRequests(storageService)
}
