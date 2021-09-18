package service

import "github.com/songjiangfeng/iris-blog/models"

type PageService struct {
	//依赖注入
}

func (s *PageService) GetPage(slug string) (models.Page, error) {
	var result models.Page
	err := orm.Raw("select * from `iris_pages` where `slug` = ?", slug).Scan(&result).Error
	return result, err
}
