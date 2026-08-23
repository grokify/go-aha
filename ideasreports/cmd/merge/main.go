package main

import (
	"fmt"
	"path/filepath"

	"github.com/grokify/mogo/log/logutil"

	"github.com/grokify/go-aha/v3/ideasreports"
)

func main() {
	f1 := "aha_list_ideas_created.xlsx"
	f2 := "aha_list_ideas_votes.xlsx"

	dir := "."
	f1 = filepath.Join(dir, f1)
	f2 = filepath.Join(dir, f2)

	is, err := ideasreports.ParseFilesXLSX([]string{f1, f2}, "example.com")
	logutil.FatalErr(err)

	err = is.WriteXLSX("ideas-merged.xlsx")
	logutil.FatalErr(err)

	fmt.Println("DONE")
}
