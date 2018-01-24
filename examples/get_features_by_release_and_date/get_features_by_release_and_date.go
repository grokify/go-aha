package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/joho/godotenv"

	au "github.com/grokify/go-aha/ahautil"
	"github.com/grokify/go-aha/client"
	tu "github.com/grokify/gotilla/time/timeutil"
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
	fmtutil.PrintJSON(featureSet.FeatureMap)
	fmt.Printf("LEN %v\n", len(featureSet.FeatureMap))
	fmt.Println("DONE")
}
