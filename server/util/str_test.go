package util

import "testing"

func TestTileXY2QuadKey(t *testing.T) {
	code := TileXY2QuadKey(2, 2, 2)
	println(code)
}

func TestQuadKey2ImgPath(t *testing.T) {
	key := "01012312"
	res := QuadKey2ImgPath(key)
	println(res)
}
