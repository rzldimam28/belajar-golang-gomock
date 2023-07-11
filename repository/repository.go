package repository

import (
	"belajar-golang-mock/model"
	"context"
	"database/sql"
)

type Repository interface {
	FindAllCats(ctx context.Context, tx *sql.Tx) ([]*model.Cat, error)
	InsertCat(ctx context.Context, tx *sql.Tx, cat model.Cat) (int64, error)
}