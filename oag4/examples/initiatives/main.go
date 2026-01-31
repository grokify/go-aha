package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/grokify/goauth"
	"github.com/grokify/mogo/encoding/jsonutil/jsonraw"
	"github.com/grokify/mogo/log/logutil"
	"github.com/grokify/mogo/net/http/httpsimple"
	flags "github.com/jessevdk/go-flags"
)

const (
	APIURLInitiatives = "/api/v1/initiatives"
)

func GetInitiatives(ctx context.Context, sc *httpsimple.Client) {
	req := httpsimple.Request{
		Method: http.MethodGet,
		URL:    APIURLInitiatives,
	}
	resp, err := sc.Do(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	b2, err := jsonraw.IndentBytes(b, "", "  ")
	logutil.FatalErr(err)
	fmt.Println(string(b2))
}

func main() {
	opts := goauth.Options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	creds, err := goauth.NewCredentialsFromSetFile(opts.CredsPath, opts.Account, true)
	logutil.FatalErr(err)

	sc, err := creds.OAuth2.NewSimpleClient(context.Background())
	logutil.FatalErr(err)

	GetInitiatives(context.Background(), sc)

	fmt.Println("DONE")
}
