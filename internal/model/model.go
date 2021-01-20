package model

import (
	"fmt"
	"github.com/go-programming-tour-book/blog-service/pkg/tracer"
	"time"

	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	StateOpen = 1
)

// todo 时间相关字段，使用时间戳！
type Model struct {
	ID         uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt  time.Time      `gorm:"created_at" json:"created_at,omitempty"`
	UpdatedAt  time.Time      `gorm:"updated_at" json:"updated_at,omitempty"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	CreatedBy  string         `gorm:"created_by" json:"created_by,omitempty"`
	ModifiedBy string         `gorm:"modified_by" json:"modified_by,omitempty"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)

	if databaseSetting.DBType == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			databaseSetting.UserName,
			databaseSetting.Password,
			databaseSetting.Host,
			databaseSetting.DBName,
			databaseSetting.Charset,
			databaseSetting.ParseTime)

		db, err = gorm.Open(mysql.New(mysql.Config{
			DSN: dsn,
			//DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			//SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
		}), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})

		if err != nil {
			return nil, err
		}
	} else {
		panic(fmt.Sprintf("not support %s database", databaseSetting.DBType))
	}

	if global.ServeSetting.RunMode == "debug" {
		db = db.Debug()
	}
	sqlDB, err := db.DB()
	if err != nil {
		//warn
		return nil, err
	}

	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	_ = db.Use(&tracer.OpenTracingPlugin{})
	return db, nil
}

// 新版本回调
//func updateTimeStampForCreateCallback(db *gorm.DB) {
//	nowTime := time.Now().Unix()
//	if createTimeField := db.Statement.Schema.LookUpField("CreatedOn"); createTimeField != nil {
//		db.Statement.SetColumn("CreatedOn", nowTime)
//	}
//
//	if createTimeField := db.Statement.Schema.LookUpField("CreatedOn"); createTimeField != nil {
//		db.Statement.SetColumn("CreatedOn", nowTime)
//	}
//}
//
//func updateTimeStampForUpdateCallback(db *gorm.DB) {
//
//}
