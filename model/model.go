package model

import "strings"

type Article struct {
	ID, Title, Category, Photographer, Cover, Body string
	ReadingMinutes                                 int
	Featured                                       bool
}
type Collection struct {
	ID, Name, Description string
	ArticleIDs            []string
	Published             bool
}
type Category struct{ ID, Name, Slug string }
type Tag struct{ ID, ArticleID, Label string }
type Homepage struct {
	Featured    Article
	Articles    []Article
	Collections []Collection
	Categories  []Category
}

func (a Article) Valid() bool {
	return a.ID != "" && strings.TrimSpace(a.Title) != "" && a.ReadingMinutes > 0
}
func (a Article) Summary(n int) string {
	if n < 1 {
		return ""
	}
	if len(a.Body) <= n {
		return a.Body
	}
	return a.Body[:n] + "..."
}
func (c Collection) Valid() bool { return c.ID != "" && c.Name != "" }
func (c Collection) Clone() Collection {
	ids := append([]string(nil), c.ArticleIDs...)
	c.ArticleIDs = ids
	return c
}
func NormalizeCategory(s string) string { return strings.ToLower(strings.TrimSpace(s)) }
func NewArticle(id, title, cat, photographer string, minutes int) Article {
	return Article{ID: id, Title: title, Category: cat, Photographer: photographer, ReadingMinutes: minutes, Cover: "cover://" + id}
}
