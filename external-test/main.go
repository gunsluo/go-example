package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/husobee/vestigo"
	"github.com/tylerb/graceful"
)

var (
	// srv is tylerb's graceful server, which allows us to turn the
	// server off at will within the code.  This is super handy for
	// us because we want to be able to end our test, in order for
	// go's testing framework to report the coverage (no good if
	// the service is interrupted or canceled or terms)
	srv = &graceful.Server{
		Timeout: 5 * time.Second,
	}
	// stop is a channel that tells the service to stop.  As seen
	// later we will make a highly destructive deathblow endpoint
	// so that in test we can conclude the test and turn the service off.
	stop chan bool
	// testMode is a bool that allows for deathpunch endpoint to exist or
	// not exist... we don't want that running in production ;)
	testMode bool = false
)

// runMain - entry-point to perform external testing of service, this is
// where go test will enter main.  we have to setup test mode in here, as
// well as the stop channel so we can stop the service
func runMain() {
	// start the stop channel
	stop = make(chan bool)
	// put the service in "testMode"
	testMode = true
	// run the main entry point
	go main()
	// watch for the stop channel
	<-stop
	// stop the graceful server
	srv.Stop(5 * time.Second)
}

// main - main entry point
func main() {
	// setup middlware stack
	n := negroni.Classic()

	// setup routes
	router := vestigo.NewRouter()

	// endpoints
	router.Post("/test", func(w http.ResponseWriter, r *http.Request) {
		if false {
			// we should see this endpoint not covered
			// if we hit the /test endpoint externally
			fmt.Println("totally never getting here")
		}
		w.WriteHeader(200)
		w.Write([]byte("done"))
	})
	// all of the above is basic boring service setup stuff.

	// only if we are in testMode should we attempt to add the death blow
	if testMode {
		// death blow endpoint - endpoint that will stop the service if stop is
		// a live channel (only live if started from RunMain)
		router.Post("/deathblow", func(w http.ResponseWriter, r *http.Request) {
			// end the graceful server if being run from RunMain()
			stop <- true
		})
	}

	// add our router to negroni
	n.UseHandler(router)

	// graceful start/stop server
	srv.Server = &http.Server{Addr: ":1234", Handler: n}
	// serve http
	srv.ListenAndServe()
}
