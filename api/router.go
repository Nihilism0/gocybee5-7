package api

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.POST("/register", register)             // 注册
	r.POST("/login", login)                   // 登录
	r.POST("/find", findpassword)             //找回密码
	r.POST("/changepassword", changepassword) //修改密码
	r.POST("/writeboard", writeboard)         //写留言
	r.POST("/seeboard", seeboard)             //看留言板
	r.POST("/addstore", addstore)             //增加商铺
	r.POST("/praise", praise)                 //用户点赞
	r.POST("/cancelpraise", cancelpraise)     //取消点赞
	r.Run(":9090")                            // 跑在 9090 端口上
}
