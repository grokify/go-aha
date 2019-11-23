package ahautil

import (
	"context"
	"strings"
)

// GetReleasesAndFeaturesForProductsAndQuarters returns releases and features
// given a product and date rate.
func GetReleasesAndFeaturesForProductsAndQuarters(
	ctx context.Context, clientAPIs ClientAPIs, products []string, yyyyq1, yyyyq2 int32) (*ReleaseSet, *FeatureSet, error) {
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
	rs2, err := rs.GetReleasesForQuarters(yyyyq1, yyyyq2)
	if err != nil {
		return &rs2, nil, err
	}
	fs, err := NewFeatureSetForReleasesIds(clientAPIs, rs2.Ids())
	return &rs2, fs, err
}
