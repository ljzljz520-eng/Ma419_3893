package query

import (
	"heritage/content"
	"testing"
)

func TestQueryTools(t *testing.T) {
	a := content.SeedArticles()
	if len(FilterArticles(a, "剪纸")) != 2 {
		t.Fatal("filter")
	}
	if len(Paginate(a, 1, 2)) != 2 {
		t.Fatal("page")
	}
	if _, ok := Featured(a); !ok {
		t.Fatal("featured")
	}
	if len(Recommend(a, "木版画", 2)) != 2 {
		t.Fatal("recommend")
	}
}
