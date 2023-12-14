package util

import "testing"

func TestResizeImg(t *testing.T) {
	img := "../static/img/img.jpg"
	ResizeImg(img)
}

func TestJoinImages(t *testing.T) {
	ul := "../static/img/fmt.jpg"
	out := "../static/img/join.jpg"
	JoinImages(ul, ul, ul, ul, out)
}
