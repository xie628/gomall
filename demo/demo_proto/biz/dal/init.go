package dal

import (
	"github.com/xie628/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/xie628/gomall/demo/demo_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
