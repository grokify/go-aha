package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/io/ioutilmore"
	"github.com/jessevdk/go-flags"

	"github.com/grokify/go-aha/v2/ahautil"
)

type Options struct {
	EnvFile   string `short:"e" long:"env" description:"Env filepath"`
	FeatureId string `short:"f" long:"feature" description:"Feature" required:"true"`
}

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

	fmt.Println(os.Getenv("AHA_ACCOUNT"))
	fmt.Println(os.Getenv("AHA_API_KEY"))
	apis := ahautil.NewClientAPIs(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))

	info, resp, err := apis.APIClient.FeaturesApi.GetFeature(
		context.Background(),
		opts.FeatureId)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode >= 300 {
		log.Fatal(fmt.Sprintf("Status Code [%v]\n", resp.StatusCode))
	}

	fmtutil.PrintJSON(info)

	fmt.Println("DONE")
}

func WriteFile(fileName string, data interface{}) {
	err := ioutilmore.WriteFileJSON(fileName, data, 0644, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("WROTE %v\n", fileName)
}
