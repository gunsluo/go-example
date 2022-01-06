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
