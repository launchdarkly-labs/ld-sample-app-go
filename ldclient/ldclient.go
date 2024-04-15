package ldclient

import (
	"fmt"
	"os"
	"sync"
	"time"

	ld "github.com/launchdarkly/go-server-sdk/v7"
)

var once sync.Once

var instance *ld.LDClient

func GetLdClient() *ld.LDClient {
	once.Do(func() {
		sdkkey := os.Getenv("LD_SDK_KEY")
		client, _ := ld.MakeClient(sdkkey, 5*time.Second)
		if client.Initialized() {
			fmt.Printf("SDK successfully initilized!\n")
		} else {
			fmt.Printf("SDK failed to initialize.\n")
			os.Exit(1)
		}
		instance = client
	})

	return instance
}
