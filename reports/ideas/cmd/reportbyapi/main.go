package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/grokify/goauth"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/log/logutil"
	"github.com/grokify/mogo/net/http/httpsimple"
)

type Options struct {
	InputFileXLSX  string `short:"i" long:"inputfile" description:"XLSX Ideas File" required:"true"`
	Domain         string `short:"d" long:"domain" description:"email domain" required:"true"`
	IdeasPortalURL string `short:"p" long:"portalurl" description:"portal url" required:"true"`
}

func main() {
	c := "/Users/johnwang/go/src/github.com/grokify/grokify-channels/goauth.json"
	a := "SYAHA"
	creds, err := goauth.NewCredentialsFromSetFile(c, a, true)
	logutil.FatalErr(err)
	fmtutil.PrintJSON(creds)

	//clt, err := creds.NewClient(context.Background())
	sc, err := creds.NewSimpleClient(context.Background())
	logutil.FatalErr(err)

	id := "7517215819393881564"

	sr := httpsimple.Request{
		Method: http.MethodGet,
		URL:    fmt.Sprintf("/api/v1/bookmarks/custom_pivots/%s", id),
	}
	resp, err := sc.Do(context.Background(), sr)
	logutil.FatalErr(err)

	b, err := io.ReadAll(resp.Body)
	logutil.FatalErr(err)
	fmt.Println(string(b))

	err = os.WriteFile("output.json", b, 0600)

	fmt.Println(id)

	os.Exit(0)
}
