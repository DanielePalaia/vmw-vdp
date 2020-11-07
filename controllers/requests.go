package controllers

import (
	"log"
	"net/http"
	"time"
	environments "vmw-vdp/envinronments"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Receive and manage the request (city and numbers of info)
func HandleMetricsRequest(w http.ResponseWriter, r *http.Request) {

	DoRequestsAndReceiveResponse(w)

}

// Receive and manage the request (city and numbers of info)
func HandlerWrapper(w http.ResponseWriter, r *http.Request) {

	log.Printf("metrics endpoing")

	promhttp.Handler()

}

func DoRequestsAndReceiveResponse(w http.ResponseWriter) {

	/* Do the request for each url*/
	urls := environments.GetUrls()

	for _, url := range urls {
		start := time.Now()
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatalf("Error calling http new request on url: %s error message: %v", url, err)

		}
		client := &http.Client{}
		/* Receive back the respose */
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Error calling client.Do on request url: %s error message: %v", url, err)
		}
		elapsed := time.Since(start).Milliseconds()
		status := resp.StatusCode

		log.Printf("called url: %v status: %v responseTime: %v", url, status, elapsed)

	}

}
