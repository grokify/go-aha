package ahautil

import (
	"context"
	"strings"

	"github.com/grokify/go-aha/client"
)

type FeatureSet struct {
	ClientAPIs   ClientAPIs
	FeatureMap   map[string]*aha.Feature
	TagFilterMap map[string]string
}

func NewFeatureSet() FeatureSet {
	return FeatureSet{
		FeatureMap:   map[string]*aha.Feature{},
		TagFilterMap: map[string]string{},
	}
}

func (fs *FeatureSet) TrimSpaceTagFilterMap() {
	newTags := map[string]string{}
	for tag, val := range fs.TagFilterMap {
		tag = strings.TrimSpace(tag)
		if len(tag) > 0 {
			newTags[tag] = val
		}
	}
	fs.TagFilterMap = newTags
}

func (fs *FeatureSet) LoadFeaturesForReleases(ctx context.Context, releaseIds []string) error {
	for _, releaseId := range releaseIds {
		err := fs.LoadFeaturesForRelease(ctx, releaseId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (fs *FeatureSet) LoadFeaturesForRelease(ctx context.Context, releaseId string) error {
	features, err := fs.ClientAPIs.GetFeaturesFullByReleaseId(ctx, releaseId)
	if err != nil {
		return err
	}
	lenTagFilters := len(fs.TagFilterMap)

	for _, feature := range features {
		if lenTagFilters == 0 {
			fs.FeatureMap[feature.Id] = feature
		} else {
		FilterTag:
			for filterTag, _ := range fs.TagFilterMap {
				for _, featureTag := range feature.Tags {
					featureTag = strings.TrimSpace(featureTag)
					if filterTag == featureTag {
						fs.FeatureMap[feature.Id] = feature
						break FilterTag
					}
				}
			}
		}
	}
	return nil
}
