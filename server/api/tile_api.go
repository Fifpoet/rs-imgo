package api

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"rs-imgo/global"
	"rs-imgo/infra"
	"rs-imgo/util"
	"strconv"
)

func GetTilePNGByZset(c *gin.Context) {
	z := util.Str2Int(c.Param("z"))
	x := util.Str2Int(c.Param("x"))
	y := util.Str2Int(c.Param("y"))
	base := global.ImgPath
	maxXY := 2<<z - 1
	if x > maxXY || y > maxXY || z > global.MaxScale {
		c.Data(http.StatusInternalServerError, "image/png", []byte{})
		c.Abort()
		return
	}
	//四进制编码
	quadKey := util.TileXY2QuadKey(x, y, z)
	imgPath := base + util.QuadKey2ImgPath(quadKey)
	key := global.ZsetKeyPrefix + strconv.Itoa(len(quadKey))

	pngs := infra.QueryPngByScore(key, quadKey)
	if len(pngs) == 1 {
		pngBytes, _ := base64.StdEncoding.DecodeString(pngs[0])
		c.Data(http.StatusOK, "image/png", pngBytes)
	} else {
		score, _ := strconv.Atoi(quadKey)
		infra.ZAddBatchPng(key, []string{imgPath}, []int{score})
		c.Redirect(http.StatusFound, c.Request.URL.Path)
	}
}

func GetTilePNGByHash(c *gin.Context) {
	z := util.Str2Int(c.Param("z"))
	x := util.Str2Int(c.Param("x"))
	y := util.Str2Int(c.Param("y"))
	base := global.ImgPath
	maxXY := 2<<z - 1
	if x > maxXY || y > maxXY || z > global.MaxScale {
		c.Data(http.StatusInternalServerError, "image/png", []byte{})
		c.Abort()
		return
	}
	//四进制编码
	quadKey := util.TileXY2QuadKey(x, y, z)
	imgPath := base + util.QuadKey2ImgPath(quadKey)
	key := global.HashKeyPrefix + strconv.Itoa(len(quadKey))

	png := infra.QueryPngByField(key, quadKey)
	if png != "" {
		pngBytes, _ := base64.StdEncoding.DecodeString(png)
		c.Data(http.StatusOK, "image/png", pngBytes)
	} else {
		infra.HAddPng(key, imgPath, quadKey)
		c.Redirect(http.StatusFound, c.Request.URL.Path)
	}
}
