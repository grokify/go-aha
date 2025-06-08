package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/grokify/gocharts/v2/data/roadmap"
	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/os/osutil"
	"github.com/grokify/mogo/type/stringsutil"
	"github.com/jessevdk/go-flags"

	"github.com/grokify/go-aha/v2/ahautil"
)

type Options struct {
	EnvFile            string `short:"e" long:"env" description:"Env filepath"`
	Products           string `short:"p" long:"products" description:"Product" required:"true"`
	ReleaseYYYYQBegin  int    `short:"b" long:"begin" description:"Begin Quarter"`
	ReleaseYYYYQFinish int    `short:"f" long:"finish" description:"Finish Quarter"`
	VerboseRaw         []bool `short:"v" long:"verbose" description:"Verbose"`
}

func (opts *Options) Verbose() bool {
	if len(opts.VerboseRaw) > 0 {
		return true
	}
	return false
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
	_, err = config.LoadDotEnv([]string{opts.EnvFile, os.Getenv("ENV_PATH")}, 1)
	if err != nil {
		log.Fatal(err)
	}

	if len(strings.TrimSpace(os.Getenv("AHA_API_KEY"))) == 0 {
		log.Printf("opts.EnvFile [%v]", opts.EnvFile)
		log.Printf("ENV_PATH [%v]", os.Getenv("ENV_PATH"))
		log.Fatal("E_NO_AHA_API_KEY")
	}
	if opts.Verbose() {
		fmt.Printf("AHA_ACCOUNT [%v]\n", os.Getenv("AHA_ACCOUNT"))
		fmt.Printf("AHA_API_KEY [%v]\n", os.Getenv("AHA_API_KEY"))
	}
	apis, err := ahautil.NewClientAPIs(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	products := stringsutil.SplitTrimSpace(strings.ToUpper(opts.Products), ",", true)

	if len(products) == 0 {
		log.Fatal("E_NO_PRODUCTS")
	}

	if opts.Verbose() {
		productSlug1 := products[0]
		info, resp, err := apis.APIClient.ProductsApi.GetProduct(context.Background(), productSlug1)
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode >= 300 {
			log.Fatal(resp.StatusCode)
		}
		fmtutil.MustPrintJSON(info)
		fmtutil.MustPrintJSON(opts)
	}

	opts.ReleaseYYYYQBegin, opts.ReleaseYYYYQFinish = roadmap.QuartersBeginEnd(
		opts.ReleaseYYYYQBegin, opts.ReleaseYYYYQFinish)

	rs, fs, err := ahautil.GetReleasesAndFeaturesForProductsAndQuarters(
		context.Background(), apis, products,
		opts.ReleaseYYYYQBegin, opts.ReleaseYYYYQFinish)
	if err != nil {
		log.Fatal(err)
	}
	if opts.Verbose() {
		fmtutil.MustPrintJSON(rs)
		fmtutil.MustPrintJSON(fs)
	}

	WriteFile(ReleasesFile, rs)
	WriteFile(FeaturesFile, fs)

	fmt.Println("DONE")
}

func WriteFile(fileName string, data interface{}) {
	err := osutil.WriteFileJSON(fileName, data, 0644, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("WROTE %v\n", fileName)
}
