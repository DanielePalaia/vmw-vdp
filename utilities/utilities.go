package utilities

import (
	"log"
	"os"
	"strconv"
)

func GetHostAndPort() (string, string) {

	host := os.Getenv("host")
	port := os.Getenv("port")

	return host, port

}

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
