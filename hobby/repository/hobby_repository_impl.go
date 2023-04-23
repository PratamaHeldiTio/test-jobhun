package repository

import (
	"context"
	"database/sql"
	"jobhun/hobby/domain"
)

type HobbyRepositoryImpl struct {
}

func NewHobbyRepository() HobbyRepository {
	return &HobbyRepositoryImpl{}
}

func (repo *HobbyRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, hobby domain.Hobby) domain.Hobby {
	query := "insert into hobi(nama_hobi) values (?)"
	result, err := tx.ExecContext(ctx, query, hobby.Name)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	hobby.ID = int(id)

	return hobby
}
