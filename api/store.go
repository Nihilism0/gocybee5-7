package api

import (
	"gin_demo/dao"
	"gin_demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func addstore(c *gin.Context) {
	if err := c.ShouldBind(&model.Store{}); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"message": "I do not know your fucking store name bro.",
		})
		return
	}
	// 传入商铺名
	storename := c.PostForm("storename")
	// 验证商铺名是否存在
	flag := dao.SelectStore(storename)
	// 存在则退出
	if flag {
		// 以 JSON 格式返回信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Store exists,you do not need add again.",
		})
		return
	}

	dao.Storeadd(storename)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "store" + storename + "have gone to the Storesum",
	})
}
