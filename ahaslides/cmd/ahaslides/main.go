// Go example that covers:
// Quickstart: https://developers.google.com/slides/quickstart/go
// Basic writing: adding a text box to slide: https://developers.google.com/slides/samples/writing
// Using SDK: https://github.com/google/google-api-go-client/blob/master/slides/v1/slides-gen.go
// Creating and Managing Presentations https://developers.google.com/slides/how-tos/presentations
// Adding Shapes and Text to a Slide: https://developers.google.com/slides/how-tos/add-shape#example
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/grokify/go-aha/ahaslides"
	au "github.com/grokify/go-aha/ahautil"
	"github.com/grokify/googleutil/slidesutil/v1"
	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/oauth2more/google"
	flags "github.com/jessevdk/go-flags"
)

type Options struct {
	EnvFile     string `short:"e" long:"env" description:"Env filepath"`
	NewTokenRaw []bool `short:"n" long:"newtoken" description:"Retrieve new token"`
}

func (opt *Options) NewToken() bool {
	if len(opt.NewTokenRaw) > 0 {
		return true
	}
	return false
}

// Post https://slides.googleapis.com/v1/presentations?alt=json: oauth2: token expired and refresh token is not set
func main() {
	opts := Options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}
	err = config.LoadDotEnvFirst(opts.EnvFile, os.Getenv("ENV_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	roadmapConfig, err := ahaslides.NewRoadmapConfigEnv()
	if err != nil {
		log.Fatal(err)
	}
	if 1 == 0 {
		fmtutil.PrintJSON(roadmapConfig)
		panic("Z")
	}

	/*
		googleClient, err := omg.NewClientFileStoreWithDefaults(
			[]byte(os.Getenv(omg.EnvGoogleAppCredentials)),
			[]string{omg.ScopeDrive, omg.ScopePresentations},
			opts.NewToken())
		if err != nil {
			log.Fatal(errors.Wrap(err, "NewClientFileStoreWithDefaults"))
		}*/

	googHTTPClient, err := google.NewClientFileStoreWithDefaultsCliEnv("", "")
	if err != nil {
		log.Fatal(err)
	}

	t := time.Now().UTC()
	slideName := fmt.Sprintf("%s (Aha!) %s\n", roadmapConfig.Title, t.Format(time.RFC3339))
	fmt.Printf("Slide Name: %s", slideName)

	fmt.Printf("FEATURES_FILEPATH [%v]\n", roadmapConfig.FeaturesFilepath)
	featureSet := au.NewFeatureSet()
	featureSet.TagFilterMap = roadmapConfig.FilterTagsMap
	featureSet.ReadFile(roadmapConfig.FeaturesFilepath)

	pc, err := slidesutil.NewPresentationCreator(googHTTPClient)
	if err != nil {
		log.Fatal(err)
	}
	presID, err := pc.CreateEmpty(slideName)
	if err != nil {
		log.Fatal(err)
	}

	res, err := ahaslides.CreateRoadmapSlide(googHTTPClient, presID, roadmapConfig, featureSet)
	if err != nil {
		log.Fatal(err)
	}

	roadmapConfig.RoadmapFormatting.Textbox.DoneBackgroundColorHex =
		roadmapConfig.RoadmapFormatting.Textbox.DefaultBackgroundColorHex
	roadmapConfig.RoadmapFormatting.Textbox.ProblemBackgroundColorHex =
		roadmapConfig.RoadmapFormatting.Textbox.DefaultBackgroundColorHex

	res, err = ahaslides.CreateRoadmapSlide(googHTTPClient, presID, roadmapConfig, featureSet)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created PresentationId [%v]\n", res.PresentationId)

	fmt.Println("DONE")
}
