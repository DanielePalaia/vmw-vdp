package utilities

import (
	"os"
	"strconv"
	"testing"
)

// Test which mock an http request. Create with POST an element and check with GET
func TestCGetHostAndPort(t *testing.T) {

	expectedHost := "localhost"
	expectedPort := "8080"

	os.Setenv("host", expectedHost)
	os.Setenv("port", expectedPort)

	actualHost, actualPort := GetHostAndPort()

	if expectedHost != actualHost {
		t.Errorf("Test failed expected result: %s actual result: %s", expectedHost, actualHost)
	}
	if expectedPort != actualPort {
		t.Errorf("Test failed expected result: %s actual result: %s", expectedPort, actualPort)
	}

}

// Test which mock an http request. Create with POST an element and check with GET
func TestCGetUrls(t *testing.T) {

	inputNumUrls := 2
	expectedUrls := make([]string, inputNumUrls)
	expectedUrls[0] = "https://httpstat.us/503"
	expectedUrls[1] = "https://httpstat.us/200"

	os.Setenv("numUrls", strconv.Itoa(inputNumUrls))
	os.Setenv("url1", expectedUrls[0])
	os.Setenv("url2", expectedUrls[1])

	urls := GetUrls()

	for i, url := range urls {
		if url != expectedUrls[i] {
			t.Errorf("Test failed expected url value:%s actual url value: %s", expectedUrls[i], url)
		}
	}

}
