package client

import (
	"fmt"

	ao "github.com/grokify/goauth/aha"
	"github.com/grokify/mogo/net/http/httpsimple"

	"github.com/grokify/go-aha/v3/oag7/aha"
)

func NewConfiguration(subdomain, apiToken string) (*aha.Configuration, error) {
	hc, err := ao.NewClient(subdomain, apiToken)
	if err != nil {
		return nil, err
	}
	ahaHost := fmt.Sprintf("%s.aha.io", subdomain)
	ahaSvrAPIURL := fmt.Sprintf("https://%s/api/v1", ahaHost)
	cfg := &aha.Configuration{
		Host:       ahaHost,
		Scheme:     "https",
		HTTPClient: hc,
		Servers: aha.ServerConfigurations{
			{URL: ahaSvrAPIURL},
		},
	}
	return cfg, nil
}

func NewSimpleClient(subdomain, apiToken string) (*httpsimple.Client, error) {
	hc, err := ao.NewClient(subdomain, apiToken)
	if err != nil {
		return nil, err
	}
	ahaSvrURL := fmt.Sprintf("https://%s.aha.io/", subdomain)
	sc := httpsimple.NewClient(hc, ahaSvrURL)
	return &sc, nil
}
