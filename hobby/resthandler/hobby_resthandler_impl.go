package resthandler

import (
	"github.com/julienschmidt/httprouter"
	"jobhun/helper"
	"jobhun/hobby/service"
	"jobhun/shareddomain/api"
	request2 "jobhun/shareddomain/request"
	"net/http"
)

type HobbyResthandlerImpl struct {
	service service.HobbyService
}

func NewHobbyReshandler(service service.HobbyService) HobbyResthandler {
	return &HobbyResthandlerImpl{service}
}

func (handler *HobbyResthandlerImpl) CreateHobby(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	hobbyRequest := request2.HobbyCreateRequest{}
	helper.ReadRequestBody(request, &hobbyRequest)

	// cal service
	hobbyResponse := handler.service.CreateHobby(request.Context(), hobbyRequest)
	response := api.ToApiResponse(201, "Created", hobbyResponse)

	helper.WriteToResponseBody(writer, response)
}
