package service

import (
	"log"

	"github.com/songjiangfeng/iris-blog/models"
)

type PageService struct {
	//依赖注入
}

func (s *PageService) GetPage(slug string) (models.IrisPage, error) {
	result, err := queries.GetPage(ctx, slug)
	if err != nil {
		log.Println(err)
	}
	//log.Println(result)
	return result, err
}
