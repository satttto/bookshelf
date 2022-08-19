package db

import (
	"context"

	"github.com/pkg/errors"

	adapterDB "github.com/satttto/bookshelf/app/adapter/db"
	"github.com/satttto/bookshelf/app/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	client *gorm.DB
}

func New(dsn string) (adapterDB.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to make a connection with postgres")
	}

	return &DB{
		client: db,
	}, nil
}

func (d *DB) Migrate() error {
	if err := d.client.AutoMigrate(
		&model.Book{},
	); err != nil {
		return err
	}

	return nil
}

func (d *DB) AddBook(ctx context.Context, book model.Book) (model.Book, error) {
	if err := d.client.Create(&book).Error; err != nil {
		return model.Book{}, err
	}

	return book, nil
}

func (d *DB) ListBooks(ctx context.Context) ([]model.Book, error) {
	var books []model.Book
	if err := d.client.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}
