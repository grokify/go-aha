package main

import (
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strings"

	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/jessevdk/go-flags"

	"github.com/grokify/go-aha/v3/reports/ideas"
)

type Options struct {
	InputFileXLSX  string `short:"i" long:"inputfile" description:"XLSX Ideas File" required:"true"`
	Domain         string `short:"d" long:"domain" description:"email domain"`
	Keyword        string `short:"k" long:"keyword" description:"keyword" required:"true"`
	IdeasPortalURL string `short:"p" long:"portalurl" description:"portal url" required:"true"`
}

func main() {
	opts := Options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	opts.Keyword = strings.TrimSpace(opts.Keyword)

	id1, err := ideas.ParseXLSX(opts.InputFileXLSX, opts.IdeasPortalURL)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(2)
	}

	// id1 = id1.FilterByNameDescKeyword(opts.Keyword, true, true, true)
	id1 = id1.FilterByNameDescRegexp(regexp.MustCompile(`(?i)\b((sys)?log(4j|s|file|ger|ging)?)\b`), true, true)

	id1 = id1.FilterByCategories([]string{
		"administration",
		"Archival \u0026 Retrieval",
		"password management",
		"security"}, true)

	fmtutil.MustPrintJSON(id1.HistogramCategories().Items)
	fmtutil.MustPrintJSON(id1.HistogramStatuses())
	fmt.Printf("Idea Count (%d)\n", len(id1))

	id1 = id1.FilterByStatusCategory([]string{"", "in progress"}, true)
	fmtutil.MustPrintJSON(id1.HistogramStatuses())
	fmt.Printf("Idea Count (%d)\n", len(id1))

	htm, tbs := id1.Report("Ideas for keyword: " + opts.Keyword)

	fileRoot := fmt.Sprintf("ideas_keyword_%s", opts.Keyword)
	fileHTML := fileRoot + ".html"
	fileXLSX := fileRoot + ".xlsx"

	err = os.WriteFile(fileHTML, []byte(htm), 0600)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(3)
	} else {
		fmt.Printf("Wrote (%s)\n", fileHTML)
	}

	err = tbs.WriteXLSX(fileXLSX, "Ideas")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(4)
	} else {
		fmt.Printf("Wrote (%s)\n", fileXLSX)
	}

	fmt.Println("DONE")
	os.Exit(0)
}
