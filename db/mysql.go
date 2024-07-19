package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"mysql-distributed-lock/config"
)

var Mysql *gorm.DB

func InitTDB() {
	log.Println("Connecting to mysql")
	var err error
	dsn := config.Config.Mysql.Dsn
	Mysql, err = gorm.Open(mysql.New(
		mysql.Config{
			DSN:                       dsn,   // DSN data source name
			DefaultStringSize:         256,   // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
		}), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	_ = Mysql.Callback().Row().After("gorm:row").Register("after_row", After)

	log.Println("Connected to mysql")

}

func After(db *gorm.DB) {
	db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	log.Println(sql)
}
