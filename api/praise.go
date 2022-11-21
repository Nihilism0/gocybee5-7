package api

import (
	"gin_demo/dao"
	"gin_demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func praise(c *gin.Context) {
	if err := c.ShouldBind(&model.Store{}); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"message": "I do not know your fucking store name bro.",
		})
		return
	}

	// 传入商铺名,用户名
	username := c.PostForm("username")
	storename := c.PostForm("storename")
	// 验证商铺名是否存在
	flag1 := dao.SelectStore(storename)
	// 不存在则退出
	if !flag1 {
		// 以 JSON 格式返回信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Store do not exists",
		})
		return
	}
	// 验证用户是否点赞
	flag2 := dao.SelectPraiseuser(storename, username)
	// 用户点过赞则退出
	if flag2 {
		// 以 JSON 格式返回信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "You have praise the store :" + storename,
		})
		return
	}
	dao.Praiseadd(storename, username)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "your praise to" + storename + "have received",
	})
}
func cancelpraise(c *gin.Context) {
	if err := c.ShouldBind(&model.Store{}); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"message": "I do not know your fucking store name bro.",
		})
		return
	}

	// 传入商铺名,用户名
	username := c.PostForm("username")
	storename := c.PostForm("storename")
	// 验证商铺名是否存在
	flag1 := dao.SelectStore(storename)
	// 不存在则退出
	if !flag1 {
		// 以 JSON 格式返回信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Store do not exists",
		})
		return
	}
	// 验证用户是否点赞
	flag2 := dao.SelectPraiseuser(storename, username)
	// 用户没点过赞则退出
	if !flag2 {
		// 以 JSON 格式返回信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "You did not praise the store :" + storename,
		})
		return
	}
	dao.CancelPraise(storename, username)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "you have canceled the praise to" + storename,
	})
}
