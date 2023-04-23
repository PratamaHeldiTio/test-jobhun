package repository

import (
	"context"
	"database/sql"
	"jobhun/hobby/domain"
)

type HobbyRepository interface {
	Create(ctx context.Context, tx *sql.Tx, hobby domain.Hobby) domain.Hobby
}
