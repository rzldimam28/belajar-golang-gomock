package service_test

import (
	"belajar-golang-mock/constant"
	repository_mocks "belajar-golang-mock/repository/mocks"
	"belajar-golang-mock/service"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func newMockDB(t *testing.T) (sqlmock.Sqlmock, *sql.DB) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("failed create new mock db: %s", err.Error())
	}
	return mock, db
}

func TestCreateNewCat_Success(t *testing.T) {
	mck, db := newMockDB(t)
	repoMock := new(repository_mocks.Repository)
	cancellationTimeOut := time.Second * 10

	req := constant.Request{
		Name: "dummy-name",
	}

	mck.ExpectBegin()

	repoMock.On("InsertCat", mock.Anything, mock.Anything, mock.AnythingOfType("model.Cat")).Return(int64(1), nil).Once()

	service := service.New(repoMock, db, cancellationTimeOut)

	id, err := service.CreateNewCat(context.Background(), req)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), id)
}