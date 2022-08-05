package types

type PostCreateDto struct {
	UserId  int    `json:"userId" binding: "required"`
	Title   string `json:"title" binding: "required"`
	Content string `json:"content" binding: "required"`
}
