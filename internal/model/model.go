package model

import (
	"blog-service/global"
	"blog-service/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Model 基础模型
type Model struct {
	ID         uint32 `gorm:"primaryKey" json:"id"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	IsDel      uint8  `json:"is_del"`
}

// NewDBEngine 创建数据源实例
func NewDBEngine(database *setting.Database) (*gorm.DB, error) {
	fmt.Println(database.Username, database.Password)
	db, err := gorm.Open(database.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			database.Username,
			database.Password,
			database.Host,
			database.DBName,
			database.Charset,
			database.ParseTime,
		))
	if err != nil {
		return nil, err
	}
	if global.Server.RunMode == "debug" {
		// 开启详细的日志模式
		db.LogMode(true)
	}
	// 开启表名单数
	db.SingularTable(true)
	sqlDB := db.DB()
	sqlDB.SetMaxIdleConns(database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(database.MaxOpenConns)
	return db, nil
}
