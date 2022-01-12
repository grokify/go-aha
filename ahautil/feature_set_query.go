package ahautil

import (
	"context"
	"log"
	"strings"

	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/time/timeutil"
)

var Debug = false

// GetReleasesAndFeaturesForProductsAndQuarters returns releases and features
// given a product and date rate.
func GetReleasesAndFeaturesForProductsAndQuarters(
	ctx context.Context, clientAPIs ClientAPIs, products []string, yyyyq1, yyyyq2 int32) (*ReleaseSet, *FeatureSet, error) {
	if yyyyq1 < 100 || yyyyq2 < 100 {
		yyyyq1, yyyyq2 = timeutil.QuartersInt32RelToAbs(yyyyq1, yyyyq2)
	}
	rs := NewReleaseSet()
	rs.ClientAPIs = clientAPIs
	for _, prod := range products {
		prod = strings.ToUpper(strings.TrimSpace(prod))
		if len(prod) == 0 {
			continue
		}
		err := rs.LoadReleasesForProduct(ctx, prod)
		if err != nil {
			return &rs, nil, err
		}
	}

	if Debug {
		fmtutil.PrintJSON(rs.ReleaseMap)
		log.Printf("PRODUCTS_RELEASES_ALL [%s]\n",
			strings.Join(rs.RefNums(), ","))
	}
	rs2, err := rs.GetReleasesForQuarters(yyyyq1, yyyyq2)
	if err != nil {
		return &rs2, nil, err
	}
	if Debug {
		log.Printf("PRODUCTS_RELEASES_ALL_QUARTERS [%s][%d][%d]\n",
			strings.Join(rs2.RefNums(), ","),
			yyyyq1, yyyyq2)
	}
	fs, err := NewFeatureSetForReleasesIds(clientAPIs, rs2.Ids())
	return &rs2, fs, err
}
