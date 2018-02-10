package tool

import (
	"log"
	"os"
)

// PathExist ， 判断文件是否存在
func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// MustCreateDir , 创建文件夹，不返回错误
func MustCreateDir(path string) {
	exist, err := PathExist(path)
	if err != nil {
		log.Fatalln(path, err)
	}
	if !exist {
		os.MkdirAll(path, 0777)
	}
}
