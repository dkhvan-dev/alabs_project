package blogs

import (
	"blogs-api/internal/users/model"
)

type Blog struct {
	Id      int        `json:"id"`
	Title   string     `json:"title"`
	Content string     `json:"content"`
	Author  model.User `json:"author"`
}

// DTOs
type BlogCreate struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type BlogUpdate struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type BlogView struct {
	Id      int            `json:"id"`
	Title   string         `json:"title"`
	Content string         `json:"content"`
	Author  model.UserView `json:"author"`
}
