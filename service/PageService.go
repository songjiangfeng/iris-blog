package service

import (
	"log"

	"github.com/songjiangfeng/iris-blog/models"
)

// PageService
type PageService struct {
	//依赖注入
}

// GetPage
//  @receiver s
//  @param slug
//  @return models.IrisPage
//  @return error
func (s *PageService) GetPage(slug string) (models.IrisPage, error) {
	result, err := queries.GetPage(ctx, slug)
	if err != nil {
		log.Println(err)
	}
	//log.Println(result)
	return result, err
}
