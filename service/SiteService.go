package service

import (
	"log"

	"github.com/songjiangfeng/iris-blog/models"
)

// SiteService
type SiteService struct {
	//依赖注入

}

// GetSite
//  @receiver s
//  @return models.IrisSite
//  @return error
func (s *SiteService) GetSite() (models.IrisSite, error) {
	result, err := queries.GetSite(ctx, 1)
	if err != nil {
		log.Println(err)
	}
	// log.Println(result)
	return result, err
}
