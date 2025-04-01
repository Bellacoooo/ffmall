package dal

import (
	"ffmall/app/frontend/biz/dal/mysql"
	"ffmall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
