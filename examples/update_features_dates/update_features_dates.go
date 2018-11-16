package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/io/ioutilmore"
	"github.com/joho/godotenv"

	au "github.com/grokify/go-aha/ahautil"
	"github.com/grokify/go-aha/client"
)

var (
	outputFile                         = "_features.json"
	updateDefaultFeatureDatesToRelease = false
)

func main() {
	releases := []string{"API-R-1"}

	updateDefaultFeatureDatesToRelease := false

	err := godotenv.Load(os.Getenv("ENV_PATH"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apis := au.NewClientAPIs(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))

	featureMap := map[string]*aha.Feature{}

	ctx := context.Background()

	for _, releaseId := range releases {
		features := []*aha.Feature{}
		var err error
		if updateDefaultFeatureDatesToRelease {
			features, err = apis.UpdateFeatureStartDueDatesToRelease(ctx, releaseId)
		} else {
			features, err = apis.GetFeaturesFullByReleaseId(ctx, releaseId)
		}

		if err != nil {
			panic(err)
		}
		fmt.Printf("Release %v has %v features\n", releaseId, len(features))
		for _, feat := range features {
			featureMap[feat.Id] = feat
		}
	}

	fmtutil.PrintJSON(featureMap)
	fmt.Println(len(featureMap))

	if 1 == 1 {
		err = ioutilmore.WriteFileJSON(outputFile, featureMap, 0644, true)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("DONE")
}
