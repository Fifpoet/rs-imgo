package util

import (
	"strconv"
	"strings"
)

func Str2Int(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}

func TileXY2QuadKey(tileX, tileY, lvl int) string {
	var quadKey strings.Builder
	for i := lvl; i > 0; i-- {
		digit := '0'
		mask := 1 << (i - 1)
		if (tileX & mask) != 0 {
			digit++
		}
		if (tileY & mask) != 0 {
			digit += 2
		}
		quadKey.WriteRune(digit)
	}
	return quadKey.String()
}

// QuadKey2ImgPath 四进制编码 转化为Img路径
// eg. 102013  -> 1/0/2/0/1/102013.png
func QuadKey2ImgPath(str string) string {
	//去除最后一位，余下用/分割
	var dir string
	if len(str) > 0 {
		dir = str[:len(str)-1]
	}
	dir = strings.Join(strings.Split(dir, ""), "/")
	return dir + "/" + str + ".png"
}
