package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/io/ioutilmore"
	"github.com/grokify/gotilla/type/stringsutil"
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"

	"github.com/grokify/go-aha/ahautil"
)

type Options struct {
	EnvFile              string `short:"e" long:"env" description:"Env filepath"`
	Products             string `short:"p" long:"products" description:"Product" required:"true"`
	ReleaseQuarterBegin  int32  `short:"b" long:"begin" description:"Begin Quarter"`
	ReleaseQuarterFinish int32  `short:"f" long:"finish" description:"Finish Quarter"`
	Verbose              []bool `short:"v" long:"verbose" description:"Verbose"`
}

const (
	ReleasesFile string = "_releases.json"
	FeaturesFile string = "_features.json"
)

func main() {
	opts := Options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}
	err = config.LoadDotEnvSkipEmpty(opts.EnvFile, os.Getenv("ENV_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	if len(strings.TrimSpace(os.Getenv("AHA_API_KEY"))) == 0 {
		log.Infof("opts.EnvFile [%v]", opts.EnvFile)
		log.Infof("ENV_PATH [%v]", os.Getenv("ENV_PATH"))
		log.Fatal("E_NO_AHA_API_KEY")
	}
	if len(opts.Verbose) > 0 {
		fmt.Printf("AHA_ACCOUNT [%v]\n", os.Getenv("AHA_ACCOUNT"))
		fmt.Printf("AHA_API_KEY [%v]\n", os.Getenv("AHA_API_KEY"))
	}
	apis := ahautil.NewClientAPIs(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))

	products := stringsutil.SplitCondenseSpace(strings.ToUpper(opts.Products), ",")

	if len(products) == 0 {
		log.Fatal("E_NO_PRODUCTS")
	}

	if len(opts.Verbose) > 0 {
		productSlug1 := products[0]
		info, resp, err := apis.APIClient.ProductsApi.GetProduct(context.Background(), productSlug1)
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode >= 300 {
			log.Fatal(resp.StatusCode)
		}
		fmtutil.PrintJSON(info)
		fmtutil.PrintJSON(opts)
	}

	rs, fs, err := ahautil.GetReleasesAndFeaturesForProductsAndQuarters(
		context.Background(), apis, products,
		opts.ReleaseQuarterBegin, opts.ReleaseQuarterFinish)
	if err != nil {
		log.Fatal(err)
	}
	fmtutil.PrintJSON(rs)
	fmtutil.PrintJSON(fs)

	WriteFile(ReleasesFile, rs)
	WriteFile(FeaturesFile, fs)

	fmt.Println("DONE")
}

func WriteFile(fileName string, data interface{}) {
	err := ioutilmore.WriteFileJSON(fileName, data, 0644, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("WROTE %v\n", fileName)
}
