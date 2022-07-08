package service

import (
	"context"
	"database/sql"
	"pace/go-rest-api/exception"
	"pace/go-rest-api/helper"
	"pace/go-rest-api/model/domain"
	"pace/go-rest-api/model/web"
	"pace/go-rest-api/repository"

	"github.com/go-playground/validator/v10"
)

type NotesServiceImpl struct {
	NotesRepository repository.NotesRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewNotesService(NotesRepository repository.NotesRepository, DB *sql.DB, validate *validator.Validate) NotesService {
	return &NotesServiceImpl{
		NotesRepository: NotesRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *NotesServiceImpl) Create(ctx context.Context, request web.NotesCreateRequest) web.NotesResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	notes := domain.Notes{
		Title:           request.Title,
		Content:         request.Content,
		BackgroundColor: request.BackgroundColor,
	}

	notes = service.NotesRepository.Save(ctx, tx, notes)

	return helper.ToNotesResponse(notes)
}

func (service *NotesServiceImpl) Update(ctx context.Context, request web.NotesUpdateRequest) web.NotesResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	notes, err := service.NotesRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	notes.Title = request.Title
	notes.Content = request.Content
	notes.BackgroundColor = request.BackgroundColor
	notes.UpdatedAt = request.UpdatedAt

	notes = service.NotesRepository.Update(ctx, tx, notes)

	return helper.ToNotesResponse(notes)
}

func (service *NotesServiceImpl) Delete(ctx context.Context, notesId string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	notes, err := service.NotesRepository.FindById(ctx, tx, notesId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.NotesRepository.Delete(ctx, tx, notes)
}

func (service *NotesServiceImpl) FindById(ctx context.Context, notesId string) web.NotesResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	notes, err := service.NotesRepository.FindById(ctx, tx, notesId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToNotesResponse(notes)
}

func (service *NotesServiceImpl) FindAll(ctx context.Context) []web.NotesResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	notes := service.NotesRepository.FindAll(ctx, tx)
	return helper.ToNotesResponses(notes)
}
