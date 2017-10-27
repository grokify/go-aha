package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/grokify/go-aha"
	"github.com/grokify/gotilla/fmt/fmtutil"
	ou "github.com/grokify/oauth2util"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ahaServerURL := "https://mysubdomain.aha.io/api/v1"
	if len(strings.TrimSpace(os.Getenv("AHA_API_BASE_URL"))) > 0 {
		ahaServerURL = os.Getenv("AHA_API_BASE_URL")
	}

	apiKey := os.Getenv("AHA_API_KEY")

	client := ou.NewClientAccessToken(apiKey)

	api := aha.NewFeaturesApiWithBasePath(ahaServerURL)
	api.Configuration.Transport = client.Transport

	info, resp, err := api.FeaturesGet("")
	if err != nil {
		log.Fatal("Error retrieving features")
	}

	fmt.Println(resp.StatusCode)
	fmtutil.PrintJSON(info)

	fmt.Println("DONE")
}
