package api

import (
	"fmt"

	"gin_demo/dao"
	"gin_demo/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"message": "verification failed",
		})
		return
	}
	// 传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 验证用户名是否重复
	flag := dao.SelectUser(username)
	fmt.Println(flag)
	if flag {
		// 以 JSON 格式返回信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user already exists",
		})
		return
	}

	dao.AddUser(username, password)
	// 以 JSON 格式返回信息
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "add user successful",
	})
}

func login(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"message": "verification failed",
		})
		return
	}
	// 传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 验证用户名是否存在
	flag := dao.SelectUser(username)
	// 不存在则退出
	if !flag {
		// 以 JSON 格式返回信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user doesn't exists",
		})
		return
	}

	// 查找正确的密码
	selectPassword := dao.SelectPasswordFromUsername(username)
	// 若不正确则传出错误
	if selectPassword != password {
		// 以 JSON 格式返回信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "wrong password",
		})
		return
	}
	// 正确则登录成功
	c.SetCookie("gin_demo_cookie", "test", 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "login successful",
	})
}

func changepassword(c *gin.Context) {
	if err := c.ShouldBind(&model.Changepassword{}); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"message": "I do not know your hole message",
		})
		return
	}
	username := c.PostForm("username")
	flag := dao.SelectUser(username)
	if !flag {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user doesn't exists",
		})
		return
	}
	newpassword := c.PostForm("newpassword")
	password := c.PostForm("password")
	selectPassword := dao.SelectPasswordFromUsername(username)
	// 若不正确则传出错误
	if selectPassword != password {
		// 以 JSON 格式返回信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "wrong password",
		})
		return
	}
	dao.ChangePassword(username, newpassword)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "change successful!",
	})
}

func findpassword(c *gin.Context) {
	if err := c.ShouldBind(&model.Findpassword{}); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"message": "I even do not know who are you",
		})
		return
	}
	// 传入用户名
	username := c.PostForm("username")
	v := dao.FindPassword(username)
	c.JSON(http.StatusOK, gin.H{
		"status":        200,
		"your password": v,
	})
}
