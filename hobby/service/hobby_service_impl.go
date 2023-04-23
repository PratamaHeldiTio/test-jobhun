package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"jobhun/helper"
	"jobhun/hobby/domain"
	"jobhun/hobby/repository"
	"jobhun/shareddomain/request"
	"jobhun/shareddomain/response"
)

type HobbyServiceImpl struct {
	repo     repository.HobbyRepository
	db       *sql.DB
	validate *validator.Validate
}

func NewHobbyService(repo repository.HobbyRepository, db *sql.DB, validate *validator.Validate) HobbyService {
	return &HobbyServiceImpl{repo, db, validate}
}

func (service *HobbyServiceImpl) CreateHobby(ctx context.Context, request request.HobbyCreateRequest) response.ResponseHobby {
	err := service.validate.Struct(request)
	if err != nil {
		panic(err)
	}

	// open connection
	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	hobby := domain.Hobby{
		Name: request.Name,
	}

	hobby = service.repo.Create(ctx, tx, hobby)

	return response.ToResponseHobby(hobby)
}
