// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Blog struct {
	ID        int     `json:"id"`
	Slug      string  `json:"slug"`
	Title     string  `json:"title"`
	Abstruct  *string `json:"abstruct"`
	Content   string  `json:"content"`
	UserID    string  `json:"userId"`
	IsPublic  bool    `json:"isPublic"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"UpdatedAt"`
}

type NewBlog struct {
	Abstruct *string `json:"abstruct"`
	Content  string  `json:"content"`
	UserID   string  `json:"userId"`
	IsPublic bool    `json:"isPublic"`
}
