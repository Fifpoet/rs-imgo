package api

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"rs-imgo/global"
	"rs-imgo/infra"
	"rs-imgo/util"
)

func GetTilePNGByDisk(c *gin.Context) {
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

	c.File(imgPath)
}

func GetTilePNGByMysql(c *gin.Context) {
	z := util.Str2Int(c.Param("z"))
	x := util.Str2Int(c.Param("x"))
	y := util.Str2Int(c.Param("y"))
	maxXY := 2<<z - 1
	if x > maxXY || y > maxXY || z > global.MaxScale {
		c.Data(http.StatusInternalServerError, "image/png", []byte{})
		c.Abort()
		return
	}
	//四进制编码
	quadKey := util.TileXY2QuadKey(x, y, z)

	png := infra.QueryB64ByQuadKey(quadKey)
	pngBytes, _ := base64.StdEncoding.DecodeString(png)
	c.Data(http.StatusOK, "image/png", pngBytes)
}
