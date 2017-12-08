package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/grokify/gotilla/fmt/fmtutil"
	hum "github.com/grokify/gotilla/net/httputilmore"
	tu "github.com/grokify/gotilla/time/timeutil"
	"github.com/joho/godotenv"

	"github.com/grokify/go-aha"
	"github.com/grokify/go-aha/ahautil"
	au "github.com/grokify/oauth2util/aha"
)

func main() {
	oldTag := "My Old Tag"
	newTag := "My New Tag"
	updateFeatureTag := true

	err := godotenv.Load(os.Getenv("ENV_PATH"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := au.NewClient(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))
	apis := ahautil.ClientAPIs{Client: client}

	featuresApi := apis.FeaturesApi()

	fsRes, res, err := featuresApi.FeaturesGet("", tu.TimeRFC3339Zero(), oldTag, "")
	if err != nil {
		panic(err)
	}
	if res.StatusCode >= 300 {
		panic(fmt.Errorf("Status Code: %v", res.StatusCode))
	}

	fmtutil.PrintJSON(fsRes)
	hum.PrintResponse(res.Response, true)

	for _, fThin := range fsRes.Features {
		fmtutil.PrintJSON(fThin)

		fFull, res, err := featuresApi.FeaturesFeatureIdGet(fThin.Id)
		if err != nil {
			panic(err)
		} else if res.StatusCode >= 300 {
			panic(fmt.Errorf("Status Code: %v", res.StatusCode))
		}

		fmtutil.PrintJSON(fFull)
		fmtutil.PrintJSON(fFull.Feature.Tags)

		if updateFeatureTag {
			newTags := []string{}
			hasOldTag := false
			for _, tag := range fFull.Feature.Tags {
				if tag == oldTag {
					newTags = append(newTags, newTag)
					hasOldTag = true
				} else {
					newTags = append(newTags, tag)
				}
			}
			if hasOldTag {
				fUpdate := aha.FeatureUpdate{Tags: strings.Join(newTags, ",")}
				updateRes, res, err := featuresApi.FeaturesFeatureIdPut(fThin.Id, fUpdate)
				if err != nil {
					panic(err)
				} else if res.StatusCode >= 300 {
					panic(fmt.Errorf("Status Code: %v", res.StatusCode))
				}
				fmtutil.PrintJSON(updateRes)
			}
		}
	}

	fmt.Printf("Found %v features\n", len(fsRes.Features))

	fmt.Println("DONE")
}
