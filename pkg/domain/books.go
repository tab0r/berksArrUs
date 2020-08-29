// common elements go in domain

package domain

type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	ISBN        string `json:"isbn"`
}
