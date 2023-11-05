package model

import (
	"html/template"
)

type Article struct {
	ID      uint
	Title   string
	Content template.HTML
}

type HomePageData struct {
	Articles    []Article
	HasPrevPage bool
	PrevPage    int
	HasNextPage bool
	NextPage    int
	Empty       bool
}
