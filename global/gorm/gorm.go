package gorm

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"portal/global/config"
	"portal/global/log"
	"time"
)

var (
	ormDB *gorm.DB
	sqlDB *sql.DB
)

// 初始化
func InitDB() {
	conf := config.GetAppConfig().Mysql
	conString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Port, conf.Db)
	var err error
	ormDB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: conString, // DSN data source name
		DefaultStringSize: 256, // string 类型字段的默认长度
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	sqlDB, err = ormDB.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func GetOrmDB() *gorm.DB {
	// conf := config.GetAppConfig()
	return ormDB.Debug()
	//if conf.Env != "prod" {
	//	return ormDB.Debug()
	//}else{
	//	return ormDB
	//}
}

func GetSqlDB() *sql.DB {
	return sqlDB
}

func CloseDB() {
	if err := sqlDB.Close();err != nil{
		log.Error("Close DB server error by",err)
	}
}


