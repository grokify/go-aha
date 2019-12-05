package ahautil

import (
	"strings"

	"github.com/grokify/go-aha/aha"
)

func FeatureCustomField(key string, feat *aha.Feature) string {
	key = strings.TrimSpace(key)
	for _, cf := range feat.CustomFields {
		name := strings.TrimSpace(cf.Name)
		if key == name {
			return strings.TrimSpace(cf.Value)
		}
	}
	return ""
}
