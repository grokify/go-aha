package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/log/logutil"
	"github.com/grokify/mogo/net/http/httpsimple"
	"github.com/grokify/mogo/os/osutil"
	"github.com/jessevdk/go-flags"

	ao "github.com/grokify/goauth/aha"
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

	ahaAccount := os.Getenv("AHA_ACCOUNT")
	ahaAccount = "saviynt-product"
	ahaAccount = "saviynt-product.aha.io/"
	ahaAccount = "saviynt-product"
	ahaSvrURL := "https://saviynt-product.aha.io/"

	ahaKey := "pDWjFwQNBtBZGrtFcPhY9EixcB1XT2-pSRf6bTO0TS4"

	ctx := context.Background()

	hc, err := ao.NewClient(ahaAccount, ahaKey)
	logutil.FatalErr(err)

	sr := httpsimple.Request{
		Method: http.MethodGet,
		// URL:    "https://saviynt-product.aha.io/api/v1/features/7489142925063719555",
		URL: "https://saviynt-product.aha.io/api/v1/ideas/7458108264465258721",
	}
	sc := httpsimple.NewClient(hc, ahaSvrURL)
	resp, err := sc.Do(ctx, sr)
	logutil.FatalErr(err)
	fmt.Printf("STATUS (%d)\n", resp.StatusCode)

	b, err := io.ReadAll(resp.Body)
	logutil.FatalErr(err)
	fmt.Println(string(b))

	/*
			        "SYAHA": {
		            "type": "oauth2",
		            "service": "aha",
		            "subdomain": "saviynt-product",
		            "oauth2": {
		                "serverURL": "https://saviynt-product.aha.io/",
		                "token": {
		                    "access_token": "pDWjFwQNBtBZGrtFcPhY9EixcB1XT2-pSRf6bTO0TS4"
		                }
		            }
		        },
	*/
	/*
		apis, err := ahautil.NewClientAPIs(ahaAccount, ahaKey)
		if err != nil {
			log.Fatal(err)
		}
		api := apis.APIClient.FeaturesApi

		//params := map[string]interface{}{}

		params := &aha.GetFeaturesOpts{}

		if 1 == 1 {
			dt, err := time.Parse(time.RFC3339, "2025-06-01T00:00:00Z")
			if err != nil {
				panic(err)
			}
			//	params["updatedSince"] = dt
			//	params["page"] = int32(2)
			//	params["perPage"] = int32(500)
			params.UpdatedSince = optional.NewTime(dt)
			params.PerPage = optional.NewInt32(int32(500))
			params.PerPage = optional.NewInt32(int32(500))
		}

		info, resp, err := api.GetFeatures(ctx, params)
		logutil.FatalErr(err)
		if resp.StatusCode > 299 {
			log.Fatalf("bad status code (%d)", resp.StatusCode)
		}
		fmtutil.PrintJSON(info)
	*/
	fmt.Println("DONE")
}

func WriteFile(fileName string, data interface{}) {
	err := osutil.WriteFileJSON(fileName, data, 0644, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("WROTE %v\n", fileName)
}
