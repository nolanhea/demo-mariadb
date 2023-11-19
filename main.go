package main

import (
	"context"
	"log"
	"time"

	"test.com/book"
	"test.com/internal/storage"
)

func main() {
	store, err := storage.NewMySqlStorage(storage.MySqlConfig{
		Username: "user",
		Password: "password",
		DbName:   "db",
		Port:     3305,
		Host:     "localhost",
	})
	if err != nil {
		log.Fatalf("Impossible to create storage, %w", err)
	}
	b, err := store.Create(context.Background(), book.Book{
		Name:       "Practical Go Lessons",
		AuthorName: "Maximilien Andile",
		CreateTime: time.Now(),
	})
	if err != nil {
		log.Fatalf("impossible to insert: %s", err)
	}
	log.Println(b)

}
