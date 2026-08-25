package main

import (
	"heritage/model"
	"heritage/service"
	"heritage/storage"
	"testing"
)

func TestWorkflowEditor(t *testing.T) {
	s, e := storage.Open(":memory:")
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	v := service.New(s)
	if e = v.PublishCollection(model.Collection{ID: "c", Name: "专题", ArticleIDs: []string{"a", "b"}}); e != nil {
		t.Fatal(e)
	}
	if e = v.EditCollection("c", []string{"b", "a"}); e != nil {
		t.Fatal(e)
	}
}
