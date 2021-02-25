package models //import "github.com/JenYata/goEx"

//Book struct
type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

//CreateBookInput struct
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

//UpdateBookInput struct
type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
