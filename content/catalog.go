package content

import "heritage/model"

func FeaturedInterview() model.Article { return SeedArticles()[0] }
func ArticleByID(id string) (model.Article, bool) {
	for _, a := range SeedArticles() {
		if a.ID == id {
			return a, true
		}
	}
	return model.Article{}, false
}
func CategoryNames() []string {
	out := []string{}
	for _, c := range SeedCategories() {
		out = append(out, c.Name)
	}
	return out
}
func CollectionByID(id string) (model.Collection, bool) {
	for _, c := range SeedCollections() {
		if c.ID == id {
			return c, true
		}
	}
	return model.Collection{}, false
}
func TagsForArticle(id string) []model.Tag {
	out := []model.Tag{}
	for _, t := range SeedTags() {
		if t.ArticleID == id {
			out = append(out, t)
		}
	}
	return out
}
func ArticleIDs() []string {
	out := []string{}
	for _, a := range SeedArticles() {
		out = append(out, a.ID)
	}
	return out
}
func IsSeedCategory(name string) bool {
	for _, c := range SeedCategories() {
		if c.Name == name || c.Slug == name {
			return true
		}
	}
	return false
}
func IsSeedArticle(id string) bool { _, ok := ArticleByID(id); return ok }
func PublishedCollections() []model.Collection {
	out := []model.Collection{}
	for _, c := range SeedCollections() {
		if c.Published {
			out = append(out, c)
		}
	}
	return out
}
func ReadingTotal() int {
	n := 0
	for _, a := range SeedArticles() {
		n += a.ReadingMinutes
	}
	return n
}
func LongArticles() []model.Article {
	out := []model.Article{}
	for _, a := range SeedArticles() {
		if a.ReadingMinutes >= 12 {
			out = append(out, a)
		}
	}
	return out
}
func PhotographerNames() []string {
	out := []string{}
	for _, a := range SeedArticles() {
		out = append(out, a.Photographer)
	}
	return out
}
func CoverURLs() []string {
	out := []string{}
	for _, a := range SeedArticles() {
		out = append(out, a.Cover)
	}
	return out
}
func ValidateSeed() bool {
	return len(SeedArticles()) >= 5 && len(SeedCategories()) == 4 && len(SeedCollections()) >= 2
}
func TopicArticles(id string) []model.Article {
	c, ok := CollectionByID(id)
	if !ok {
		return nil
	}
	out := []model.Article{}
	for _, aid := range c.ArticleIDs {
		if a, found := ArticleByID(aid); found {
			out = append(out, a)
		}
	}
	return out
}
func CategoryArticleCount(name string) int {
	n := 0
	for _, a := range SeedArticles() {
		if a.Category == name {
			n++
		}
	}
	return n
}
func HasFeatured() bool {
	for _, a := range SeedArticles() {
		if a.Featured {
			return true
		}
	}
	return false
}
