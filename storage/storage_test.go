package storage

import (
	"heritage/content"
	"os"
	"testing"
)

func TestStoreRoundTrip(t *testing.T) {
	p := t.TempDir() + "/x.db"
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if e = s.SaveArticle(content.SeedArticles()[0]); e != nil {
		t.Fatal(e)
	}
	if s.ArticleCount() != 1 {
		t.Fatal("count")
	}
	if e = s.SaveCategory(content.SeedCategories()[0]); e != nil {
		t.Fatal(e)
	}
	if _, e = s.GetArticle("a-interview"); e != nil {
		t.Fatal(e)
	}
	_ = os.Remove(p)
}
