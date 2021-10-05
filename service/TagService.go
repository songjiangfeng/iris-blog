package service

import (
	"log"

	"github.com/songjiangfeng/iris-blog/models"
)

type TagService struct {
	//依赖注入
}

func (*TagService) GetAll() ([]models.IrisTag, error) {
	result, err := queries.GetTags(ctx)
	if err != nil {
		log.Println(err)
	}
	// log.Println(result)
	return result, err
}

func (*TagService) GetTag(id int64) (models.IrisTag, error) {
	result, err := queries.GetTag(ctx, id)
	if err != nil {
		log.Println(err)
	}
	// log.Println(result)
	return result, err
}

func (*TagService) GetTagPostId(id int64) ([]int64, error) {

	result, err := queries.GetTagPostId(ctx, id)
	if err != nil {
		log.Println(err)
	}
	// log.Println(result)
	return result, err
}
