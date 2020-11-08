package controllers

import (
	"testing"
)

func TestCGetUrls(t *testing.T) {

	url1 := "https://httpstat.us/503"
	url2 := "https://httpstat.us/200"
	expectedPrometheusCode1 := 0
	expectedPrometheusCode2 := 1

	prometheusCode := DoRequestAndReceiveResponse(url1)
	if prometheusCode != expectedPrometheusCode1 {
		t.Errorf("Test failed Expected prometheuscode: %d Actual Result: %d", expectedPrometheusCode1, prometheusCode)
	}
	prometheusCode = DoRequestAndReceiveResponse(url2)
	if prometheusCode != expectedPrometheusCode2 {
		t.Errorf("Test failed Expected prometheuscode: %d Actual Result: %d", expectedPrometheusCode1, prometheusCode)
	}

}
