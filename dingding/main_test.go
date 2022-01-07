package main

import (
	"context"
	"fmt"
	"testing"

	dingtalk "github.com/gunsluo/godingtalk"
)

func TestA(t *testing.T) {
	ctx := context.Background()
	config := dingtalk.ISVConfig{
		CorpId:      corpId,
		AppId:       appId,
		SuiteKey:    suiteKey,
		SuiteSecret: suiteSecret,
		AESKey:      encodingAESKey,
		Token:       token,
	}
	client = dingtalk.NewISVClient(config)
	if _, err := client.GetAndRefreshSuiteAccessToken(ctx); err != nil {
		panic(err)
	}

	p := RegisterCallbackReq{
		Signature:    "4195f3a915c38cdb68548cd722c3ea061b1ef0d5",
		MsgSignature: "4195f3a915c38cdb68548cd722c3ea061b1ef0d5",
		Timestamp:    "1639730923428",
		Nonce:        "lcXMQEAO",
		Encrypt:      `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`,
	}

	notification, err := client.DecryptAndUnmarshalPushNotification(p.Signature, p.Timestamp, p.Nonce, p.Encrypt)
	if err != nil {
		panic(err)
	}

	fmt.Println("-->", notification)

	var random string
	if fn, ok := eventTypesHandleFuncs[notification.EventType]; ok {
		v, err := fn(ctx, notification)
		if err != nil {
			panic(err)
		}
		random = v
	} else {
		fmt.Println("not support event type " + notification.EventType)
		panic(err)
	}

	fmt.Println("-->", random)

}

/*
func TestB(t *testing.T) {

	p := RegisterCallbackReq{
		Signature:    "a5144443588e0dc51eb40480af2a056c62a85442",
		MsgSignature: "a5144443588e0dc51eb40480af2a056c62a85442",
		Timestamp:    "1639489786105",
		Nonce:        "de8eBj29",
		Encrypt:      `DOft5JRfH1FK8KilsOGRhQtlPc0+xiG9BoILD1qgLoB1POMy5SvaeSVZe1adRdosdU9QhiOEa02SnZjDS4SYxJpSK3ZAXqvQ5S3Hb2tnwjOikgJEY0Ri1s/MtWFHGn5rZaLy5wvPzK3h2Lx1lrKnFYm7zhN/TthMjC0+9z4I01O+2ju4BqDu/ugS6rgIQ4ltirLnIIlQo+WYgOw4cJBC6YDr2BA08pwcLU3QMoSamhfuuMO0vx9NfAzfB2jvrUficx3INtQ3+NLmc2tiSSOO79W2HECjhNaVUmCEgoSqNDQQYCp5ChlNKzLUgvlLTlwiqZvPkyQBNbt25qGA7+9FfX41CLFJLsE9FiKvl+KzzLXiODQkPQ1KYXEQnblkLsTcrtyPLoFqVv/H7zj37JC0bvbHHKIoXEu/9aDE+pDL6VQeTmXEKT1ZLAjtEWPw4RKJduukVROrMV+63ezAPKVXGO55lJDTMFDMLuIHfpYlXnVdeQcEnsjESApdgp0XfIZvbtIXPiwfAbhlaPvaD4pUBgT3TjnOqheSzG8j4qECmPWgnccZGCbpcX/dedQw8iLYRqzeaMrrzPXU9c3odCWqZlJ+gEL2zzw1pgMi075HD9wjbRMrvLzH1zPdG6kE2b9pqcjzWu+O7hEU8b/n2pNRQfTmggEnc/9YTUmjBT2w9/QiE+BJaYVxyUnLVbGR87Z6aJXC2kAv/WulAfTdbJlURFb3RfB5oyi+NGG8OxwkjXGFXc8Q7bOmQsgy1Wa0Re4yhXxmHmU3hJyAun0rsNUfYnr4Usd5/E2v7hpCE5/l9iJqnG/GAjARUot9N5e8TeCgi0JgrsvcG10txwnmRtYdtYqujP/Ry5BhbcKDs8qhfZW8yt50MoFp2EZqlO4K6134lCIp6+0Jv71on7vOC49LiSgac1zC6l1IfKatFT86SPG4lZVvrBWjcLCrl9nhkRKl4sifJjNjVO3YJXb2JW0Oq++mnjB/5egMKrOTF04C9BlDECWP1mvRTsIgEflZDG43`,
	}

	dtCrypto := dingtalk.NewDingTalkCrypto(token, encodingAESKey, suiteKey)
	plainMsg, err := dtCrypto.GetDecryptMsg(p.Signature, p.Timestamp, p.Nonce, p.Encrypt)
	if err != nil {
		panic(err)
	}

	var pd PushData
	if err := json.Unmarshal([]byte(plainMsg), &pd); err != nil {
		panic(err)
	}

	for _, item := range pd.BizData {
		switch item.BizType {
		case 2:
		case 4:
		case 16:
			var d Biz16Data
			if err := json.Unmarshal([]byte(item.BizData), &d); err != nil {
				panic(err)
			}
		}
	}
}

*/

