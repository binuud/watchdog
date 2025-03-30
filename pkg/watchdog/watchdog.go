package watchDogServer

import (
	"crypto/x509"
	"time"

	"github.com/binuud/watchdog/gen/go/v1/watchdog"
	"github.com/binuud/watchdog/pkg/domainUtils"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WatchDogService struct {
	DomainWatch *watchdog.DomainWatch
	Data        []*watchdog.DomainRow
}

func NewWatchDogService(fileName string) *WatchDogService {

	serverObj := &WatchDogService{}

	return serverObj.initFromConfig(fileName)

}

// read contents from the config yaml file
// build the in memory database
// this allows us to add/delete domains once the server is running
func (s *WatchDogService) initFromConfig(fileName string) *WatchDogService {

	domainWatch := &watchdog.DomainWatch{}
	readYaml(fileName, domainWatch)

	domainEntries := make([]*watchdog.DomainRow, 0)

	for _, item := range domainWatch.Domains {
		domainEntry := s.getDomainEntry(item.Name)
		if domainEntry == nil {
			// domain not in memory
			// create a new entry for same
			domainEntry = &watchdog.DomainRow{
				Domain: item,
				Info:   &watchdog.DomainInfo{},
				Summary: &watchdog.DomainSummary{
					Domain: item,
				},
			}
		} else {
			logrus.Fatalf("Remove duplicate domain entry %s", item.Name)
		}
		domainEntries = append(domainEntries, domainEntry)

	}
	s.DomainWatch = domainWatch
	s.Data = domainEntries
	return s
}

func (s *WatchDogService) CheckDomains() error {

	for _, domainEntry := range s.Data {

		// get domain details
		// persist data in domain Entry for summarization and storage
		s.getDomainDetails(domainEntry)

		// convert summarization to a different loop later
		s.summarize(domainEntry)
	}

	return nil

}

// get the domain row corresponding to the given domain name
// from the data object
func (s *WatchDogService) getDomainEntry(domainName string) *watchdog.DomainRow {
	for _, item := range s.Data {
		if domainName == item.Domain.Name {
			return item
		}
	}
	return nil
}

// get domain details like certificates
// ip resolv and reachablilty of a domain
func (s *WatchDogService) getDomainDetails(domainEntry *watchdog.DomainRow) {

	domain := domainEntry.Domain
	endpointStatuses := make([]*watchdog.EndpointStatus, 0)
	for _, item := range domain.Endpoints {
		responseCode, err := domainUtils.CheckHttpHead(item)
		if err != nil {
			// handle err
		}
		endPointStatus := &watchdog.EndpointStatus{
			Endpoint:   item,
			StatusCode: int64(responseCode),
		}
		endpointStatuses = append(endpointStatuses, endPointStatus)
		logrus.Printf("Received response %d %s", responseCode, item)
	}
	domainEntry.Info.EndpointStatuses = endpointStatuses

	// fetch the certs
	certs, err := domainUtils.GetCertificates(domain.Name, 443)
	if err != nil {
		logrus.Errorf("error when getting certificate of domain %s, %v", domain.Name, err)
	}
	for _, item := range certs {
		domainEntry.Info.Certificates = append(domainEntry.Info.Certificates, item.Raw)
	}

	ipAddresses, err := domainUtils.ResolveIP(domain.Name)
	if err != nil {
		logrus.Errorf("error when resolving ip of domain %s, %v", domain.Name, err)
	}

	// save the ip addresses in the domainRow object

	domainEntry.Info.IpAddresses = nil
	for _, ipAddr := range ipAddresses {
		logrus.Printf("Ip address of domain %s %v", domain.Name, ipAddr.String())
		domainEntry.Info.IpAddresses = append(domainEntry.Info.IpAddresses, ipAddr.String())
	}

}

func (s *WatchDogService) summarize(domainEntry *watchdog.DomainRow) {
	// fetch and analyse the certificates
	domainEntry.Summary.CreatedAt = timestamppb.Now()

	if len(domainEntry.Info.IpAddresses) > 0 {
		domainEntry.Summary.Resolvable = true
	} else {
		domainEntry.Summary.Resolvable = false
	}
	domainEntry.Summary.NumIp = int64(len(domainEntry.Info.IpAddresses))

	// reachable endpoints
	reachable := 0
	for _, item := range domainEntry.Info.EndpointStatuses {
		if item.StatusCode >= 100 && item.StatusCode < 400 {
			reachable++
		}
	}
	domainEntry.Summary.NumEndpoints = int64(len(domainEntry.Domain.Endpoints))
	domainEntry.Summary.ValidEndpoints = int64(reachable)

	if reachable == len(domainEntry.Info.EndpointStatuses) && reachable > 0 {
		domainEntry.Summary.Reachable = true
	} else {
		domainEntry.Summary.Reachable = false
	}

	err := s.analyseCertificates(domainEntry)
	if err != nil {
		logrus.Errorf("error when checking certificated of domain %s, %v", domainEntry.Domain.Name, err)
	}

}

func (s *WatchDogService) analyseCertificates(domainEntry *watchdog.DomainRow) error {

	domainEntry.Summary.CertsStatus = nil
	validCertCount := 0
	expiringCertCount := 0
	leastExpiry := 64000

	for _, certRaw := range domainEntry.Info.Certificates {

		cert, err := x509.ParseCertificate(certRaw)
		if err != nil {
			logrus.Errorf("cannot parse certificate for domain %s", domainEntry.Domain.Name)
		}

		now := time.Now()
		diff := cert.NotAfter.Sub(now)
		certExpiry := int64(diff.Hours() / 24)

		certSummary := &watchdog.CertificateStatus{
			CertExpiry: certExpiry,
			CertValid:  true,
			Status:     watchdog.CertificateStatus_Valid,
		}

		err = cert.VerifyHostname(domainEntry.Domain.Name)
		if err != nil {
			certSummary.CertValid = false
			certSummary.Status = watchdog.CertificateStatus_WrongCertificate
		} else {
			validCertCount++
		}
		if certExpiry < int64(leastExpiry) {
			leastExpiry = int(certExpiry)
		}
		if certExpiry < 10 {
			// need to warn
			expiringCertCount++
			certSummary.Status = watchdog.CertificateStatus_Expiring
			logrus.Warnf("Certificate expires soon for domain %s, expires in %d days", domainEntry.Domain.Name, certExpiry)
		} else {
			// logrus.Printf("Certificate valid domain %s, expires in %d days", domainEntry.Domain.Name, certExpiry)
		}
		domainEntry.Summary.CertsStatus = append(domainEntry.Summary.CertsStatus, certSummary)

	}
	domainEntry.Summary.NumCerts = int64(len(domainEntry.Info.Certificates))
	domainEntry.Summary.NumValidCerts = int64(validCertCount)
	domainEntry.Summary.NumExpiringCerts = int64(expiringCertCount)
	domainEntry.Summary.LeastExpiryInDays = int64(leastExpiry)
	return nil

}

func (s *WatchDogService) GetByNameOrUUID(name string, uuid string) *watchdog.DomainRow {
	for _, item := range s.Data {
		if item.Domain.Name == name || item.Domain.Uuid == uuid {
			return item
		}
	}
	return nil
}
