// @Author zhangjiaozhu 2023/7/4 9:37:00
package SendEmail

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"net/textproto"
)

func SendEmail(toUserEmail, code string) error {
	host := "smtp.qq.com" // qq邮箱smtp服务器地址
	//port := "25"  //非SSL协议端口
	port := "465" //SSl协议端口
	userName := "2963754309@qq.com"
	//password := "tkrihxlwpjphbadd" // qq邮箱填授权码（非密码）
	password := "gljbmfzqlnwadcgh" // qq邮箱填授权码（非密码）

	e := &email.Email{
		To:      []string{toUserEmail},
		From:    userName,
		Subject: "Email Send Test",                                 // 发送内容标题
		Text:    []byte("Text Body is, of course, supported!"),     //信息内容
		HTML:    []byte("<h1>This is your code:" + code + "</h1>"), //信息内容 （html格式）
		Headers: textproto.MIMEHeader{},
	}
	// 非SSL协议端口25
	//err := e.Send(host+":"+port, smtp.PlainAuth("", userName, password, host))
	//if err != nil {
	//	panic(err)
	//}
	// 使用SSl协议端口 465/587
	return e.SendWithTLS(host+":"+port, smtp.PlainAuth("", userName, password, host), &tls.Config{InsecureSkipVerify: true, ServerName: host})

}
