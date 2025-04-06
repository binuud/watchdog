package watchDogServer

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func (s *WatchDogService) PrintSummary() {
	fmt.Print("\n\nPrintSummary")
	fmt.Printf("\n%s", strings.Repeat("-", 80))
	fmt.Printf("\n%-30s %30s  %4s  %12s", "Domain", "Certs           ", "IP ", "Reachability")
	fmt.Printf("\n%-30s %4s %4s %4s %4s %8s   %4s   (%4s/%-4s)", "", "Tot", "Val", "Exp", "Days", "Status", "Num", "Vld", "Tot")
	fmt.Printf("\n%s", strings.Repeat("-", 80))
	for _, domainEntry := range s.Data {
		fmt.Printf("\n%-30s %4d %4d %4d %4d %8s   %4d   (%4d/%-4d)", domainEntry.Domain.Name,
			domainEntry.Summary.NumCerts,
			domainEntry.Summary.NumValidCerts,
			domainEntry.Summary.NumExpiringCerts,
			domainEntry.Summary.LeastCertExpiryInDays,
			domainEntry.Summary.CertsStatus[0].Status.String(),
			len(domainEntry.Info.IpAddresses),
			domainEntry.Summary.ValidEndpoints,
			domainEntry.Summary.NumEndpoints,
		)
	}
	fmt.Printf("\n\n%s\n\n", strings.Repeat("-", 80))
}

func (s *WatchDogService) PrintSummaryTable() {

	green := "\033[32m"
	orange := "\033[38;5;202m"
	reset := "\033[0m" // Reset color to default

	// Unicode character for check mark

	alertTick := fmt.Sprintf("%s⚠️%s\n", orange, reset)
	greenTick := fmt.Sprintf("%s✓%s\n", green, reset)

	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}

	t := table.NewWriter()
	// total 10 columns
	t.AppendHeader(table.Row{"Domains", "Domains", "Domains", "Certs", "Certs", "Certs", "Certs", "IP", "Reachable", "Reachable"}, rowConfigAutoMerge)
	t.AppendHeader(table.Row{"Name/Sub", "Mutated", "Expiry", "Tot", "Valid", "Expiry", "Validity", "(num)", "Valid", "Total"}, rowConfigAutoMerge)
	t.AppendHeader(table.Row{"", "(Days)", "(Days)", "(Counts)", "(Counts)", "(Days)", "", "(Counts)", "(Counts)", "(Counts)"}, rowConfigAutoMerge)
	for _, domainEntry := range s.Data {

		domainMutatedString := ""
		domainExpiryString := ""
		if domainEntry.Domain.DomainName != "" {
			tick := greenTick
			if domainEntry.Summary.WhoIsMutatedDays < 10 {
				tick = alertTick
			}
			domainMutatedString = fmt.Sprintf("%4d %2s", domainEntry.Summary.WhoIsMutatedDays, tick)

			tick = greenTick
			if domainEntry.Summary.DomainExpiryDays < 10 {
				tick = alertTick
			}
			domainExpiryString = fmt.Sprintf("%4d %2s", domainEntry.Summary.DomainExpiryDays, tick)
		}

		tick := greenTick
		if domainEntry.Summary.LeastCertExpiryInDays < 10 {
			tick = alertTick
		}
		certExpiryString := fmt.Sprintf("%4d %2s", domainEntry.Summary.LeastCertExpiryInDays, tick)

		t.AppendRow(table.Row{domainEntry.Domain.Name,
			domainMutatedString,
			domainExpiryString,
			domainEntry.Summary.NumCerts,
			domainEntry.Summary.NumValidCerts,
			certExpiryString,
			domainEntry.Summary.CertsStatus[0].Status.String(),
			len(domainEntry.Info.IpAddresses),
			domainEntry.Summary.ValidEndpoints,
			domainEntry.Summary.NumEndpoints,
		})
	}

	t.SetAutoIndex(true)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AutoMerge: true},
		{Number: 2, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		{Number: 3, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		{Number: 4, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		{Number: 6},
		{Number: 7, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		{Number: 8, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		{Number: 9, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		{Number: 10, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
	})
	t.AppendFooter(table.Row{"", "", "", "", "", "", "", "", "", ""})
	t.SetStyle(table.StyleLight)
	t.Style().Options.SeparateRows = true
	fmt.Println(t.Render())

}
