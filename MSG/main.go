package main

/*
@Time    : 2021/11/5 11:58
@Author  : austsxk
@Email   : austsxk@163.com
@File    : main.go
@Software: GoLand
*/

import (
	"github.com/cloopen/go-sms-sdk/cloopen"
	"log"
)

func main() {
	cfg := cloopen.DefaultConfig().
		// 开发者主账号,登陆云通讯网站后,可在控制台首页看到开发者主账号ACCOUNT SID和主账号令牌AUTH TOKEN
		WithAPIAccount("8a216dab2bc78f016b41a624210c35").
		// 主账号令牌 TOKEN,登陆云通讯网站后,可在控制台首页看到开发者主账号ACCOUNT SID和主账号令牌AUTH TOKEN
		WithAPIToken("51a75c98512644419d5a23da4d184")
	sms := cloopen.NewJsonClient(cfg).SMS()
	// 下发包体参数
	input := &cloopen.SendRequest{
		// 应用的APPID
		AppId: "8a216da86b2bc78f016b41a624810erc",
		// 手机号码
		To: "13155470000",
		// 模版ID
		TemplateId: "1",
		// 模版变量内容 非必填
		Datas: []string{"123", "5"},
	}
	// 下发
	resp, err := sms.Send(input)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Response MsgId: %s \n", resp.TemplateSMS.SmsMessageSid)

}
