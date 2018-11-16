package ahautil

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/grokify/go-aha/client"
)

type FeatureSet struct {
	ClientAPIs   ClientAPIs
	FeatureMap   map[string]*aha.Feature
	TagFilterMap map[string]string
}

func NewFeatureSet() *FeatureSet {
	return &FeatureSet{
		FeatureMap:   map[string]*aha.Feature{},
		TagFilterMap: map[string]string{},
	}
}

func (fs *FeatureSet) ReadFile(featuresPath string) error {
	fs.TrimSpaceTagFilterMap()
	featuresMap := map[string]*aha.Feature{}
	bytes, err := ioutil.ReadFile(featuresPath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &featuresMap)
	if err != nil {
		return err
	}
	for _, feature := range featuresMap {
		if fs.LoadFeatureCheckTags(feature) {
			fs.FeatureMap[feature.Id] = feature
		}
	}
	return nil
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

	for _, feature := range features {
		if fs.LoadFeatureCheckTags(feature) {
			fs.FeatureMap[feature.Id] = feature
		}
	}
	return nil
}

func (fs *FeatureSet) LoadFeatureCheckTags(feature *aha.Feature) bool {
	lenTagFilters := len(fs.TagFilterMap)
	if lenTagFilters == 0 {
		return true
	} else {
		for filterTag := range fs.TagFilterMap {
			for _, featureTag := range feature.Tags {
				featureTag = strings.TrimSpace(featureTag)
				if filterTag == featureTag {
					return true
				}
			}
		}
	}
	return false
}

func (fs *FeatureSet) GetFeaturesMapByTag() map[string]map[string]*aha.Feature {
	fs.TrimSpaceTagFilterMap()
	featuresMap2 := map[string]map[string]*aha.Feature{}

FEATS:
	for id, feat := range fs.FeatureMap {
		for _, tagTry := range feat.Tags {
			if _, ok := fs.TagFilterMap[tagTry]; ok {
				if _, ok2 := featuresMap2[tagTry]; !ok2 {
					featuresMap2[tagTry] = map[string]*aha.Feature{}
				}
				featuresMap2[tagTry][id] = feat
				continue FEATS
			}
		}
	}
	return featuresMap2
}
