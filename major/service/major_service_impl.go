package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"jobhun/helper"
	"jobhun/major/domain"
	"jobhun/major/repository"
	"jobhun/shareddomain/request"
	"jobhun/shareddomain/response"
)

type MajorServiceImpl struct {
	repo     repository.MajorRepository
	db       *sql.DB
	validate *validator.Validate
}

func NewMajorService(repo repository.MajorRepository, db *sql.DB, validate *validator.Validate) MajorService {
	return &MajorServiceImpl{repo, db, validate}
}

func (service *MajorServiceImpl) CreateMajor(ctx context.Context, request request.MajorCreateRequest) response.ResponseMajor {
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

	major := domain.Major{
		Name: request.Name,
	}

	major = service.repo.Create(ctx, tx, major)

	return response.ToResponseMajor(major)
}
