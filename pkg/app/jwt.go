package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-programing-tour-book/blog-service/global"
	"github.com/go-programing-tour-book/blog-service/pkg/util"
	"time"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 根据 Claims 结构体创建 Token 实例 第一个形参是加密算法方案
	token, err := tokenClaims.SignedString(GetJWTSecret())           // tokenClaims.SignedString 生成签名字符串 根据传入的 Secret 进行签名并返回标准的 Token
	return token, err
}

// ParseToken 解析和效验 Token
// 流程是解析传入的 Token 根据 Claims 的相关属性要求进行效验
func ParseToken(token string) (*Claims, error) {
	// ParseWithClaims 用于解析鉴权的声明 返回 *Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		// tokenClaims.Valid 验证基于时间的声明 如过期时间 ExpiresAt 签发者 Issuer 生效时间 Not Before 即便令牌中没有任何声明 也仍然被认为是有效的
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
