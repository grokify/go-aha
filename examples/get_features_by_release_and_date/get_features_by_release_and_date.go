package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	au "github.com/grokify/go-aha/ahautil"
	//"github.com/grokify/go-aha/client"
)

var (
	outputFile                         = "_features.json"
	updateDefaultFeatureDatesToRelease = false
)

func main() {
	releases := []string{"API-R-1", "API-R-2"}

	err := godotenv.Load(os.Getenv("ENV_PATH"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apis := au.NewClientAPIs(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))

	featureSet := au.NewFeatureSet()
	featureSet.ClientAPIs = apis

	err = featureSet.LoadFeaturesForReleases(context.Background(), releases)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("---")

	if 1 == 1 {
		fmtutil.PrintJSON(featureSet)
	}

	if 1 == 1 {
		bytes, err := json.MarshalIndent(featureSet, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("LEN %v\n", len(featureSet.FeatureMap))

		err = ioutil.WriteFile(outputFile, bytes, 0644)
		if err != nil {
			log.Fatal(err)
		}
		log.Info(fmt.Sprintf("Wrote [%v]\n", outputFile))
	}
	fmt.Println("DONE")
}
