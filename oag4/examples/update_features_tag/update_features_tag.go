package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/antihax/optional"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/net/http/httputilmore"
	"github.com/joho/godotenv"

	"github.com/grokify/go-aha/v3/oag4/aha"
	au "github.com/grokify/go-aha/v3/oag4/ahautil"
)

func main() {
	oldTag := "My Old Tag"
	newTag := "My New Tag"
	updateFeatureTag := false

	err := godotenv.Load(os.Getenv("ENV_PATH"))
	if err != nil {
		slog.Error("error loading .env file", "msg", err.Error())
		os.Exit(1)
	}

	apis, err := au.NewClientAPIs(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(2)
	}
	featuresApi := apis.APIClient.FeaturesApi
	ctx := context.Background()

	params := aha.GetFeaturesOpts{
		Tag:     optional.NewString(oldTag),
		PerPage: optional.NewInt32(int32(500))}
	/*
		   	Q              optional.String
		   	UpdatedSince   optional.Time
		   	Tag            optional.String
		   	AssignedToUser optional.String
		   	Page           optional.Int32
		   	PerPage        optional.Int32
		   }


		fsRes, resp, err := featuresApi.FeaturesGet(ctx, map[string]interface{}{
			"tag":      oldTag,
			"per_page": 500,
		})*/
	fsRes, resp, err := featuresApi.GetFeatures(ctx, &params)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(3)
	}

	if resp.StatusCode >= 300 {
		slog.Error("invalid status code", "status_code", resp.StatusCode)
		os.Exit(4)
	}

	fmtutil.MustPrintJSON(fsRes)
	err = httputilmore.PrintResponse(resp, true)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(4)
	}

	for i, fThin := range fsRes.Features {
		fmtutil.MustPrintJSON(fThin)

		fFull, resp, err := featuresApi.GetFeature(ctx, fThin.Id)
		if err != nil {
			slog.Error(err.Error(), "feature_index", i)
			os.Exit(4)
		} else if resp.StatusCode >= 300 {
			slog.Error("invalid status code", "status_code", resp.StatusCode, "feature_index", i)
			os.Exit(5)
		}

		fmtutil.MustPrintJSON(fFull)
		fmtutil.MustPrintJSON(fFull.Feature.Tags)

		if updateFeatureTag {
			newTags := []string{}
			hasOldTag := false
			for _, tag := range fFull.Feature.Tags {
				if tag == oldTag {
					newTags = append(newTags, newTag)
					hasOldTag = true
				} else {
					newTags = append(newTags, tag)
				}
			}
			if hasOldTag {
				fUpdate := aha.FeatureUpdate{Tags: strings.Join(newTags, ",")}
				updateRes, resp, err := featuresApi.UpdateFeature(ctx, fThin.Id, fUpdate)
				if err != nil {
					slog.Error(err.Error(), "feature_index", i)
					os.Exit(6)
				} else if resp.StatusCode >= 300 {
					slog.Error("invalid status code", "status_code", resp.StatusCode, "feature_index", i)
					os.Exit(7)
				}
				fmtutil.MustPrintJSON(updateRes)
			}
		}
	}
	slog.Info("found_features", "feature_count", len(fsRes.Features))

	fmt.Println("DONE")
	os.Exit(0)
}
