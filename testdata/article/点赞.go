package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/service/redis_ser"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	global.Redis = core.ConnectRedis()
	redis_ser.Digg("ddbht4cBnMkaexhcvNHZ")
}
