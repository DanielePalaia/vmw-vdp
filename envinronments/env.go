// Package environments contains the project environment
package environments

type urlToRequest struct {
	urls []string
}

var Urls urlToRequest

func SetUrls(urls []string) {

	Urls.urls = urls
}

func GetUrls() []string {

	return Urls.urls
}
