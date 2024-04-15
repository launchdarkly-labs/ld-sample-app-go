package main

import "LdSampleAppGo/api"

func main() {
	a := api.New("3000")

	a.Run()
}
