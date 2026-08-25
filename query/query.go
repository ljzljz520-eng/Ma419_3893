package query

import "heritage/model"

func FilterArticles(in []model.Article, category string) []model.Article {
	out := []model.Article{}
	for _, a := range in {
		if category == "" || a.Category == category {
			out = append(out, a)
		}
	}
	return out
}
func Featured(in []model.Article) (model.Article, bool) {
	for _, a := range in {
		if a.Featured {
			return a, true
		}
	}
	if len(in) > 0 {
		return in[0], true
	}
	return model.Article{}, false
}
func Paginate(in []model.Article, page, size int) []model.Article {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	start := (page - 1) * size
	if start >= len(in) {
		return []model.Article{}
	}
	end := start + size
	if end > len(in) {
		end = len(in)
	}
	return append([]model.Article(nil), in[start:end]...)
}
func ByReadingTime(in []model.Article, max int) []model.Article {
	out := []model.Article{}
	for _, a := range in {
		if a.ReadingMinutes <= max {
			out = append(out, a)
		}
	}
	return out
}
func Related(in []model.Article, category string, exclude string) []model.Article {
	out := []model.Article{}
	for _, a := range in {
		if a.Category == category && a.ID != exclude {
			out = append(out, a)
		}
	}
	return out
}
func SortByTitle(in []model.Article) []model.Article {
	out := append([]model.Article(nil), in...)
	for i := 0; i < len(out); i++ {
		for j := i + 1; j < len(out); j++ {
			if out[j].Title < out[i].Title {
				out[i], out[j] = out[j], out[i]
			}
		}
	}
	return out
}
