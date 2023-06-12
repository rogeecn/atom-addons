package captcha

import (
	"errors"
	"log"
	"time"

	"github.com/mojocn/base64Captcha"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
	"github.com/spf13/viper"
)

type CaptchaResponse struct {
	CaptchaId     string `json:"captcha_id,omitempty"`
	PicPath       string `json:"pic_path,omitempty"`
	CaptchaLength uint   `json:"captcha_length,omitempty"`
	OpenCaptcha   uint   `json:"open_captcha,omitempty"`
}

type Captcha struct {
	captcha *base64Captcha.Captcha
}

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	var conf Config
	if err := o.UnmarshalConfig(&conf); err != nil {
		return err
	}

	return container.Container.Provide(func() (*Captcha, error) {
		driver := base64Captcha.NewDriverDigit(
			int(conf.Width),
			int(conf.Height),
			int(conf.Long),
			conf.MaxScrew,
			conf.DotCount,
		)

		store := base64Captcha.DefaultMemStore
		return &Captcha{
			captcha: base64Captcha.NewCaptcha(driver, store),
		}, nil
	}, o.DiOptions()...)
}

func (c *Captcha) OpenCaptchaTimeOutDuration() time.Duration {
	d, err := time.ParseDuration(viper.GetString("CAPTCHA_IMG_OPEN_TIMEOUT"))
	if err != nil {
		log.Panic(err)
	}
	return d
}

func (c *Captcha) Generate() (*CaptchaResponse, error) {
	id, b64s, err := c.captcha.Generate()
	if err != nil {
		return nil, errors.New("验证码获取失败")
	}

	return &CaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: viper.GetUint("CAPTCHA_IMG_KEY_LONG"),
		OpenCaptcha:   viper.GetUint("CAPTCHA_IMG_OPEN"),
	}, nil
}

func (c *Captcha) Verify(id, answer string) bool {
	return c.captcha.Verify(id, answer, false)
}
