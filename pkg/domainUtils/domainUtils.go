package domainUtils

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

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

// get all the certificates associated with a domain, on the given port
func GetCertificates(domain string, sslPort int) ([]*x509.Certificate, error) {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", domain, sslPort), conf)
	if err != nil {
		logrus.Print("Error in Dial", err)
		return nil, err
	}
	defer conn.Close()
	certs := conn.ConnectionState().PeerCertificates

	return certs, nil

}
