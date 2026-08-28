package storage

import (
	"context"
	"database/sql"
	"encoding/json"
	"heritage/model"
	_ "modernc.org/sqlite"
)

type Store struct{ db *sql.DB }

func Open(path string) (*Store, error) {
	db, e := sql.Open("sqlite", path)
	if e != nil {
		return nil, e
	}
	s := &Store{db}
	if e = s.migrate(); e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) migrate() error {
	_, e := s.db.Exec(`CREATE TABLE IF NOT EXISTS articles(id TEXT PRIMARY KEY,title TEXT,category TEXT,photographer TEXT,minutes INTEGER,cover TEXT,body TEXT,featured INTEGER); CREATE TABLE IF NOT EXISTS collections(id TEXT PRIMARY KEY,name TEXT,description TEXT,article_ids TEXT,published INTEGER); CREATE TABLE IF NOT EXISTS categories(id TEXT PRIMARY KEY,name TEXT,slug TEXT); CREATE TABLE IF NOT EXISTS tags(id TEXT PRIMARY KEY,article_id TEXT,label TEXT)`)
	return e
}
func (s *Store) Close() error                   { return s.db.Close() }
func (s *Store) Ping(ctx context.Context) error { return s.db.PingContext(ctx) }
func (s *Store) SaveArticle(a model.Article) error {
	_, e := s.db.Exec(`INSERT INTO articles VALUES(?,?,?,?,?,?,?,?) ON CONFLICT(id) DO UPDATE SET title=excluded.title,category=excluded.category,photographer=excluded.photographer,minutes=excluded.minutes,cover=excluded.cover,body=excluded.body,featured=excluded.featured`, a.ID, a.Title, a.Category, a.Photographer, a.ReadingMinutes, a.Cover, a.Body, boolInt(a.Featured))
	return e
}
func (s *Store) GetArticle(id string) (model.Article, error) {
	var a model.Article
	var f int
	e := s.db.QueryRow(`SELECT id,title,category,photographer,minutes,cover,body,featured FROM articles WHERE id=?`, id).Scan(&a.ID, &a.Title, &a.Category, &a.Photographer, &a.ReadingMinutes, &a.Cover, &a.Body, &f)
	a.Featured = f == 1
	return a, e
}
func (s *Store) ListArticles(category string) ([]model.Article, error) {
	q := `SELECT id,title,category,photographer,minutes,cover,body,featured FROM articles`
	args := []any{}
	if category != "" {
		q += ` WHERE category=?`
		args = append(args, category)
	}
	q += ` ORDER BY id`
	rows, e := s.db.Query(q, args...)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []model.Article{}
	for rows.Next() {
		var a model.Article
		var f int
		if e = rows.Scan(&a.ID, &a.Title, &a.Category, &a.Photographer, &a.ReadingMinutes, &a.Cover, &a.Body, &f); e != nil {
			return nil, e
		}
		a.Featured = f == 1
		out = append(out, a)
	}
	return out, rows.Err()
}
func (s *Store) SaveCollection(c model.Collection) error {
	b, _ := json.Marshal(c.ArticleIDs)
	_, e := s.db.Exec(`INSERT INTO collections VALUES(?,?,?,?,?) ON CONFLICT(id) DO UPDATE SET name=excluded.name,description=excluded.description,article_ids=excluded.article_ids,published=excluded.published`, c.ID, c.Name, c.Description, string(b), boolInt(c.Published))
	return e
}
func (s *Store) ListCollections() ([]model.Collection, error) {
	rows, e := s.db.Query(`SELECT id,name,description,article_ids,published FROM collections ORDER BY id`)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []model.Collection{}
	for rows.Next() {
		var c model.Collection
		var b string
		var p int
		if e = rows.Scan(&c.ID, &c.Name, &c.Description, &b, &p); e != nil {
			return nil, e
		}
		_ = json.Unmarshal([]byte(b), &c.ArticleIDs)
		c.Published = p == 1
		out = append(out, c)
	}
	return out, rows.Err()
}
func (s *Store) SaveCategory(c model.Category) error {
	_, e := s.db.Exec(`INSERT INTO categories VALUES(?,?,?) ON CONFLICT(id) DO UPDATE SET name=excluded.name,slug=excluded.slug`, c.ID, c.Name, c.Slug)
	return e
}
func (s *Store) ListCategories() ([]model.Category, error) {
	rows, e := s.db.Query(`SELECT id,name,slug FROM categories ORDER BY id`)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []model.Category{}
	for rows.Next() {
		var c model.Category
		if e = rows.Scan(&c.ID, &c.Name, &c.Slug); e != nil {
			return nil, e
		}
		out = append(out, c)
	}
	return out, rows.Err()
}
func (s *Store) SaveTag(t model.Tag) error {
	_, e := s.db.Exec(`INSERT INTO tags VALUES(?,?,?) ON CONFLICT(id) DO UPDATE SET article_id=excluded.article_id,label=excluded.label`, t.ID, t.ArticleID, t.Label)
	return e
}
func (s *Store) Count(table string) (int, error) {
	var n int
	if table != "articles" && table != "collections" && table != "categories" && table != "tags" {
		return 0, sql.ErrNoRows
	}
	e := s.db.QueryRow(`SELECT count(*) FROM ` + table).Scan(&n)
	return n, e
}
func boolInt(v bool) int {
	if v {
		return 1
	}
	return 0
}
