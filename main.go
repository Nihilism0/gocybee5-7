package main

import (
	"gin_demo/api"
	"gin_demo/dao"
)

func main() {
	dao.RedisLoad()
	dao.DataLoad()
	dao.BoardLoad()
	api.InitRouter()
}
