package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

type API struct {
	Port     string
	LdSdkKey string
}

func (api API) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func (api API) Run() {
	ipaddr := "0.0.0.0"
	apiPort := fmt.Sprintf("%s:%s", ipaddr, api.Port)
	log.Println("Started API, listening on " + apiPort)
	router := httprouter.New()

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
