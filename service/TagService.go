package service

import "github.com/songjiangfeng/iris-blog/models"

type TagService struct {
	//依赖注入
}

func (*TagService) GetAll() ([]models.Tag, error) {
	var result []models.Tag
	err := orm.Raw("select * from `iris_tags`").Scan(&result).Error
	return result, err
}

func (*TagService) GetTag(id int64) (models.Tag, error) {
	var result models.Tag
	err := orm.Raw("select * from `iris_tags` where id = ?", id).Scan(&result).Error
	return result, err
}

func (*TagService) GetTagPost(id int64) ([]models.Tag, error) {
	var result []models.Tag
	err := orm.Raw("select t.*, tp.post_id,p.title,p.content,p.created_at,p.views from `iris_tags` as t inner join `iris_tags_posts` as tp on t.id = tp.tag_id inner join `iris_posts` as p on tp.post_id = p.id where t.id = ? limit 10", id).Scan(&result).Error
	return result, err
}
