package file

import (
	"fmt"
	"os"
)

// Put 将数据写入文件，会自动创建该文件
func Put(data []byte, to string) error {
	err := os.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Exists 判断文件是否存在
func Exists(filePath string) bool {
	_, err := os.Stat(filePath)
	fmt.Printf("os.Stat 返回：%#v/n", err)
	fmt.Printf("os.IsNotExist 返回：%#v/n", os.IsNotExist(err))
	if os.IsNotExist(err) {
		return false
	}
	return true
}
