package repository

import (
	"belajar-golang-mock/model"
	"context"
	"database/sql"
)

type repository struct{}

func New() Repository {
	return &repository{}
}

var (
	SelectAllCatsQuery = "SELECT c.id, c.name FROM cats c;"
	InsertCatQuery = "INSERT INTO cats(name) VALUES($1) returning id;"
)

func (r *repository) FindAllCats(ctx context.Context, tx *sql.Tx) ([]*model.Cat, error) {
	
	rows, err := tx.QueryContext(ctx, SelectAllCatsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var cats []*model.Cat
	for rows.Next() {
		cat := &model.Cat{}
		err := rows.Scan(&cat.ID, &cat.Name)
		if err != nil {
			return nil, err
		}
		cats = append(cats, cat)
	}

	return cats, nil
}

func (r *repository) InsertCat(ctx context.Context, tx *sql.Tx, cat model.Cat) (int64, error) {
	
	var id int64
	err := tx.QueryRowContext(ctx, InsertCatQuery, cat.Name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}