package repository

import (
	"context"
	"database/sql"
	"jobhun/student/domain"
)

type StudentRepository interface {
	Create(ctx context.Context, tx *sql.Tx, student domain.Student) domain.Student
	Update(ctx context.Context, tx *sql.Tx, student domain.Student) domain.Student
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Student
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Student, error)
	Delete(ctx context.Context, tx *sql.Tx, id int)
}
