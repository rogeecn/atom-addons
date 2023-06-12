package captcha

import (
	"log"
	"time"

	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

const DefaultPrefix = "Captcha"

func DefaultProvider() container.ProviderContainer {
	return container.ProviderContainer{
		Provider: Provide,
		Options: []opt.Option{
			opt.Prefix(DefaultPrefix),
		},
	}
}

type Config struct {
	Long        uint    // 验证码长度
	Width       uint    // 验证码宽度
	Height      uint    // 验证码高度
	Open        uint    // 防爆破验证码开启此数，0代表每次登录都需要验证码，其他数字代表错误密码此数，如3代表错误三次后出现验证码
	OpenTimeOut string  // 防爆破验证码超时时间，单位：s(秒)
	MaxScrew    float64 // MaxSkew max absolute skew factor of a single digit.
	DotCount    int     // Number of background circles.
}

func (c *Config) OpenCaptchaTimeOutDuration() time.Duration {
	d, err := time.ParseDuration(c.OpenTimeOut)
	if err != nil {
		log.Panic(err)
	}
	return d
}
