package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime/multipart"

	"github.com/fishjar/gin-boilerplate/config"
)

// MD5Pwd 密码哈希函数
func MD5Pwd(username string, password string) string {
	salt := config.Config.PasswordSalt
	m := md5.New()
	m.Write([]byte(username))
	m.Write([]byte(password))
	m.Write([]byte(password))
	m.Write([]byte(salt))
	m.Write([]byte(salt))
	m.Write([]byte(salt))
	return hex.EncodeToString(m.Sum(nil))
}

// MD5File 计算文件MD5
func MD5File(f *multipart.FileHeader) (string, error) {
	src, err := f.Open()
	defer src.Close()
	if err != nil {
		return "", err
	}

	h := md5.New()
	if _, err := io.Copy(h, src); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
