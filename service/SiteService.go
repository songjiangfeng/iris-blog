package service

import (
	"encoding/json"
	"log"

	"github.com/songjiangfeng/iris-blog/cache"
	"github.com/songjiangfeng/iris-blog/models"
)

// SiteService
type SiteService struct {
	//依赖注入
	cache cache.RedisCache
}

// GetSite
//  @receiver s
//  @return models.IrisSite
//  @return error
func (s *SiteService) GetSite() (models.IrisSite, error) {

	key := "site"
	keycheck := s.cache.RedisKeyExists(key)
	var result models.IrisSite
	if !keycheck {
		result, err = queries.GetSite(ctx, 1)
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
