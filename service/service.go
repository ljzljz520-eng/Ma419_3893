package service

import (
	"errors"
	"heritage/model"
	"heritage/query"
	"heritage/storage"
)

type Service struct{ Store *storage.Store }

func New(s *storage.Store) *Service { return &Service{Store: s} }
func (s *Service) Seed() error {
	for _, c := range []model.Category{{"cat-paper", "剪纸", "paper"}, {"cat-wood", "木版画", "wood"}, {"cat-costume", "戏曲服饰", "costume"}, {"cat-instrument", "地方乐器", "instrument"}} {
		if e := s.Store.SaveCategory(c); e != nil {
			return e
		}
	}
	return nil
}
func (s *Service) PublishArticle(a model.Article) error {
	if !a.Valid() {
		return errors.New("invalid article")
	}
	return s.Store.SaveArticle(a)
}
func (s *Service) PublishCollection(c model.Collection) error {
	if !c.Valid() {
		return errors.New("invalid collection")
	}
	return s.Store.SaveCollection(c)
}
// CopyArticleIDs returns an independent copy of the article-ID list so that a
// later edit to the draft (reordering, adding, removing) never rewrites the
// previous list's contents. Returning the slice directly would alias the
// caller's backing array and let one business list overwrite another.
func (s *Service) CopyArticleIDs(ids []string) []string {
	return append([]string(nil), ids...)
}
func (s *Service) BuildHomepage(category string) (model.Homepage, error) {
	as, e := s.Store.ListArticles(category)
	if e != nil {
		return model.Homepage{}, e
	}
	cs, e := s.Store.ListCollections()
	if e != nil {
		return model.Homepage{}, e
	}
	cats, e := s.Store.ListCategories()
	if e != nil {
		return model.Homepage{}, e
	}
	f, _ := query.Featured(as)
	return model.Homepage{Featured: f, Articles: query.Paginate(as, 1, 20), Collections: cs, Categories: cats}, nil
}
func (s *Service) EditCollection(id string, ids []string) error {
	cs, e := s.Store.ListCollections()
	if e != nil {
		return e
	}
	for _, c := range cs {
		if c.ID == id {
			c.ArticleIDs = s.CopyArticleIDs(ids)
			return s.Store.SaveCollection(c)
		}
	}
	return errors.New("collection not found")
}
func (s *Service) AddTag(t model.Tag) error {
	if t.ID == "" || t.ArticleID == "" {
		return errors.New("invalid tag")
	}
	return s.Store.SaveTag(t)
}
