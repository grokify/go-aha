package ahautil

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/antihax/optional"
	"github.com/grokify/mogo/errors/errorsutil"
	tu "github.com/grokify/mogo/time/timeutil"

	"github.com/grokify/go-aha/v3/oagv4/aha"
)

type ReleaseSet struct {
	ClientAPIs ClientAPIs
	ReleaseMap map[string]aha.Release `json:"releaseMap"`
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

func (rs *ReleaseSet) ReleaseCount() int {
	return len(rs.ReleaseMap)
}

// ReleaseIds returns the list of Release Ids.
func (rs *ReleaseSet) RefNums() []string {
	ids := []string{}
	for _, rel := range rs.ReleaseMap {
		rel.ReferenceNum = strings.TrimSpace(rel.ReferenceNum)
		if len(rel.ReferenceNum) > 0 {
			ids = append(ids, rel.ReferenceNum)
		}
	}
	sort.Strings(ids)
	return ids
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
		return newErrorBadStatusCode(resp.StatusCode)
	}
	//fmtutil.PrintJSON(info)
	//panic("Z")
	for _, rel := range info.Releases {
		rs.ReleaseMap[rel.Id] = rel
		if 1 == 0 { // not needed
			relMore, resp, err := api.GetRelease(context.Background(), rel.Id)
			if err != nil {
				return err
			} else if resp.StatusCode >= 300 {
				return errorsutil.Wrap(err, fmt.Sprintf("Aha.GetRelease.StatusCode [%v]", resp.StatusCode))
			}
			rs.ReleaseMap[rel.Id] = relMore.Release
		}
	}
	//fmtutil.PrintJSON(rs.ReleaseMap)
	//panic("Z")
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
func (rs *ReleaseSet) GetReleasesForQuarters(yyyyq1, yyyyq2 int) (ReleaseSet, error) {
	newSet := NewReleaseSet()
	if yyyyq1 == 0 && yyyyq2 == 0 {
		newSet.ReleaseMap = rs.ReleaseMap
		return newSet, nil
	} else if yyyyq2 == 0 {
		dtQ1, _, err := tu.ParseQuarterInt32StartEndTimes(yyyyq1)
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
		dtQ2, _, err := tu.ParseQuarterInt32StartEndTimes(yyyyq2)
		if err != nil {
			return newSet, err
		}
		dtQ2Next := tu.QuarterAdd(dtQ2, 1)
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
		dtQ1, _, err := tu.ParseQuarterInt32StartEndTimes(int(yyyyq1))
		if err != nil {
			return newSet, err
		}
		dtQ2, _, err := tu.ParseQuarterInt32StartEndTimes(int(yyyyq2))
		if err != nil {
			return newSet, err
		}
		dtQ1, dtQ2 = tu.MinMax(dtQ1, dtQ2)
		dtQ2Next := tu.QuarterAdd(dtQ2, 1)
		for id, rel := range rs.ReleaseMap {
			rel.ReleaseDate = strings.TrimSpace(rel.ReleaseDate)
			if len(rel.ReleaseDate) == 0 {
				continue
			}
			releaseDate, err := time.Parse(tu.RFC3339FullDate, rel.ReleaseDate)
			if err != nil {
				continue
			}
			if Debug {
				fmt.Printf("REL_RELEASE_DAT [%v]\n", releaseDate.Format(time.RFC3339))
				fmt.Printf("FILTER_DATE_BEG [%v]\n", dtQ1.Format(time.RFC3339))
				fmt.Printf("FILTER_DATE_END [%v]\n", dtQ2Next.Format(time.RFC3339))
				inRange := tu.InRange(dtQ1, dtQ2Next, releaseDate, true, false)
				fmt.Printf("IN_RANGE [%v]\n", inRange)
			}

			if tu.InRange(dtQ1, dtQ2Next, releaseDate, true, false) {
				newSet.ReleaseMap[id] = rel
			}
		}
	}
	return newSet, nil
}
