package dao

func SelectStore(storename string) bool {
	flag, _ := redisDb.SIsMember("storesum", storename).Result()
	return flag
}
func Storeadd(storename string) {
	redisDb.SAdd("storesum", storename)
}
