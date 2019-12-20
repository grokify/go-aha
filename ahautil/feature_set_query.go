package ahautil

import (
	"context"
	"strings"

	"github.com/grokify/gotilla/fmt/fmtutil"
	log "github.com/sirupsen/logrus"
)

var Debug = false

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
	//fmtutil.PrintJSON(rs.ReleaseMap)
	//fmt.Printf("REL_COUNT [%v]\n", rs.ReleaseCount())
	//fmtutil.PrintJSON(rs.Ids())
	//fmtutil.PrintJSON(rs.RefNums())
	//panic("Z")
	if Debug {
		fmtutil.PrintJSON(rs.ReleaseMap)
		log.Infof("PRODUCTS_RELEASES_ALL [%s]\n",
			strings.Join(rs.RefNums(), ","))
		//panic("Z")
	}
	rs2, err := rs.GetReleasesForQuarters(yyyyq1, yyyyq2)
	if err != nil {
		return &rs2, nil, err
	}
	if Debug {
		log.Infof("PRODUCTS_RELEASES_ALL_QUARTERS [%s][%d][%d]\n",
			strings.Join(rs2.RefNums(), ","),
			yyyyq1, yyyyq2)
	}
	fs, err := NewFeatureSetForReleasesIds(clientAPIs, rs2.Ids())
	return &rs2, fs, err
}
