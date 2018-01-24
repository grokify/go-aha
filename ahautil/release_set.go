package ahautil

import (
	"context"
	"fmt"
	"time"

	"github.com/grokify/go-aha/client"
	tu "github.com/grokify/gotilla/time/timeutil"
)

type ReleaseSet struct {
	ClientAPIs ClientAPIs
	ReleaseMap map[string]aha.Release
}

func NewReleaseSet() ReleaseSet {
	return ReleaseSet{
		ReleaseMap: map[string]aha.Release{},
	}
}

func (rs *ReleaseSet) AddReleases(releases []aha.Release) {
	for _, rel := range releases {
		rs.ReleaseMap[rel.Id] = rel
	}
}

func (rs *ReleaseSet) LoadReleasesForProducts(ctx context.Context, productSlugs []string) error {
	for _, productSlug := range productSlugs {
		err := rs.LoadReleasesForProduct(ctx, productSlug)
		if err != nil {
			return err
		}
	}
	return nil
}

func (rs *ReleaseSet) LoadReleasesForProduct(ctx context.Context, productSlug string) error {
	api := rs.ClientAPIs.APIClient.ReleasesApi

	params := map[string]interface{}{
		"page":     int32(1),
		"pagePage": int32(500),
	}

	info, resp, err := api.GetProductReleases(ctx, productSlug, params)

	if err != nil {
		return err
	} else if resp.StatusCode >= 300 {
		return fmt.Errorf("Aha! API Error Status Code %v", resp.StatusCode)
	}

	for _, rel := range info.Releases {
		rs.ReleaseMap[rel.Id] = rel
	}

	return nil
}

func (rs *ReleaseSet) GetReleasesForDates(ctx context.Context, beg time.Time, end time.Time) ([]aha.Release, error) {
	rels := []aha.Release{}

	for _, rel := range rs.ReleaseMap {
		dtaString := rel.StartDate
		dtzString := rel.ReleaseDate
		gteDta := false
		lteDtz := false
		if len(dtaString) > 0 {
			dta, err := time.Parse(tu.RFC3339YMD, dtaString)
			if err != nil {
				return rels, err
			}
			if tu.IsGreaterThan(dta.UTC(), beg.UTC(), true) {
				gteDta = true
			}
		}
		if len(dtzString) > 0 {
			dtz, err := time.Parse(tu.RFC3339YMD, dtzString)
			if err != nil {
				return rels, err
			}
			if tu.IsLessThan(dtz.UTC(), end.UTC(), true) {
				lteDtz = true
			}
		}
		if gteDta && lteDtz {
			rels = append(rels, rel)
		}
	}
	return rels, nil
}
