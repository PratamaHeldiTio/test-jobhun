package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"jobhun/app"
	"jobhun/hobby/repository"
	"jobhun/hobby/resthandler"
	"jobhun/hobby/service"
	repository2 "jobhun/major/repository"
	resthandler2 "jobhun/major/resthandler"
	service2 "jobhun/major/service"
	repository3 "jobhun/student/repository"
	resthandler3 "jobhun/student/resthandler"
	service3 "jobhun/student/service"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	//hobby
	hobbyRepository := repository.NewHobbyRepository()
	hobbyService := service.NewHobbyService(hobbyRepository, db, validate)
	hobbyResthandler := resthandler.NewHobbyReshandler(hobbyService)

	//major
	majorRepository := repository2.NewMajorRepository()
	majorService := service2.NewMajorService(majorRepository, db, validate)
	majorResthandler := resthandler2.NewMajorReshandler(majorService)

	// student
	studentRepository := repository3.NewStudentRepositoryImpl()
	studentService := service3.NewStudentService(studentRepository, db, validate)
	studentResthandler := resthandler3.NewStudentResthandler(studentService)
	router := httprouter.New()

	router.POST("/api/v1/hobby", hobbyResthandler.CreateHobby)
	router.POST("/api/v1/major", majorResthandler.CreateMajor)
	router.POST("/api/v1/student", studentResthandler.CreateStudent)
	router.PUT("/api/v1/student/:id", studentResthandler.UpdateStudent)
	router.GET("/api/v1/student", studentResthandler.FindAllStudent)
	router.GET("/api/v1/student/:id", studentResthandler.FindByIDStudent)
	router.DELETE("/api/v1/student/:id", studentResthandler.DeleteStudent)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
