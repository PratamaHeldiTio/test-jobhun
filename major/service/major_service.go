package service

import (
	"context"
	"jobhun/shareddomain/request"
	"jobhun/shareddomain/response"
)

type MajorService interface {
	CreateMajor(ctx context.Context, request request.MajorCreateRequest) response.ResponseMajor
}
