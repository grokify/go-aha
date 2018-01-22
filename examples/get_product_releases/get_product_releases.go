package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/joho/godotenv"

	au "github.com/grokify/go-aha/ahautil"
)

func main() {
	err := godotenv.Load(os.Getenv("ENV_PATH"))
	if err != nil {
		panic(err)
	}

	apis := au.NewClientAPIs(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))

	api := apis.APIClient.ReleasesApi
	ctx := context.Background()

	params := map[string]interface{}{
		"page":     int32(1),
		"pagePage": int32(500),
	}

	info, resp, err := api.GetProductReleases(ctx, "GLIP", params)

	if err != nil {
		log.Fatal("Error retrieving features")
	}

	fmt.Println(resp.StatusCode)
	fmtutil.PrintJSON(info)
	fmt.Printf("Found %v releases\n", len(info.Releases))
	fmt.Println("===")

	/*
		for _, f := range info.Features {
			fmtutil.PrintJSON(f)

			feat, resp, err := api.GetFeature(ctx, f.Id)
			if err != nil {
				log.Fatal("Error retrieving feature")
			}

			fmt.Println(resp.StatusCode)
			fmtutil.PrintJSON(feat)

			fmt.Println("ESFeature")
			f2 := au.AhaToEsFeature(feat.Feature)
			fmtutil.PrintJSON(f2)

			break
		}
	*/
	fmt.Println("DONE")
}
