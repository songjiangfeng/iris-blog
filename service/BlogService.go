package service

import "github.com/songjiangfeng/iris-blog/models"

type BlogService struct {
	//依赖注入
}

func (s *BlogService) GetAll() ([]models.Post, error) {
	var result []models.Post
	err := orm.Raw("select * from `iris_posts` order by `created_at` desc ").Scan(&result).Error
	return result, err
}

func (s *BlogService) GetPostByPage(page int64, pagesize int64) ([]models.Post, error) {
	var result []models.Post
	start := (page - 1) * pagesize
	err := orm.Raw("select * from `iris_posts` order by `created_at` desc  limit   ? ,  ? ", start, pagesize).Scan(&result).Error
	return result, err
}

func (s *BlogService) GetPost(id int64) (models.Post, error) {
	var result models.Post
	err := orm.Raw("select * from `iris_posts` where `id` = ?", id).Scan(&result).Error
	return result, err
}
func (s *BlogService) Limit(id int64) ([]models.Post, error) {
	var result []models.Post
	err := orm.Raw("select * from `iris_posts` order by `created_at` desc  limit  ? ", id).Scan(&result).Error
	return result, err
}
