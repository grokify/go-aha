package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/log/logutil"
	"github.com/grokify/mogo/net/http/httpsimple"
	"github.com/jessevdk/go-flags"

	"github.com/grokify/go-aha/v3/oag7/aha"
	"github.com/grokify/go-aha/v3/oag7/client"
	"github.com/grokify/go-aha/v3/oag7/features"
	"github.com/grokify/go-aha/v3/oag7/ideas"
)

type Options struct {
	EnvFile string `short:"e" long:"env" description:"Env filepath"`
}

func main() {
	opts := Options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}
	_, err = config.LoadDotEnv([]string{opts.EnvFile, os.Getenv("ENV_PATH")}, -1)
	if err != nil {
		log.Fatal(err)
	}
	/*
			ahaAccount := os.Getenv("AHA_ACCOUNT")
			ahaAccount = "saviynt-product"
			ahaAccount = "saviynt-product.aha.io/"
			ahaAccount = "saviynt-product"
		 	ahaSvrURL := "https://saviynt-product.aha.io/"
			ahaSvrAPIURL := "https://saviynt-product.aha.io/api/v1"
			ahaHost := "saviynt-product.aha.io"
	*/
	ahaSubdomain := "saviynt-product"

	apiToken := "pDWjFwQNBtBZGrtFcPhY9EixcB1XT2-pSRf6bTO0TS4"

	ctx := context.Background()
	cfg, err := client.NewConfiguration(ahaSubdomain, apiToken)
	logutil.FatalErr(err)

	sc, err := client.NewSimpleClient(ahaSubdomain, apiToken)
	logutil.FatalErr(err)

	sr := httpsimple.Request{
		Method: http.MethodGet,
		// URL:    "https://saviynt-product.aha.io/api/v1/features/7489142925063719555",
		URL: "https://saviynt-product.aha.io/api/v1/ideas/7458108264465258721",
	}

	resp, err := sc.Do(ctx, sr)
	logutil.FatalErr(err)
	fmt.Printf("STATUS (%d)\n", resp.StatusCode)

	b, err := io.ReadAll(resp.Body)
	logutil.FatalErr(err)
	fmt.Println(string(b))

	clt := aha.NewAPIClient(cfg)
	ahaIdeasPortalURL := "https://ideas.saviynt.com/"
	ahaAdminURL := "https://saviynt-product.aha.io/"
	//ideaID := "7458108264465258721"
	//ideaID = "EIC-I-6538"

	ideaIDs := []string{
		"EIC-I-4513",
		"EIC-I-6250",
		"EIC-I-3510",
		"EIC-I-5683",
		"EIC-I-5623",
		"EIC-I-6178",
		"EIC-I-6998",
		"EIC-I-6178"}

	is, err := ideas.GetIdeaStatusSet(clt, ideaIDs, ahaIdeasPortalURL, ahaAdminURL)
	logutil.FatalErr(err)
	fmtutil.PrintJSON(is)

	tbl := is.Table()
	err = tbl.WriteXLSX("ideas.xlsx", "ideas")
	logutil.FatalErr(err)

	f, err := features.GetFeatureRaw(ctx, sc.HTTPClient, ahaAdminURL, "IN-710", true)
	logutil.FatalErr(err)
	fmt.Println(string(f))

	fmt.Println("DONE")
}
