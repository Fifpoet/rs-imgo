package util

import (
	"fmt"
	"golang.org/x/image/draw"
	"golang.org/x/image/tiff"
	"image"
	"image/png"
	"log"
	"os"
)

type Coordinate struct {
	x      int
	y      int
	output string
}

func ExtractMaxSquare(tiffPath, out string) (*image.RGBA, error) {
	file, err := os.Open(tiffPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, err := tiff.Decode(file)
	if err != nil {
		return nil, err
	}
	// 获取原始图像尺寸
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	// 计算最大正方形的边长
	size := 256
	for size <= width && size <= height {
		size <<= 1
	}
	size >>= 1
	// 提取最大正方形
	square := image.NewRGBA(image.Rect(0, 0, size, size))
	log.Printf("TIFF提取最大正方形：%dpx\n", size)
	draw.Draw(square, square.Bounds(), img, image.Point{0, 0}, draw.Src)
	outfile, _ := os.Create(out)
	err = png.Encode(outfile, square)
	if err != nil {
		log.Fatalf("写入大png失败")
		return nil, err
	}
	return square, nil
}

// DecomposeSquare 一张256 * 2^n的正方形，分解到文件夹
func DecomposeSquare(img *image.RGBA, size, x0, y0 int, curPath string, code string) {
	subsize := size >> 1
	if size == 256 {
		return
	}
	//存储本层的略缩图
	StoreSubSquare(img, size, x0, y0, curPath, code)
	//递归四个子图像
	if subsize != 256 {
		for i := 0; i < 4; i++ {
			CreateDir(fmt.Sprintf("%s%d", curPath, i))
		}
	}
	DecomposeSquare(img, subsize, 0, 0, curPath+"0/", code+"0")
	DecomposeSquare(img, subsize, subsize, 0, curPath+"1/", code+"1")
	DecomposeSquare(img, subsize, 0, subsize, curPath+"2/", code+"2")
	DecomposeSquare(img, subsize, subsize, subsize, curPath+"3/", code+"3")
	return
}

// StoreSubSquare 根据size取coors右下的正方形并重采样存储成256*256的略缩图
// eg. code=0132 的文件夹下，递归有四个子文件夹和四个子略缩图, 本函数用来存储01320.png 01321.png ...
func StoreSubSquare(img *image.RGBA, size, x0, y0 int, curPath string, code string) {
	subsize := size >> 1
	coors := []Coordinate{{x: x0, y: y0}, {x: x0 + subsize, y: y0}, {x: x0, y: y0 + subsize}, {x: x0 + subsize, y: y0 + subsize}}
	for i := range coors {
		//TODO 不能直接使用copy对象
		coors[i].output = fmt.Sprintf("%s%s%d.png", curPath, code, i)
	}
	for _, c := range coors {
		subSquare := img.SubImage(image.Rect(c.x, c.y, c.x+subsize, c.y+subsize)).(*image.RGBA)
		//最邻近像元重采样
		dstNearest := image.NewRGBA(image.Rect(0, 0, 256, 256))
		draw.NearestNeighbor.Scale(dstNearest, dstNearest.Bounds(), subSquare, subSquare.Bounds(), draw.Src, nil)
		file, err := os.Create(c.output)
		if err != nil {
			log.Fatalf("创建%s文件失败：%s\n", c.output, err.Error())
		}
		err = png.Encode(file, dstNearest)
		if err != nil {
			log.Fatalf("绘制略缩图失败：%s\n", err.Error())
		}
		file.Close()
	}
}
