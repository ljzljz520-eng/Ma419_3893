package main

import (
	"heritage/model"
	"heritage/storage"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := t.TempDir() + "/persist.db"
	s, e := storage.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	if e = s.SaveArticle(model.Article{ID: "persist", Title: "持久", Category: "剪纸", ReadingMinutes: 5, Cover: "c"}); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = storage.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	a, e := s.GetArticle("persist")
	if e != nil || a.Title != "持久" {
		t.Fatal("reopen")
	}
}
