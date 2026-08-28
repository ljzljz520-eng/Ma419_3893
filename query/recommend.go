package query

import "heritage/model"

func Recommend(in []model.Article, category string, limit int) []model.Article {
	if limit < 1 {
		return []model.Article{}
	}
	out := []model.Article{}
	for _, a := range in {
		if a.Category == category {
			out = append(out, a)
			if len(out) == limit {
				return out
			}
		}
	}
	for _, a := range in {
		if a.Category != category {
			out = append(out, a)
			if len(out) == limit {
				return out
			}
		}
	}
	return out
}
func ByPhotographer(in []model.Article, name string) []model.Article {
	out := []model.Article{}
	for _, a := range in {
		if a.Photographer == name {
			out = append(out, a)
		}
	}
	return out
}
func FeaturedFirst(in []model.Article) []model.Article {
	out := []model.Article{}
	for _, a := range in {
		if a.Featured {
			out = append(out, a)
		}
	}
	for _, a := range in {
		if !a.Featured {
			out = append(out, a)
		}
	}
	return out
}
func DistinctCategories(in []model.Article) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, a := range in {
		if !seen[a.Category] {
			seen[a.Category] = true
			out = append(out, a.Category)
		}
	}
	return out
}
func AverageReading(in []model.Article) int {
	if len(in) == 0 {
		return 0
	}
	n := 0
	for _, a := range in {
		n += a.ReadingMinutes
	}
	return n / len(in)
}
func Longest(in []model.Article) model.Article {
	var out model.Article
	for _, a := range in {
		if a.ReadingMinutes > out.ReadingMinutes {
			out = a
		}
	}
	return out
}
func Shortest(in []model.Article) model.Article {
	if len(in) == 0 {
		return model.Article{}
	}
	out := in[0]
	for _, a := range in[1:] {
		if a.ReadingMinutes < out.ReadingMinutes {
			out = a
		}
	}
	return out
}
func PageCount(total, size int) int {
	if size < 1 {
		return 0
	}
	n := total / size
	if total%size != 0 {
		n++
	}
	return n
}
func ValidPage(page, size int) bool { return page > 0 && size > 0 && size <= 100 }
func ClampSize(size int) int {
	if size < 1 {
		return 1
	}
	if size > 50 {
		return 50
	}
	return size
}
func Search(in []model.Article, term string) []model.Article {
	out := []model.Article{}
	for _, a := range in {
		if contains(a.Title, term) || contains(a.Body, term) {
			out = append(out, a)
		}
	}
	return out
}
func contains(s, t string) bool {
	if t == "" {
		return true
	}
	return len(s) >= len(t) && stringIndex(s, t) >= 0
}
func stringIndex(s, t string) int {
	for i := 0; i+len(t) <= len(s); i++ {
		if s[i:i+len(t)] == t {
			return i
		}
	}
	return -1
}
