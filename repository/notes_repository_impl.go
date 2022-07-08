package repository

import (
	"context"
	"database/sql"
	"errors"
	"pace/go-rest-api/helper"
	"pace/go-rest-api/model/domain"

	"github.com/google/uuid"
)

type NotesRepositoryImpl struct {
}

func NewNotesRepository() NotesRepository {
	return &NotesRepositoryImpl{}
}

func (repository *NotesRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, notes domain.Notes) domain.Notes {
	SQL := "INSERT INTO notes(id, title, content, background_color) values (?, ?, ?, IFNULL(?, NULL))"
	uuid := uuid.New()
	_, err := tx.ExecContext(ctx, SQL, uuid, notes.Title, notes.Content, notes.BackgroundColor)
	helper.PanicIfError(err)

	notes.Id = uuid.String()
	return notes
}

func (repository *NotesRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, notes domain.Notes) domain.Notes {
	SQL := "UPDATE notes SET title = IFNULL(?, NULL), content = IFNULL(?, NULL), background_color = IFNULL(?, NULL), updated_at = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, notes.Title, notes.Content, notes.BackgroundColor, notes.UpdatedAt, notes.Id)
	helper.PanicIfError(err)

	return notes
}

func (repository *NotesRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, notes domain.Notes) {
	SQL := "DELETE FROM notes WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, notes.Id)
	helper.PanicIfError(err)
}

func (repository *NotesRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, notesId string) (domain.Notes, error) {
	SQL := "SELECT id, title, content, background_color, created_at, updated_at FROM notes WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, notesId)
	helper.PanicIfError(err)
	defer rows.Close()

	notes := domain.Notes{}
	if rows.Next() {
		err := rows.Scan(&notes.Id, &notes.Title, &notes.Content, &notes.BackgroundColor, &notes.CreatedAt, &notes.UpdatedAt)
		helper.PanicIfError(err)
		return notes, nil
	} else {
		return notes, errors.New("notes not found")
	}

}

func (repository *NotesRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Notes {
	SQL := "SELECT id, title, content, background_color, created_at, updated_at FROM notes ORDER BY updated_at DESC"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var notes []domain.Notes
	for rows.Next() {
		note := domain.Notes{}
		err := rows.Scan(&note.Id, &note.Title, &note.Content, &note.BackgroundColor, &note.CreatedAt, &note.UpdatedAt)
		helper.PanicIfError(err)
		notes = append(notes, note)
	}

	return notes
}
