package app

import (
	"blog/global"
	"blog/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	APPSecret string `json:"app_secret"`
	jwt.StandardClaims
}

// GetJWTSecret 获取JWT的Secret
func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

// GenerateToken 生成 JWT Token 的行为
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		APPSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	tokeClaims := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token, err := tokeClaims.SignedString(GetJWTSecret())
	return token, err
}

// ParseToken 解析和校验 Token
func ParseToken(token string) (*Claims, error) {
	//ParseWithClaims：用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回 *Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		//Valid：验证基于时间的声明，例如：过期时间（ExpiresAt）、签发者（Issuer）、生效时间（Not Before）
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
