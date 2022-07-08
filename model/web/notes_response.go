package web

type NotesResponse struct {
	Id              string  `json:"id"`
	Title           *string `json:"title"`
	Content         *string `json:"content"`
	BackgroundColor *string `json:"background_color"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}
