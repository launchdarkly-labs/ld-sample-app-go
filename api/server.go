package api

import (
	"LdSampleAppGo/ldclient"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/launchdarkly/go-sdk-common/v3/ldcontext"
	"github.com/launchdarkly/go-sdk-common/v3/ldvalue"
	ld "github.com/launchdarkly/go-server-sdk/v7"
)

type API struct {
	Port     string
	LdSdkKey string
}

func CurrentContext() ldcontext.Context {
	context := ldcontext.NewBuilder("018ee35c-5014-75d4-9efd-a3dcd1f31341").
		Kind("device").
		Name("Linux").
		Build()

	return context
}

func MonitorFlagChange(client *ld.LDClient, flagKey string, context ldcontext.Context) {
	updateChannel := client.GetFlagTracker().AddFlagValueChangeListener(flagKey, context, ldvalue.Null())
	go func() {
		for event := range updateChannel {
			log.Printf("Flag value changed to '%s'", event.NewValue)
		}
	}()
}

func (api API) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
	client := ldclient.GetLdClient()
	context := CurrentContext()

	flagValue, _ := client.BoolVariation("test-flag", context, false)
	fmt.Fprint(w, "The flag value is: '"+strconv.FormatBool(flagValue)+"'")
}

func (api API) Run() {
	ipaddr := "0.0.0.0"
	apiPort := fmt.Sprintf("%s:%s", ipaddr, api.Port)
	log.Println("Started API, listening on " + apiPort)
	router := httprouter.New()

	client := ldclient.GetLdClient()
	MonitorFlagChange(client, "test-flag", CurrentContext())

	router.GET("/", api.Index)
	log.Fatal(http.ListenAndServe(apiPort, router))
}

func New(port string) *API {
	api := &API{
		Port:     port,
		LdSdkKey: os.Getenv("LD_SDK_KEY"),
	}

	return api
}
