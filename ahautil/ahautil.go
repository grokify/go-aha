package ahautil

import (
	"fmt"
	"net/http"

	"github.com/grokify/go-aha"
	tu "github.com/grokify/gotilla/time/timeutil"
)

type ClientAPIs struct {
	Client *http.Client
}

func (apis *ClientAPIs) FeaturesApi() *aha.FeaturesApi {
	api := aha.NewFeaturesApi()
	api.Configuration.Transport = apis.Client.Transport
	return api
}

func (apis *ClientAPIs) ReleasesApi() *aha.ReleasesApi {
	api := aha.NewReleasesApi()
	api.Configuration.Transport = apis.Client.Transport
	return api
}

func (apis *ClientAPIs) ReleasesFeaturesApi() *aha.ReleasesFeaturesApi {
	api := aha.NewReleasesFeaturesApi()
	api.Configuration.Transport = apis.Client.Transport
	return api
}

func (apis *ClientAPIs) GetReleaseById(releaseId string) (aha.Release, error) {
	rel := aha.Release{}
	api := apis.ReleasesApi()
	rinfo, resp, err := api.ReleasesReleaseIdGet(releaseId)
	if err != nil {
		return rel, err
	}
	if resp.StatusCode > 299 {
		return rel, fmt.Errorf("Bad response: %v", resp.StatusCode)
	}
	return rinfo.Release, nil
}

func (apis *ClientAPIs) GetFeatureById(featureId string) (aha.Feature, error) {
	feat := aha.Feature{}
	fapi := apis.FeaturesApi()
	fthick, res, err := fapi.FeaturesFeatureIdGet(featureId)
	if err != nil {
		return feat, err
	}
	if res.Response.StatusCode > 299 {
		return feat, fmt.Errorf("Bad response: %v", res.Response.StatusCode)
	}
	return fthick.Feature, nil
}

func (apis *ClientAPIs) GetFeaturesMetaByRelease(releaseId string) ([]aha.FeatureMeta, error) {
	features := []aha.FeatureMeta{}
	frapi := apis.ReleasesFeaturesApi()
	finfo, resp, err := frapi.ReleasesReleaseIdFeaturesGet(releaseId)
	if err != nil {
		return features, err
	}
	if resp.StatusCode > 299 {
		return features, fmt.Errorf("API Bad Status: %v", resp.StatusCode)
	}
	return finfo.Features, nil
}

func (apis *ClientAPIs) UpdateFeatureStartDueDatesToRelease(releaseId string) ([]aha.Feature, error) {
	features := []aha.Feature{}

	rel, err := apis.GetReleaseById(releaseId)
	if err != nil {
		return features, err
	}

	featureMetas, err := apis.GetFeaturesMetaByRelease(releaseId)
	if err != nil {
		return features, err
	}

	fapi := apis.FeaturesApi()

	for _, fthin := range featureMetas {
		fthick, res, err := fapi.FeaturesFeatureIdGet(fthin.Id)
		if err != nil {
			return features, err
		}
		if res.Response.StatusCode > 299 {
			return features, fmt.Errorf("CODE %v", res.Response.StatusCode)
		}

		feat := fthick.Feature

		featureId := fthick.Feature.Id
		if fthick.Feature.StartDate.Unix() == tu.RFC3339YMDZeroUnix ||
			fthick.Feature.DueDate.Unix() == tu.RFC3339YMDZeroUnix {

			featureUpdate := aha.FeatureUpdate{}

			if fthick.Feature.StartDate.Unix() == tu.RFC3339YMDZeroUnix {
				featureUpdate.StartDate = rel.StartDate
			}
			if fthick.Feature.DueDate.Unix() == tu.RFC3339YMDZeroUnix {
				featureUpdate.DueDate = rel.ReleaseDate
			}

			// Update Feature with Dates
			_, res, err := fapi.FeaturesFeatureIdPut(featureId, featureUpdate)
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

func (apis *ClientAPIs) GetFeaturesByRelease(releaseId string) ([]aha.Feature, error) {
	features := []aha.Feature{}

	featureMetas, err := apis.GetFeaturesMetaByRelease(releaseId)
	if err != nil {
		return features, err
	}

	fapi := apis.FeaturesApi()

	for _, fthin := range featureMetas {
		fthick, res, err := fapi.FeaturesFeatureIdGet(fthin.Id)
		if err != nil {
			return features, err
		}
		if res.Response.StatusCode > 299 {
			return features, fmt.Errorf("API Bad Status: %v", res.Response.StatusCode)
		}
		features = append(features, fthick.Feature)
	}
	return features, nil
}
