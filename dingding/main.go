package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	dingtalk "github.com/gunsluo/godingtalk"
)

var (
	client *dingtalk.DingTalkClient
)

var (
	corpId         = "corpid"
	suiteId        = "20001"
	suiteKey       = "suite"
	suiteSecret    = "quRSU"
	appId          = "91"
	token          = "fortesting"
	encodingAESKey = "p6y"
)

func main() {
	port := 12345

	engine := gin.New()

	engine.POST("/dd/register/callback", callback)
	engine.GET("/", index)

	config := dingtalk.ISVConfig{
		CorpId:      corpId,
		AppId:       appId,
		SuiteKey:    suiteKey,
		SuiteSecret: suiteSecret,
		AESKey:      encodingAESKey,
		Token:       token,
	}
	client = dingtalk.NewISVClient(config)
	if _, err := client.GetAndRefreshSuiteAccessToken(); err != nil {
		panic(err)
	}

	if err := engine.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatal("HTTP RESTFul Server exceptions")
	}
}

type RegisterCallbackReq struct {
	Signature    string `json:"signature" form:"signature"`
	MsgSignature string `json:"msg_signature" form:"msg_signature"`
	Timestamp    string `json:"timestamp" form:"timestamp"`
	Nonce        string `json:"nonce" form:"nonce"`
	Encrypt      string `json:"encrypt" form:"encrypt"`
}

func index(ctx *gin.Context) {
	fmt.Println("index.html->", ctx.Query("corpId"), ctx.Request.URL)
	ctx.String(http.StatusOK, ctx.Query("corpId"))
}

func callback(ctx *gin.Context) {
	var p RegisterCallbackReq
	if err := ctx.ShouldBindQuery(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	if err := ctx.ShouldBindJSON(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	fmt.Println("callback-->", p)

	notification, err := client.DecryptAndUnmarshalPushNotification(p.Signature, p.Timestamp, p.Nonce, p.Encrypt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	fmt.Println("--->", notification)

	var random string
	if fn, ok := eventTypesHandleFuncs[notification.EventType]; ok {
		v, err := fn(notification)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		random = v
	} else {
		fmt.Println("not support event type " + notification.EventType)
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	encryRandom, sigRandom, err := client.Encrypt(random, p.Timestamp, p.Nonce)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg_signature": sigRandom,
		"timeStamp":     p.Timestamp,
		"nonce":         p.Nonce,
		"encrypt":       encryRandom,
	})
}

var eventTypesHandleFuncs = map[string]func(*dingtalk.PushNotification) (string, error){
	dingtalk.CheckCreateSuiteURLEventType: checkCreateSuiteURL,
	dingtalk.CheckUpdateSuiteUrlEventType: checkCreateSuiteURL,
	dingtalk.CheckUrlEventType:            checkURL,
	dingtalk.SyncHTTPPushHighEventType:    syncHTTPPushHigh,
	dingtalk.SyncHTTPPushMediumEventType:  syncHTTPPushMedium,
}

func checkCreateSuiteURL(n *dingtalk.PushNotification) (string, error) {
	if suiteKey != n.TestSuiteKey {
		return "", errors.New("invalid suite key")
	}
	return n.Random, nil
}

func checkURL(n *dingtalk.PushNotification) (string, error) {
	return "success", nil
}

func syncHTTPPushHigh(n *dingtalk.PushNotification) (string, error) {
	var authCorpids []string
	var permanentCodes []string

	for _, item := range n.BizItems {
		switch v := item.BizData.(type) {
		case dingtalk.Biz2Data:
			client.SetSuiteTicket(v.SuiteTicket)
		case dingtalk.Biz4Data:
			authCorpids = append(authCorpids, v.AuthCorpInfo.CorpID)
			permanentCodes = append(permanentCodes, v.PermanentCode)
		case dingtalk.Biz16Data:
			// TODO: remove org
		default:
			fmt.Printf("-->%d, %T\n", item.BizType, item.BizData)
		}
	}

	// active
	for i, corpId := range authCorpids {
		resp, err := client.IsvActivateSuite(corpId, permanentCodes[i])
		if err != nil {
			return "", err
		}
		fmt.Println("--->", resp)
	}

	return "success", nil
}

func syncHTTPPushMedium(n *dingtalk.PushNotification) (string, error) {

	return "success", nil
}
