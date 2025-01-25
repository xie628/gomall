package main

import (
	"github.com/joho/godotenv"
	"github.com/xie628/gomall/demo/demo_proto/biz/dal"
	"github.com/xie628/gomall/demo/demo_proto/biz/dal/model"
	"github.com/xie628/gomall/demo/demo_proto/biz/dal/mysql"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()
	//mysql.DB.Create(&model.User{Email: "test@test.com", Password: "123456"})
	mysql.DB.Model(&model.User{}).Where("email = ?", "test@test.com").Update("password", "test")
}
