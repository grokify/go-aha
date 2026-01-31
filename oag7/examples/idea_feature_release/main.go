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
			ahaAccount = "company"
			ahaAccount = "company.aha.io/"
			ahaAccount = "company"
		 	ahaSvrURL := "https://company.aha.io/"
			ahaSvrAPIURL := "https://company.aha.io/api/v1"
			ahaHost := "company.aha.io"
	*/
	ahaSubdomain := "company"

	apiToken := "REDACTED-TOKEN"

	ctx := context.Background()
	cfg, err := client.NewConfiguration(ahaSubdomain, apiToken)
	logutil.FatalErr(err)

	sc, err := client.NewSimpleClient(ahaSubdomain, apiToken)
	logutil.FatalErr(err)

	sr := httpsimple.Request{
		Method: http.MethodGet,
		// URL:    "https://company.aha.io/api/v1/features/7489142925063719555",
		URL: "https://company.aha.io/api/v1/ideas/7458108264465258721",
	}

	resp, err := sc.Do(ctx, sr)
	logutil.FatalErr(err)
	fmt.Printf("STATUS (%d)\n", resp.StatusCode)

	b, err := io.ReadAll(resp.Body)
	logutil.FatalErr(err)
	fmt.Println(string(b))

	clt := aha.NewAPIClient(cfg)
	//ahaIdeasPortalURL := "https://ideas.example.com/"
	//ahaAdminURL := "https://company.aha.io/"
	//ideaID := "7458108264465258721"
	//ideaID = "EIC-I-6538"

	ideaID := "EIC-I-3065"
	ideaID = "EIC-I-3558"

	res, err := ideas.QueryIdeaFeatureReleaseMetadata(clt, ideaID)
	logutil.FatalErr(err)
	fmtutil.PrintJSON(res)

	fmt.Println("DONE")
}
