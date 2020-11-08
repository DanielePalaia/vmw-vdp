// Package controllers contains the main get hanlde function of the project to handle /service
package controllers

import (
	"log"
	"net/http"
	"time"
	environments "vmw-vdp/envinronments"
	"vmw-vdp/prometheus"
)

// HandleServiceRequest Receives and manage the request (city and numbers of info)
func HandleServiceRequest(w http.ResponseWriter, r *http.Request) {

	log.Print("HandleServiceRequest: Receiving a /service request")

	/* Do the request for each url*/
	urls := environments.GetUrls()

	for _, url := range urls {

		log.Printf("new request to: %v", url)
		DoRequestAndReceiveResponse(url)

	}

	// Reply back to the client request
	w.WriteHeader(http.StatusOK)

}

// DoRequestAndReceiveResponse receive a manage a single url
func DoRequestAndReceiveResponse(url string) int {

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

	prometheusCode := 0
	if status == 200 {
		prometheusCode = 1
		prometheus.LinkUp.WithLabelValues(url).Set(1)
	} else if status == 503 {
		prometheusCode = 0
		prometheus.LinkUp.WithLabelValues(url).Set(0)
	}
	prometheus.ResponseMs.WithLabelValues(url).Set(float64(elapsed))

	log.Printf("called url: %v status: %v responseTime: %v", url, status, elapsed)

	return prometheusCode

}
