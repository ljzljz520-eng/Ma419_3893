package model

func ReadingLabel(m int) string {
	if m <= 0 {
		return "未标注"
	}
	if m < 5 {
		return "短读"
	}
	if m < 12 {
		return "中读"
	}
	return "长访谈"
}
func CategoryColor(c string) string {
	switch c {
	case "剪纸":
		return "朱砂"
	case "木版画":
		return "墨色"
	case "戏曲服饰":
		return "绛红"
	case "地方乐器":
		return "青黛"
	}
	return "素白"
}
func IsLongForm(a Article) bool { return a.ReadingMinutes >= 12 }
func HasCover(a Article) bool   { return a.Cover != "" }
func WordCount(a Article) int   { return len([]rune(a.Body)) }
func CardKey(a Article) string  { return a.Category + ":" + a.ID }
func DisplayPhotographer(a Article) string {
	if a.Photographer == "" {
		return "影像志编辑部"
	}
	return a.Photographer
}
func CollectionSize(c Collection) int { return len(c.ArticleIDs) }
func CollectionLabel(c Collection) string {
	if c.Published {
		return "已发布"
	}
	return "草稿"
}
func CategoryMatch(a Article, c Category) bool { return a.Category == c.Name || a.Category == c.Slug }
func SafeMinutes(m int) int {
	if m < 1 {
		return 1
	}
	if m > 120 {
		return 120
	}
	return m
}
func TruncateTitle(s string, n int) string {
	r := []rune(s)
	if len(r) <= n {
		return s
	}
	return string(r[:n]) + "…"
}
func ArticleEqual(a, b Article) bool {
	return a.ID == b.ID && a.Title == b.Title && a.Category == b.Category
}
func CategoryEqual(a, b Category) bool { return a.ID == b.ID && a.Slug == b.Slug }
func TagEqual(a, b Tag) bool           { return a.ID == b.ID && a.ArticleID == b.ArticleID }
func CloneArticles(in []Article) []Article {
	out := make([]Article, len(in))
	copy(out, in)
	return out
}
func CloneCategories(in []Category) []Category {
	out := make([]Category, len(in))
	copy(out, in)
	return out
}
func CloneTags(in []Tag) []Tag { out := make([]Tag, len(in)); copy(out, in); return out }
func FeaturedCount(in []Article) int {
	n := 0
	for _, a := range in {
		if a.Featured {
			n++
		}
	}
	return n
}
func CategoryCounts(in []Article) map[string]int {
	m := map[string]int{}
	for _, a := range in {
		m[a.Category]++
	}
	return m
}
func ContainsID(ids []string, id string) bool {
	for _, v := range ids {
		if v == id {
			return true
		}
	}
	return false
}
func AppendUnique(ids []string, id string) []string {
	if ContainsID(ids, id) {
		return ids
	}
	return append(ids, id)
}
func RemoveID(ids []string, id string) []string {
	out := []string{}
	for _, v := range ids {
		if v != id {
			out = append(out, v)
		}
	}
	return out
}
func ReverseIDs(ids []string) []string {
	out := append([]string(nil), ids...)
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return out
}
