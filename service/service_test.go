package service

import (
	"heritage/content"
	"heritage/model"
	"heritage/storage"
	"testing"
)

func TestServicePublish(t *testing.T) {
	s, e := storage.Open(":memory:")
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	v := New(s)
	if e = v.PublishArticle(content.SeedArticles()[0]); e != nil {
		t.Fatal(e)
	}
	if e = v.PublishCollection(model.Collection{ID: "c", Name: "n", ArticleIDs: []string{"a-interview"}}); e != nil {
		t.Fatal(e)
	}
}
