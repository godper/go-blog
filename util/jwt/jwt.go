package ujwt

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var signKey string = "wgreqfrefq"

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

//NewJWT 新建一个jwt实例
func NewJWT(secret string) *JWT {
	SetSignKey(secret)
	return &JWT{
		[]byte(GetSignKey()),
	}
}

//GetSignKey rdasf
func GetSignKey() string {
	return signKey
}

//SetSignKey  fa
func SetSignKey(key string) string {
	signKey = key
	return signKey
}

//CustomClaims 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	ID       uint   `json:"Id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

//GenerateToken 生成令牌
func (j *JWT) GenerateToken(data map[string]interface{}) (token string, err error) {
	ID, _ := data["ID"].(uint)
	Name, _ := data["Name"].(string)
	Nickname, _ := data["Nickname"].(string)
	Role, _ := data["Role"].(string)

	claims := CustomClaims{
		ID,
		Name,
		Nickname,
		Role,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),  // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 14400), // 过期时间四小时
			Issuer:    "godper.com",                     //签名的发行者
		},
	}
	token, err = j.CreateToken(claims)
	return
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 一些常量
var (
	ErrTokenExpired     error = errors.New("Token is expired")
	ErrTokenNotValidYet error = errors.New("Token not active yet")
	ErrTokenMalformed   error = errors.New("That's not even a token")
	ErrTokenInvalid     error = errors.New("Couldn't handle this token")
)

//ParseToken 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				if claims, ok := token.Claims.(*CustomClaims); ok {
					return claims, ErrTokenExpired
				}
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

//RefreshToken 更新token
func (j *JWT) RefreshToken(claims *CustomClaims) (string, error) {
	// jwt.TimeFunc = func() time.Time {
	// 	return time.Unix(0, 0)
	// }
	// token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
	// 	return j.SigningKey, nil
	// })
	// if err != nil {
	// 	return "", err
	// }
	// if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
	// 	jwt.TimeFunc = time.Now
	claims.StandardClaims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
	return j.CreateToken(*claims)
	// }
	// return "", ErrTokenInvalid
}
