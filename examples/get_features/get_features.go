package main

import (
	"fmt"
	"log"
	"os"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/time/timeutil"
	"github.com/joho/godotenv"

	"github.com/grokify/go-aha/ahautil"
	au "github.com/grokify/oauth2util/aha"
)

func main() {
	err := godotenv.Load(os.Getenv("ENV_PATH"))
	if err != nil {
		panic(err)
	}

	client := au.NewClient(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))
	apis := ahautil.ClientAPIs{Client: client}
	api := apis.FeaturesApi()

	info, resp, err := api.FeaturesGet("", timeutil.TimeRFC3339Zero(), "", "", 1, 500)
	if err != nil {
		log.Fatal("Error retrieving features")
	}

	fmt.Println(resp.StatusCode)
	fmtutil.PrintJSON(info)
	fmt.Println("===")

	for _, f := range info.Features {
		fmtutil.PrintJSON(f)

		feat, resp, err := api.FeaturesFeatureIdGet(f.Id)
		if err != nil {
			log.Fatal("Error retrieving feature")
		}

		fmt.Println(resp.StatusCode)
		fmtutil.PrintJSON(feat)
		break
	}

	fmt.Println("DONE")
}
