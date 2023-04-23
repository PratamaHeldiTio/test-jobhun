package resthandler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type StudentResthandler interface {
	CreateStudent(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateStudent(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllStudent(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByIDStudent(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteStudent(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
