package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/time/timeutil"
	"github.com/grokify/swaggman/swagger2"
	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"

	"github.com/grokify/elastirad-go"
	"github.com/grokify/elastirad-go/models"
	"github.com/grokify/elastirad-go/models/v5"

	"github.com/grokify/go-aha/ahautil"
	"github.com/grokify/go-aha/client"
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

func indexFeature(api *aha.FeaturesApi, esClient elastirad.Client, featureId string) error {
	feat, resp, err := api.FeaturesFeatureIdGet(featureId)
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

func indexFeaturesPage(api *aha.FeaturesApi, esClient elastirad.Client, pageNum int32) (*aha.FeaturesResponse, *aha.APIResponse, error) {
	info, resp, err := api.FeaturesGet("", timeutil.TimeRFC3339Zero(), "", "", pageNum, 500)
	if err != nil || resp.StatusCode >= 400 {
		return info, resp, err
		//log.Fatal("Error retrieving features: %v", err.Error())
	}
	for i, f := range info.Features {
		fmt.Printf("PAGE %v IDX %v\n", pageNum, i)
		err := indexFeature(api, esClient, f.Id)
		if err != nil {
			panic(err)
		}
	}
	return info, resp, err
}

func main() {
	if 1 == 0 {
		specpath := filepath.Join(os.Getenv("GOPATH"), "src", "github.com/grokify/go-aha/codegen/aha_api-v1_swagger-v2.0.json")
		spec, err := swagger2.ReadSwagger2Spec(specpath)
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
	ahaClient := ahaoauth.NewClient(os.Getenv("AHA_ACCOUNT"), os.Getenv("AHA_API_KEY"))
	apis := ahautil.ClientAPIs{Client: ahaClient}

	doCreateIndex := false
	doUpdateIndex := false
	doUpdateIndexByLastUpdate := true

	if doCreateIndex {
		createIndex(esClient)
	}

	if doUpdateIndex {
		api := apis.FeaturesApi()
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
