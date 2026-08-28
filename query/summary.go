package query

import "heritage/model"

func Summary(in []model.Article) map[string]int {
	f := 0
	for _, a := range in {
		if a.Featured {
			f++
		}
	}
	return map[string]int{"total": len(in), "featured": f, "long": len(ByReadingTime(in, 120))}
}
