package features

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
		if intf.GetServiceName() == IntegrationServiceNameJira && intf.GetName() == IntegrationServiceNameJiraKey {
			return strings.TrimSpace(intf.GetValue())
		}
	}
	return ""
}

func (f Feature) JiraURL() string {
	for _, intf := range f.IntegrationFields {
		if intf.GetServiceName() == IntegrationServiceNameJira && intf.GetName() == IntegrationServiceNameJiraURL {
			return strings.TrimSpace(intf.GetValue())
		}
	}
	return ""
}
