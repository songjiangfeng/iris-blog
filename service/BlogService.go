package service

import (
	"github.com/songjiangfeng/iris-blog/models"
)

type BlogService struct {
	//依赖注入
}

const postnumber = 10

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

func (s *BlogService) GetLatest() ([]models.Post, error) {
	var result []models.Post
	err := orm.Raw("select * from `iris_posts` order by `created_at` desc limit ?", postnumber).Scan(&result).Error
	return result, err
}

func (s *BlogService) GetHot() ([]models.Post, error) {
	var result []models.Post
	err := orm.Raw("select * from `iris_posts` order by `views` desc limit ?", postnumber).Scan(&result).Error
	return result, err
}

func (s *BlogService) ViewPlus(id int64) {
	orm.Exec("update `iris_posts` set `views` = `views` + 1 where id = ?", id)
}
