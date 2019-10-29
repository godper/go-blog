package dao

import (
	"blog/conf"
	"blog/model"

	"github.com/jinzhu/gorm"
)

//Dao 层
type Dao struct {
	c  *conf.Config
	DB *gorm.DB
}

//NewDao 实例化 Dao
func NewDao(c *conf.Config) *Dao {
	return &Dao{
		c:  c,
		DB: model.NewDB(c.Database),
	}
}
