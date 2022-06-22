package main

import (
	"fmt"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

const (
	accessKeyId     = "LTAI5tHkCvrvTMBGHQbWUHjm"
	accessKeySecret = "3V6orVzIOMdLF5Ptq6BTHsImBNHVuM"
)

func main() {
	client, err := CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if err != nil {
		panic(err)
	}

	needAddedTemps := []string{}
	newAddedTemps := []string{"SMS_243566548"}
	for _, code := range newAddedTemps {
		querySmsTemplateRequest := &dysmsapi20170525.QuerySmsTemplateRequest{
			TemplateCode: tea.String(code),
		}

		resp, err := client.QuerySmsTemplateWithOptions(querySmsTemplateRequest, &util.RuntimeOptions{})
		if err != nil {
			panic(err)
		}

		if resp.Body.Code == nil {
			panic(fmt.Errorf("invalid code"))
		}

		if *resp.Body.Code != "OK" {
			if *resp.Body.Code == "isv.SMS_TEMPLATE_ILLEGAL" {
				needAddedTemps = append(needAddedTemps, code)
				continue
			}
			panic(fmt.Errorf("invalid code"))
		}

		if resp.Body.TemplateStatus == nil {
			continue
		}

		if *resp.Body.TemplateStatus != 0 {
			fmt.Printf("code: %s audit failure\n", code)
		}

	}

	fmt.Printf("--%#v\n", needAddedTemps)

	for _, code := range needAddedTemps {
		_ = code
		addSmsTemplateRequest := &dysmsapi20170525.AddSmsTemplateRequest{
			TemplateType:    tea.Int32(0),
			TemplateName:    tea.String("test1"),
			TemplateContent: tea.String("this s a test, code: ${code}."),
			Remark:          tea.String("testing"),
		}

		_, err = client.AddSmsTemplateWithOptions(addSmsTemplateRequest, &util.RuntimeOptions{})
		if err != nil {
			panic(err)
		}
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:     tea.String("阿里云短信测试"),
		TemplateCode: tea.String("SMS_154950909"),
		//TemplateCode:  tea.String(newAddedTemps[0]),
		PhoneNumbers:  tea.String("18980501737"),
		TemplateParam: tea.String("{\"code\":\"1234\"}"),
	}

	// 复制代码运行请自行打印 API 的返回值
	resp, err := client.SendSmsWithOptions(sendSmsRequest, &util.RuntimeOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("--%#v\n", *resp)
}

func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, err error) {
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, err = dysmsapi20170525.NewClient(config)
	return _result, err
}
