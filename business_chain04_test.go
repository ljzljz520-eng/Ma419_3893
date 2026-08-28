package main

import (
	"heritage/model"
	"heritage/service"
	"heritage/storage"
	"testing"
)

func TestBusinessChain04(t *testing.T) {
	s, e := storage.Open(":memory:")
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	v := service.New(s)
	first := []string{"a", "b", "c"}
	second := v.CopyArticleIDs(first)
	second[0] = "changed"
	if first[0] == "changed" {
		t.Fatalf("lists share backing storage: %v", first)
	}
	_ = model.Collection{ID: "c", Name: "n", ArticleIDs: second}
}
