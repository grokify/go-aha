package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/antihax/optional"
	"github.com/grokify/gotilla/fmt/fmtutil"
	hum "github.com/grokify/gotilla/net/httputilmore"
	"github.com/joho/godotenv"

	au "github.com/grokify/go-aha/ahautil"
	"github.com/grokify/go-aha/client"
)

func main() {
	oldTag := "My Old Tag"
	newTag := "My New Tag"
	updateFeatureTag := false

	err := godotenv.Load(os.Getenv("ENV_PATH"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apis := au.NewClientAPIs(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))
	featuresApi := apis.APIClient.FeaturesApi
	ctx := context.Background()

	params := aha.GetFeaturesOpts{
		Tag:     optional.NewString(oldTag),
		PerPage: optional.NewInt32(int32(500))}
	/*
		   	Q              optional.String
		   	UpdatedSince   optional.Time
		   	Tag            optional.String
		   	AssignedToUser optional.String
		   	Page           optional.Int32
		   	PerPage        optional.Int32
		   }


		fsRes, resp, err := featuresApi.FeaturesGet(ctx, map[string]interface{}{
			"tag":      oldTag,
			"per_page": 500,
		})*/
	fsRes, resp, err := featuresApi.GetFeatures(ctx, &params)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode >= 300 {
		panic(fmt.Errorf("Status Code: %v", resp.StatusCode))
	}

	fmtutil.PrintJSON(fsRes)
	hum.PrintResponse(resp, true)

	for _, fThin := range fsRes.Features {
		fmtutil.PrintJSON(fThin)

		fFull, resp, err := featuresApi.GetFeature(ctx, fThin.Id)
		if err != nil {
			panic(err)
		} else if resp.StatusCode >= 300 {
			panic(fmt.Errorf("Status Code: %v", resp.StatusCode))
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
				updateRes, resp, err := featuresApi.UpdateFeature(ctx, fThin.Id, fUpdate)
				if err != nil {
					panic(err)
				} else if resp.StatusCode >= 300 {
					panic(fmt.Errorf("Status Code: %v", resp.StatusCode))
				}
				fmtutil.PrintJSON(updateRes)
			}
		}
	}

	fmt.Printf("Found %v features\n", len(fsRes.Features))

	fmt.Println("DONE")
}
