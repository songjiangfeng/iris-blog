package service

import (
	"log"

	"github.com/songjiangfeng/iris-blog/models"
)

type UserService struct {
	//依赖注入
}

func (s *UserService) GetAll() ([]models.GoadminUser, error) {

	result, err := queries.GetUsers(ctx)
	if err != nil {
		log.Println(err)
	}
	// log.Println(result)
	return result, err
}

func (s *UserService) GetUser(id int64) (models.GoadminUser, error) {

	result, err := queries.GetUser(ctx, int32(id))
	if err != nil {
		log.Println(err)
	}
	// log.Println(result)
	return result, err
}
