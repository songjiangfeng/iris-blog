package service

import (
	"context"
	"database/sql"

	"github.com/GoAdminGroup/go-admin/modules/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/songjiangfeng/iris-blog/models"
)

var (
	queries *models.Queries
	ctx     context.Context
	err     error
)

type MysqlService struct {
	//依赖注入
}

func (s *MysqlService) Init(c db.Connection) {

	db, err := sql.Open("mysql", "root:root@/iris_blog?charset=utf8&parseTime=True&loc=Local")
	queries = models.New(db)
	ctx = context.Background()

	if err != nil {
		panic("initialize orm failed")
	}
}
