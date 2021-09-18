package service

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/jinzhu/gorm"
)

var (
	orm *gorm.DB
	err error
)

type MysqlService struct {
	//依赖注入
}

func (s *MysqlService) Init(c db.Connection) {

	orm, err = gorm.Open("mysql", c.GetDB("default"))
	orm.LogMode(true)
	if err != nil {
		panic("initialize orm failed")
	}
}
