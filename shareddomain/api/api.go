package api

type APIResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func ToApiResponse(code int, status string, data interface{}) APIResponse {
	return APIResponse{
		Code:   code,
		Status: status,
		Data:   data,
	}
}
