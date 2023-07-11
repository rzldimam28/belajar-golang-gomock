package service

import (
	"belajar-golang-mock/constant"
	"context"
)

type Service interface {
	ListAllCats(ctx context.Context) ([]*constant.Response, error) 
	CreateNewCat(ctx context.Context, request constant.Request) (int64, error)
}