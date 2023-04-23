package resthandler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type HobbyResthandler interface {
	CreateHobby(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
