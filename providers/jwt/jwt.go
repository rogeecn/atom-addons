package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"

	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/sync/singleflight"
)

const (
	CtxKey     = "claims"
	HttpHeader = "Authorization"
)

type BaseClaims struct {
	UID  uint64 `json:"uid,omitempty"`
	Role uint64 `json:"role,omitempty"`
}

// Custom claims structure
type Claims struct {
	BaseClaims
	jwt.RegisteredClaims
}

const TOKEN_PREFIX = "Bearer "

type JWT struct {
	singleflight *singleflight.Group
	config       *Config
	SigningKey   []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	var config Config
	if err := o.UnmarshalConfig(&config); err != nil {
		return err
	}
	return container.Container.Provide(func() (*JWT, error) {
		return &JWT{
			config:     &config,
			SigningKey: []byte(config.SigningKey),
		}, nil
	}, o.DiOptions()...)
}

func (j *JWT) CreateClaims(baseClaims BaseClaims) *Claims {
	ep, _ := time.ParseDuration(j.config.ExpiresTime)
	claims := Claims{
		BaseClaims: baseClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(-time.Second * 10)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),                // 过期时间 7天  配置文件
			Issuer:    j.config.Issuer,                                       // 签名的发行者
		},
	}
	return &claims
}

// 创建一个token
func (j *JWT) CreateToken(claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims *Claims) (string, error) {
	v, err, _ := j.singleflight.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*Claims, error) {
	tokenString = strings.TrimPrefix(tokenString, TOKEN_PREFIX)
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
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
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}
