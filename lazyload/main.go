package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"gitlab.com/tesgo/kit/acsdk"
)

func main() {
	lazyLoadingSDK("127.0.0.1:7001")

	select {}
}

var (
	sdk      *acsdk.ACtlSDK
	sdkMetux sync.Mutex
	//sdkOnce  sync.Once
)

const (
	waitTimes               = 3
	firstWaitSecond         = 10
	cycleWaitTimes          = 5
	firstCycleWaitSecond    = 10
	cycleWaitIntervalSecond = 5
)

func lazyLoadingSDK(grpcAddr string) {
	go func() {
		// waiting for ac grpc startup
		start := time.Now()
		fmt.Println("start==>", start)
		time.Sleep(firstWaitSecond * time.Second)
		intervalSecond := firstCycleWaitSecond
		var csdk *acsdk.ACtlSDK

	LOOP:
		for i := 0; i < waitTimes; i++ {
			for j := 0; j < cycleWaitTimes; j++ {
				csdk = loadSDK(grpcAddr)
				if csdk != nil {
					break LOOP
				}

				time.Sleep(time.Duration(intervalSecond) * time.Second)
			}

			intervalSecond += cycleWaitIntervalSecond
		}

		// ok
		end := time.Now()
		fmt.Println("end==>", start, end, end.Sub(start), csdk)
	}()
}

func loadSDK(grpcAddr string) *acsdk.ACtlSDK {
	sdkMetux.Lock()
	defer sdkMetux.Unlock()

	if sdk != nil {
		return sdk
	}

	nsdk, err := acsdk.New(grpcAddr)
	if err != nil {
		return nil
	}

	if nsdk == nil {
		return nil
	}

	err = nsdk.Ping(context.Background())
	if err != nil {
		return nil
	}
	sdk = nsdk

	return sdk
}

func closeSDK() {
	sdkMetux.Lock()
	defer sdkMetux.Unlock()

	if sdk != nil {
		sdk.Close()
		sdk = nil
	}
}
