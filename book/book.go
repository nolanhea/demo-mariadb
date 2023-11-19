package book

import "time"

type Book struct {
	ID         int64
	Name       string
	AuthorName string
	CreateTime time.Time
}
