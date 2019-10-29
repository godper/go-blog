package service

import (
	"blog/cache"
	"blog/conf"
	"blog/dao"
	ujwt "blog/util/jwt"
	"blog/util/page"
	"blog/util/upload"
)

//Service 服务
type Service struct {
	C           *conf.Config
	Dao         *dao.Dao
	Cache       *cache.Cache
	Jwt         *ujwt.JWT
	Page        *page.Page
	ImageUpload *upload.ImageUpload
}

//Srv 初始化Service实例
var Srv *Service

//Init 初始化Srv
func Init() {
	c := conf.Cfg
	Srv = NewService(c)
}

//NewService 生成新服务实例
func NewService(c *conf.Config) *Service {
	return &Service{
		C:           c,
		Dao:         dao.NewDao(c),
		Cache:       cache.NewCache(c),
		Jwt:         ujwt.NewJWT(c.App.JwtSecret),
		Page:        page.NewPage(c.App.PageSize),
		ImageUpload: upload.NewImageUplaod(c.App),
	}
}
