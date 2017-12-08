package main

import (
	"fmt"
	"os"
	"time"

	"github.com/grokify/gotilla/fmt/fmtutil"
	tu "github.com/grokify/gotilla/time/timeutil"
	"github.com/joho/godotenv"

	"github.com/grokify/go-aha"
	"github.com/grokify/go-aha/ahautil"
	au "github.com/grokify/oauth2util/aha"
)

func main() {
	productId := "PROD"
	releaseId := "PROD-R-1"
	beginDate := "2018-01-01"
	endDate := "2018-03-31"

	err := godotenv.Load(os.Getenv("ENV_PATH"))
	if err != nil {
		panic(err)
	}

	client := au.NewClient(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))
	apis := ahautil.ClientAPIs{Client: client}

	// Get Current Release Info
	rel, err := apis.GetReleaseById(releaseId)
	if err != nil {
		panic(err)
	}

	fmt.Println("Current Release Info")
	fmtutil.PrintJSON(rel)

	// Update Release
	dt1, _ := time.Parse(tu.RFC3339YMD, beginDate)
	dt2, _ := time.Parse(tu.RFC3339YMD, endDate)

	rapi := apis.ReleasesApi()
	body := aha.ReleaseUpdateWrap{
		Release: aha.ReleaseUpdate{
			StartDate:            tu.RFC3339YMDTime{Time: dt1},
			DevelopmentStartedOn: tu.RFC3339YMDTime{Time: dt1},
			ReleaseDate:          tu.RFC3339YMDTime{Time: dt2},
			ExternalReleaseDate:  tu.RFC3339YMDTime{Time: dt2},
			ParkingLot:           false,
		},
	}

	fmt.Println("Release Request Body")
	fmtutil.PrintJSON(body)

	relUpdate, resp, err := rapi.ProductsProductIdReleasesReleaseIdPut(productId, releaseId, body)
	if err != nil {
		panic(err)
	} else if resp.StatusCode > 299 {
		panic(fmt.Errorf("CODE %v", resp.StatusCode))
	}

	// Print Updated Release Info
	fmt.Println("Updated Release Info")
	fmtutil.PrintJSON(relUpdate)

	fmt.Println("DONE")
}
