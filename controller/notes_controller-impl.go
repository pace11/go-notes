package controller

import (
	"encoding/json"
	"net/http"
	"pace/go-rest-api/helper"
	"pace/go-rest-api/model/web"
	"pace/go-rest-api/service"

	"github.com/julienschmidt/httprouter"
)

type NotesControllerImpl struct {
	NotesService service.NotesService
}

func NewNotesController(notesService service.NotesService) NotesController {
	return &NotesControllerImpl{
		NotesService: notesService,
	}
}

func (controller *NotesControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)

	notesCreateRequest := web.NotesCreateRequest{}
	err := decoder.Decode(&notesCreateRequest)
	helper.PanicIfError(err)

	notesResponse := controller.NotesService.Create(request.Context(), notesCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   notesResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)

}

func (controller *NotesControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)

	notesUpdateRequest := web.NotesUpdateRequest{}
	err := decoder.Decode(&notesUpdateRequest)
	helper.PanicIfError(err)

	notesId := params.ByName("notesId")
	notesUpdateRequest.Id = notesId

	categoryResponse := controller.NotesService.Update(request.Context(), notesUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *NotesControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	notesId := params.ByName("notesId")

	controller.NotesService.Delete(request.Context(), notesId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *NotesControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	notesId := params.ByName("notesId")

	notesResponnse := controller.NotesService.FindById(request.Context(), notesId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   notesResponnse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
func (controller *NotesControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	notesResponses := controller.NotesService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   notesResponses,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
