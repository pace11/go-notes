package web

type NotesUpdateRequest struct {
	Id              string  `validate:"required"`
	Title           *string `validate:"required,max=100,min=1" json:"title"`
	Content         *string `json:"content"`
	BackgroundColor *string `json:"background_color"`
	UpdatedAt       string  `json:"updated_at"`
}
