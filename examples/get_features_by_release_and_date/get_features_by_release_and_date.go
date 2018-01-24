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

func GetBeginData(feature aha.Feature) (time.Time, error) {
	if PossibleDate(feature.StartDate) {
		return time.Parse(tu.RFC3339YMD, feature.StartDate)
	} else if PossibleDate(feature.Release.StartDate) {
		return time.Parse(tu.RFC3339YMD, feature.Release.StartDate)
	}
	return time.Now(), fmt.Errorf("Date Not Found")
}

func GetEndDate(feature aha.Feature) (time.Time, error) {
	if PossibleDate(feature.DueDate) {
		return time.Parse(tu.RFC3339YMD, feature.DueDate)
	} else if PossibleDate(feature.Release.ReleaseDate) {
		return time.Parse(tu.RFC3339YMD, feature.Release.ReleaseDate)
	}
	return time.Now(), fmt.Errorf("Date Not Found")
}

func PossibleDate(dateString string) bool {
	dateString = strings.TrimSpace(dateString)
	if len(dateString) == 0 {
		return false
	} else if strings.Index(dateString, "0") == 0 {
		return false
	}
	return true
}

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
