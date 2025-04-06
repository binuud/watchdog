package watchDogServer

import (
	"time"

	"github.com/binuud/watchdog/gen/go/v1/watchdog"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
	log "github.com/sirupsen/logrus"
)

func (s *WatchDogService) getWhoIsRecord(d *watchdog.DomainRow) error {

	var err error
	if d.Domain.DomainName == "" {
		d.Info.Whois = ""
		return nil
	}

	d.Info.Whois, err = whois.Whois(d.Domain.DomainName)
	// log.Printf("Who is for %s \n", d.Domain.DomainName)
	if err != nil {
		return err
	}

	return nil

}

func (s *WatchDogService) summarizeWhoIs(d *watchdog.DomainRow) error {

	w, err := whoisparser.Parse(d.Info.Whois)
	if err != nil {
		return err
	}
	log.Println("###################################")
	log.Printf("Who is summary for %s \n", d.Domain.DomainName)
	log.Println("###################################")

	// Get the current time
	currentTime := time.Now()

	// Calculate domain expiry
	domainExpiryDuration := w.Domain.ExpirationDateInTime.Sub(currentTime)

	// Convert duration to days
	d.Summary.DomainExpiring = false
	domainExpiry := int(domainExpiryDuration.Hours() / 24)
	if domainExpiry < 10 {
		d.Summary.DomainExpiring = true
	}

	// Convert domain updated time
	d.Summary.WhoIsMutated = false
	domainUpdatedDuration := currentTime.Sub(*w.Domain.UpdatedDateInTime)

	domainUpdatedDays := int(domainUpdatedDuration.Hours() / 24)
	if domainUpdatedDays < 10 {
		// who is data was changed in last 10 days,
		// set who is mutated to true
		d.Summary.WhoIsMutated = true
	}

	d.Summary.DomainExpiryAt = w.Domain.ExpirationDateInTime.Unix()
	d.Summary.WhoIsMutatedAt = w.Domain.UpdatedDateInTime.Unix()
	d.Summary.DomainExpiryDays = int64(domainExpiry)
	d.Summary.WhoIsMutatedDays = int64(domainUpdatedDays)

	return nil
}
