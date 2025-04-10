package ideas

import (
	"fmt"

	"github.com/grokify/gocharts/v2/charts/google"
	"github.com/grokify/gocharts/v2/charts/google/piechart"
	"github.com/grokify/gocharts/v2/data/table"
	"github.com/grokify/gocharts/v2/data/table/tabulator"
)

func (ids Ideas) Report(domain string) (string, *table.Table) {
	id2 := ids.FilterByEmailDomain(domain, true, true)
	id2.SortVotes(true)
	cats := id2.HistogramCategories()
	statuses := id2.HistogramStatus()

	htm := fmt.Sprintf(`<html><head><script src="%s"></script>%s</head><body><h1>%s</h2>`,
		google.ChartsLoaderJS,
		tabulator.HTMLURLs,
		fmt.Sprintf("Ideas for %s", domain),
	)

	htm += fmt.Sprintf(`<p>%d ideas</p>`, len(id2))

	cht1 := piechart.NewPieChartMaterialInts("Ideas by Status", "Status", "Count", statuses.Bins)
	cht1.ChartDiv = "chart1"
	htm += cht1.HTML()

	cht2 := piechart.NewPieChartMaterialInts("Ideas by Category", "Category", "Count", cats.Bins)
	cht2.ChartDiv = "chart2"
	htm += cht2.HTML()

	tbs := id2.Table(ColumnsShort())

	htm += tbs.ToHTML(false)

	htm += "</body></html>"
	return htm, tbs
}
