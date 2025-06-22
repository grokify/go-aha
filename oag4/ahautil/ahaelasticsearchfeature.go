package ahautil

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/grokify/goelastic"
	"github.com/grokify/goelastic/models/es5"
	"github.com/grokify/mogo/encoding/jsonutil/jsonraw"
	"github.com/grokify/mogo/net/http/httpsimple"
	"github.com/grokify/mogo/time/timeutil"

	"github.com/grokify/go-aha/v3/oag4/aha"
)

var rxProductId = regexp.MustCompile(`^([^-]+)`)

// Feature is struct with additional properties to search on, notably ReferencePrefix.
type Feature struct {
	Id           string    `json:"id,omitempty"`
	ReferenceNum string    `json:"reference_num,omitempty"`
	Name         string    `json:"name,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	// Start date in YYYY-MM-DD format.
	StartDate string `json:"start_date,omitempty"`
	// Due date in YYYY-MM-DD format.
	DueDate         string   `json:"due_date,omitempty"`
	Url             string   `json:"url,omitempty"`
	Resource        string   `json:"resource,omitempty"`
	Tags            []string `json:"tags,omitempty"`
	ReferencePrefix string   `json:"reference_prefix,omitempty"`
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

func AhaFeatureSearch(ctx context.Context, esClient httpsimple.Client, refPrefix string, dt time.Time) ([]Feature, error) {
	features := []Feature{}

	body := es5.QueryBody{
		Query: es5.Query{
			Match: map[string]es5.MatchQuery{
				"reference_prefix": {
					Query: refPrefix,
				},
			},
			Range: map[string]es5.Range{
				"due_date": {
					GTE: dt.Format(timeutil.RFC3339FullDate),
				},
			},
		},
	}

	esReq := httpsimple.Request{
		Method:   http.MethodGet,
		URL:      strings.Join([]string{"/aha/feature", goelastic.SlugSearch}, "/"),
		BodyType: httpsimple.BodyTypeJSON,
		Body:     body}

	resp, err := esClient.Do(ctx, esReq)
	if err != nil {
		log.Fatal(err)
	}
	bodybytes, err := jsonraw.Indent(resp.Body, "", "  ")
	if err != nil {
		return features, err
	}

	fmt.Println(string(bodybytes))

	esRes := es5.ResponseBody{}
	err = json.Unmarshal(bodybytes, &esRes)

	if err != nil {
		fmt.Printf("U_ERR: %v\n", err)
	} else {
		if 1 == 1 {
			fmt.Printf("U_RES_BODY: %v\n", string(bodybytes))
		}
		fmt.Printf("U_RES_STATUS: %v\n", resp.StatusCode)
	}

	return features, err
}
