package dao

func Praiseadd(storename, username string) {
	redisDb.SAdd(storename, username)
}

func SelectPraiseuser(storename, username string) bool {
	flag, _ := redisDb.SIsMember(storename, username).Result()
	return flag
}
func CancelPraise(storename, username string) {
	redisDb.SRem(storename, username)
}
