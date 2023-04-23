package repository

import (
	"context"
	"database/sql"
	"errors"
	"jobhun/student/domain"
)

type StudentRepositoryImpl struct {
}

func NewStudentRepositoryImpl() StudentRepository {
	return &StudentRepositoryImpl{}
}

func (repo *StudentRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, student domain.Student) domain.Student {
	query := "insert into mahasiswa(nama, usia, gender, tanggal_registrasi) values (?,?,?,?)"
	result, err := tx.ExecContext(ctx, query, student.Name, student.Age, student.Gender, student.CreatedAt)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	student.ID = int(id)

	for _, hobby := range student.Hobbies {
		query := "insert into mahasiswa_hobi (id_mahasiswa, id_hobi) values (?, ?)"
		_, err := tx.ExecContext(ctx, query, student.ID, hobby)
		if err != nil {
			panic(err)
		}
	}

	return student
}

func (repo *StudentRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, student domain.Student) domain.Student {
	query := "update mahasiswa set nama = ?, usia = ?, gender = ? where id = ?"
	_, err := tx.ExecContext(ctx, query, student.Name, student.Age, student.Gender, student.ID)
	if err != nil {
		panic(err)
	}

	return student
}

func (repo *StudentRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Student {
	query := "select id, nama from mahasiswa"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	var students []domain.Student
	for rows.Next() {
		student := domain.Student{}
		err := rows.Scan(&student.ID, &student.Name)
		if err != nil {
			panic(err)
		}
		students = append(students, student)
	}

	return students
}

func (repo *StudentRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Student, error) {
	query := "select id, nama, usia, gender, tanggal_registrasi from mahasiswa where id = ?"
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		panic(err)
	}

	student := domain.Student{}
	if rows.Next() {
		err := rows.Scan(&student.ID, &student.Name, &student.Age, &student.Gender, &student.CreatedAt)
		if err != nil {
			panic(err)
		}
		return student, nil
	} else {
		return student, errors.New("mahasiswa tidak ditemukan")
	}
}

func (repo *StudentRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	query := "DELETE FROM mahasiswa WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		panic(err)
	}
}
