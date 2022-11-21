package dao

import (
	"gin_demo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var GlobalDb2 *gorm.DB

func BoardLoad() {
	db, _ := gorm.Open(mysql.New(mysql.Config{ //配置
		DSN: "root:123456@tcp(127.0.0.1:3306)/gindemo?charset=utf8mb4&parseTime=True&loc=Local",
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "gindemo_",
			SingularTable: false,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(10) //数据池
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	GlobalDb2 = db   //全局变量GlobalDb赋值
	TestBoardCreat() //若无表单自动创建表单
}
func Board(username string, say string) {
	GlobalDb2.Model(&model.RealBoard{}).Create(&model.RealBoard{Username: username, Board: say})
}
