package service

import (
	"log"

	"github.com/songjiangfeng/iris-blog/models"
)

type MenuService struct {
	//依赖注入
}

func (s *MenuService) GetAll() ([]models.IrisMenu, error) {
	result, err := queries.GetMenu(ctx)
	if err != nil {
		log.Println(err)
	}
	// log.Println(result)
	return result, err
}
