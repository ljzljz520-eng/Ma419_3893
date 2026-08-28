package content

import "heritage/model"

func SeedCategories() []model.Category {
	return []model.Category{{"cat-paper", "剪纸", "paper"}, {"cat-wood", "木版画", "wood"}, {"cat-costume", "戏曲服饰", "costume"}, {"cat-instrument", "地方乐器", "instrument"}}
}
func SeedArticles() []model.Article {
	return []model.Article{
		{ID: "a-interview", Title: "守住一张纸的呼吸：周师傅长访谈", Category: "剪纸", Photographer: "林晓", ReadingMinutes: 18, Cover: "cover://interview", Body: "在小城北巷，周师傅把一把剪刀用成了时间。" + stringsBlock(), Featured: true},
		{ID: "a-paper", Title: "窗花里的四季", Category: "剪纸", Photographer: "陈默", ReadingMinutes: 7, Cover: "cover://paper", Body: "纸张在指尖折叠出四季。"},
		{ID: "a-wood", Title: "木版画的回声", Category: "木版画", Photographer: "赵青", ReadingMinutes: 9, Cover: "cover://wood", Body: "老木版记录着河岸的纹理。"},
		{ID: "a-costume", Title: "戏台之后的衣裳", Category: "戏曲服饰", Photographer: "沈秋", ReadingMinutes: 11, Cover: "cover://costume", Body: "一针一线缝起戏曲记忆。"},
		{ID: "a-instrument", Title: "一把月琴的地方声", Category: "地方乐器", Photographer: "顾野", ReadingMinutes: 8, Cover: "cover://instrument", Body: "月琴声从庙会传到巷口。"},
	}
}
func stringsBlock() string { return " 访谈记录关注手艺、家庭与地方记忆。" }
func SeedCollections() []model.Collection {
	return []model.Collection{{ID: "col-featured", Name: "专题推荐：手艺人的一天", Description: "从清晨到灯下的影像", ArticleIDs: []string{"a-interview", "a-paper"}, Published: true}, {ID: "col-season", Name: "季节里的非遗", Description: "四时风物", ArticleIDs: []string{"a-wood", "a-costume", "a-instrument"}, Published: true}}
}
func SeedTags() []model.Tag {
	return []model.Tag{{"t1", "a-interview", "口述史"}, {"t2", "a-paper", "节气"}, {"t3", "a-wood", "版印"}, {"t4", "a-costume", "戏台"}, {"t5", "a-instrument", "声音"}}
}
