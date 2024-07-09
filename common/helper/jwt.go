package helper

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"time"
)

type MapClaims struct {
	Key        string            `json:"key"`         // 唯一值
	Time       int64             `json:"time"`        //当前时间
	ExpireTime int64             `json:"expire_time"` //到期时间
	Nonce      string            `json:"nonce"`       //随机值
	Data       map[string]string `json:"data"`        //自定义参数
	jwt.RegisteredClaims
}

/**
 * key 唯一索引值
 * data 自定义参数
 * jwtSecret 自定义密钥
 * ttl有效期
**/
func GenToken(jwtSecret string, key string, data map[string]string, ttl time.Duration) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(ttl * time.Second)

	claims := MapClaims{
		Key:        key,
		Time:       nowTime.Unix(),
		Data:       data,
		ExpireTime: expireTime.Unix(),
		Nonce:      Md5(nowTime.String() + key + expireTime.String()),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "verify",                                        //签发者
			Subject:   "any",                                           //签发给谁，例如某个人
			Audience:  jwt.ClaimStrings{"any"},                         //签发给谁，例如某个服务
			ExpiresAt: jwt.NewNumericDate(expireTime),                  //过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)), //最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                  //创建时间
			ID:        key,                                             //随机数 尽量唯一
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	return tokenString, err
}

func ParseToken(tokenString string, jwtSecret string) (*MapClaims, error) {
	claims := new(MapClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token失效")
	}
	return claims, nil
}
