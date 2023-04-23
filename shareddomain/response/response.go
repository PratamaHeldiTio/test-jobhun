package response

import (
	"jobhun/hobby/domain"
	domain2 "jobhun/major/domain"
	domain3 "jobhun/student/domain"
	"time"
)

type ResponseHobby struct {
	ID   int    `json:"id"`
	Name string `json:"nama"`
}

func ToResponseHobby(hobby domain.Hobby) ResponseHobby {
	return ResponseHobby{
		ID:   hobby.ID,
		Name: hobby.Name,
	}
}

type ResponseMajor struct {
	ID   int    `json:"id"`
	Name string `json:"nama"`
}

type StudentResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"nama"`
	Age       int       `json:"usia"`
	Gender    int       `json:"gender"`
	CreatedAt time.Time `json:"tanggal_registrasi"`
	Hobbies   []int     `json:"hobi"`
}

type StudentResponseUpdate struct {
	ID     int    `json:"id"`
	Name   string `json:"nama"`
	Age    int    `json:"usia"`
	Gender int    `json:"gender"`
}

type StudentResponseBasic struct {
	ID   int    `json:"id"`
	Name string `json:"nama"`
}

func ToResponseMajor(major domain2.Major) ResponseMajor {
	return ResponseMajor{
		ID:   major.ID,
		Name: major.Name,
	}
}

func ToResponseStudent(student domain3.Student) StudentResponse {
	return StudentResponse{
		ID:        student.ID,
		Name:      student.Name,
		Age:       student.Age,
		Gender:    student.Gender,
		CreatedAt: student.CreatedAt,
		Hobbies:   student.Hobbies,
	}
}

func ToResponseUpdateStudent(student domain3.Student) StudentResponseUpdate {
	return StudentResponseUpdate{
		ID:     student.ID,
		Name:   student.Name,
		Age:    student.Age,
		Gender: student.Gender,
	}
}

func ToResponseStudentBasic(student domain3.Student) StudentResponseBasic {
	return StudentResponseBasic{
		ID:   student.ID,
		Name: student.Name,
	}
}
