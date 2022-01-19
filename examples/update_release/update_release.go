package main

import (
	"context"
	"fmt"
	"os"

	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/joho/godotenv"

	"github.com/grokify/go-aha/v2/aha"
	au "github.com/grokify/go-aha/v2/ahautil"
)

func main() {
	productId := "PROD"
	releaseId := "PROD-R-1"
	beginDate := "2018-01-01"
	endDate := "2018-03-31"
	updateRelease := false

	err := godotenv.Load(os.Getenv("ENV_PATH"))
	if err != nil {
		panic(err)
	}

	apis := au.NewClientAPIs(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))
	releasesApi := apis.APIClient.ReleasesApi
	ctx := context.Background()

	// Get Current Release Info
	rel, resp, err := releasesApi.GetRelease(ctx, releaseId)
	if err != nil {
		panic(err)
	} else if resp.StatusCode >= 300 {
		panic(fmt.Errorf("Status Code %v", resp.StatusCode))
	}

	fmt.Println("Current Release Info")
	fmtutil.PrintJSON(rel)

	// Update Release
	if updateRelease {
		body := aha.ReleaseUpdateWrap{
			Release: aha.ReleaseUpdate{
				StartDate:            beginDate,
				DevelopmentStartedOn: beginDate,
				ReleaseDate:          endDate,
				ExternalReleaseDate:  endDate,
				ParkingLot:           false,
			},
		}

		fmt.Println("Release Request Body")
		fmtutil.PrintJSON(body)

		relUpdate, resp, err := releasesApi.UpdateProductRelease(
			ctx, productId, releaseId, body)
		if err != nil {
			panic(err)
		} else if resp.StatusCode > 299 {
			panic(fmt.Errorf("CODE %v", resp.StatusCode))
		}

		// Print Updated Release Info
		fmt.Println("Updated Release Info")
		fmtutil.PrintJSON(relUpdate)
	}

	fmt.Println("DONE")
}
