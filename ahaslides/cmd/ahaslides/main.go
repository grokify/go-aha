// Go example that covers:
// Quickstart: https://developers.google.com/slides/quickstart/go
// Basic writing: adding a text box to slide: https://developers.google.com/slides/samples/writing
// Using SDK: https://github.com/google/google-api-go-client/blob/master/slides/v1/slides-gen.go
// Creating and Managing Presentations https://developers.google.com/slides/how-tos/presentations
// Adding Shapes and Text to a Slide: https://developers.google.com/slides/how-tos/add-shape#example
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/grokify/go-aha/v2/ahaslides"
	au "github.com/grokify/go-aha/v2/ahautil"
	"github.com/grokify/goauth/google"
	"github.com/grokify/googleutil/slidesutil/v1"
	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/log/logutil"
	flags "github.com/jessevdk/go-flags"
)

type Options struct {
	EnvFile     string `short:"e" long:"env" description:"Env filepath"`
	NewTokenRaw []bool `short:"n" long:"newtoken" description:"Retrieve new token"`
}

func (opt *Options) NewToken() bool {
	return len(opt.NewTokenRaw) > 0
}

// Post https://slides.googleapis.com/v1/presentations?alt=json: oauth2: token expired and refresh token is not set
func main() {
	opts := Options{}
	_, err := flags.Parse(&opts)
	logutil.FatalErr(err)

	_, err = config.LoadDotEnv([]string{opts.EnvFile, os.Getenv("ENV_PATH")}, -1)
	logutil.FatalErr(err)

	roadmapConfig, err := ahaslides.NewRoadmapConfigEnv()
	logutil.FatalErr(err)

	if 1 == 0 {
		fmtutil.MustPrintJSON(roadmapConfig)
		panic("Z")
	}

	/*
		googleClient, err := omg.NewClientFileStoreWithDefaults(
			[]byte(os.Getenv(omg.EnvGoogleAppCredentials)),
			[]string{omg.ScopeDrive, omg.ScopePresentations},
			opts.NewToken())
		if err != nil {
			log.Fatal(errorsutil.Wrap(err, "NewClientFileStoreWithDefaults"))
		}*/

	googHTTPClient, err := google.NewClientFileStoreWithDefaultsCliEnv("", "")
	logutil.FatalErr(err)

	t := time.Now().UTC()
	slideName := fmt.Sprintf("%s (Aha!) %s\n", roadmapConfig.Title, t.Format(time.RFC3339))
	fmt.Printf("Slide Name: %s", slideName)

	fmt.Printf("FEATURES_FILEPATH [%v]\n", roadmapConfig.FeaturesFilepath)
	featureSet := au.NewFeatureSet()
	featureSet.TagFilterMap = roadmapConfig.FilterTagsMap
	err = featureSet.ReadFile(roadmapConfig.FeaturesFilepath)
	logutil.FatalErr(err)

	pc, err := slidesutil.NewPresentationCreator(googHTTPClient)
	logutil.FatalErr(err)

	presID, err := pc.CreateEmpty(slideName)
	logutil.FatalErr(err)

	res, err := ahaslides.CreateRoadmapSlide(googHTTPClient, presID, roadmapConfig, featureSet)
	logutil.FatalErr(err)

	roadmapConfig.RoadmapFormatting.Textbox.DoneBackgroundColorHex =
		roadmapConfig.RoadmapFormatting.Textbox.DefaultBackgroundColorHex
	roadmapConfig.RoadmapFormatting.Textbox.ProblemBackgroundColorHex =
		roadmapConfig.RoadmapFormatting.Textbox.DefaultBackgroundColorHex

	res, err = ahaslides.CreateRoadmapSlide(googHTTPClient, presID, roadmapConfig, featureSet)
	logutil.FatalErr(err)

	fmt.Printf("Created PresentationId [%v]\n", res.PresentationId)

	fmt.Println("DONE")
}
