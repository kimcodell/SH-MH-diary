package types

type PostCreateDto struct {
	UserId  int    `json:"userId"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
