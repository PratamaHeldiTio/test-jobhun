package resthandler

import (
	"github.com/julienschmidt/httprouter"
	"jobhun/helper"
	"jobhun/major/service"
	"jobhun/shareddomain/api"
	request2 "jobhun/shareddomain/request"
	"net/http"
)

type MajorResthandlerImpl struct {
	service service.MajorService
}

func NewMajorReshandler(service service.MajorService) MajorResthandler {
	return &MajorResthandlerImpl{service}
}

func (handler *MajorResthandlerImpl) CreateMajor(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	MajorRequest := request2.MajorCreateRequest{}
	helper.ReadRequestBody(request, &MajorRequest)

	// cal service
	MajorResponse := handler.service.CreateMajor(request.Context(), MajorRequest)
	response := api.ToApiResponse(201, "Created", MajorResponse)

	helper.WriteToResponseBody(writer, response)
}
