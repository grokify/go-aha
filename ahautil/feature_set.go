package ahautil

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/grokify/go-aha/aha"
	"github.com/pkg/errors"
)

type FeatureSet struct {
	ClientAPIs   ClientAPIs              `json:"-"`
	FeatureMap   map[string]*aha.Feature `json:"featureMap,omitempty"`
	TagFilterMap map[string]string       `json:"tagFilterMap,omitempty"`
}

func ReadFeatureSet(featuresPath string) (*FeatureSet, error) {
	featuresSet := NewFeatureSet()
	bytes, err := ioutil.ReadFile(featuresPath)
	if err != nil {
		return nil, err
	}
	return featuresSet, json.Unmarshal(bytes, featuresSet)
}

/*
func ReadFeatures(featuresPath string) (map[string]aha.Feature, error) {
	featuresMap := map[string]aha.Feature{}
	bytes, err := ioutil.ReadFile(featuresPath)
	if err != nil {
		return featuresMap, err
	}
	err = json.Unmarshal(bytes, &featuresMap)
	return featuresMap, err
}
*/

func NewFeatureSet() *FeatureSet {
	return &FeatureSet{
		FeatureMap:   map[string]*aha.Feature{},
		TagFilterMap: map[string]string{},
	}
}

// NewFeatureSetForReleasesIds returns a new FeatureSet given a list of
// ReleaseIds.
func NewFeatureSetForReleasesIds(clientAPIs ClientAPIs, ids []string) (*FeatureSet, error) {
	fs := NewFeatureSet()
	fs.ClientAPIs = clientAPIs
	err := fs.LoadFeaturesForReleases(context.Background(), ids)
	return fs, err
}

func (fs *FeatureSet) ReadFile(featuresPath string) error {
	if 1 == 1 {
		fs2, err := ReadFeatureSet(featuresPath)
		if err != nil {
			return errors.Wrap(err, "FeatureSet.ReadFile - ReadFeatureSet")
		}
		for id2, feat2 := range fs2.FeatureMap {
			fs.FeatureMap[id2] = feat2
		}
	}
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
		//if fs.LoadFeatureCheckTags(feature) {
		if FeatureTagMatch(feature, fs.TagFilterMap) {
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
		//if fs.LoadFeatureCheckTags(feature) {
		if FeatureTagMatch(feature, fs.TagFilterMap) {
			fs.FeatureMap[feature.Id] = feature
		}
	}
	return nil
}

// FeatureTagMatch takes an `*aha.Feature` and cmopares it to a
// tagFilterMap, returning `true` if there's a match and `false`
// if there isn't.
func FeatureTagMatch(feature *aha.Feature, tagFilterMap map[string]string) bool {
	lenTagFilters := len(tagFilterMap)
	if lenTagFilters == 0 {
		return true
	} else {
		for filterTag := range tagFilterMap {
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

// GetFeaturesMapByTag returns a map[string]map[string] where
// the first key is a filter tag and the second key is the
// feature id.
func (fs *FeatureSet) GetFeaturesMapByTag() (map[string]map[string]*aha.Feature, error) {
	fs.TrimSpaceTagFilterMap()
	featuresMap2 := map[string]map[string]*aha.Feature{}
	if len(fs.FeatureMap) == 0 {
		return featuresMap2, fmt.Errorf("ahautil.FeatureSet.GetFeaturesMapByTag - E_NO_FEATURES")
	}
	if len(fs.TagFilterMap) == 0 {
		return featuresMap2, fmt.Errorf("ahautil.FeatureSet.GetFeaturesMapByTag - E_NO_FILTER_TAGS")
	}

FEATS:
	for id, feat := range fs.FeatureMap {
		for _, tagTry := range feat.Tags {
			//fmt.Printf("FEAT [%v] TAG_TRY [%v]\n", feat.Name, tagTry)
			if _, ok := fs.TagFilterMap[tagTry]; ok {
				if _, ok2 := featuresMap2[tagTry]; !ok2 {
					featuresMap2[tagTry] = map[string]*aha.Feature{}
				}
				featuresMap2[tagTry][id] = feat
				continue FEATS
			}
		}
	}
	return featuresMap2, nil
}
