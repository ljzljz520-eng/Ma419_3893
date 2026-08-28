package service

import (
	"errors"
	"heritage/model"
)

func ValidateCollection(c model.Collection) error {
	if c.ID == "" {
		return errors.New("missing id")
	}
	if c.Name == "" {
		return errors.New("missing name")
	}
	if len(c.ArticleIDs) == 0 {
		return errors.New("empty collection")
	}
	return nil
}
func ValidateArticle(a model.Article) error {
	if a.ID == "" {
		return errors.New("missing id")
	}
	if a.Title == "" {
		return errors.New("missing title")
	}
	if a.Category == "" {
		return errors.New("missing category")
	}
	if a.ReadingMinutes <= 0 {
		return errors.New("invalid reading time")
	}
	return nil
}
func NormalizeArticle(a model.Article) model.Article {
	a.Title = model.TruncateTitle(a.Title, 120)
	a.ReadingMinutes = model.SafeMinutes(a.ReadingMinutes)
	return a
}
func CanPublish(a model.Article) bool { return ValidateArticle(a) == nil && a.Cover != "" }
func CanFeature(a model.Article) bool { return a.Valid() && a.ReadingMinutes >= 5 }
func MergeTags(existing []model.Tag, incoming []model.Tag) []model.Tag {
	out := append([]model.Tag(nil), existing...)
	for _, t := range incoming {
		found := false
		for i := range out {
			if out[i].ID == t.ID {
				out[i] = t
				found = true
			}
		}
		if !found {
			out = append(out, t)
		}
	}
	return out
}
func RemoveTag(tags []model.Tag, id string) []model.Tag {
	out := []model.Tag{}
	for _, t := range tags {
		if t.ID != id {
			out = append(out, t)
		}
	}
	return out
}
func CollectionContains(c model.Collection, id string) bool {
	return model.ContainsID(c.ArticleIDs, id)
}
func AddToCollection(c model.Collection, id string) model.Collection {
	c.ArticleIDs = model.AppendUnique(c.ArticleIDs, id)
	return c
}
func RemoveFromCollection(c model.Collection, id string) model.Collection {
	c.ArticleIDs = model.RemoveID(c.ArticleIDs, id)
	return c
}
func Publish(c model.Collection) model.Collection   { c.Published = true; return c }
func Unpublish(c model.Collection) model.Collection { c.Published = false; return c }
func Feature(a model.Article) model.Article         { a.Featured = true; return a }
func Unfeature(a model.Article) model.Article       { a.Featured = false; return a }
func ArticleSummary(a model.Article) string         { return a.Summary(140) }
func CollectionSummary(c model.Collection) string   { return c.Name + "：" + c.Description }
func EnsureCover(a model.Article) model.Article {
	if a.Cover == "" {
		a.Cover = "cover://" + a.ID
	}
	return a
}
func PrepareForSave(a model.Article) model.Article { return EnsureCover(NormalizeArticle(a)) }
