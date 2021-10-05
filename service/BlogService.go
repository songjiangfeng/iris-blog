package service

import (
	"log"

	"github.com/songjiangfeng/iris-blog/models"
)

type BlogService struct {
	//依赖注入
}

const postnumber = 10

func (s *BlogService) GetPost(id int64) (models.IrisPost, error) {
	result, err := queries.GetPost(ctx, id)
	if err != nil {
		log.Println(err)
	}
	// log.Println(result)
	return result, err
}

func (s *BlogService) GetPrevPost(id int64) (models.IrisPost, error) {
	result, err := queries.GetPrevPost(ctx, id)
	if err != nil {
		log.Println(err)
	}
	// log.Println(result)
	return result, err
}

func (s *BlogService) GetNextPost(id int64) (models.IrisPost, error) {
	result, err := queries.GetNextPost(ctx, id)
	if err != nil {
		log.Println(err)
	}
	// log.Println(result)
	return result, err
}

func (s *BlogService) GetLatest() ([]models.IrisPost, error) {
	result, err := queries.GetLatestPost(ctx, postnumber)
	if err != nil {
		log.Println(err)
	}
	// log.Println(result)
	return result, err
}

func (s *BlogService) GetHot() ([]models.IrisPost, error) {
	result, err := queries.GetHotPost(ctx, postnumber)
	if err != nil {
		log.Println(err)
	}
	// log.Println(result)
	return result, err
}

func (s *BlogService) GetPostByPage(page int64, pagesize int64) ([]models.IrisPost, error) {
	start := (page - 1) * pagesize
	arg := models.GetPostByPageParams{}
	arg.Limit = int32(start)
	arg.Offset = int32(pagesize)
	result, err := queries.GetPostByPage(ctx, arg)
	if err != nil {
		log.Println(err)
	}
	// log.Println(result)
	return result, err
}

func (s *BlogService) GetPostByTagID(tagid int64) ([]models.IrisPost, error) {

	postid, _ := queries.GetTagPostId(ctx, tagid)

	var result []models.IrisPost

	for _, v := range postid {
		item, err := queries.GetPost(ctx, v)
		if err != nil {
			log.Println(err)
		}
		result = append(result, item)
	}
	// log.Println(result)
	return result, err
}

func (s *BlogService) ViewPlus(id int64) error {
	err := queries.ViewPlus(ctx, id)
	if err != nil {
		log.Println(err)
	}
	return err
}
