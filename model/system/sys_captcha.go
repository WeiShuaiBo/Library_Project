package system

import "time"

type SysCaptcha struct {
	CaptchaId  string    `json:"captchaId" gorm:"primarykey;not null;comment:校验ID"` //校验ID 在检测验证码是否正确时候使用
	Code       string    `json:"code" gorm:"not null;comment:验证码值"`                 //验证码值
	ExpireTime time.Time `json:"expireTime" gorm:"not null;comment:过期时间"`           // 过期时间
}

func (receiver *SysCaptcha) TableName() string {
	return "sys_captcha"
}
