package types

type SimplePost struct {
	Id        int
	Title     string
	CreatedAt string
}

type Post struct {
	SimplePost
	Content   string
	Author    string
	UpdatedAt string
}

type Posts []SimplePost