func TestB(t *testing.T) {
	ctx := context.Background()
	config := dingtalk.ISVConfig{
		CorpId:      corpId,
		AppId:       appId,
		SuiteKey:    suiteKey,
		SuiteSecret: suiteSecret,
		AESKey:      encodingAESKey,
		Token:       token,
	}
	client = dingtalk.NewISVClient(config)
	if _, err := client.GetAndRefreshSuiteAccessToken(ctx); err != nil {
		panic(err)
	}

	resp, err := client.IsvListUnactivateSuites(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("-->", resp.CorpList)
}

func TestC(t *testing.T) {
	ctx := context.Background()
	config := dingtalk.ISVConfig{
		CorpId:      corpId,
		AppId:       appId,
		SuiteKey:    suiteKey,
		SuiteSecret: suiteSecret,
		AESKey:      encodingAESKey,
		Token:       token,
	}
	client = dingtalk.NewISVClient(config)
	if _, err := client.GetAndRefreshSuiteAccessToken(ctx); err != nil {
		panic(err)
	}

	code := "06b83551866b3258b651f5c8be3df29f"
	resp, err := client.SNSGetUserInfoByCode(ctx, code)
	if err != nil {
		panic(err)
	}
	fmt.Println("-->", resp.UserInfo)
	unionId := resp.UserInfo.UnionId
	//unionId := "ImEJLvh65cfliPMCWiS7iP43AiEiE"

	// other corp id, using our corp to test
	authCorpId := corpId
	userResp, err := client.GetCorpUserByUnionId(ctx, authCorpId, unionId)
	fmt.Println("-->", userResp, err)
	if err != nil {
		panic(err)
	}
	fmt.Println("userId-->", userResp.UserInfo)
	userId := userResp.UserInfo.UserId

	userDetailResp, err := client.GetCorpUserDetailByUserId(ctx, authCorpId, userId)
	if err != nil {
		panic(err)
	}

	fmt.Println("userInfo-->", userDetailResp)
}

func TestD(t *testing.T) {
	ctx := context.Background()
	config := dingtalk.ISVConfig{
		CorpId:      corpId,
		AppId:       appId,
		SuiteKey:    suiteKey,
		SuiteSecret: suiteSecret,
		AESKey:      encodingAESKey,
		Token:       token,
	}
	client = dingtalk.NewISVClient(config)
	if _, err := client.GetAndRefreshSuiteAccessToken(ctx); err != nil {
		panic(err)
	}

	code := "c532ea9a72be389ca2c9d1273e745a2b"
	authCorpId := corpId
	resp, err := client.GetCorpUserInfoByCode(ctx, authCorpId, code)
	if err != nil {
		panic(err)
	}
	fmt.Println("-->", resp.UserInfo)

	userId := resp.UserInfo.UserId
	userDetailResp, err := client.GetCorpUserDetailByUserId(ctx, authCorpId, userId)
	if err != nil {
		panic(err)
	}

	fmt.Println("userInfo-->", userDetailResp)
}

func TestE(t *testing.T) {
	ctx := context.Background()
	config := dingtalk.ISVConfig{
		CorpId:      corpId,
		AppId:       appId,
		SuiteKey:    suiteKey,
		SuiteSecret: suiteSecret,
		AESKey:      encodingAESKey,
		Token:       token,
	}

	client = dingtalk.NewISVClient(config)
	if _, err := client.GetAndRefreshSuiteAccessToken(ctx); err != nil {
		panic(err)
	}

	authCorpId := "ding5fd1afa8a7c9bf97acaaa37764f94726"
	agentId := 1408568213
	//
	resp, err := client.SendMessageByTemplate(ctx, authCorpId,
		dingtalk.SendMessageByTemplateRequest{
			AgentId:    agentId,
			TemplateId: "be308c698ff140728ae538ee7c887eca",
			//TemplateId: "b651be75b2404f4f86e448304344fd2b",
			UserIds: []string{`manager7002`, `16393884365594803`, `094428093926199868`},
			DeptIds: []string{},
			Data:    `{"name":"` + "luoji" + `"}`,
		})
	if err != nil {
		fmt.Println("userInfo-->", resp)
		panic(err)
	}

	fmt.Println("userInfo-->", resp)
}

func TestF(t *testing.T) {
	ctx := context.Background()
	config := dingtalk.ISVConfig{
		CorpId:      corpId,
		AppId:       appId,
		SuiteKey:    suiteKey,
		SuiteSecret: suiteSecret,
		AESKey:      encodingAESKey,
		Token:       token,
	}
	client = dingtalk.NewISVClient(config)
	if _, err := client.GetAndRefreshSuiteAccessToken(ctx); err != nil {
		panic(err)
	}

	p := RegisterCallbackReq{
		Signature:    "57717955a044ae7c81f4fd93c79d65d50370efab",
		MsgSignature: "57717955a044ae7c81f4fd93c79d65d50370efab",
		Timestamp:    "1641457369837",
		Nonce:        "KfyRauUG",
		Encrypt:      `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`,
	}

	notification, err := client.DecryptAndUnmarshalPushNotification(p.Signature, p.Timestamp, p.Nonce, p.Encrypt)
	if err != nil {
		panic(err)
	}

	fmt.Println("-->", notification)

	var random string
	if fn, ok := eventTypesHandleFuncs[notification.EventType]; ok {
		v, err := fn(ctx, notification)
		if err != nil {
			panic(err)
		}
		random = v
	} else {
		fmt.Println("not support event type " + notification.EventType)
		panic(err)
	}

	fmt.Println("-->", random)

}
