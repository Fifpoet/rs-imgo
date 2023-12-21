package util

import (
	"fmt"
	"os"
)

func CreateDir(path string) error {
	// 检查路径是否已经存在
	fileInfo, err := os.Stat(path)
	if err == nil {
		// 如果已经存在同名文件或文件夹，则返回错误
		if fileInfo.IsDir() {
			return fmt.Errorf("文件夹已经存在：%s", path)
		} else {
			return fmt.Errorf("同名文件已经存在：%s", path)
		}
	}

	// 创建文件夹
	err = os.Mkdir(path, 0755)
	if err != nil {
		return fmt.Errorf("创建文件夹失败：%s", err.Error())
	}

	return nil
}
