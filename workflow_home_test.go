package main

import (
	"heritage/content"
	"heritage/service"
	"heritage/storage"
	"testing"
)

func TestWorkflowHomepage(t *testing.T) {
	s, e := storage.Open(":memory:")
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	v := service.New(s)
	for _, c := range content.SeedCategories() {
		_ = s.SaveCategory(c)
	}
	for _, a := range content.SeedArticles() {
		_ = v.PublishArticle(a)
	}
	for _, c := range content.SeedCollections() {
		_ = v.PublishCollection(c)
	}
	h, e := v.BuildHomepage("")
	if e != nil || h.Featured.ID != "a-interview" || len(h.Articles) < 5 || len(h.Collections) < 2 {
		t.Fatal("homepage")
	}
}
