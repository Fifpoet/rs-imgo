package core

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"rs-imgo/api"
)

func RunGinServer() {
	r := gin.Default()
	v1 := r.Group("/v1")
	v1.GET("/map/:z/:x/:y", api.GetTilePNGByDisk)
	v1.GET("/cache/:z/:x/:y", api.UpdateCache)
	pprof.Register(r)
	r.Run(":8888")
}
