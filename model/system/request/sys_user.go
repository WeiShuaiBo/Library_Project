package request

// User login structure
type Login struct {
	UserId   int    `json:"UserId" form:"UserId"`     // 用户名
	Password string `json:"Password" form:"Password"` // 密码
	//Captcha   string `json:"captcha"`   // 验证码
	//CaptchaId string `json:"captchaId"` // 验证码ID

}
