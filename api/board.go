package api

import (
	"gin_demo/dao"
	"gin_demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func writeboard(c *gin.Context) {
	if err := c.ShouldBind(&model.Board{}); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"message": "Want to speak anonymoisly?no way!",
		})
		return
	}
	// 传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")
	say := c.PostForm("board")
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
	dao.Board(username, say)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "your message\"" + say + "\"have gone to the board",
	})
}

func seeboard(c *gin.Context) {
	type Info struct {
		ID       int
		Username string
		Board    string
	}
	i := 0
	var b []Info
	dao.GlobalDb2.Model(&model.RealBoard{}).Find(&b)
	for {
		if b[i].Username == "" {
			break
		} else {
			c.JSON(http.StatusOK, gin.H{
				"ID":      b[i].ID,
				"User":    b[i].Username,
				"context": b[i].Board,
			})
		}
		i++
	}
}
