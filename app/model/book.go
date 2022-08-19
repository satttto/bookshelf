package model

type Category string

const (
	CATEGORY_UNKNOWN Category = "UNKNOWN"
	CATEGORY_TECH    Category = "TECH"
)

type Book struct {
	ID       int64
	Title    string
	Category Category
	Auther   string
	Timestamp
}
