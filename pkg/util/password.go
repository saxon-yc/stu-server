package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

const (
	saltPassword    = "qkhPAGA13HocW3GAEWwb"
	defaultPassword = "123456"
)

func GeneratePassword(str string) (password string) {
	// 使用md5加密
	m := md5.New()
	m.Write([]byte(str))
	mByte := m.Sum(nil)

	// hmac 认证
	h := hmac.New(sha256.New, []byte(saltPassword))
	h.Write(mByte)
	password = hex.EncodeToString(h.Sum(nil))

	return
}

func ResetPassword() (password string) {
	m := md5.New()
	m.Write([]byte(defaultPassword))
	s := hex.EncodeToString(m.Sum(nil))

	password = s

	return
}

func GenerateToken(id uint32) (token string) {
	m := md5.New()
	m.Write([]byte(fmt.Sprintf("%d%s", id, saltPassword)))
	token = hex.EncodeToString(m.Sum(nil))

	return
}
