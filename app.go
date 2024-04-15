package main

import "github.com/launchdarkly-labs/ld-sample-app-go/api"

func main() {
	a := api.New("3000")

	a.Run()
}
