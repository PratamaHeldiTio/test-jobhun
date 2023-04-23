package resthandler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type MajorResthandler interface {
	CreateMajor(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
