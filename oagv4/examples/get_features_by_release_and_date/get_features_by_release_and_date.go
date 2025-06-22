package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/joho/godotenv"

	au "github.com/grokify/go-aha/v3/oagv4/ahautil"
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

	apis, err := au.NewClientAPIs(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	featureSet := au.NewFeatureSet()
	featureSet.ClientAPIs = apis

	err = featureSet.LoadFeaturesForReleases(context.Background(), releases)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("---")

	if 1 == 1 {
		fmtutil.MustPrintJSON(featureSet)
	}

	if 1 == 1 {
		bytes, err := json.MarshalIndent(featureSet, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("LEN %v\n", len(featureSet.FeatureMap))

		err = os.WriteFile(outputFile, bytes, 0644)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Wrote [%v]\n", outputFile)
	}
	fmt.Println("DONE")
}
