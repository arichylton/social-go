package store

import (
	"context"
	"database/sql"
	"errors"
)

var ErrorNotFound = errors.New("resource not found")

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, int64) (*Post, error)
		Delete(context.Context, int64) error
	}
	Users interface {
		Create(context.Context, *User) error
	}
	Comments interface {
		GetPostByID(context.Context, int64) ([]Comment, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:    &PostStore{db},
		Users:    &UserStore{db},
		Comments: &CommentStore{db},
	}
}
