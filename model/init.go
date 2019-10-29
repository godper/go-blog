package model

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Database 数据库配置
type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

//Model 基础共用数据表字段
type Model struct {
	ID        uint   `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt string `json:"created_at,omitempty" godper:"2006-01-02"`
	UpdatedAt string `json:"updated_at,omitempty" godper:"2006-01-02"`
	DeletedAt string `gorm:"default:NULL" sql:"index" json:"deleted_at,omitempty" godper:"2006-01-02"`
}

//NewDB 生成数据库实例
func NewDB(d *Database) *gorm.DB {
	connType := d.Type
	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		d.User,
		d.Password,
		d.Host,
		d.Name,
	)

	db, err := gorm.Open(connType, connString)

	db.LogMode(true)
	// Error
	if err != nil {
		panic(err)
	}
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(20)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	//迁移数据库
	migration(db)
	return db
}

func migration(db *gorm.DB) {
	// 自动迁移模式
	db.Set("gorm:table_options", "charset=utf8").
		AutoMigrate(&User{}).
		AutoMigrate(&Article{}).
		AutoMigrate(&Topic{}).
		AutoMigrate(&Admin{}).
		AutoMigrate(&ArticleTag{}).
		AutoMigrate(&Tag{})
}

// updateTimeStampForCreateCallback will set `CreatedAt`, `UpdatedAt` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := strconv.FormatInt(time.Now().Unix(), 10)
		if CreatedAtField, ok := scope.FieldByName("CreatedAt"); ok {
			if CreatedAtField.IsBlank {
				CreatedAtField.Set(nowTime)
			}
		}

		if UpdatedAtField, ok := scope.FieldByName("UpdatedAt"); ok {
			if UpdatedAtField.IsBlank {
				UpdatedAtField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `UpdatedAt` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	// if _, ok := scope.Get("gorm:update_column"); !ok {
	// 	scope.SetColumn("UpdatedAt", time.Now().Unix())
	// }
	if !scope.HasError() {
		nowTime := strconv.FormatInt(time.Now().Unix(), 10)
		if UpdatedAtField, ok := scope.FieldByName("UpdatedAt"); ok {
			if UpdatedAtField.IsBlank {
				UpdatedAtField.Set(nowTime)
			}
		}
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedAt")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
