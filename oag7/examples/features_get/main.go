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
	ahaSubdomain := os.Getenv("AHA_ACCOUNT")
	apiToken := os.Getenv("AHA_API_KEY")

	ctx := context.Background()
	cfg, err := client.NewConfiguration(ahaSubdomain, apiToken)
	logutil.FatalErr(err)

	sc, err := client.NewSimpleClient(ahaSubdomain, apiToken)
	logutil.FatalErr(err)

	ahaAdminURL := fmt.Sprintf("https://%s.aha.io/", ahaSubdomain)

	sr := httpsimple.Request{
		Method: http.MethodGet,
		URL:    ahaAdminURL + "api/v1/ideas/IDEA-1",
	}

	resp, err := sc.Do(ctx, sr)
	logutil.FatalErr(err)
	fmt.Printf("STATUS (%d)\n", resp.StatusCode)

	b, err := io.ReadAll(resp.Body)
	logutil.FatalErr(err)
	fmt.Println(string(b))

	clt := aha.NewAPIClient(cfg)
	ahaIdeasPortalURL := "https://ideas.example.com/"

	ideaIDs := []string{
		"IDEA-1",
		"IDEA-2",
		"IDEA-3"}

	is, err := ideas.GetIdeaStatusSet(clt, ideaIDs, ahaIdeasPortalURL, ahaAdminURL)
	logutil.FatalErr(err)
	fmtutil.PrintJSON(is)

	tbl := is.Table()
	err = tbl.WriteXLSX("ideas.xlsx", "ideas")
	logutil.FatalErr(err)

	f, err := features.GetFeatureRaw(ctx, sc.HTTPClient, ahaAdminURL, "FEAT-1", true)
	logutil.FatalErr(err)
	fmt.Println(string(f))

	fmt.Println("DONE")
}
