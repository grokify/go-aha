package featuresutil

import (
	"strings"

	"github.com/grokify/go-aha/v3/oag7/aha"
)

const (
	IntegrationServiceNameJira    = "jira"
	IntegrationServiceNameJiraKey = "key"
	IntegrationServiceNameJiraURL = "url"
)

type Feature aha.Feature

func (f Feature) JiraKey() string {
	for _, intf := range f.IntegrationFields {
		if intf.ServiceName == IntegrationServiceNameJira && intf.Name == IntegrationServiceNameJiraKey {
			return strings.TrimSpace(intf.Value)
		}
	}
	return ""
}

func (f Feature) JiraURL() string {
	for _, intf := range f.IntegrationFields {
		if intf.ServiceName == IntegrationServiceNameJira && intf.Name == IntegrationServiceNameJiraURL {
			return strings.TrimSpace(intf.Value)
		}
	}
	return ""
}
