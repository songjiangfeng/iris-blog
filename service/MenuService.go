package service

import "github.com/songjiangfeng/iris-blog/models"

type MenuService struct {
	//依赖注入
}

func (s *MenuService) GetAll() ([]models.Menu, error) {
	var result []models.Menu
	err := orm.Raw("select * from `iris_menus` order by weight,id asc").Scan(&result).Error
	return result, err
}
