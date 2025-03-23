package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/antihax/optional"
	"github.com/grokify/goelastic"
	"github.com/grokify/goelastic/models"
	"github.com/grokify/goelastic/models/es5"
	"github.com/grokify/mogo/encoding/jsonutil/jsonraw"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/log/logutil"
	"github.com/grokify/mogo/net/http/httpsimple"
	"github.com/grokify/mogo/os/osutil"
	"github.com/grokify/mogo/time/timeutil"
	"github.com/grokify/spectrum/openapi2"
	"github.com/joho/godotenv"

	"github.com/grokify/go-aha/v2/aha"
	"github.com/grokify/go-aha/v2/ahautil"
	ahaoauth "github.com/grokify/goauth/aha"
)

func createIndex(esClient httpsimple.Client) {
	body := es5.CreateIndexBody{
		Mappings: map[string]es5.Mapping{
			"feature": {
				All: es5.All{Enabled: true},
				Properties: map[string]es5.Property{
					"id":            {Type: "string", Index: "not_analyzed"},
					"reference_num": {Type: "keyword", Index: "not_analyzed"},
					//"product_id":    v5.Property{Type: "keyword", Index: "not_analyzed"},
					"reference_prefix": {Type: "keyword", Index: "not_analyzed"},
					"name":             {Type: "string"},
					"start_date":       {Type: "date", Format: "yyyy-MM-dd"},
					"due_date":         {Type: "date", Format: "yyyy-MM-dd"},
				},
			},
		},
	}
	fmtutil.PrintJSON(body)
	esReq := httpsimple.Request{
		Method:   http.MethodPut,
		URL:      "/aha",
		BodyType: httpsimple.BodyTypeJSON,
		Body:     body}

	resp, err := esClient.Do(esReq)
	if err != nil {
		fmt.Printf("U_ERR: %v\n", err)
	} else {
		body, err := jsonraw.Indent(resp.Body, "", "  ")
		logutil.FatalErr(err)
		fmt.Printf("U_RES_BODY: %v\n", string(body))
		fmt.Printf("U_RES_STATUS: %v\n", resp.StatusCode)
	}
}

func indexFeature(api *aha.FeaturesApiService, esClient httpsimple.Client, featureId string) error {
	feat, resp, err := api.GetFeature(context.Background(), featureId)
	if err != nil {
		return err
	} else if resp.StatusCode >= 300 {
		return fmt.Errorf("AHA API Error: %v", resp.StatusCode)
	}

	update := true

	esReq := httpsimple.Request{
		Method:   http.MethodPost,
		URL:      strings.Join([]string{"/aha/feature", featureId, goelastic.SlugUpdate}, "/"),
		BodyType: httpsimple.BodyTypeJSON}

	if update {
		esReq.Body = models.UpdateIndexDoc{Doc: feat.Feature, DocAsUpsert: true}
	} else {
		esReq.Body = feat.Feature
	}

	resp, err = esClient.Do(esReq)
	if err != nil {
		fmt.Printf("U_ERR: %v\n", err)
		panic(err)
	} else {
		body, err := jsonraw.Indent(resp.Body, "", "  ")
		logutil.FatalErr(err)
		fmt.Printf("U_RES_BODY: %v\n", string(body))
		fmt.Printf("U_RES_STATUS: %v\n", resp.StatusCode)
	}
	return nil
}

func indexFeaturesPage(api *aha.FeaturesApiService, esClient httpsimple.Client, pageNum int32) (*aha.FeaturesResponse, *http.Response, error) {
	// func indexFeaturesPage(api *aha.FeaturesApiService, esClient goelastic.Client, pageNum int32) (*aha.FeaturesResponse, *aha.APIResponse, error) {
	opts := aha.GetFeaturesOpts{
		Page:    optional.NewInt32(pageNum),
		PerPage: optional.NewInt32(int32(500))}

	info, resp, err := api.GetFeatures(context.Background(), &opts)
	if err != nil || resp.StatusCode >= 400 {
		return &info, resp, err
		//log.Fatal("Error retrieving features: %v", err.Error())
	}
	for i, f := range info.Features {
		fmt.Printf("PAGE %v IDX %v\n", pageNum, i)
		err := indexFeature(api, esClient, f.Id)
		if err != nil {
			panic(err)
		}
	}
	return &info, resp, err
}

func main() {
	if 1 == 0 {
		specpath := osutil.GoPath("src", "github.com/grokify/go-aha/codegen/aha_api-v1_swagger-v2.0.json")
		spec, err := openapi2.ReadOpenAPI2SpecFile(specpath)
		if err != nil {
			panic(err)
		}
		fmtutil.PrintJSON(spec)

		if feat, ok := spec.Definitions["Feature"]; ok {
			fmtutil.PrintJSON(feat)
			fmt.Println("GOTIT")
		}
	}
	err := godotenv.Load(os.Getenv("ENV_PATH"))
	if err != nil {
		panic(err)
	}

	esClient, err := goelastic.NewSimpleClient("", "", "", true)
	if err != nil {
		panic(err)
	}
	ahaHttpClient, err := ahaoauth.NewClient(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	apis := ahautil.NewClientAPIsHTTPClient(
		fmt.Sprintf(ahaoauth.BaseURLFormat, os.Getenv("AHA_ACCOUNT")), ahaHttpClient)

	doCreateIndex := false
	doUpdateIndex := false
	doUpdateIndexByLastUpdate := true

	if doCreateIndex {
		createIndex(esClient)
	}

	if doUpdateIndex {
		api := apis.APIClient.FeaturesApi
		nxtPage := int32(1)
		maxPage := int32(1)
		idx := 0
		for nxtPage <= maxPage {
			fmt.Printf("PAGE %v\n", nxtPage)
			fmt.Printf("BEG %v NXT %v MAX %v\n", idx, nxtPage, maxPage)
			features, resp, err := indexFeaturesPage(api, esClient, nxtPage)
			if err != nil {
				panic(err)
			} else if resp.StatusCode >= 400 {
				panic(resp.StatusCode)
			}
			nxtPage = int32(features.Pagination.CurrentPage) + 1
			maxPage = int32(features.Pagination.TotalPages)
			fmt.Printf("FIN %v NXT %v MAX %v\n", idx, nxtPage, maxPage)
			idx += 1
		}
	}

	if doUpdateIndexByLastUpdate {
		ahaes := ahautil.AhaElasticsearch{
			AhaAPIs:  &apis,
			EsClient: esClient,
		}

		t, err := timeutil.TimeDeltaDow(time.Now().UTC(), time.Sunday, -10, false, false)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", t.Format(time.RFC3339))

		err = ahaes.IndexFeaturesUpdatedSince(t)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("DONE")
}
