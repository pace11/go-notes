package service

import (
	"context"
	"pace/go-rest-api/model/web"
)

type NotesService interface {
	Create(ctx context.Context, request web.NotesCreateRequest) web.NotesResponse
	Update(ctx context.Context, request web.NotesUpdateRequest) web.NotesResponse
	Delete(ctx context.Context, notesId string)
	FindById(ctx context.Context, notesId string) web.NotesResponse
	FindAll(ctx context.Context) []web.NotesResponse
}
