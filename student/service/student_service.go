package service

import (
	"context"
	"jobhun/shareddomain/request"
	"jobhun/shareddomain/response"
)

type StudentService interface {
	CreateStudent(ctx context.Context, request request.StudentRequest) response.StudentResponse
	UpdateStudent(ctx context.Context, request request.StudentRequest) response.StudentResponseUpdate
	FindAllStudent(ctx context.Context) []response.StudentResponseBasic
	FindByIDStudent(ctx context.Context, id int) response.StudentResponse
	DeleteStudent(ctx context.Context, id int)
}
