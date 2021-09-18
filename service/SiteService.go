package service

import "github.com/songjiangfeng/iris-blog/models"

type SiteService struct {
	//依赖注入
}

func (s *SiteService) GetSite() (models.Site, error) {
	var result models.Site
	err := orm.Raw("select * from `iris_site` where id = ? ", 1).Scan(&result).Error
	return result, err
}
