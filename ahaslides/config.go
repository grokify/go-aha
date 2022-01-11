package ahaslides

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/grokify/mogo/time/timeutil"
	"github.com/grokify/mogo/type/stringsutil"
	"github.com/pkg/errors"
)

type RoadmapConfig struct {
	Title                string            `json:"title"`
	FeaturesFilepath     string            `json:"featuresFilepath"`
	AddAhaLinks          bool              `json:"addAhaLinks"`
	DimensionUnit        string            `json:"dimensionsUnit"`
	DisclaimerText       string            `json:"disclaimerText"`
	FilterTags           []string          `json:filterTags"`
	FilterTagsMap        map[string]string `json:"filterTagsMap"`
	TagPrefixStripRx     *regexp.Regexp
	FeaturePrefixStripRx *regexp.Regexp
	FeatureNameSepRx     *regexp.Regexp
	FeatureSnapToQuarter bool  `json:"featureSnapToQuarter"`
	QuarterStartInt32    int32 `json:"quarterStart"`
	QuarterEndInt32      int32 `json:"quarterEnd"`
	QuarterStartTime     time.Time
	QuarterCount         int32             `json:"quarterCount"`
	RoadmapFormatting    RoadmapFormatting `json:"roadmapFormatting"`

	roadmapConfigRaw   string
	filterTagsRaw      string
	featurePrefixStrip string
	featureNameSep     string
	tagPrefixStrip     string
	quarterStartString string
	quarterCountString string
}

func NewRoadmapConfigEnv() (RoadmapConfig, error) {
	cfg := RoadmapConfig{
		AddAhaLinks:          stringsutil.ToBool(os.Getenv("ROADMAP_ADD_AHA_LINKS")),
		Title:                os.Getenv("ROADMAP_TITLE"),
		DisclaimerText:       os.Getenv("ROADMAP_DISCLAIMER_TEXT"),
		DimensionUnit:        "PT",
		FeaturesFilepath:     os.Getenv("ROADMAP_FEATURES_FILEPATH"),
		FeatureSnapToQuarter: stringsutil.ToBool(os.Getenv("ROADMAP_FEATURE_SNAP_TO_QUARTER")),
		FilterTags:           []string{},
		FilterTagsMap:        map[string]string{},
		filterTagsRaw:        os.Getenv("ROADMAP_FILTER_TAGS"),
		quarterStartString:   os.Getenv("ROADMAP_QUARTER_START"),
		quarterCountString:   os.Getenv("ROADMAP_QUARTER_COUNT"),
		tagPrefixStrip:       os.Getenv("ROADMAP_FILTER_TAG_PREFIX_STRIP"),
		featurePrefixStrip:   os.Getenv("ROADMAP_FEATURE_NAME_PREFIX_STRIP"),
		featureNameSep:       os.Getenv("ROADMAP_FEATURE_NAME_SEP"),
		roadmapConfigRaw:     os.Getenv("ROADMAP_FORMATTING"),
	}
	err := cfg.inflate()
	return cfg, err
}

func (cfg *RoadmapConfig) inflate() error {
	if len(cfg.tagPrefixStrip) > 0 {
		pre := fmt.Sprintf(`^%v`, cfg.tagPrefixStrip)
		rx, err := regexp.Compile(pre)
		if err != nil {
			return err
		}
		cfg.TagPrefixStripRx = rx
	}

	if len(cfg.featureNameSep) > 0 {
		sep := `\s+` + regexp.QuoteMeta(cfg.featureNameSep) + `\s+`
		rx, err := regexp.Compile(sep)
		if err != nil {
			return err
		}
		cfg.FeatureNameSepRx = rx
	}

	if len(cfg.featurePrefixStrip) > 0 {
		pre := fmt.Sprintf(`^%v`, cfg.featurePrefixStrip)
		rx, err := regexp.Compile(pre)
		if err != nil {
			return err
		}
		cfg.FeaturePrefixStripRx = rx
	}

	cfg.quarterStartString = strings.TrimSpace(cfg.quarterStartString)
	if len(cfg.quarterStartString) > 0 {
		i, err := strconv.Atoi(cfg.quarterStartString)
		if err != nil {
			return err
		}
		cfg.QuarterStartInt32 = int32(i)
		if !timeutil.IsQuarterInt32(cfg.QuarterStartInt32) {
			return fmt.Errorf("Start Quarter is invalid [%v] [%v]", cfg.QuarterStartInt32, err.Error())
		}
		qtrStartDt, err := timeutil.QuarterStringStartTime(cfg.quarterStartString)
		if err != nil {
			return err
		}
		cfg.QuarterStartTime = qtrStartDt
	}

	cfg.quarterCountString = strings.TrimSpace(cfg.quarterCountString)
	if len(cfg.quarterCountString) > 0 {
		quarterCount, err := strconv.Atoi(cfg.quarterCountString)
		if err != nil {
			return err
		}
		cfg.QuarterCount = int32(quarterCount)
		quarterEnd, err := timeutil.DeltaQuarterInt32(cfg.QuarterStartInt32, quarterCount-1)
		if err != nil {
			return errors.Wrap(err, "Calculate_Quarter_End")
		}
		cfg.QuarterEndInt32 = quarterEnd
	}
	cfg.inflateTagFilters()

	cfg.roadmapConfigRaw = strings.TrimSpace(cfg.roadmapConfigRaw)
	if len(cfg.roadmapConfigRaw) == 0 {
		cfg.RoadmapFormatting = DefaultFormatting
	}
	return nil
}

func (cfg *RoadmapConfig) inflateTagFilters() {
	cfg.FilterTags = stringsutil.SplitCondenseSpace(cfg.filterTagsRaw, ",")
	filterMap := map[string]string{}
	for _, filterTag := range cfg.FilterTags {
		filterTag = strings.TrimSpace(filterTag)
		filterMap[filterTag] = ""
	}
	cfg.FilterTagsMap = filterMap
}
