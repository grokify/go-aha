package main

import (
	"fmt"
	"log"
	"os"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/io/ioutilmore"
	"github.com/joho/godotenv"

	"github.com/grokify/go-aha"
	"github.com/grokify/go-aha/ahautil"
	au "github.com/grokify/oauth2util/aha"
)

func main() {
	releases := []string{"PROD-R-1"}
	updateDefaultFeatureDatesToRelease := true

	err := godotenv.Load(os.Getenv("ENV_PATH"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := au.NewClient(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))
	apis := ahautil.ClientAPIs{Client: client}

	featureMap := map[string]aha.Feature{}

	for _, releaseId := range releases {
		features := []aha.Feature{}
		var err error
		if updateDefaultFeatureDatesToRelease {
			features, err = apis.UpdateFeatureStartDueDatesToRelease(releaseId)
		} else {
			features, err = apis.GetFeaturesByRelease(releaseId)
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

	filename := "_features.json"

	err = ioutilmore.WriteJSON(filename, featureMap, 0644, true)
	if err != nil {
		panic(err)
	}

	fmt.Println("DONE")
}
