package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/joho/godotenv"

	"github.com/grokify/go-aha/v2/aha"
	au "github.com/grokify/go-aha/v2/ahautil"
)

func main() {
	featureId := "API-1"
	startDate := "2018-01-01"
	dueDate := "2018-06-30"

	err := godotenv.Load(os.Getenv("ENV_PATH"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apis := au.NewClientAPIs(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))
	fapi := apis.APIClient.FeaturesApi

	featureUpdate := aha.FeatureUpdate{
		StartDate: startDate,
		DueDate:   dueDate}

	// Update Feature with Dates
	_, res, err := fapi.UpdateFeature(context.Background(), featureId, featureUpdate)
	if err != nil {
		panic(err)
	}
	if res.StatusCode > 299 {
		panic(fmt.Sprintf("Bad Status Code %v", res.StatusCode))
	}

	feat, err := apis.GetFeatureById(featureId)
	if err != nil {
		panic(err)
	}
	fmtutil.PrintJSON(feat)

	fmt.Println("DONE")
}
