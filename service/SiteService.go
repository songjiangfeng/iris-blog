package service

import (
	"log"

	"github.com/songjiangfeng/iris-blog/models"
)

type SiteService struct {
	//依赖注入

}

func (s *SiteService) GetSite() (models.IrisSite, error) {
	result, err := queries.GetSite(ctx, 1)
	if err != nil {
		log.Println(err)
	}
	// log.Println(result)
	return result, err
}
