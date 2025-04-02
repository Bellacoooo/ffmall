package dal

import (
	"ffmall/app/product/biz/dal/mysql"
	"ffmall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
