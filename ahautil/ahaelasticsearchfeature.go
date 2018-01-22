package ahautil

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	//"github.com/grokify/gotilla/net/httputilmore"
	"github.com/grokify/gotilla/time/timeutil"
	"github.com/valyala/fasthttp"

	"github.com/grokify/elastirad-go"
	"github.com/grokify/elastirad-go/models"
	"github.com/grokify/elastirad-go/models/v5"

	"github.com/grokify/go-aha/client"
)

var rxProductId = regexp.MustCompile(`^([^-]+)`)

// Feature is struct with addtional properties to search on, notably ReferencePrefix.
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

/*

WORKS

{
    "query": {
        "range" : {
            "due_date" : {
                "gte": "2017-12-01"
            }
        }
    }
}

GOOD curl -XGET http://localhost:9200/aha/feature/_search -d  '{"query":{"range":{"due_date":{"gte":"2017-12-01"}}}}'

GOOD curl -XGET http://localhost:9200/aha/feature/_search -d '{"query": { "match" : { "referece_prefix" : "API"  } }}'

curl -XGET http://localhost:9200/aha/feature/_search -d '{"query":{"bool":{"must":[{"match":{"referece_prefix":"API"}},{"range":{"due_date":{"gte":"2017-12-01"}}}]}}}'


{
  "query": {
    "bool" : {
      "must" : [
        {
          "match" : { "referece_prefix" : "API" }
        },
        {
          "range" : { "due_date" : "due_date":{"gte":"2017-12-01"} }
        },
      ]
    }
  }
}


curl -XGET http://localhost:9200/aha/feature/_search -d  '{"query":{"filter":{"term":{"referece_prefix":"API"}}}}'

curl -XGET http://localhost:9200/aha/feature/_search -d '{"query":{"bool":{"filter":{"term":{"referece_prefix":"API"}},"must":{"range":{"due_date":"2017-12-01"}},"boost":1.0}}}'

curl -XGET http://localhost:9200/aha/feature/_search -d '{"query":{"bool":{"filter":{"term":{"referece_prefix":"API"}},"boost":1.0}}}'

====




{"aha":{"mappings":{"feature":{"_all":{"enabled":true},"properties":{"created_at":{"type":"date"},"due_date":{"type":"date","format":"yyyy-MM-dd"},"id":{"type":"keyword"},"initiative":{"properties":{"created_at":{"type":"date"},"id":{"type":"text","fields":{"keyword":{"type":"keyword","ignore_above":256}}},"name":{"type":"text","fields":{"keyword":{"type":"keyword","ignore_above":256}}},"resource":{"type":"text","fields":{"keyword":{"type":"keyword","ignore_above":256}}},"url":{"type":"text","fields":{"keyword":{"type":"keyword","ignore_above":256}}}}},"name":{"type":"text"},"referece_prefix":{"type":"text","fields":{"keyword":{"type":"keyword","ignore_above":256}}},"reference_num":{"type":"keyword"},"reference_prefix":{"type":"keyword"},"resource":{"type":"text","fields":{"keyword":{"type":"keyword","ignore_above":256}}},"start_date":{"type":"date","format":"yyyy-MM-dd"},"tags":{"type":"text","fields":{"keyword":{"type":"keyword","ignore_above":256}}},"url":{"type":"text","fields":{"keyword":{"type":"keyword","ignore_above":256}}}}}}}}



curl -XGET http://localhost:9200/aha/feature/_search -d  '{"query":{"range":{"due_date":{"gte":"2017-12-01"}}}}'

curl -XGET http://localhost:9200/aha/feature/_search -d  '{"query":{"bool":{"must":{"range":{"due_date":"2017-12-01"}},"minimum_should_match":1,"boost":1.0}}}'

{
  "query": {
    "bool" : {
      "filter" : {
        "term" : { "referece_prefix" : "API" }
      },
      "must": {
        "range" : {
          "due_date" : "2017-12-01"
        }
      },
      "minimum_should_match" : 1,
      "boost" : 1.0
    }
  }
}

curl -XGET http://localhost:9200/aha/feature/_search -d '{"query":{"bool":{"filter":{"term":{"referece_prefix":"API"}},"must":{"range":{"due_date":"2017-12-01"}},"boost":1.0}}}'

curl -XGET http://localhost:9200/aha/feature/_search -d '{"query":{"bool":{"filter":{"range":{"due_date":"2017-12-01"}},"minimum_should_match":1,"boost":1.0}}}'


curl -XGET http://localhost:9200/aha/feature/_search -d '{"query":{"bool":{"filter":{"range":{"due_date":"12/01/2017"}},"minimum_should_match":1,"boost":1.0}}}'

{
  "query": {
    "bool" : {
      "must" : {
        "term" : { "referece_prefix" : "API" },
        "range" : {
          "age" : { "due_date" : "2017-12-01" }
        }
      },
      "minimum_should_match" : 1,
      "boost" : 1.0
    }
  }
}

curl -XGET http://localhost:9200/aha/feature/_search -d  '{"query":{"bool":{"must":{"term":{"referece_prefix":"API"},"range":{"age":{"due_date":"2017-12-01"}}},"boost":1.0}}}'


curl -XGET http://localhost:9200/aha/feature/_search -d '{"query":{"bool":{"filter":{"term":{"referece_prefix":"API"}},"must":{"range":{"age":{"due_date":"2017-12-01"}}},"minimum_should_match":1,"boost":1.0}}}'


{
  "query": {
    "bool" : {
      "must" : {
        "term" : { "user" : "kimchy" }
      },
      "filter": {
        "term" : { "tag" : "tech" }
      },
      "must_not" : {
        "range" : {
          "age" : { "gte" : 10, "lte" : 20 }
        }
      },
      "should" : [
        { "term" : { "tag" : "wow" } },
        { "term" : { "tag" : "elasticsearch" } }
      ],
      "minimum_should_match" : 1,
      "boost" : 1.0
    }
  }
}

curl -XGET http://localhost:9200/aha/feature/_search -d '{"query": { "match" : { "referece_prefix" : "API"  }, "range": { "due_date": { "gte": "2017-12-01"} } }}'


curl -XGET http://localhost:9200/aha/feature/_search -d '{"query": { "match" : { "referece_prefix" : "API"  } }}'

curl -XGET http://localhost:9200/aha/feature/_search -d '{"query": {  "range": { "due_date": { "gte": "2017-12-01"} } }}'



curl -XGET http://localhost:9200/aha/feature/_search -d '{"query": { "match" : { "referece_prefix" : "API"  } , "range": { "due_date": { "gte": "2017-12-01"} }}}'


{"took":103,"timed_out":false,"_shards":{"total":5,"successful":5,"failed":0},"hits":{"total":171,"max_score":1.0,"hits":[{"_index":"aha","_type":"feature","_id":"6415263541142796262","_score":1.0,"_source":{"id":"6415263541142796262","reference_num":"GLIP-16","name":"Glip Form Support","created_at":"2017-05-01T20:18:59.493Z","url":"https://ringcentral.aha.io/features/GLIP-16","resource":"https://ringcentral.aha.io/api/v1/features/GLIP-16","initiative":{"created_at":"0001-01-01T00:00:00Z"},"due_date":"2018-12-31","start_date":"2018-07-01","referece_prefix":"GLIP","tags":["BOCA","Forms"]}},{"_index":"aha","_type":"feature","_id":"6413380988347890947","_score":1.0,"_source":{"id":"6413380988347890947","reference_num":"API-6","name":"Voicemail Drop APIs","created_at":"2017-04-26T18:33:43.512Z","url":"https://ringcentral.aha.io/features/API-6","resource":"https://ringcentral.aha.io/api/v1/features/API-6","initiative":{"created_at":"0001-01-01T00:00:00Z"},"due_date":"2018-12-31","start_date":"2018-07-01","referece_prefix":"API","tags":["Platform","Telco"]}},{"_index":"aha","_type":"feature","_id":"6437153540550256600","_score":1.0,"_source":{"id":"6437153540550256600","reference_num":"API-50","name":"Access Delegation (API)","created_at":"2017-06-29T20:03:22.142Z","url":"https://ringcentral.aha.io/features/API-50","resource":"https://ringcentral.aha.io/api/v1/features/API-50","initiative":{"created_at":"0001-01-01T00:00:00Z"},"due_date":"2018-12-31","start_date":"2018-07-01","referece_prefix":"API"}},{"_index":"aha","_type":"feature","_id":"6450432641029624428","_score":1.0,"_source":{"id":"6450432641029624428","reference_num":"SDK-38","name":"Quickstart - SMS - Python","created_at":"2017-08-04T14:53:03.654Z","url":"https://ringcentral.aha.io/features/SDK-38","resource":"https://ringcentral.aha.io/api/v1/features/SDK-38","initiative":{"created_at":"0001-01-01T00:00:00Z"},"due_date":"2018-12-31","start_date":"2018-07-01","referece_prefix":"SDK"}},{"_index":"aha","_type":"feature","_id":"6412771692923127469","_score":1.0,"_source":{"id":"6412771692923127469","reference_num":"DPW-30","name":"Intercom Support Chat","created_at":"2017-04-25T03:09:20.865Z","url":"https://ringcentral.aha.io/features/DPW-30","resource":"https://ringcentral.aha.io/api/v1/features/DPW-30","initiative":{"id":"6412771945909400418","name":"Customer Support","url":"https://ringcentral.aha.io/initiatives/6412771945909400418","resource":"https://ringcentral.aha.io/api/v1/initiatives/6412771945909400418","created_at":"2017-04-25T03:10:19.767Z"},"due_date":"2018-12-31","start_date":"2018-07-01","referece_prefix":"DPW","tags":["Customer Support"]}},{"_index":"aha","_type":"feature","_id":"6434246221889140662","_score":1.0,"_source":{"id":"6434246221889140662","reference_num":"API-41","name":"Contact Viewing Permissions","created_at":"2017-06-22T00:01:29.27Z","url":"https://ringcentral.aha.io/features/API-41","resource":"https://ringcentral.aha.io/api/v1/features/API-41","initiative":{"created_at":"0001-01-01T00:00:00Z"},"due_date":"2017-12-15","start_date":"2017-07-09","referece_prefix":"API","reference_prefix":"API"}},{"_index":"aha","_type":"feature","_id":"6434199483936190474","_score":1.0,"_source":{"id":"6434199483936190474","reference_num":"DPW-61","name":"SW: API Delegation (Pro Serve, Support)","created_at":"2017-06-21T21:00:07.241Z","url":"https://ringcentral.aha.io/features/DPW-61","resource":"https://ringcentral.aha.io/api/v1/features/DPW-61","initiative":{"created_at":"0001-01-01T00:00:00Z"},"due_date":"2018-12-31","start_date":"2018-07-01","referece_prefix":"DPW"}},{"_index":"aha","_type":"feature","_id":"6434206097057995332","_score":1.0,"_source":{"id":"6434206097057995332","reference_num":"SDK-30","name":"Widget Feature: Local Presence","created_at":"2017-06-21T21:25:46.975Z","url":"https://ringcentral.aha.io/features/SDK-30","resource":"https://ringcentral.aha.io/api/v1/features/SDK-30","initiative":{"created_at":"0001-01-01T00:00:00Z"},"due_date":"2018-12-31","start_date":"2018-07-01","referece_prefix":"SDK"}},{"_index":"aha","_type":"feature","_id":"6411536893062657387","_score":1.0,"_source":{"id":"6411536893062657387","reference_num":"DPW-13","name":"Deprecating Legacy API's","created_at":"2017-04-21T19:17:41.621Z","url":"https://ringcentral.aha.io/features/DPW-13","resource":"https://ringcentral.aha.io/api/v1/features/DPW-13","initiative":{"id":"6411539465330075841","name":"Maintenance","url":"https://ringcentral.aha.io/initiatives/6411539465330075841","resource":"https://ringcentral.aha.io/api/v1/initiatives/6411539465330075841","created_at":"2017-04-21T19:27:40.523Z"},"due_date":"2017-12-31","start_date":"2017-11-20","referece_prefix":"DPW","reference_prefix":"DPW"}},{"_index":"aha","_type":"feature","_id":"6411542379176694556","_score":1.0,"_source":{"id":"6411542379176694556","reference_num":"DPW-15","name":"Upgrade Developer Accounts from Free to Premium","created_at":"2017-04-21T19:38:58.956Z","url":"https://ringcentral.aha.io/features/DPW-15","resource":"https://ringcentral.aha.io/api/v1/features/DPW-15","initiative":{"id":"6411543063665978880","name":"Upsell","url":"https://ringcentral.aha.io/initiatives/6411543063665978880","resource":"https://ringcentral.aha.io/api/v1/initiatives/6411543063665978880","created_at":"2017-04-21T19:41:38.324Z"},"due_date":"2017-12-31","start_date":"2017-11-20","referece_prefix":"DPW","tags":["Account","Free Account","Upsell"],"reference_prefix":"DPW"}}]}}

*/
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
