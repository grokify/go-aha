package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"

	"github.com/grokify/go-aha/v3/reports/ideas"
)

type Options struct {
	InputFileXLSX  string `short:"i" long:"inputfile" description:"XLSX Ideas File" required:"true"`
	Domain         string `short:"d" long:"domain" description:"email domain" required:"true"`
	IdeasPortalURL string `short:"p" long:"portalurl" description:"portal url" required:"true"`
}

func main() {
	opts := Options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	id1, err := ideas.ParseXLSX(opts.InputFileXLSX, opts.IdeasPortalURL)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(2)
	}

	domain := strings.ToLower(strings.TrimSpace(opts.Domain))

	htm, tbs := id1.Report(domain)

	err = os.WriteFile(fmt.Sprintf("ideas_%s.html", domain), []byte(htm), 0600)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(3)
	}

	err = tbs.WriteXLSX(fmt.Sprintf("ideas_%s_url.xlsx", domain), "Ideas Q1")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(4)
	}

	fmt.Println("DONE")
	os.Exit(0)
}
