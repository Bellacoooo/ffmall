package mysql

import (
	"ffmall/app/frontend/biz/dal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/ffmall?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&model.User{}) //自动创建表
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&model.Product{})
	if err != nil {
		panic(err)
	}
	products := []model.Product{
		{Name: "T-shirt-1", Price: 100, Picture: "/static/image/t-shirt-1.jpeg"},
		{Name: "T-shirt-2", Price: 110, Picture: "/static/image/t-shirt-1.jpeg"},
		{Name: "T-shirt-3", Price: 120, Picture: "/static/image/t-shirt-2.jpeg"},
	}
	for _, product := range products {
		err = DB.Create(&product).Error
		if err != nil {
			panic(err)
		}
	}
}
