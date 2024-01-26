package main

import (
	"rs-imgo/core"
	"rs-imgo/infra"
)

func main() {

	infra.InitRedis()
	infra.InitMysql()
	//启动http服务
	core.RunGinServer()
}
