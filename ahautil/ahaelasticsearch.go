package ahautil

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/grokify/goelastic"
	"github.com/grokify/goelastic/models"
	"github.com/grokify/mogo/encoding/jsonutil/jsonraw"
	"github.com/grokify/mogo/net/http/httpsimple"
)

type AhaElasticsearch struct {
	AhaAPIs  *ClientAPIs
	EsClient httpsimple.Client
}

func (ae *AhaElasticsearch) IndexFeaturesUpdatedSince(updatedSince time.Time) error {
	api := ae.AhaAPIs.APIClient.FeaturesApi

	fres, res, err := api.GetFeatures(context.Background(), nil)

	if err != nil {
		return err
	} else if res.StatusCode >= 300 {
		return fmt.Errorf("AHA API Status Code: %v", res.StatusCode)
	}

	for _, f := range fres.Features {
		fmt.Println(f.Id)
		err := ae.IndexFeatureId(f.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ae *AhaElasticsearch) IndexFeatureId(featureID string) error {
	featuresApi := ae.AhaAPIs.APIClient.FeaturesApi

	feat, resp, err := featuresApi.GetFeature(context.Background(), featureID)
	if err != nil {
		return err
	} else if resp.StatusCode >= 300 {
		return fmt.Errorf("AHA API Error: %v", resp.StatusCode)
	}

	update := true

	esFeature := AhaToEsFeature(&feat.Feature)

	esReq := httpsimple.Request{
		Method:   http.MethodPost,
		URL:      strings.Join([]string{"/aha/feature", featureID, goelastic.SlugUpdate}, "/"),
		BodyType: httpsimple.BodyTypeJSON}

	if update {
		esReq.Body = models.UpdateIndexDoc{Doc: esFeature, DocAsUpsert: true}
	} else {
		esReq.Body = esFeature
	}

	resp, err = ae.EsClient.Do(esReq)

	if err != nil {
		return err
	} else {
		body, err := jsonraw.Indent(resp.Body, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf("U_RES_BODY: %v\n", string(body))
		fmt.Printf("U_RES_STATUS: %v\n", resp.StatusCode)
	}
	return nil
}
