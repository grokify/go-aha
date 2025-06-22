package featuresort

import (
	"sort"
	"strings"

	"github.com/grokify/go-aha/v3/oagv4/aha"
)

/*
func featureEndDate(feat *aha.Feature) time.Time {
	return timeutil.ParseFirstValueOrZero(timeutil.RFC3339FullDate,
		[]string{feat.DueDate, feat.Release.ReleaseDate})
}

func featureEndDateStringDisplay(feat *aha.Feature) string {
	dt := timeutil.ParseFirstValueOrZero(timeutil.RFC3339FullDate,
		[]string{feat.DueDate, feat.Release.ReleaseDate})
	if timeutil.IsZeroAny(dt) {
		return "0/0"
	}
	return dt.Format(timeutil.MonthDay)
}*/

func featureEndDateOrLastSortString(feat *aha.Feature) string {
	fDate := strings.TrimSpace(feat.DueDate)
	rDate := strings.TrimSpace(feat.Release.ReleaseDate)
	if len(fDate) > 0 {
		return fDate
	}
	if len(rDate) > 0 {
		return rDate
	}
	return "9999-99-99"
}

func SortFeatures(features []*aha.Feature) []*aha.Feature {
	items := []SortItem{}
	itemsMap := map[string]*aha.Feature{}
	for _, feat := range features {
		itemsMap[feat.Id] = feat
		items = append(items, SortItem{
			Id:    feat.Id,
			Value: featureEndDateOrLastSortString(feat),
		})
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].Value < items[j].Value
	})
	sorted := []*aha.Feature{}
	for _, item := range items {
		if feat, ok := itemsMap[item.Id]; ok {
			sorted = append(sorted, feat)
		}
	}
	return sorted
}

type SortItem struct {
	Value string
	Id    string
}
