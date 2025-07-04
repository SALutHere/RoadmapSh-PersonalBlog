package models

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	Id        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Text      string
}

func NewArticle(title, text string) Article {
	return Article{
		Id:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title:     title,
		Text:      text,
	}
}
