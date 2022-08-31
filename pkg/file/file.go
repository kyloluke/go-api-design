package file

import (
	"os"
	"path/filepath"
	"strings"
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
	_, err := os.Stat(filePath) // 文件不存在时 err = &fs.PathError{Op:"CreateFile", Path:"app/cmd/test_command.go", Err:0x2}
	if os.IsNotExist(err) {     // 文件不存在时， true
		return false
	}
	return true
}

func FileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
