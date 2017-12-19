package ahautil

import (
	"context"
	"fmt"
	"strings"

	"github.com/grokify/go-aha"
	ao "github.com/grokify/oauth2util/aha"
)

type ClientAPIs struct {
	//Client    *http.Client
	APIClient *aha.APIClient
	Config    *aha.Configuration
}

func NewClientAPIs(account, apiKey string) ClientAPIs {
	clientAPIs := ClientAPIs{}

	cfg := aha.NewConfiguration()

	httpClient := ao.NewClient(account, apiKey)

	cfg.HTTPClient = httpClient

	clientAPIs.APIClient = aha.NewAPIClient(cfg)

	return clientAPIs
}

/*
func NewAPIClient(account, apiKey string) *aha.APIClient {
	client := ao.NewClient(account, apiKey)
	cfg := aha.NewConfiguration()
	cfg.HTTPClient = client
	return aha.NewAPIClient(cfg)
}
*/
/*
func (apis *ClientAPIs) FeaturesApi() *aha.FeaturesApi {
	api := apis.APIClient.FeaturesApi
	api.Configuration.Transport = apis.Client.Transport
	return api
}

func (apis *ClientAPIs) ReleasesApi() *aha.ReleasesApi {
	api := apis.APIClient.ReleasesApi
	api.Configuration.Transport = apis.Client.Transport
	return api
}
*/

/*
func (apis *ClientAPIs) ReleasesFeaturesApi() *aha.ReleasesFeaturesApi {
	api := aha.NewReleasesFeaturesApi()
	api.Configuration.Transport = apis.Client.Transport
	return api
}
*/

func (apis *ClientAPIs) GetReleaseById(releaseId string) (*aha.Release, error) {
	//rel := &aha.Release{}
	api := apis.APIClient.ReleasesApi
	rinfo, resp, err := api.ReleasesReleaseIdGet(context.Background(), releaseId)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("Bad response: %v", resp.StatusCode)
	}
	return rinfo.Release, nil
}

func (apis *ClientAPIs) GetFeatureById(featureId string) (*aha.Feature, error) {
	//feat := &aha.Feature{}
	//fapi := apis.FeaturesApi()
	fapi := apis.APIClient.FeaturesApi
	fthick, res, err := fapi.FeaturesFeatureIdGet(context.Background(), featureId)
	if err != nil {
		return nil, err
	}
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Bad response: %v", res.StatusCode)
	}
	return fthick.Feature, nil
}

func (apis *ClientAPIs) GetFeaturesMetaByRelease(releaseId string) ([]aha.FeatureMeta, error) {
	features := []aha.FeatureMeta{}
	//frapi := apis.ReleasesFeaturesApi()
	fapi := apis.APIClient.FeaturesApi
	finfo, resp, err := fapi.ReleasesReleaseIdFeaturesGet(context.Background(), releaseId)
	if err != nil {
		return features, err
	}
	if resp.StatusCode > 299 {
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

func (apis *ClientAPIs) UpdateFeatureStartDueDatesToRelease(releaseId string) ([]*aha.Feature, error) {
	features := []*aha.Feature{}

	rel, err := apis.GetReleaseById(releaseId)
	if err != nil {
		return features, err
	}

	featureMetas, err := apis.GetFeaturesMetaByRelease(releaseId)
	if err != nil {
		return features, err
	}

	fapi := apis.APIClient.FeaturesApi

	for _, fthin := range featureMetas {
		fthick, res, err := fapi.FeaturesFeatureIdGet(context.Background(), fthin.Id)
		if err != nil {
			return features, err
		}
		if res.StatusCode > 299 {
			return features, fmt.Errorf("CODE %v", res.StatusCode)
		}

		feat := fthick.Feature

		featureId := fthick.Feature.Id
		//if fthick.Feature.StartDate.Unix() == tu.RFC3339YMDZeroUnix ||
		//	fthick.Feature.DueDate.Unix() == tu.RFC3339YMDZeroUnix {
		if ymdIsNotSet(fthick.Feature.StartDate) || ymdIsNotSet(fthick.Feature.DueDate) {

			featureUpdate := aha.FeatureUpdate{}

			if ymdIsNotSet(fthick.Feature.StartDate) {
				featureUpdate.StartDate = rel.StartDate
			}
			if ymdIsNotSet(fthick.Feature.DueDate) {
				featureUpdate.DueDate = rel.ReleaseDate
			}

			// Update Feature with Dates
			_, res, err := fapi.FeaturesFeatureIdPut(context.Background(), featureId, featureUpdate)
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

func (apis *ClientAPIs) GetFeaturesByRelease(releaseId string) ([]*aha.Feature, error) {
	features := []*aha.Feature{}

	featureMetas, err := apis.GetFeaturesMetaByRelease(releaseId)
	if err != nil {
		return features, err
	}

	fapi := apis.APIClient.FeaturesApi

	for _, fthin := range featureMetas {
		fthick, res, err := fapi.FeaturesFeatureIdGet(context.Background(), fthin.Id)
		if err != nil {
			return features, err
		}
		if res.StatusCode > 299 {
			return features, fmt.Errorf("API Bad Status: %v", res.StatusCode)
		}
		features = append(features, fthick.Feature)
	}
	return features, nil
}
