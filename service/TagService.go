package service

import (
	"log"

	"github.com/songjiangfeng/iris-blog/models"
)

// TagService
type TagService struct {
	//依赖注入
}

// GetAll
//  @receiver *TagService
//  @return []models.IrisTag
//  @return error
func (*TagService) GetAll() ([]models.IrisTag, error) {
	result, err := queries.GetTags(ctx)
	if err != nil {
		log.Println(err)
	}
	// log.Println(result)
	return result, err
}

// GetTag
//  @receiver *TagService
//  @param id
//  @return models.IrisTag
//  @return error
func (*TagService) GetTag(id int64) (models.IrisTag, error) {
	result, err := queries.GetTag(ctx, id)
	if err != nil {
		log.Println(err)
	}
	// log.Println(result)
	return result, err
}

// GetTagPostId
//  @receiver *TagService
//  @param id
//  @return []int64
//  @return error
func (*TagService) GetTagPostId(id int64) ([]int64, error) {

	result, err := queries.GetTagPostId(ctx, id)
	if err != nil {
		log.Println(err)
	}
	// log.Println(result)
	return result, err
}
