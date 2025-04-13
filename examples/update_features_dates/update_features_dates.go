package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/os/osutil"
	"github.com/joho/godotenv"

	"github.com/grokify/go-aha/v2/aha"
	au "github.com/grokify/go-aha/v2/ahautil"
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

	apis, err := au.NewClientAPIs(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}

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

	fmtutil.MustPrintJSON(featureMap)
	fmt.Println(len(featureMap))

	if 1 == 1 {
		err = osutil.WriteFileJSON(outputFile, featureMap, os.FileMode(0644), "", "  ")
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("DONE")
}
