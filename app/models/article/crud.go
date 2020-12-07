package article

import (
	"goblog/pkg/model"
	"goblog/pkg/types"
)




// Get 通过 ID 获取文章
func Get(idstr string) (*Article, error) {
	var articled Article
	id := types.StringToInt(idstr)
	if err := model.DB.First(&articled, id).Error; err != nil {
		return &articled, err
	}
	return &articled, nil
}

func GetAll()([]Article,error)  {
	var articles  []Article
	if err := model.DB.Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}
