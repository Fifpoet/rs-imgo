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

func TestDecomposeTiff(t *testing.T) {
	tiff := "/home/fifpoet/Downloads/files/default/file_example_TIFF_10MB.tiff"
	out := "/home/fifpoet/Desktop/go_project/rs-imgo/server/static/output/"
	DecomposeTiff(tiff, out)
}
