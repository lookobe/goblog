package article

import (
	"goblog/app/models"
	"goblog/app/models/user"
	"goblog/pkg/route"
)

// Article 文章模型
type Article struct {
	models.BaseModel
	// ID int64
	Title string
	Body string

	UserID uint64 `gorm:"not null;index"`
	User user.User
	CategoryID uint64 `gorm:"not null;default:8;index"`
}
// Link 方法用来生成文章链接 
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}

// CreatedAtDate 创建日期
func (a Article) CreatedAtDate() string {
	return a.CreatedAt.Format("2006-01-02")
}