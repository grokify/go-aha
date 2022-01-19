package ahautil

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/antihax/optional"
	"github.com/grokify/go-aha/v2/aha"
	ao "github.com/grokify/goauth/aha"
)

const (
	FeatureStatusShippedLc          string = "shipped"
	FeatureStatusWillNotImplementLc string = "will not implement"
)

type ClientAPIs struct {
	APIClient *aha.APIClient
	Config    *aha.Configuration
}

func NewClientAPIsHTTPClient(httpClient *http.Client) ClientAPIs {
	cfg := aha.NewConfiguration()
	cfg.HTTPClient = httpClient
	return ClientAPIs{
		Config:    cfg,
		APIClient: aha.NewAPIClient(cfg)}
}

func NewClientAPIs(account, apiKey string) ClientAPIs {
	httpClient := ao.NewClient(account, apiKey)
	return NewClientAPIsHTTPClient(httpClient)
}

func (apis *ClientAPIs) GetReleaseById(releaseId string) (*aha.Release, error) {
	api := apis.APIClient.ReleasesApi
	rinfo, resp, err := api.GetRelease(context.Background(), releaseId)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("Bad response: %v", resp.StatusCode)
	}
	return &rinfo.Release, nil
}

func (apis *ClientAPIs) GetFeatureById(featureId string) (*aha.Feature, error) {
	fapi := apis.APIClient.FeaturesApi
	fthick, res, err := fapi.GetFeature(context.Background(), featureId)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 300 {
		return nil, fmt.Errorf("Bad response: %v", res.StatusCode)
	}
	return &fthick.Feature, nil
}

func (apis *ClientAPIs) GetFeaturesMetaByReleaseId(ctx context.Context, releaseId string) ([]aha.FeatureMeta, error) {
	features := []aha.FeatureMeta{}
	fapi := apis.APIClient.FeaturesApi

	params := &aha.GetReleaseFeaturesOpts{
		Page:    optional.NewInt32(int32(1)),
		PerPage: optional.NewInt32(int32(500))}

	finfo, resp, err := fapi.GetReleaseFeatures(ctx, releaseId, params)
	if err != nil {
		return features, err
	}
	if resp.StatusCode >= 300 {
		return features, fmt.Errorf("API Bad Status: %v", resp.StatusCode)
	}
	return finfo.Features, nil
}

func ymdIsNotSet(ymd string) bool {
	if len(ymd) == 0 {
		return true
	} else if strings.Index(ymd, "0") == 0 {
		return true
	}
	return false
}

func (apis *ClientAPIs) UpdateFeatureStartDueDatesToRelease(ctx context.Context, releaseId string) ([]*aha.Feature, error) {
	features := []*aha.Feature{}

	rel, err := apis.GetReleaseById(releaseId)
	if err != nil {
		return features, err
	}

	featureMetas, err := apis.GetFeaturesMetaByReleaseId(ctx, releaseId)
	if err != nil {
		return features, err
	}

	fapi := apis.APIClient.FeaturesApi

	for _, fthin := range featureMetas {
		fthick, res, err := fapi.GetFeature(ctx, fthin.Id)
		if err != nil {
			return features, err
		}
		if res.StatusCode > 299 {
			return features, fmt.Errorf("CODE %v", res.StatusCode)
		}

		feat := &fthick.Feature

		featureId := fthick.Feature.Id
		if ymdIsNotSet(fthick.Feature.StartDate) || ymdIsNotSet(fthick.Feature.DueDate) {

			featureUpdate := aha.FeatureUpdate{}

			if ymdIsNotSet(fthick.Feature.StartDate) {
				featureUpdate.StartDate = rel.StartDate
			}
			if ymdIsNotSet(fthick.Feature.DueDate) {
				featureUpdate.DueDate = rel.ReleaseDate
			}

			// Update Feature with Dates
			_, res, err := fapi.UpdateFeature(context.Background(), featureId, featureUpdate)
			if err != nil {
				return features, err
			}
			if res.StatusCode > 299 {
				return features, fmt.Errorf("Bad Status Code %v", res.StatusCode)
			}

			feat, err = apis.GetFeatureById(featureId)
			if err != nil {
				return features, err
			}
		}
		features = append(features, feat)
	}
	return features, nil
}

func (apis *ClientAPIs) GetFeaturesFullByReleaseId(ctx context.Context, releaseId string) ([]*aha.Feature, error) {
	features := []*aha.Feature{}

	featureMetas, err := apis.GetFeaturesMetaByReleaseId(ctx, releaseId)
	if err != nil {
		return features, err
	}

	fapi := apis.APIClient.FeaturesApi

	for _, fthin := range featureMetas {
		fthick, res, err := fapi.GetFeature(ctx, fthin.Id)
		if err != nil {
			return features, err
		}
		if res.StatusCode > 299 {
			return features, fmt.Errorf("API Bad Status: %v", res.StatusCode)
		}
		features = append(features, &fthick.Feature)
	}
	return features, nil
}
