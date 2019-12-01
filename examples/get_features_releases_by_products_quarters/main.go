package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/type/stringsutil"
	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"

	"github.com/grokify/go-aha/ahautil"
)

type Options struct {
	EnvFile              string `short:"e" long:"env" description:"Env filepath"`
	Products             string `short:"p" long:"products" description:"Product" required:"true"`
	ReleaseQuarterBegin  int32  `short:"b" long:"begin" description:"Begin Quarter"`
	ReleaseQuarterFinish int32  `short:"f" long:"finish" description:"Finish Quarter"`
}

func init() {
	err := godotenv.Load(os.Getenv("ENV_PATH"))
	if err != nil {
		log.Fatal("$ENV_PATH not found")
	}
}

func main() {
	opts := Options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	apis := ahautil.NewClientAPIs(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))

	products := stringsutil.SplitCondenseSpace(strings.ToUpper(opts.Products), ",")

	rs, fs, err := ahautil.GetReleasesAndFeaturesForProductsAndQuarters(
		context.Background(), apis, products,
		opts.ReleaseQuarterBegin, opts.ReleaseQuarterFinish)
	if err != nil {
		log.Fatal(err)
	}
	fmtutil.PrintJSON(rs)
	fmtutil.PrintJSON(fs)

	fmt.Println("DONE")
}
