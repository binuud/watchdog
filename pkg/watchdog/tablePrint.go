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
			domainEntry.Summary.LeastExpiryInDays,
			domainEntry.Summary.CertsStatus[0].Status.String(),
			len(domainEntry.Info.IpAddresses),
			domainEntry.Summary.ValidEndpoints,
			domainEntry.Summary.NumEndpoints,
		)
	}
	fmt.Printf("\n\n%s\n\n", strings.Repeat("-", 80))
}

func (s *WatchDogService) PrintSummaryTable() {

	// rowConfigAutoMerge := table.RowConfig{AutoMerge: true}

	// t := table.NewWriter()
	// t.AppendHeader(table.Row{"Node IP", "Pods", "Namespace", "Container", "RCE", "RCE"}, rowConfigAutoMerge)
	// t.AppendHeader(table.Row{"Node IP", "Pods", "Namespace", "Container", "EXE", "RUN"})
	// t.AppendRow(table.Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 1", "Y", "Y"}, rowConfigAutoMerge)
	// t.AppendRow(table.Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "N"}, rowConfigAutoMerge)
	// t.AppendRow(table.Row{"1.1.1.1", "Pod 1A", "NS 1B", "C 3", "N", "N"}, rowConfigAutoMerge)
	// t.AppendRow(table.Row{"1.1.1.1", "Pod 1B", "NS 2", "C 4", "N", "N"}, rowConfigAutoMerge)
	// t.AppendRow(table.Row{"1.1.1.1", "Pod 1B", "NS 2", "C 5", "Y", "N"}, rowConfigAutoMerge)
	// t.AppendRow(table.Row{"2.2.2.2", "Pod 2", "NS 3", "C 6", "Y", "Y"}, rowConfigAutoMerge)
	// t.AppendRow(table.Row{"2.2.2.2", "Pod 2", "NS 3", "C 7", "Y", "Y"}, rowConfigAutoMerge)
	// t.AppendFooter(table.Row{"", "", "", 7, 5, 3})
	// t.SetAutoIndex(true)
	// t.SetColumnConfigs([]table.ColumnConfig{
	// 	{Number: 1, AutoMerge: true},
	// 	{Number: 2, AutoMerge: true},
	// 	{Number: 3, AutoMerge: true},
	// 	{Number: 4, AutoMerge: true},
	// 	{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
	// 	{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
	// })
	// t.SetStyle(table.StyleLight)
	// t.Style().Options.SeparateRows = true
	// fmt.Println(t.Render())

	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}

	t := table.NewWriter()

	t.AppendHeader(table.Row{"Domains", "Certs", "Certs", "Certs", "Certs", "Certs", "IP", "Reachable", "Reachable"}, rowConfigAutoMerge)
	t.AppendHeader(table.Row{"Domains", "Tot", "Valid", "Exp", "Days", "Validity", "(num)", "Valid", "Total"}, rowConfigAutoMerge)
	for _, domainEntry := range s.Data {
		t.AppendRow(table.Row{domainEntry.Domain.Name,
			domainEntry.Summary.NumCerts,
			domainEntry.Summary.NumValidCerts,
			domainEntry.Summary.NumExpiringCerts,
			domainEntry.Summary.LeastExpiryInDays,
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
	})
	t.AppendFooter(table.Row{"", "", "", "", "", "", "", "", ""})
	t.SetStyle(table.StyleLight)
	t.Style().Options.SeparateRows = true
	fmt.Println(t.Render())

}
