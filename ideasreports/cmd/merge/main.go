package main

import (
	"fmt"
	"path/filepath"

	"github.com/grokify/mogo/log/logutil"

	"github.com/grokify/go-aha/v3/ideasreports"
)

func main() {
	f1 := "aha_list_ideas_250115041716_axa-created.xlsx"
	f2 := "aha_list_ideas_250115024951_axa-votes.xlsx"

	dir := "/Users/johnwang/Downloads"
	f1 = filepath.Join(dir, f1)
	f2 = filepath.Join(dir, f2)

	is, err := ideasreports.ParseFilesXLSX([]string{f1, f2}, "axa.com")
	logutil.FatalErr(err)

	err = is.WriteXLSX("axa-ideas-2025-01-14.xlsx")
	logutil.FatalErr(err)

	fmt.Println("DONE")
}
