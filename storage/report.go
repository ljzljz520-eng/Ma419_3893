package storage

import (
	"context"
	"heritage/model"
)

func (s *Store) SaveAll(ctx context.Context, articles []model.Article, collections []model.Collection, categories []model.Category, tags []model.Tag) error {
	for _, c := range categories {
		if err := s.SaveCategory(c); err != nil {
			return err
		}
	}
	for _, a := range articles {
		if err := s.SaveArticle(a); err != nil {
			return err
		}
	}
	for _, c := range collections {
		if err := s.SaveCollection(c); err != nil {
			return err
		}
	}
	for _, t := range tags {
		if err := s.SaveTag(t); err != nil {
			return err
		}
	}
	return nil
}
func (s *Store) ArticleExists(id string) bool { _, e := s.GetArticle(id); return e == nil }
func (s *Store) CollectionExists(id string) bool {
	cs, e := s.ListCollections()
	if e != nil {
		return false
	}
	for _, c := range cs {
		if c.ID == id {
			return true
		}
	}
	return false
}
func (s *Store) CategoryExists(id string) bool {
	cs, e := s.ListCategories()
	if e != nil {
		return false
	}
	for _, c := range cs {
		if c.ID == id {
			return true
		}
	}
	return false
}
func (s *Store) ArticleCount() int    { n, _ := s.Count("articles"); return n }
func (s *Store) CollectionCount() int { n, _ := s.Count("collections"); return n }
func (s *Store) CategoryCount() int   { n, _ := s.Count("categories"); return n }
func (s *Store) TagCount() int        { n, _ := s.Count("tags"); return n }
func (s *Store) DeleteArticle(id string) error {
	_, e := s.db.Exec("DELETE FROM articles WHERE id=?", id)
	return e
}
func (s *Store) DeleteCollection(id string) error {
	_, e := s.db.Exec("DELETE FROM collections WHERE id=?", id)
	return e
}
func (s *Store) DeleteCategory(id string) error {
	_, e := s.db.Exec("DELETE FROM categories WHERE id=?", id)
	return e
}
func (s *Store) DeleteTag(id string) error {
	_, e := s.db.Exec("DELETE FROM tags WHERE id=?", id)
	return e
}
func (s *Store) Backup(ctx context.Context) error {
	_, e := s.db.ExecContext(ctx, "PRAGMA wal_checkpoint")
	return e
}
