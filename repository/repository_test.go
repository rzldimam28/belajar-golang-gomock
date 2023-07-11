package repository_test

import (
	"belajar-golang-mock/model"
	"belajar-golang-mock/repository"
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func newMockRepository(t *testing.T) (repository.Repository, *sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("failed opening a stub database connection: %s", err.Error())
	}
	return repository.New(), db, mock
}

func TestInsertCat_Success(t *testing.T) {

	mockRepo, db, mck := newMockRepository(t)
	defer db.Close()

	cat := model.Cat{
		Name: "dummy-name",
	}

	mck.ExpectBegin()

  tx, err := db.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	mck.
		ExpectQuery(repository.InsertCatQuery).
		WithArgs(cat.Name).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	id, err := mockRepo.InsertCat(context.Background(), tx, cat)

	assert.Nil(t, err)
	assert.Equal(t, id, int64(1))

}