package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/antihax/optional"
	"github.com/grokify/simplego/fmt/fmtutil"
	"github.com/grokify/simplego/time/timeutil"
	"github.com/grokify/swaggman/swagger2"
	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"

	"github.com/grokify/elastirad-go"
	"github.com/grokify/elastirad-go/models"
	v5 "github.com/grokify/elastirad-go/models/v5"

	"github.com/grokify/go-aha/aha"
	"github.com/grokify/go-aha/ahautil"
	ahaoauth "github.com/grokify/oauth2more/aha"
)

func createIndex(esClient elastirad.Client) {
	body := v5.CreateIndexBody{
		Mappings: map[string]v5.Mapping{
			"feature": {
				All: v5.All{Enabled: true},
				Properties: map[string]v5.Property{
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
	esReq := models.Request{
		Method: "PUT",
		Path:   []interface{}{"/aha"},
		Body:   body}

	res, req, err := esClient.SendFastRequest(esReq)

	if err != nil {
		fmt.Printf("U_ERR: %v\n", err)
	} else {
		fmt.Printf("U_RES_BODY: %v\n", string(res.Body()))
		fmt.Printf("U_RES_STATUS: %v\n", res.StatusCode())
	}
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)
}

func indexFeature(api *aha.FeaturesApiService, esClient elastirad.Client, featureId string) error {
	feat, resp, err := api.GetFeature(context.Background(), featureId)
	if err != nil {
		return err
	} else if resp.StatusCode >= 300 {
		return fmt.Errorf("AHA API Error: %v", resp.StatusCode)
	}

	update := true

	esReq := models.Request{
		Method: "POST",
		Path:   []interface{}{"/aha/feature", featureId, elastirad.UpdateSlug}}

	if update {
		esReq.Body = models.UpdateIndexDoc{Doc: feat.Feature, DocAsUpsert: true}
	} else {
		esReq.Body = feat.Feature
	}

	res, req, err := esClient.SendFastRequest(esReq)

	if err != nil {
		fmt.Printf("U_ERR: %v\n", err)
		panic(err)
	} else {
		fmt.Printf("U_RES_BODY: %v\n", string(res.Body()))
		fmt.Printf("U_RES_STATUS: %v\n", res.StatusCode())
	}
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)
	return err
}

//func indexFeaturesPage(api *aha.FeaturesApiService, esClient elastirad.Client, pageNum int32) (*aha.FeaturesResponse, *aha.APIResponse, error) {
func indexFeaturesPage(api *aha.FeaturesApiService, esClient elastirad.Client, pageNum int32) (*aha.FeaturesResponse, *http.Response, error) {
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
		specpath := filepath.Join(os.Getenv("GOPATH"), "src", "github.com/grokify/go-aha/codegen/aha_api-v1_swagger-v2.0.json")
		spec, err := swagger2.ReadOpenAPI2SpecFile(specpath)
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

	esClient := elastirad.NewClient(url.URL{})
	ahaHttpClient := ahaoauth.NewClient(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))
	apis := ahautil.NewClientAPIsHTTPClient(ahaHttpClient)

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
			EsClient: &esClient,
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
