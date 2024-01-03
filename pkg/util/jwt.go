package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	UserName             string `json:"username"`
	jwt.RegisteredClaims        // v5版本新加的方法
}

const secretKey = "qwertyuiop"

// 生成JWT
func GenToken(username string) (string, error) {
	// 创建一个我们自己声明的数据
	claims := User{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 定义过期时间
			Issuer:    "somebody",                                         // 签发人
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	return token.SignedString([]byte(secretKey))
}

// 解析JWT
func ParseToken(tokenString string) (*User, error) {

	// 解析token
	var mc = new(User)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
