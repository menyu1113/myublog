package utiljwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"myublog/global/vipers"
)

type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(vipers.JwtKey),
	}
}

type MyClaims struct {
	Username string `json:"username"`
	//Id uint `json:"id"`
	jwt.StandardClaims
}

// 定义错误
var (
	TokenErrorsExist   error = errors.New("Token不存在")
	TokenExpired       error = errors.New("Token已过期,请重新登录")
	TokenNotValidYet   error = errors.New("Token无效,请重新登录")
	TokenMalformed     error = errors.New("Token不正确,请重新登录")
	TokenInvalid       error = errors.New("这不是一个Token,请重新登录")
	TokenNoCurrentUser error = errors.New("Token与ID不匹配")
)

// CreateToken 生成token
func (j *JWT) CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

// ParserToken 解析token
func (j *JWT) ParserToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}
	return nil, TokenInvalid
}
