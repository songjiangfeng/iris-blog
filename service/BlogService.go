package service

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/songjiangfeng/iris-blog/cache"
	"github.com/songjiangfeng/iris-blog/models"
)

// BlogService
type BlogService struct {
	//依赖注入
	cache cache.RedisCache
}

// const
const postnumber = 10

// GetPost
//  @receiver s
//  @param id
//  @return models.IrisPost
//  @return error

func (s *BlogService) GetPost(id int64) (models.IrisPost, error) {

	key := "post" + strconv.FormatInt(id, 16)
	keycheck := s.cache.RedisKeyExists(key)
	var result models.IrisPost

	if !keycheck {
		result, err = queries.GetPost(ctx, id)
		if err != nil {
			log.Println(err)
		}
		err = s.cache.RedisSetCache(key, result)
	}
	data, _ := s.cache.RedisGetCache(key)
	err = json.Unmarshal([]byte(data), &result)
	// log.Println(result)
	return result, err
}

// GetPrevPost
//  @receiver s
//  @param id
//  @return models.IrisPost
//  @return error
func (s *BlogService) GetPrevPost(id int64) (models.IrisPost, error) {
	result, err := queries.GetPrevPost(ctx, id)
	if err != nil {
		log.Println(err)
	}
	// log.Println(result)
	return result, err
}

// GetNextPost
//  @receiver s
//  @param id
//  @return models.IrisPost
//  @return error
func (s *BlogService) GetNextPost(id int64) (models.IrisPost, error) {
	result, err := queries.GetNextPost(ctx, id)
	if err != nil {
		log.Println(err)
	}
	// log.Println(result)
	return result, err
}

// GetLatest
//  @receiver s
//  @return []models.IrisPost
//  @return error
func (s *BlogService) GetLatest() ([]models.IrisPost, error) {
	result, err := queries.GetLatestPost(ctx, postnumber)
	if err != nil {
		log.Println(err)
	}
	// log.Println(result)
	return result, err
}

// GetHot
//  @receiver s
//  @return []models.IrisPost
//  @return error
func (s *BlogService) GetHot() ([]models.IrisPost, error) {
	result, err := queries.GetHotPost(ctx, postnumber)
	if err != nil {
		log.Println(err)
	}
	// log.Println(result)
	return result, err
}

// GetPostByPage
//  @receiver s
//  @param page
//  @param pagesize
//  @return []models.IrisPost
//  @return error
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

// GetPostByTagID
//  @receiver s
//  @param tagid
//  @return []models.IrisPost
//  @return error
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

// ViewPlus
//  @receiver s
//  @param id
//  @return error
func (s *BlogService) ViewPlus(id int64) error {
	err := queries.ViewPlus(ctx, id)
	if err != nil {
		log.Println(err)
	}
	return err
}
