package repository

import (
	"context"
	"database/sql"
	"pace/go-rest-api/model/domain"
)

type NotesRepository interface {
	Save(ctx context.Context, tx *sql.Tx, notes domain.Notes) domain.Notes
	Update(ctx context.Context, tx *sql.Tx, notes domain.Notes) domain.Notes
	Delete(ctx context.Context, tx *sql.Tx, notes domain.Notes)
	FindById(ctx context.Context, tx *sql.Tx, notesId string) (domain.Notes, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Notes
}
