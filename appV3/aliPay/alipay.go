package aliPay

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"library/appV3/model"
	"log"
	"net/http"
)

var client *alipay.Client

const (
	kAppId      = "9021000123609903" //应用ID
	kPrivateKey = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCSoEVt9Nl++OQq8YcS1VrmyvzrRh8ZC8ZQv3dNDpm2i5hAg4iIlCzl/ChSy+sfR2qxEDB/cJwq5U0sw0hEq2jAqvVC1rgtj6S84C+6hXkOiwtIXJrVMupgW1MkLCq0Zbm69V6r6VG1T+8jD9rMtfDJfWnMHcBJ+qwDRJMwTxXccVd5o3S+7boxvLJsPhyWKfmbe6xmB6TRBzxWUDZM8rkWFiAr4nwnPutcP5k9fpVL6Dk3d42PxNh4Ucm7Q+ieg/YThdCyGKSPMVGHBVaZm5JUtaqt8GIprvwvrbMM8N5dDCOEyVlqWqUnR09Mq/yikB01NkwkRCAn8SelHNturkeFAgMBAAECggEAdzAS4AvNlJoLFyFYNCX8i2jr6PAKLVjV7yOEfb3lk26r551EKgmQ6a5stMkQKk/qWV4Yni9SssfMURu4riFLuHn/fkJ+WoLOXb467fq7aef61up37eBChusVjWzdleCu9luohkPV6HW+pRipOgiXX6IzkvmIKlq64rkmkHlpAtRfE9iiy3+vMhQ7QTioProa0X4/+H9E6zjhaXTply5H4fFBYLFV8dfpMVFDJ4IN6gZpcbSQhKins0MZNz60SP3N4fudtqqD7FM2umaqGnim4qXGUBiBwHgVkJMoagQ1G7Ol8iXwwgUtLyuPPirg0efrsE6Kutkig4LMqghV2HkXLQKBgQDT0/NVlFHGO5c67QkCpSkQF3KSSwVzRdEOaOyJD+v/2Cm1tKvoZIK3Urv9iblclUHDqtvKnUwqrESnOEM7StrppMxXvfSyY0kNuyb5j8vSsgHdOVFLDJNko0tlo29kslvlRtwan7H8mRlhPYTi5qctU66jHEvUjgnRXYBINg9PuwKBgQCxM6MbBMnBS0vwACqiAqNJBVvvlB+mG123KgSbEY6Rc6l6u3nNW5eMF0PbZj5IaZa0Z17+5hFXUGTcTXrmmXV48LbF9Z1VUWr48aWlsiRu3UXL9PYKY5cDv8x20dpShuOdbJVQ3DpR345PgVCdYAuuqlG0ZOStg36RbQRjM40xvwKBgEnu3i1ueSQxRFViyhRMRQrCxFBfMuXK5m6bHIOyNPK1Jcmv55hTDHSjwc16NmIkDjIW/mO3hxAV1FhxALY/KC0IQfIV8MQadzL9sVrFX6SIULJAASmql/82J2iwJH8G6aAanVQFjP/XB86yxCDV1F+zp25yv9zOPor+kXmitLFlAoGAHUTjNw5GaPgP9feBEzuOTvxkoCD+TUiN5Tg6hIaU3u+U2eHnj4UGdixNmAq+VOWj7+53IXFNAfgUgNMHbtmALtbLycz1DOei3LXFX6YaIHnKEpNGpJaolgTzN9kXz7PaGuGZlD6cH3PmpLk+YJBBvbsCPeLAZuymVk0EgYI9Wy8CgYEAnYaJXkT0vDEffaXP1BxBkVZV1jVHnjZwwd6VRTh7Gipu0bQPRg7b5HX3OzLEtj8nGsHlTixIAezP9C+3NL8ylUXGg7zilrgEojvgCoBXP0boVo3mTpUx9DTmAn7bdlYNB4xHaXpZEx7qpr1bve5E8obtYA8H0cLPzbjF3sWAgMU="
	//kServerPort = "8080"
	// TODO 设置回调地址域名
	kServerDomain = "http://uznr4z.natappfree.cc"
)

