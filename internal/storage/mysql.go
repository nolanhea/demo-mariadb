package storage

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"test.com/book"
)

type MySqlStorage struct {
	db *sql.DB
}

type MySqlConfig struct {
	Username string
	Password string
	DbName   string
	Port     uint
	Host     string
}

func NewMySqlStorage(conf MySqlConfig) (MySqlStorage, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		conf.Username, conf.Password,
		conf.Host, conf.Port, conf.DbName,
	)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return MySqlStorage{}, fmt.Errorf("Impossible to open: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return MySqlStorage{}, fmt.Errorf("Impossibel to open sql connection, %w", err)
	}
	return MySqlStorage{
		db: db,
	}, nil
}

func (s MySqlStorage) Create(ctx context.Context, b book.Book) (book.Book, error) {
	query := "INSERT INTO `book` (`create_time`, `name`, `author_name`) VALUES (?, ?, ?)"
	insertResult, err := s.db.ExecContext(ctx, query, b.CreateTime, b.Name, b.AuthorName)
	if err != nil {
		return b, fmt.Errorf("error while insert: %w", err)
	}
	id, err := insertResult.LastInsertId()
	if err != nil {
		return b, fmt.Errorf("error while get last insert id: %w", err)
	}
	b.ID = id

	return b, nil
}

//go get -u github.com/go-sql-driver/mysql
//docker-compose exec database mariadb -uuser -ppassword
//show databases;
//use db;
