package ahautil

import (
	"context"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/grokify/elastirad-go"
	"github.com/grokify/elastirad-go/models"
)

type AhaElasticsearch struct {
	AhaAPIs  *ClientAPIs
	EsClient *elastirad.Client
}

func (ae *AhaElasticsearch) IndexFeaturesUpdatedSince(updatedSince time.Time) error {
	api := ae.AhaAPIs.APIClient.FeaturesApi

	fres, res, err := api.FeaturesGet(context.Background(), map[string]interface{}{})

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

func (ae *AhaElasticsearch) IndexFeatureId(featureId string) error {
	featuresApi := ae.AhaAPIs.APIClient.FeaturesApi

	feat, resp, err := featuresApi.FeaturesFeatureIdGet(context.Background(), featureId)
	if err != nil {
		return err
	} else if resp.StatusCode >= 300 {
		return fmt.Errorf("AHA API Error: %v", resp.StatusCode)
	}

	update := true

	esFeature := AhaToEsFeature(feat.Feature)

	esReq := models.Request{
		Method: "POST",
		Path:   []interface{}{"/aha/feature", featureId, elastirad.UpdateSlug}}

	if update {
		esReq.Body = models.UpdateIndexDoc{Doc: esFeature, DocAsUpsert: true}
	} else {
		esReq.Body = esFeature
	}

	res, req, err := ae.EsClient.SendFastRequest(esReq)

	if err != nil {
		return err
	} else {
		fmt.Printf("U_RES_BODY: %v\n", string(res.Body()))
		fmt.Printf("U_RES_STATUS: %v\n", res.StatusCode())
	}
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)
	return nil
}
