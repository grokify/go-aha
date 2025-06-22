package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/os/osutil"
	"github.com/jessevdk/go-flags"

	"github.com/grokify/go-aha/v3/oagv4/ahautil"
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
	_, err = config.LoadDotEnv([]string{opts.EnvFile, os.Getenv("ENV_PATH")}, -1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(os.Getenv("AHA_ACCOUNT"))
	fmt.Println(os.Getenv("AHA_API_KEY"))
	apis, err := ahautil.NewClientAPIs(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	info, resp, err := apis.APIClient.FeaturesApi.GetFeature(
		context.Background(),
		opts.FeatureId)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode >= 300 {
		log.Fatal(fmt.Sprintf("Status Code [%v]\n", resp.StatusCode))
	}

	fmtutil.MustPrintJSON(info)

	fmt.Println("DONE")
}

func WriteFile(fileName string, data interface{}) {
	err := osutil.WriteFileJSON(fileName, data, 0644, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("WROTE %v\n", fileName)
}
