package api

import (
	"encoding/json"
	"heritage/model"
	"heritage/service"
	"net/http"
)

type Server struct{ S *service.Service }

func New(s *service.Service) *Server { return &Server{S: s} }
func (h *Server) Handler() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/", h.home)
	m.HandleFunc("/api/home", h.home)
	m.HandleFunc("/api/articles", h.articles)
	return m
}
func (h *Server) home(w http.ResponseWriter, r *http.Request) {
	cat := r.URL.Query().Get("category")
	v, e := h.S.BuildHomepage(cat)
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
func (h *Server) articles(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method", 405)
		return
	}
	var v struct {
		ID, Title, Category, Photographer string
		ReadingMinutes                    int
		Body                              string
	}
	if e := json.NewDecoder(r.Body).Decode(&v); e != nil {
		http.Error(w, e.Error(), 400)
		return
	}
	e := h.S.PublishArticle(modelArticle(v))
	if e != nil {
		http.Error(w, e.Error(), 400)
		return
	}
	w.WriteHeader(201)
}
func modelArticle(v struct {
	ID, Title, Category, Photographer string
	ReadingMinutes                    int
	Body                              string
}) model.Article {
	return model.Article{ID: v.ID, Title: v.Title, Category: v.Category, Photographer: v.Photographer, ReadingMinutes: v.ReadingMinutes, Body: v.Body, Cover: "cover://" + v.ID}
}
