package service

import (
	"context"
	"jobhun/shareddomain/request"
	"jobhun/shareddomain/response"
)

type HobbyService interface {
	CreateHobby(ctx context.Context, request request.HobbyCreateRequest) response.ResponseHobby
}
