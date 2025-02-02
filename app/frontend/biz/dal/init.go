package dal

import (
	"github.com/xie628/gomall/app/frontend/biz/dal/mysql"
	"github.com/xie628/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
