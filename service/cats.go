package service

import (
	"belajar-golang-mock/constant"
	"belajar-golang-mock/model"
	"belajar-golang-mock/repository"
	"context"
	"database/sql"
	"log"
	"time"
)

type service struct {
	repository repository.Repository
	db                  *sql.DB
	cancellationTimeOut time.Duration
}

func New(repository repository.Repository, db *sql.DB, cancellationTimeOut time.Duration) Service {
	return &service{
		repository: repository,
		db: db,
		cancellationTimeOut: cancellationTimeOut,
	}
}

func (s *service) ListAllCats(ctx context.Context) ([]*constant.Response, error) {

	ctx, cancel:= context.WithTimeout(ctx, s.cancellationTimeOut)
	defer cancel()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatalf("error creating transactions: %s", err.Error())
	}

	cats, err := s.repository.FindAllCats(ctx, tx)	
	if err != nil {
		tx.Rollback()
		return []*constant.Response{}, nil
	}

	tx.Commit()

	var responses []*constant.Response
	for _, cat := range cats {
		response := constant.Response{}
		response.ID = cat.ID
		response.Name = cat.Name
		responses = append(responses, &response)
	}

	return responses, nil
}

func (s *service) CreateNewCat(ctx context.Context, request constant.Request) (int64, error) {
	ctx, cancel:= context.WithTimeout(ctx, s.cancellationTimeOut)
	defer cancel()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatalf("error creating transactions: %s", err.Error())
	}
	
	catToCreate := model.Cat{
		Name: request.Name,
	}

	id, err := s.repository.InsertCat(ctx, tx, catToCreate)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()

	return id, nil
}