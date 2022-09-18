package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/pkg/app"
	"gohub/pkg/auth"
	"gohub/pkg/helpers"
	"mime/multipart"
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

// SaveUploadAvatar 保存用户头像
func SaveUploadAvatar(c *gin.Context, file *multipart.FileHeader) (string, error) {
	var avatar string
	// 确保目录存在不存在则创建
	publicPath := "public"
	dirName := fmt.Sprintf("/uploads/avatars/%s/%s/", app.TimenowInTimezone().Format("2006/01/02"), auth.CurrentUID(c))
	os.MkdirAll(publicPath+dirName, 0755)
	// 保存文件
	fileName := helpers.RandomString(16) + filepath.Ext(file.Filename)

	avatarPath := publicPath + dirName + fileName

	if err := c.SaveUploadedFile(file, avatarPath); err != nil {
		return avatar, err
	}

	return avatarPath, nil
}
