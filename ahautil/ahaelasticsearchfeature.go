package ahautil

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/grokify/gotilla/time/timeutil"
	"github.com/valyala/fasthttp"

	"github.com/grokify/elastirad-go"
	"github.com/grokify/elastirad-go/models"
	"github.com/grokify/elastirad-go/models/v5"

	"github.com/grokify/go-aha/client"
)

var rxProductId = regexp.MustCompile(`^([^-]+)`)

// Feature is struct with additional properties to search on, notably ReferencePrefix.
type Feature struct {
	Id string `json:"id,omitempty"`

	ReferenceNum string `json:"reference_num,omitempty"`

	Name string `json:"name,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	// Start date in YYYY-MM-DD format.
	StartDate string `json:"start_date,omitempty"`

	// Due date in YYYY-MM-DD format.
	DueDate string `json:"due_date,omitempty"`

	Url string `json:"url,omitempty"`

	Resource string `json:"resource,omitempty"`

	Tags []string `json:"tags,omitempty"`

	ReferencePrefix string `json:"reference_prefix,omitempty"`
}

func AhaToEsFeature(f *aha.Feature) Feature {
	f2 := Feature{
		Id:              f.Id,
		ReferenceNum:    f.ReferenceNum,
		Name:            f.Name,
		CreatedAt:       f.CreatedAt,
		StartDate:       f.StartDate, //JCW CHECK THIS
		DueDate:         f.DueDate,
		Url:             f.Url,
		Resource:        f.Resource,
		Tags:            f.Tags,
		ReferencePrefix: ReferencePrefixFromReferenceNum(f.ReferenceNum),
	}
	return f2
}

func ReferencePrefixFromReferenceNum(refNum string) string {
	m := rxProductId.FindAllStringSubmatch(refNum, -1)
	if len(m) > 0 && len(m[0]) == 2 {
		return m[0][1]
	}
	return ""
}

func AhaFeatureSearch(esClient elastirad.Client, refPrefix string, dt time.Time) ([]Feature, error) {
	features := []Feature{}

	body := v5.QueryBody{
		Query: v5.Query{
			Match: map[string]v5.MatchQuery{
				"reference_prefix": {
					Query: refPrefix,
				},
			},
			Range: map[string]v5.Range{
				"due_date": {
					GTE: dt.Format(timeutil.RFC3339YMD),
				},
			},
		},
	}

	esReq := models.Request{
		Method: "GET",
		Path:   []interface{}{"/aha/feature", elastirad.SearchSlug},
		Body:   body}

	res, req, err := esClient.SendFastRequest(esReq)

	fmt.Println(string(res.Body()))

	esRes := v5.ResponseBody{}
	err = json.Unmarshal(res.Body(), &esRes)

	if err != nil {
		fmt.Printf("U_ERR: %v\n", err)
	} else {
		if 1 == 1 {
			fmt.Printf("U_RES_BODY: %v\n", string(res.Body()))
		}
		fmt.Printf("U_RES_STATUS: %v\n", res.StatusCode())
	}
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)

	return features, err
}