func init() {
	var err error

	if client, err = alipay.New(kAppId, kPrivateKey, false); err != nil {
		log.Println("初始化支付宝失败", err)
		return
	}

	// 应用公钥证书,该证书用于验证支付宝发送的异步通知和回调请求中的签名。
	if err = client.LoadAppCertPublicKeyFromFile("./appV2/aliPay/appPublicCert.crt"); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}
	//从指定的文件中加载支付宝根证书 (alipayRootCert.crt) ，该证书用于验证支付宝发出的接口请求的签名。
	if err = client.LoadAliPayRootCertFromFile("./appV2/aliPay/alipayRootCert.crt"); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}
	//从指定的文件中加载支付宝公钥证书 (alipayPublicCert.crt) ，该证书用于对敏感信息进行加密。
	if err = client.LoadAlipayCertPublicKeyFromFile("./appV2/aliPay/alipayPublicCert.crt"); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}
	//使用指定的内容加密密钥 (“iotxR/d99T9Awom/UaSqiQ==”) 进行内容加密或解密操作。
	if err = client.SetEncryptKey("kTobkn3+Zyf2rt0OxsvqPg=="); err != nil {
		log.Println("加载内容加密密钥发生错误", err)
		return
	}
	if err != nil {
		fmt.Println("jia密出错")
	}

	//http.HandleFunc("/alipay/pay", pay)
	//http.HandleFunc("/alipay/callback", callback)
	//http.HandleFunc("/alipay/notify", notify)
	//
	//http.ListenAndServe(":"+kServerPort, nil)//开启一个服务器
}

func AliPay(writer http.ResponseWriter, request *http.Request, ordersId string) bool {
	fmt.Println("pay进入成功")
	var tradeNo = ordersId
	if !model.GetPay(tradeNo) {
		fmt.Println("查询订单失败")
		return false
	}
	//tradeNo = fmt.Sprintf("%d", xid.Next())

	var p = alipay.TradePagePay{}
	p.NotifyURL = kServerDomain + "/alipay/notify"
	p.ReturnURL = kServerDomain + "/alipay/callback"
	p.Subject = "支付测试:" + tradeNo
	p.OutTradeNo = tradeNo
	p.TotalAmount = "10.00" //价格
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, _ := client.TradePagePay(p)
	fmt.Println("准备重定向")
	http.Redirect(writer, request, url.String(), http.StatusTemporaryRedirect)
	return true
}

func Callback(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	if err := client.VerifySign(request.Form); err != nil {
		log.Println("回调验证签名发生错误", err)
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("回调验证签名发生错误"))
		return
	}

	log.Println("回调验证签名通过")

	var outTradeNo = request.Form.Get("out_trade_no")
	var p = alipay.TradeQuery{}
	p.OutTradeNo = outTradeNo

	rsp, err := client.TradeQuery(p)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(fmt.Sprintf("验证订单 %s 信息发生错误: %s", outTradeNo, err.Error())))
		return
	}

	if rsp.IsFailure() {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(fmt.Sprintf("验证订单 %s 信息发生错误: %s-%s", outTradeNo, rsp.Msg, rsp.SubMsg)))
		return
	}
	writer.WriteHeader(http.StatusOK)
	if model.Pay(outTradeNo) {
		writer.Write([]byte(fmt.Sprintf("订单 %s 支付成功", outTradeNo)))
		//outTradeNo = request.Form.Get("out_trade_no")
	} else {
		writer.Write([]byte(fmt.Sprintf("订单 %s 支付失败", outTradeNo)))

	}
}

func Notify(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	var notification, err = client.DecodeNotification(request.Form)
	if err != nil {
		log.Println("解析异步通知发生错误", err)
		return
	}

	log.Println("解析异步通知成功:", notification.NotifyId)

	var p = alipay.NewPayload("alipay.trade.query")
	p.AddBizField("out_trade_no", notification.OutTradeNo)

	var rsp *alipay.TradeQueryRsp
	if err = client.Request(p, &rsp); err != nil {
		log.Printf("异步通知验证订单 %s 信息发生错误: %s \n", notification.OutTradeNo, err.Error())
		return
	}
	if rsp.IsFailure() {
		log.Printf("异步通知验证订单 %s 信息发生错误: %s-%s \n", notification.OutTradeNo, rsp.Msg, rsp.SubMsg)
		return
	}

	log.Printf("订单 %s 支付成功 \n", notification.OutTradeNo)

	client.ACKNotification(writer)
}
