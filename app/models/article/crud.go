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
