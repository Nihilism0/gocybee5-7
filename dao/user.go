package dao

import (
	"time"

	"gin_demo/model"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var GlobalDb1 *gorm.DB
var redisDb *redis.Client

func DataLoad() {
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
	GlobalDb1 = db  //全局变量GlobalDb赋值
	TestUserCreat() //若无表单自动创建表单
}

// 根据redis配置初始化一个客户端
func initClient() (err error) {
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "49.234.42.190:6379", // redis地址
		Password: "000415",             // redis密码，没有则留空
		DB:       0,                    // 默认数据库，默认是0
	})

	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	_, err = redisDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func RedisLoad() {
	err := initClient()
	if err != nil {
		//redis连接错误
		panic(err)
	}
}

func TestUserCreat() {
	GlobalDb1.AutoMigrate(&model.User{})
}
func TestBoardCreat() {
	GlobalDb2.AutoMigrate(&model.RealBoard{})
}

func AddUser(username, password string) {
	GlobalDb1.Model(&model.User{}).Create(&model.User{
		Username: username,
		Password: password,
	})
}

// 若没有这个用户返回 false，反之返回 true
func SelectUser(username string) bool {
	var u struct {
		Username string
	}
	GlobalDb1.Model(&model.User{}).Where("username = ?", username).Find(&u)
	if u.Username == "" {
		return false
	} else {
		return true
	}
}

func SelectPasswordFromUsername(username string) string {
	var u struct {
		Password string
	}
	GlobalDb1.Model(&model.User{}).Where("username = ?", username).Find(&u)
	return u.Password
}

func ChangePassword(username string, newpassword string) {
	GlobalDb1.Model(&model.User{}).Where("username = ?", username).Update("password", newpassword)
}

func FindPassword(username string) string {
	var u struct {
		Password string
	}
	GlobalDb1.Model(&model.User{}).Where("username = ?", username).Find(&u)
	return u.Password
}
