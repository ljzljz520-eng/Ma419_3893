package service

import (
	"reflect"
	"testing"

	"heritage/model"
	"heritage/storage"
)

// openMemoryStore opens an in-memory SQLite store that works in restricted
// sandboxes (file-based SQLite needs a writable temp dir that may be absent).
func openMemoryStore(t *testing.T) *storage.Store {
	t.Helper()
	st, err := storage.Open("file::memory:?cache=shared")
	if err != nil {
		t.Fatalf("open in-memory store: %v", err)
	}
	return st
}

// TestWorkflowEditor covers the editor business chain described in the contract:
//   创建文章和分类 → 复制专题文章列表 → 编辑草稿顺序 → 独立保存并发布专题。
//
// The list copied for editing must stay independent from the source list: a
// later edit to the draft must not rewrite the previous list's contents.
func TestWorkflowEditor(t *testing.T) {
	st := openMemoryStore(t)
	defer st.Close()
	s := New(st)

	// 创建文章和分类
	if err := st.SaveCategory(model.Category{ID: "cat-paper", Name: "剪纸", Slug: "paper"}); err != nil {
		t.Fatal(err)
	}
	for _, a := range []model.Article{
		model.NewArticle("a-1", "甲", "剪纸", "林", 6),
		model.NewArticle("a-2", "乙", "剪纸", "林", 7),
		model.NewArticle("a-3", "丙", "剪纸", "林", 8),
	} {
		if err := s.PublishArticle(a); err != nil {
			t.Fatal(err)
		}
	}

	// 第一份专题列表（源）
	srcIDs := []string{"a-1", "a-2", "a-3"}
	if err := s.PublishCollection(model.Collection{
		ID: "col-src", Name: "源专题", ArticleIDs: srcIDs, Published: true,
	}); err != nil {
		t.Fatal(err)
	}

	// 复制专题文章列表 → 草稿
	draft := s.CopyArticleIDs(srcIDs)
	// 编辑草稿顺序（in-place swap, simulating a reorder）
	draft[0], draft[len(draft)-1] = draft[len(draft)-1], draft[0]

	// 独立保存并发布第二份专题
	if err := s.PublishCollection(model.Collection{
		ID: "col-draft", Name: "草稿专题", ArticleIDs: draft, Published: true,
	}); err != nil {
		t.Fatal(err)
	}

	// 前一份列表的内容不应被后一次编辑改写
	if want := []string{"a-1", "a-2", "a-3"}; !reflect.DeepEqual(srcIDs, want) {
		t.Fatalf("source list rewritten by later edit: got %v want %v", srcIDs, want)
	}
	// 草稿应反映编辑后的顺序
	if want := []string{"a-3", "a-2", "a-1"}; !reflect.DeepEqual(draft, want) {
		t.Fatalf("draft not edited independently: got %v want %v", draft, want)
	}

	// 两份专题在存储中各自独立保存
	if got := mustCollection(t, st, "col-src"); !reflect.DeepEqual(got.ArticleIDs, []string{"a-1", "a-2", "a-3"}) {
		t.Fatalf("stored first list changed: %v", got.ArticleIDs)
	}
	if got := mustCollection(t, st, "col-draft"); !reflect.DeepEqual(got.ArticleIDs, []string{"a-3", "a-2", "a-1"}) {
		t.Fatalf("stored second list unexpected: %v", got.ArticleIDs)
	}
}

// TestCopyArticleIDsIndependence pins the slice-alias bug directly: copying a
// list must yield an independent slice, so editing the copy never mutates the
// original — different business lists stay independently saved.
func TestCopyArticleIDsIndependence(t *testing.T) {
	s := &Service{}
	src := []string{"a-1", "a-2", "a-3"}
	draft := s.CopyArticleIDs(src)
	draft[0] = "x"
	if !reflect.DeepEqual(src, []string{"a-1", "a-2", "a-3"}) {
		t.Fatalf("CopyArticleIDs shared backing array; source rewritten to %v", src)
	}
	if draft[0] != "x" {
		t.Fatalf("draft not editable independently: %v", draft)
	}
}

func mustCollection(t *testing.T, st *storage.Store, id string) model.Collection {
	t.Helper()
	cs, err := st.ListCollections()
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range cs {
		if c.ID == id {
			return c
		}
	}
	t.Fatalf("collection %s not found", id)
	return model.Collection{}
}
