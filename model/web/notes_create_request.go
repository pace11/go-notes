package web

type NotesCreateRequest struct {
	Title           *string `validate:"required,max=100,min=1" json:"title"`
	Content         *string `validate:"required" json:"content"`
	BackgroundColor *string `json:"background_color"`
}
