package domainUtils

import "net/http"

// inspects the http header
// returns the status code (200)
// status codes for reference
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status
// we will treat 200-299 as success in general
func CheckHttpHead(url string) (int, error) {
	r, err := http.Head(url)
	if err != nil {

		return 0, err
	}
	// log.Printf("URL Check (%s) (%d)", url, r.StatusCode)
	return r.StatusCode, nil
}
