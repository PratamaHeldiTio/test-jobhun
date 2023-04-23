package repository

import (
	"context"
	"database/sql"
	"jobhun/major/domain"
)

type MajorRepository interface {
	Create(ctx context.Context, tx *sql.Tx, major domain.Major) domain.Major
}
