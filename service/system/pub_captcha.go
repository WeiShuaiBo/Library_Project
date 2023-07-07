package system

import (
	"Library_Project/global"
	"Library_Project/model/system"
	"errors"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"time"
)

type CaptchaService struct {
}

var result = base64Captcha.NewMemoryStore(20240, 3*time.Minute)

// mathConfig 生成图形化算术验证码配置
func mathConfig() *base64Captcha.DriverMath {
	mathType := &base64Captcha.DriverMath{
		Height:          50,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowHollowLine,
		BgColor: &color.RGBA{
			R: 40,
			G: 30,
			B: 89,
			A: 29,
		},
		Fonts: nil,
	}
	return mathType
}

// digitConfig 生成图形化数字验证码配置
func digitConfig() *base64Captcha.DriverDigit {
	digitType := &base64Captcha.DriverDigit{
		Height:   50,
		Width:    100,
		Length:   5,
		MaxSkew:  0.45,
		DotCount: 100,
	}
	return digitType
}

// stringConfig 生成图形化字符串验证码配置
func stringConfig() *base64Captcha.DriverString {
	stringType := &base64Captcha.DriverString{
		Height:          100,
		Width:           50,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine,
		Length:          5,
		Source:          "123456789qwertyuiopasdfghjklzxcvb",
		BgColor: &color.RGBA{
			R: 40,
			G: 30,
			B: 89,
			A: 29,
		},
		Fonts: nil,
	}
	return stringType
}

// chineseConfig 生成图形化汉字验证码配置
func chineseConfig() *base64Captcha.DriverChinese {
	chineseType := &base64Captcha.DriverChinese{
		Height:          50,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowSlimeLine,
		Length:          2,
		Source:          "设想,你在,处理,消费者,的音,频输,出音,频可,能无,论什,么都,没有,任何,输出,或者,它可,能是,单声道,立体声,或是,环绕立,体声的,不想要,的值",
		BgColor: &color.RGBA{
			R: 40,
			G: 30,
			B: 89,
			A: 29,
		},
		Fonts: nil,
	}
	return chineseType
}

// autoConfig 生成图形化数字音频验证码配置
func autoConfig() *base64Captcha.DriverAudio {
	chineseType := &base64Captcha.DriverAudio{
		Length:   4,
		Language: "zh",
	}
	return chineseType
}

// @Result id 验证码id
// @Result bse64s 图片base64编码
// @Result err 错误
func (s *CaptchaService) GetCaptcha() (string, string, error) {
	var driver base64Captcha.Driver
	// switch case分支中的方法为目录3的配置
	// switch case分支中的方法为目录3的配置
	// switch case分支中的方法为目录3的配置
	switch global.FAST_VP.GetString("captcha.type") {
	case "audio":
		driver = autoConfig()
	case "string":
		driver = stringConfig()
	case "math":
		driver = mathConfig()
	case "chinese":
		driver = chineseConfig()
	case "digit":
		driver = digitConfig()
	}
	if driver == nil {
		driver = digitConfig()
		global.FAST_LOG.Info("未配置 ' code.captcha_type ' ,使用默认数字验证码配置")
	}
	// 创建验证码并传入创建的类型的配置，以及存储的对象
	c := base64Captcha.NewCaptcha(driver, result)
	id, b64s, err := c.Generate()
	return id, b64s, err
}

// VerifyCaptcha
// @Pram id 验证码id
// @Pram VerifyValue 用户输入的答案
// @Result true：正确，false：失败
func (s *CaptchaService) VerifyCaptcha(id, VerifyValue string) (isVerify bool, err error) {
	var captcha system.SysCaptcha
	captcha.CaptchaId = id
	err = global.FAST_DB.First(&captcha).Error
	if err != nil {
		global.FAST_LOG.Error("验证码不正确")
		return false, err
	}
	if time.Now().After(captcha.ExpireTime) {
		global.FAST_DB.Delete(&captcha)
		err = errors.New("验证码已过期，请重新获取")
		return false, err
	}
	// result 为步骤1 创建的图片验证码存储对象
	if !result.Verify(id, VerifyValue, true) {
		err = errors.New("验证码不正确")
		return false, err
	}
	global.FAST_DB.Delete(&captcha)
	return true, nil
}

// GetCodeAnswer
// @Pram codeId 验证码id
// @Result 验证码答案
func (s *CaptchaService) GetCodeAnswer(codeId string) string {
	// result 为步骤1 创建的图片验证码存储对象
	return result.Get(codeId, false)
}
