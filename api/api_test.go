package api

import (
	"heritage/content"
	"heritage/service"
	"heritage/storage"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPHome(t *testing.T) {
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
	r := httptest.NewRequest(http.MethodGet, "/api/home", nil)
	w := httptest.NewRecorder()
	New(v).Handler().ServeHTTP(w, r)
	if w.Code != 200 {
		t.Fatal(w.Code)
	}
}
