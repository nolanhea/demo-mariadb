package storage

import (
	"context"

	"test.com/book"
)

type Storage interface {
	Create(ctx context.Context, b book.Book) error
}
