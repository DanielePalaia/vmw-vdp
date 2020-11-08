// Package environments contains the project environment
package environments

type urlToRequest struct {
	urls []string
}

// Urls This is a global variable to keep configuration
var Urls urlToRequest

// SetUrls is a setter for Urls
func SetUrls(urls []string) {

	Urls.urls = urls
}

// GetUrls is a getter
func GetUrls() []string {

	return Urls.urls
}
