// Package utilities provide general utilities functions which can be reused in other projects
package utilities

import (
	"log"
	"os"
	"strconv"
)

// GetHostAndPort returns the host and the port defined in the OS variables host and port
func GetHostAndPort() (string, string) {

	host := os.Getenv("host")
	port := os.Getenv("port")

	return host, port

}

// GetUrls returns the number of Urls which this project should target and url1, url2....
func GetUrls() []string {

	numOfUrls := os.Getenv("numUrls")
	numOfUrlsInt := 0
	var err error

	if numOfUrlsInt, err = strconv.Atoi(numOfUrls); err != nil {
		log.Fatalf("error converting numUrls env variable to int probably bad values was set for it error: %v", err)
	}

	urls := make([]string, numOfUrlsInt)

	for i := 0; i < numOfUrlsInt; i++ {
		url := os.Getenv("url" + strconv.Itoa(i+1))
		urls[i] = url
	}

	return urls

}

// PrintLogs prints general logs prints for main
func PrintLogs(host string, port string, urls []string) {
	log.Printf("Starting service at host: %v and port: %v targetting urls: %v", host, port, urls)
}
