package ideas

import (
	"fmt"

	"github.com/grokify/gocharts/v2/charts/google"
	"github.com/grokify/gocharts/v2/charts/google/piechart"
	"github.com/grokify/gocharts/v2/data/table"
	"github.com/grokify/gocharts/v2/data/table/tabulator"
)

func (ids Ideas) ReportForKeyword(keyword string) (string, *table.Table) {
	id2 := ids.FilterByNameDescKeyword(keyword, true, true, true)
	id2.SortVotes(true)
	return id2.Report(fmt.Sprintf("Ideas for Keyword: %s", keyword))
}

func (ids Ideas) ReportForDomain(domain string) (string, *table.Table) {
	id2 := ids.FilterByEmailDomain(domain, true, true)
	id2.SortVotes(true)
	return id2.Report(fmt.Sprintf("Ideas for Domain: %s", domain))
}

func (ids Ideas) Report(name string) (string, *table.Table) {
	cats := ids.HistogramCategories()
	statuses := ids.HistogramStatus()

	htm := fmt.Sprintf(`<html><head><script src="%s"></script>%s</head><body><h1>%s</h2>`,
		google.ChartsLoaderJS,
		tabulator.HTMLURLs,
		name,
	)

	htm += fmt.Sprintf(`<p>%d ideas</p>`, len(ids))

	cht1 := piechart.NewPieChartMaterialInts("Ideas by Status", "Status", "Count", statuses.Items)
	cht1.ChartDiv = "chart1"
	htm += cht1.HTML()

	cht2 := piechart.NewPieChartMaterialInts("Ideas by Category", "Category", "Count", cats.Items)
	cht2.ChartDiv = "chart2"
	htm += cht2.HTML()

	tbs := ids.Table(ColumnsShort())

	htm += tbs.ToHTML(false)

	htm += "</body></html>"
	return htm, tbs
}
