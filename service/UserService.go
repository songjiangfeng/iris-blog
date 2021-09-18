package service

import "github.com/songjiangfeng/iris-blog/models"

type UserService struct {
	//依赖注入
}

func (s *UserService) GetAll() ([]models.User, error) {
	var result []models.User
	err := orm.Raw("select * from `goadmin_users`").Scan(&result).Error
	return result, err
}

func (s *UserService) GetUser(id int64) (models.User, error) {
	var result models.User
	err := orm.Raw("select * from `goadmin_users` where `id` = ? ", id).Scan(&result).Error
	return result, err
}
