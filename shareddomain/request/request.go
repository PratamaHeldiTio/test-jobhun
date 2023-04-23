package request

type HobbyCreateRequest struct {
	Name string `validate:"required,max=100"`
}

type MajorCreateRequest struct {
	Name string `validate:"required,max=100"`
}

type StudentRequest struct {
	ID      int
	Name    string `json:"nama" validate:"required,max=100"`
	Age     int    `json:"usia" validate:"required"`
	Gender  int    `json:"gender" validate:"gte=0,lte=1"`
	Hobbies []int  `json:"hobi"`
}
