package roadmap

import "encoding/json"

type RoadmapConfig2 struct {
	QuarterStart int32                `json:"quarterStart"`
	QuarterCount int32                `json:"quarterCount"`
	AhaConfig    RoadmapConfig2Aha    `json:"ahaConfig"`
	SlidesConfig RoadmapConfig2Slides `json:"slidesConfig"`
}

type RoadmapConfig2Aha struct {
	AddAhaLinks bool     `json:"addAhaLinks"`
	FilterTags  []string `json:"filterTags"`
}

type RoadmapConfig2Slides struct {
	Title                string `json:"title"`
	DisclaimerText       string `json:"disclaimerText"`
	FeatureSnapToQuarter bool   `json:"featureSnapToQuarter"`
	DimensionUnit        string `json:"dimensionsUnit"`
}

func ParseConfig(jsonCfg []byte) (RoadmapConfig2, error) {
	rc2 := RoadmapConfig2{}
	err := json.Unmarshal(jsonCfg, &rc2)
	return rc2, err
}
