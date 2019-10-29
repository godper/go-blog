package conf

import (
	"blog/model"
	uredis "blog/util/redis"
	"log"
	"time"

	"github.com/go-ini/ini"
)

//Config 配置
type Config struct {
	App      *App
	Server   *Server
	Database *model.Database
	Redis    *uredis.RedisConf
}

//配置文件.ini
var cfgfile *ini.File

//Cfg 初始conf实例
var Cfg *Config

//Init 初始化Conf
func Init() {
	Cfg = NewConf()
}

//NewConf  Setup initialize the configuration instance
func NewConf() *Config {
	c := Config{
		App:      &App{},
		Server:   &Server{},
		Database: &model.Database{},
		Redis:    &uredis.RedisConf{},
	}

	var err error
	cfgfile, err = ini.Load("conf/conf.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/conf.ini': %v", err)
	}

	mapTo("app", c.App)
	mapTo("server", c.Server)
	mapTo("database", c.Database)
	mapTo("redis", c.Redis)

	c.App.ImageMaxSize = c.App.ImageMaxSize * 1024 * 1024
	c.Server.ReadTimeout = c.Server.ReadTimeout * time.Second
	c.Server.WriteTimeout = c.Server.WriteTimeout * time.Second
	//c.Redis.IdleTimeout = c.Redis.IdleTimeout * time.Second
	return &c
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfgfile.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("cfgfile.MapTo %s err: %v", section, err)
	}
}
