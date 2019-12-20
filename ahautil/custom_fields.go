package ahautil

import (
	"strings"

	"github.com/grokify/go-aha/aha"
)

func FeatureCustomField(want string, feat *aha.Feature) string {
	want = strings.TrimSpace(want)
	for _, cf := range feat.CustomFields {
		key := strings.TrimSpace(cf.Key)
		name := strings.TrimSpace(cf.Name)
		if want == key || want == name {
			return strings.TrimSpace(cf.Value)
		}
	}
	return ""
}
