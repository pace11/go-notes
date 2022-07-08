package helper

import (
	"pace/go-rest-api/model/domain"
	"pace/go-rest-api/model/web"
)

func ToNotesResponse(notes domain.Notes) web.NotesResponse {
	return web.NotesResponse{
		Id:              notes.Id,
		Title:           notes.Title,
		Content:         notes.Content,
		BackgroundColor: notes.BackgroundColor,
		CreatedAt:       notes.CreatedAt,
		UpdatedAt:       notes.UpdatedAt,
	}
}

func ToNotesResponses(notes []domain.Notes) []web.NotesResponse {
	var notesResponses []web.NotesResponse
	for _, notes := range notes {
		notesResponses = append(notesResponses, ToNotesResponse(notes))
	}

	return notesResponses
}
