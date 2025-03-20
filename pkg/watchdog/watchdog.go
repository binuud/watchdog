package watchDogServer

import (
	"time"

	"github.com/binuud/watchdog/gen/go/v1/watchdog"
	"github.com/binuud/watchdog/pkg/domainUtils"
	"github.com/sirupsen/logrus"
)

func (s *WatchDogServer) checkDomains() error {
	for _, item := range s.DomainWatch.Domains {
		s.checkDomain(item)
	}
	return nil
}

func (s *WatchDogServer) checkDomain(domain *watchdog.DomainItem) {

	for _, item := range domain.Endpoints {
		responseCode, err := domainUtils.CheckHttpHead(item)
		if err != nil {
			// handle err
		}
		logrus.Printf("Received response %d %s", responseCode, item)
	}
	// fetch and analyse the certificates
	err := s.analyseCertificates(domain)
	if err != nil {
		logrus.Errorf("error when checking certificated of domain %s, %v", domain.Name, err)
	}

	ipAddresses, err := domainUtils.ResolveIP(domain.Name)
	if err != nil {
		logrus.Errorf("error when resolving ip of domain %s, %v", domain.Name, err)
	}
	logrus.Printf("Ip address of domain %s %v", domain.Name, ipAddresses)

}

func (s *WatchDogServer) analyseCertificates(domain *watchdog.DomainItem) error {

	certs, err := domainUtils.GetCertificates(domain.Name, 443)
	if err != nil {
		return err
	}

	if len(certs) > 0 {
		now := time.Now()
		diff := certs[0].NotAfter.Sub(now)
		certExpiry := int32(diff.Hours() / 24)

		err := certs[0].VerifyHostname(domain.Name)
		if err == nil && certExpiry > 10 {
			logrus.Printf("Certificate valid for domain %s, expires in %d days", domain.Name, certExpiry)
		} else {
			// need to warn
			logrus.Warnf("Certificate expires soon for domain %s, expires in %d days", domain.Name, certExpiry)
		}
	}
	return nil
}
