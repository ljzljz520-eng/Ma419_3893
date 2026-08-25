package api

import (
	"net/http"
	"strings"
)

func MethodAllowed(r *http.Request, methods ...string) bool {
	for _, m := range methods {
		if r.Method == m {
			return true
		}
	}
	return false
}
func QueryCategory(r *http.Request) string { return strings.TrimSpace(r.URL.Query().Get("category")) }
func IsJSON(r *http.Request) bool {
	return strings.Contains(r.Header.Get("Content-Type"), "application/json")
}
func StatusForError(err error) int {
	if err == nil {
		return 200
	}
	return 400
}
func RequestID(r *http.Request) string { return r.Header.Get("X-Request-ID") }
func LimitParam(r *http.Request, def int) int {
	v := r.URL.Query().Get("limit")
	if v == "" {
		return def
	}
	n := 0
	for _, c := range v {
		if c < '0' || c > '9' {
			return def
		}
		n = n*10 + int(c-'0')
	}
	if n < 1 {
		return def
	}
	if n > 100 {
		return 100
	}
	return n
}
func PageParam(r *http.Request) int {
	v := r.URL.Query().Get("page")
	if v == "" {
		return 1
	}
	n := 0
	for _, c := range v {
		if c < '0' || c > '9' {
			return 1
		}
		n = n*10 + int(c-'0')
	}
	if n < 1 {
		return 1
	}
	return n
}
func HeaderValue(r *http.Request, key string) string { return strings.TrimSpace(r.Header.Get(key)) }
func IsHealthPath(r *http.Request) bool              { return r.URL.Path == "/healthz" }
func IsHomePath(r *http.Request) bool                { return r.URL.Path == "/" || r.URL.Path == "/api/home" }
func IsArticlePath(r *http.Request) bool             { return strings.HasPrefix(r.URL.Path, "/api/articles") }
func ContentTypeJSON(w http.ResponseWriter)          { w.Header().Set("Content-Type", "application/json") }
func NoCache(w http.ResponseWriter)                  { w.Header().Set("Cache-Control", "no-store") }
func AllowMethods(w http.ResponseWriter, methods ...string) {
	w.Header().Set("Allow", strings.Join(methods, ", "))
}
func IsEmptyBody(body []byte) bool { return len(strings.TrimSpace(string(body))) == 0 }
func NormalizePath(path string) string {
	if path == "" {
		return "/"
	}
	if path[0] != '/' {
		return "/" + path
	}
	return path
}
