package ahautil

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/antihax/optional"
	"github.com/grokify/go-aha/aha"
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

// ReleaseIds returns the list of Release Ids.
func (rs *ReleaseSet) Ids() []string {
	ids := []string{}
	for _, rel := range rs.ReleaseMap {
		rel.Id = strings.TrimSpace(rel.Id)
		if len(rel.Id) > 0 {
			ids = append(ids, rel.Id)
		}
	}
	sort.Strings(ids)
	return ids
}

// ReferenceNumbers returns the list of Release Reference Numbers.
func (rs *ReleaseSet) ReferenceNumbers() []string {
	refs := []string{}
	for _, rel := range rs.ReleaseMap {
		rel.ReferenceNum = strings.TrimSpace(rel.ReferenceNum)
		if len(rel.ReferenceNum) > 0 {
			refs = append(refs, rel.ReferenceNum)
		}
	}
	sort.Strings(refs)
	return refs
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

	params := &aha.GetProductReleasesOpts{
		Page:    optional.NewInt32(int32(1)),
		PerPage: optional.NewInt32(int32(500))}

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
			dta, err := time.Parse(tu.RFC3339FullDate, dtaString)
			if err != nil {
				return rels, err
			}
			if tu.IsGreaterThan(dta.UTC(), beg.UTC(), true) {
				gteDta = true
			}
		}
		if len(dtzString) > 0 {
			dtz, err := time.Parse(tu.RFC3339FullDate, dtzString)
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

// GetReleasesForQuarters retrives a filter of products within 2 quarters.
func (rs *ReleaseSet) GetReleasesForQuarters(yyyyq1, yyyyq2 int32) (ReleaseSet, error) {
	newSet := NewReleaseSet()
	if yyyyq1 == 0 && yyyyq2 == 0 {
		newSet.ReleaseMap = rs.ReleaseMap
		return newSet, nil
	} else if yyyyq2 == 0 {
		dtQ1, err := tu.QuarterInt32StartTime(yyyyq1)
		if err != nil {
			return newSet, err
		}
		for id, rel := range rs.ReleaseMap {
			relDt, err := time.Parse(tu.RFC3339FullDate, rel.ReleaseDate)
			if err != nil {
				continue
			}
			if relDt.After(dtQ1) || relDt.Equal(dtQ1) {
				newSet.ReleaseMap[id] = rel
			}
		}
	} else if yyyyq1 == 0 {
		dtQ2, err := tu.QuarterInt32StartTime(yyyyq2)
		if err != nil {
			return newSet, err
		}
		dtQ2Next := tu.NextQuarter(dtQ2)
		for id, rel := range rs.ReleaseMap {
			relDt, err := time.Parse(tu.RFC3339FullDate, rel.ReleaseDate)
			if err != nil {
				continue
			}
			if relDt.Before(dtQ2Next) {
				newSet.ReleaseMap[id] = rel
			}
		}
	} else {
		dtQ1, err := tu.QuarterInt32StartTime(yyyyq1)
		if err != nil {
			return newSet, err
		}
		dtQ2, err := tu.QuarterInt32StartTime(yyyyq2)
		if err != nil {
			return newSet, err
		}
		dtQ1, dtQ2 = tu.MinMax(dtQ1, dtQ2)
		dtQ2Next := tu.NextQuarter(dtQ2)
		for id, rel := range rs.ReleaseMap {
			relDt, err := time.Parse(tu.RFC3339FullDate, rel.ReleaseDate)
			if err != nil {
				continue
			}
			if tu.InRange(relDt, dtQ1, dtQ2Next, true, false) {
				newSet.ReleaseMap[id] = rel
			}
		}
	}
	return newSet, nil
}
