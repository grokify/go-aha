package ahautil

import (
	"strings"
	"time"

	tu "github.com/grokify/mogo/time/timeutil"

	"github.com/grokify/go-aha/v3/oagv4/aha"
)

func GetBeginTimeForFeatureOrRelease(feature *aha.Feature) (time.Time, error) {
	if possibleDate(feature.StartDate) {
		return time.Parse(tu.RFC3339FullDate, feature.StartDate)
	} else if possibleDate(feature.Release.StartDate) {
		return time.Parse(tu.RFC3339FullDate, feature.Release.StartDate)
	}
	return time.Now(), newErrorDateNotFound([]string{feature.StartDate, feature.Release.StartDate})
}

func GetEndTimeForFeatureOrRelease(feature *aha.Feature) (time.Time, error) {
	if possibleDate(feature.DueDate) {
		return time.Parse(tu.RFC3339FullDate, feature.DueDate)
	} else if possibleDate(feature.Release.ReleaseDate) {
		return time.Parse(tu.RFC3339FullDate, feature.Release.ReleaseDate)
	}
	return time.Now(), newErrorDateNotFound([]string{feature.DueDate, feature.Release.ReleaseDate})
}

func possibleDate(dateString string) bool {
	dateString = strings.TrimSpace(dateString)
	if len(dateString) == 0 {
		return false
	} else if strings.Index(dateString, "0") == 0 {
		return false
	}
	return true
}
