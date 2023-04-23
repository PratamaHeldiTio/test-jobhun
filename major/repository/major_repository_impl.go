package repository

import (
	"context"
	"database/sql"
	"jobhun/major/domain"
)

type MajorRepositoryImpl struct{}

func NewMajorRepository() MajorRepository {
	return &MajorRepositoryImpl{}
}

func (repo *MajorRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, major domain.Major) domain.Major {
	query := "insert into jurusan(nama_jurusan) values (?)"
	result, err := tx.ExecContext(ctx, query, major.Name)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	major.ID = int(id)

	return major
}
