package app

import (
	"fmt"
	"net/http"
	"pace/go-rest-api/app"
	"pace/go-rest-api/controller"
	"pace/go-rest-api/helper"
	"pace/go-rest-api/middleware"
	"pace/go-rest-api/repository"
	"pace/go-rest-api/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func Start() {

	db := app.NewDB()
	validate := validator.New()
	notesRepository := repository.NewNotesRepository()
	notesService := service.NewNotesService(notesRepository, db, validate)
	notesController := controller.NewNotesController(notesService)

	router := httprouter.New()

	router.GET("/api/notes", notesController.FindAll)
	router.GET("/api/notes/:notesId", notesController.FindById)
	router.POST("/api/notes", notesController.Create)
	router.PUT("/api/notes/:notesId", notesController.Update)
	router.DELETE("/api/notes/:notesId", notesController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	fmt.Println("Running api in", server.Addr)

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
