package system

import (
	"time"

	"github.com/codebdy/entify-core/model/meta"
	"github.com/codebdy/entify-core/shared"
)

type Meta struct {
	Id               shared.ID    `json:"id"`
	Name             string       `json:"name"`
	Content          meta.UMLMeta `json:"content"`
	PublishedContent meta.UMLMeta `json:"publishedContent"`
	PublishedAt      time.Time    `json:"publishedAt"`
	CreatedAt        time.Time    `json:"createdAt"`
	UpdatedAt        time.Time    `json:"updatedAt"`
}

type App struct {
	Id        shared.ID `json:"id"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	ImageUrl  string    `json:"imageUrl"`
	MetaId    shared.ID `json:"metaId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Service struct {
	Id        shared.ID `json:"id"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	ImageUrl  string    `json:"imageUrl"`
	MetaId    shared.ID `json:"metaId"`
	IsSystem  bool      `json:"isSystem"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
