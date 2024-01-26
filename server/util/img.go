package util

import (
	"encoding/base64"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

const fmtSize = 256

// JoinImages 合成四张256的矩形图片
func JoinImages(ulP, urP, dlP, drP, outputPath string) error {
	// 打开并解码图片文件
	ul, err := openAndDecodeImage(ulP)
	ur, err := openAndDecodeImage(urP)
	dl, err := openAndDecodeImage(dlP)
	dr, err := openAndDecodeImage(drP)
	// 检查图片尺寸
	if ul.Bounds().Size() != image.Pt(256, 256) ||
		ur.Bounds().Size() != image.Pt(256, 256) ||
		dl.Bounds().Size() != image.Pt(256, 256) ||
		dr.Bounds().Size() != image.Pt(256, 256) {
		return fmt.Errorf("image size should be 256x256")
	}
	// 创建输出图像
	outputImg := image.NewRGBA(image.Rect(0, 0, fmtSize*2, fmtSize*2))

	// 在指定位置绘制图片
	draw.Draw(outputImg, ul.Bounds(), ul, image.ZP, draw.Src)
	draw.Draw(outputImg, ur.Bounds().Add(image.Pt(256, 0)), ur, image.ZP, draw.Src)
	draw.Draw(outputImg, dl.Bounds().Add(image.Pt(0, 256)), dl, image.ZP, draw.Src)
	draw.Draw(outputImg, dr.Bounds().Add(image.Pt(256, 256)), dr, image.ZP, draw.Src)

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// 根据输出文件的扩展名选择合适的编码器
	switch ext := getFileExtension(outputPath); ext {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(outputFile, outputImg, &jpeg.Options{Quality: 100})
	case ".png":
		err = png.Encode(outputFile, outputImg)
	default:
		err = fmt.Errorf("unsupported output file format")
	}
	if err != nil {
		return err
	}
	return nil
}

// 打开并解码图片文件
func openAndDecodeImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		panic("cannot open img.jpg")
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// 获取文件的扩展名
func getFileExtension(path string) string {
	for i := len(path) - 1; i >= 0 && path[i] != '/'; i-- {
		if path[i] == '.' {
			return path[i:]
		}
	}
	return ""
}
func ResizeImg(imgP string) {
	// 加载原始图像
	srcImage, err := imaging.Open(imgP)
	if err != nil {
		fmt.Println("无法打开图像:", err)
		return
	}
	targetSize := image.Point{X: 256, Y: 256}
	resizedImage := imaging.Resize(srcImage, targetSize.X, targetSize.Y, imaging.Lanczos)
	// 保存
	err = imaging.Save(resizedImage, "fmt.jpg")
	if err != nil {
		fmt.Println("无法保存图像:", err)
		return
	}
	fmt.Println("图像重采样完成")
}

func DecomposeTiff(tiffPath string, outPath string) {
	rgba, err := ExtractMaxSquare(tiffPath, outPath+"tiff.png")
	if err != nil {
		return
	}
	DecomposeSquare(rgba, rgba.Bounds().Dx(), 0, 0, outPath, "")
}

func PathToB64(path string) string {
	bs, _ := os.ReadFile(path)
	b64 := base64.StdEncoding.EncodeToString(bs)
	return b64
}
