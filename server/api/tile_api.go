package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rs-imgo/global"
	"rs-imgo/infra"
	"rs-imgo/util"
	"strconv"
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
	key := global.ZsetKeyPrefix + strconv.Itoa(len(quadKey))
	pngs := infra.QueryPngByScore(key, quadKey)
	if len(pngs) == 1 {
		c.String(http.StatusOK, pngs[0])
	} else {
		score, _ := strconv.Atoi(quadKey)
		infra.ZAddBatchPng(key, []string{imgPath}, []int{score})
	}
}
