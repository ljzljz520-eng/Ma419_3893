package content

func Manifest() map[string]int {
	return map[string]int{"articles": len(SeedArticles()), "categories": len(SeedCategories()), "collections": len(SeedCollections()), "tags": len(SeedTags())}
}
