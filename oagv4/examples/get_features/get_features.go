package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/antihax/optional"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/joho/godotenv"

	"github.com/grokify/go-aha/v3/oagv4/aha"
	au "github.com/grokify/go-aha/v3/oagv4/ahautil"
)

func main() {
	err := godotenv.Load(os.Getenv("ENV_PATH"))
	if err != nil {
		panic(err)
	}

	apis, err := au.NewClientAPIs(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	api := apis.APIClient.FeaturesApi
	ctx := context.Background()

	//params := map[string]interface{}{}

	params := &aha.GetFeaturesOpts{}

	if 1 == 1 {
		dt, err := time.Parse(time.RFC3339, "2017-12-01T00:00:00Z")
		if err != nil {
			panic(err)
		}
		/*
			params["updatedSince"] = dt
			params["page"] = int32(2)
			params["perPage"] = int32(500)
		*/
		params.UpdatedSince = optional.NewTime(dt)
		params.PerPage = optional.NewInt32(int32(500))
		params.PerPage = optional.NewInt32(int32(500))
	}

	info, resp, err := api.GetFeatures(ctx, params)

	if err != nil {
		log.Fatal("Error retrieving features")
	}

	fmt.Println(resp.StatusCode)
	fmtutil.MustPrintJSON(info)
	fmt.Printf("Found %v features\n", len(info.Features))
	fmt.Println("===")

	for _, f := range info.Features {
		fmtutil.MustPrintJSON(f)

		feat, resp, err := api.GetFeature(ctx, f.Id)
		if err != nil {
			log.Fatal("Error retrieving feature")
		}

		fmt.Println(resp.StatusCode)
		fmtutil.MustPrintJSON(feat)

		fmt.Println("ESFeature")
		f2 := au.AhaToEsFeature(&feat.Feature)
		fmtutil.MustPrintJSON(f2)

		break
	}

	fmt.Println("DONE")
}
