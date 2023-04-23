package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"jobhun/helper"
	"jobhun/shareddomain/request"
	"jobhun/shareddomain/response"
	"jobhun/student/domain"
	"jobhun/student/repository"
	"time"
)

type StudentServiceImpl struct {
	repo     repository.StudentRepository
	db       *sql.DB
	validate *validator.Validate
}

func NewStudentService(repo repository.StudentRepository, db *sql.DB, validate *validator.Validate) StudentService {
	return &StudentServiceImpl{repo, db, validate}
}

func (service *StudentServiceImpl) CreateStudent(ctx context.Context, request request.StudentRequest) response.StudentResponse {
	err := service.validate.Struct(request)
	if err != nil {
		panic(err)
	}

	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	student := domain.Student{
		Name:      request.Name,
		Age:       request.Age,
		Gender:    request.Gender,
		Hobbies:   request.Hobbies,
		CreatedAt: time.Now(),
	}

	student = service.repo.Create(ctx, tx, student)

	return response.ToResponseStudent(student)
}

func (service *StudentServiceImpl) UpdateStudent(ctx context.Context, request request.StudentRequest) response.StudentResponseUpdate {
	err := service.validate.Struct(request)
	if err != nil {
		panic(err)
	}

	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	_ = service.FindByIDStudent(ctx, request.ID)
	if err != nil {
		panic(err)
	}

	studentRequest := domain.Student{
		ID:     request.ID,
		Name:   request.Name,
		Age:    request.Age,
		Gender: request.Gender,
	}

	student := service.repo.Update(ctx, tx, studentRequest)

	return response.ToResponseUpdateStudent(student)
}

func (service *StudentServiceImpl) FindAllStudent(ctx context.Context) []response.StudentResponseBasic {
	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	categories := service.repo.FindAll(ctx, tx)

	categoriesResponse := []response.StudentResponseBasic{}
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, response.ToResponseStudentBasic(category))
	}

	return categoriesResponse
}

func (service *StudentServiceImpl) FindByIDStudent(ctx context.Context, id int) response.StudentResponse {
	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	student, err := service.repo.FindById(ctx, tx, id)
	if err != nil {
		panic(err)
	}

	return response.ToResponseStudent(student)
}

func (service *StudentServiceImpl) DeleteStudent(ctx context.Context, id int) {
	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	_ = service.FindByIDStudent(ctx, id)
	if err != nil {
		panic(err)
	}

	service.repo.Delete(ctx, tx, id)
}
