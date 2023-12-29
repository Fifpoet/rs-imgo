package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"rs-imgo/global"
	"rs-imgo/util"
)

func GetTilePNG(c *gin.Context) {
	z := util.Str2Int(c.Param("z"))
	x := util.Str2Int(c.Param("x"))
	y := util.Str2Int(c.Param("y"))
	base := global.ImgPath
	maxXY := 2<<z - 1
	if x > maxXY || y > maxXY {
		log.Printf("请求越界： %d %d %d", z, x, y)
		c.Abort()
	}
	//四进制编码
	quadKey := util.TileXY2QuadKey(x, y, z)
	imgPath := base + util.QuadKey2ImgPath(quadKey)
	c.File(imgPath)
}
