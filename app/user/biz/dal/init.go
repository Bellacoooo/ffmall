package dal

import (
	"ffmall/app/user/biz/dal/mysql"
	"ffmall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
