package article

import (
	"goblog/pkg/route"
	"goblog/app/models"
)
// Article 文章模型
type Article struct {
	models.BaseModel
	// ID int64
	Title string
	Body string
}
// Link 方法用来生成文章链接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}