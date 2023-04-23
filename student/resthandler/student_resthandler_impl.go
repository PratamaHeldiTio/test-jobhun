package resthandler

import (
	"github.com/julienschmidt/httprouter"
	"jobhun/helper"
	"jobhun/shareddomain/api"
	request2 "jobhun/shareddomain/request"
	"jobhun/student/service"
	"net/http"
	"strconv"
)

type StudentResthandlerImpl struct {
	service service.StudentService
}

func NewStudentResthandler(service service.StudentService) StudentResthandler {
	return &StudentResthandlerImpl{service}
}

func (handler *StudentResthandlerImpl) CreateStudent(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	student := request2.StudentRequest{}
	helper.ReadRequestBody(request, &student)

	studentResponse := handler.service.CreateStudent(request.Context(), student)
	responseAPI := api.ToApiResponse(201, "Created", studentResponse)
	helper.WriteToResponseBody(writer, responseAPI)
}

func (handler *StudentResthandlerImpl) UpdateStudent(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	student := request2.StudentRequest{}
	helper.ReadRequestBody(request, &student)

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		panic(err)
	}

	student.ID = id

	studentResponse := handler.service.UpdateStudent(request.Context(), student)
	responseAPI := api.ToApiResponse(200, "OK", studentResponse)

	helper.WriteToResponseBody(writer, responseAPI)
}

func (handler *StudentResthandlerImpl) FindAllStudent(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentResponses := handler.service.FindAllStudent(request.Context())
	responseAPI := api.ToApiResponse(200, "OK", studentResponses)

	helper.WriteToResponseBody(writer, responseAPI)
}

func (handler *StudentResthandlerImpl) FindByIDStudent(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		panic(err)
	}

	studentResponse := handler.service.FindByIDStudent(request.Context(), id)
	responseAPI := api.ToApiResponse(200, "OK", studentResponse)

	helper.WriteToResponseBody(writer, responseAPI)
}

func (handler *StudentResthandlerImpl) DeleteStudent(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		panic(err)
	}

	handler.service.DeleteStudent(request.Context(), id)
	responseAPI := api.ToApiResponse(200, "OK", nil)

	helper.WriteToResponseBody(writer, responseAPI)
}
