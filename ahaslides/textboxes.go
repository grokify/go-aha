package ahaslides

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/grokify/gocharts/data/roadmap"
	su "github.com/grokify/googleutil/slidesutil/v1"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/math/mathutil"
	"github.com/grokify/gotilla/time/timeutil"
	"github.com/grokify/gotilla/type/stringsutil"
	"google.golang.org/api/slides/v1"

	"github.com/grokify/go-aha/aha"
	"github.com/grokify/go-aha/ahautil"
)

var Unit = "PT"

func getUniqueId() string {
	return strconv.Itoa(int(time.Now().Unix()))
}

func RoadmapTextBoxRequests(rmCfg RoadmapConfig, featureSet *ahautil.FeatureSet, pageId string) []*slides.Request {
	uid := getUniqueId()

	featuresMap2, err := featureSet.GetFeaturesMapByTag()
	if err != nil {
		log.Fatal(err)
	}

	if len(featuresMap2) == 0 {
		fmtutil.PrintJSON(featuresMap2)
		panic("aha-slides/textboxes.go/RoadmapTextBoxRequests - featureSet.GetFeaturesMapByTag - E_NO_FEATURES_IN_MAP2")
	}

	if 1 == 0 {
		fmtutil.PrintJSON(featuresMap2)
		panic("ZZ")
	}

	ahaTagFeaturesArr := []AhaTagFeatures{}

	// featureMapsArr is used to generate an array of features
	featureMapsArr := []map[string]*aha.Feature{}
	for _, filter := range rmCfg.FilterTags {
		if fMap, ok := featuresMap2[filter]; ok {
			featureMapsArr = append(featureMapsArr, fMap) // V1
			ahaTagFeaturesArr = append(ahaTagFeaturesArr, // V2
				AhaTagFeatures{
					Tag:      filter,
					Features: fMap,
				},
			)
		}
	}

	itemsRM := []roadmap.Item{}
	requests := []*slides.Request{}

	outCan := CanvasFloat64{
		MinX: 150.0,
		MinY: 115.0,
		MaxX: CanvasMaxX,
		MaxY: CanvasMaxY,
	}

	slideCanvas := SlideCanvasInfo{
		BoxFgColor:      su.MustParseRgbColorHex(rmCfg.RoadmapFormatting.Textbox.DefaultForegroundColorHex),
		BoxBgColor:      su.MustParseRgbColorHex(rmCfg.RoadmapFormatting.Textbox.DefaultBackgroundColorHex),
		BoxHeight:       20.0,
		BoxMarginBottom: 5.0,
		Canvas:          outCan,
	}

	totalBoxHeight := slideCanvas.BoxHeight + slideCanvas.BoxMarginBottom
	elBaseIndex := 0

	horzLabelRequests := []*slides.Request{}
	tagTextBoxInfo := su.CreateShapeTextBoxRequestInfo{
		PageId:             pageId,
		Height:             20.0,
		Width:              140.0,
		DimensionUnit:      rmCfg.DimensionUnit,
		LocationX:          20.0,
		LocationUnit:       rmCfg.DimensionUnit,
		FontSize:           rmCfg.RoadmapFormatting.Row.Heading.FontSize,
		FontSizeUnit:       rmCfg.RoadmapFormatting.Row.Heading.FontSizeUnit,
		FontBold:           rmCfg.RoadmapFormatting.Row.Heading.FontBold,
		FontItalic:         rmCfg.RoadmapFormatting.Row.Heading.FontItalic,
		ForegroundColorHex: rmCfg.RoadmapFormatting.Row.Heading.ForegroundColorHex,
	}
	tagTextBoxBgInfo := su.CreateShapeTextBoxRequestInfo{
		PageId:             pageId,
		Width:              680.0,
		DimensionUnit:      rmCfg.DimensionUnit,
		LocationX:          15.0,
		LocationUnit:       rmCfg.DimensionUnit,
		BackgroundColorHex: rmCfg.RoadmapFormatting.Row.RowOddBackgroundColorHex,
	}

	for i, ahaTagFeatures := range ahaTagFeaturesArr {
		fMap := ahaTagFeatures.Features

		srcCan := roadmap.Canvas{}

		err := srcCan.SetMinMaxQuarter(
			rmCfg.QuarterStartInt32, rmCfg.QuarterEndInt32)
		if err != nil {
			panic(err)
		}
		srcCan.SetRangeCells(200)
		if 1 == 0 {
			fmtutil.PrintJSON(srcCan)
			panic("Z")
		}
		if 1 == 0 {
			rng, err := srcCan.Range.CellRange()
			if err != nil {
				panic(err)
			}
			fmt.Printf("Cell Range %v", rng)
		}

		for _, feat := range fMap {
			if strings.ToLower(feat.WorkflowStatus.Name) == ahautil.FeatureStatusWillNotImplementLc {
				continue
			}
			fmt.Printf("STATUS [%v]\n", feat.WorkflowStatus.Name)
			item := ahaFeatureToRoadmapItem(rmCfg, srcCan, feat)
			itemsRM = append(itemsRM, item)
			srcCan.AddItem(item)
		}
		//panic("Z")

		err = srcCan.InflateItems()
		if err != nil {
			fmtutil.PrintJSON(srcCan)
			panic(err)
		}
		srcCan.BuildRows()
		if 1 == 0 {
			fmtutil.PrintJSON(srcCan)
			fmt.Println(len(srcCan.Rows))
		}
		requestsRoadmap, err := googleSlideDrawRoadmap(rmCfg, pageId, elBaseIndex, srcCan, slideCanvas, rmCfg.DimensionUnit, rmCfg.AddAhaLinks)
		if err != nil {
			panic(err)
		}
		elBaseIndex += len(fMap)
		requests = append(requests, requestsRoadmap...)

		rowCount := len(srcCan.Rows)
		rowHeight := float64(rowCount) * (slideCanvas.BoxHeight + slideCanvas.BoxMarginBottom)

		objectIdTag := su.FormatObjectIdSimple(ahaTagFeatures.Tag)

		if mathutil.IsEven(i) {
			tagTextBoxBgInfo.ObjectId = fmt.Sprintf("TAGLABELBG-%v-%v", objectIdTag, uid)
			tagTextBoxBgInfo.Height = totalBoxHeight * float64(len(srcCan.Rows))
			tagTextBoxBgInfo.LocationY = slideCanvas.Canvas.MinY
			tagTextBoxBgReqs, err := tagTextBoxBgInfo.Requests()
			if err != nil {
				panic(err)
			}
			horzLabelRequests = append(horzLabelRequests, tagTextBoxBgReqs...)
		}

		tagTextBoxInfo.ObjectId = fmt.Sprintf("TAGLABEL-%v-%v", objectIdTag, uid)
		tagTextBoxInfo.Text = rmCfg.TagPrefixStripRx.ReplaceAllString(ahaTagFeatures.Tag, "")
		tagTextBoxInfo.LocationY = slideCanvas.Canvas.MinY
		tagTextBoxReqs, err := tagTextBoxInfo.Requests()
		if err != nil {
			panic(err)
		}
		horzLabelRequests = append(horzLabelRequests, tagTextBoxReqs...)

		slideCanvas.Canvas.MinY += rowHeight
	}

	verts := getVerticalLinesAndHeadings(
		rmCfg, pageId,
		outCan.MinX, outCan.MaxX, outCan.MinY,
		rmCfg.DimensionUnit,
		rmCfg.QuarterStartTime, rmCfg.QuarterCount)
	requests = append(verts, requests...)
	requests = append(horzLabelRequests, requests...)

	return requests
}

func ahaFeatureToRoadmapItem(rmCfg RoadmapConfig, srcCan roadmap.Canvas, feat *aha.Feature) roadmap.Item {
	featStartTime, err := ahautil.GetBeginTimeForFeatureOrRelease(feat)
	if err != nil {
		panic(err)
	}
	featEndTime, err := ahautil.GetEndTimeForFeatureOrRelease(feat)
	if err != nil {
		panic(err)
	}
	if rmCfg.FeatureSnapToQuarter {
		featStartTime = timeutil.QuarterStart(featStartTime)
		featEndTime = timeutil.QuarterEnd(featEndTime)
	}

	featName := feat.Name
	featNameShort := featName
	if rmCfg.FeaturePrefixStripRx != nil {
		featName = rmCfg.FeaturePrefixStripRx.ReplaceAllString(featName, "")
		featNameShort = featName
	}
	if rmCfg.FeatureNameSepRx != nil {
		parts := rmCfg.FeatureNameSepRx.Split(featName, -1)
		if len(parts) > 0 {
			featNameShort = parts[0]
		}
		featName = rmCfg.FeatureNameSepRx.ReplaceAllString(featName, " ")
	}
	tboxFgColor, tboxBgColor := textboxColorsForAhaFeature(rmCfg, feat)

	return roadmap.Item{
		MinTime:            timeutil.MinTime(featStartTime, srcCan.MinTime),
		MaxTime:            timeutil.MaxTime(featEndTime, srcCan.MaxTime),
		Name:               featName,
		NameShort:          featNameShort,
		URL:                feat.Url,
		ForegroundColorHex: tboxFgColor,
		BackgroundColorHex: tboxBgColor}
}

func textboxColorsForAhaFeature(rmCfg RoadmapConfig, feat *aha.Feature) (string, string) {
	status := strings.ToLower(strings.TrimSpace(feat.WorkflowStatus.Name))
	switch status {
	case ahautil.FeatureStatusShippedLc:
		return rmCfg.RoadmapFormatting.Textbox.DoneForegroundColorHex,
			rmCfg.RoadmapFormatting.Textbox.DoneBackgroundColorHex
	case ahautil.FeatureStatusWillNotImplementLc:
		return rmCfg.RoadmapFormatting.Textbox.DeadForegroundColorHex,
			rmCfg.RoadmapFormatting.Textbox.DeadBackgroundColorHex
	default:
		if stringsutil.SliceIndexOfLcTrimSpace("problem", feat.Tags) > -1 {
			return rmCfg.RoadmapFormatting.Textbox.ProblemForegroundColorHex,
				rmCfg.RoadmapFormatting.Textbox.ProblemBackgroundColorHex
		} else {
			return rmCfg.RoadmapFormatting.Textbox.DefaultForegroundColorHex,
				rmCfg.RoadmapFormatting.Textbox.DefaultBackgroundColorHex
		}
	}
}

func googleSlideDrawRoadmap(rmCfg RoadmapConfig, pageId string, elBaseIndex int, srcCan roadmap.Canvas, outCan SlideCanvasInfo, dimensionUnit string, addAhaLinks bool) ([]*slides.Request, error) {
	uid := getUniqueId()
	requests := []*slides.Request{}
	err := srcCan.InflateItems()
	if err != nil {
		return requests, err
	}
	srcCan.BuildRows()

	idx := 0
	rowYWatermark := outCan.Canvas.MinY

	textBoxInfo := su.CreateShapeTextBoxRequestInfo{
		PageId:             pageId,
		DimensionUnit:      dimensionUnit,
		LocationUnit:       dimensionUnit,
		FontSize:           10.0,
		FontSizeUnit:       dimensionUnit,
		ForegroundColorRgb: outCan.BoxFgColor,
		BackgroundColorRgb: outCan.BoxBgColor}

	for iRow, row := range srcCan.Rows {
		for _, el := range row {
			fmtutil.PrintJSON(el)
			srcBoxWdtX := el.Max - el.Min
			srcAllWdtX := srcCan.MaxX - srcCan.MinX
			srcBoxMinX := el.Min
			srcBoxMaxX := el.Max
			srcPctWdtX := float64(srcBoxWdtX) / float64(srcAllWdtX)

			srcAllMinX := srcCan.MinX
			outAllWdtX := outCan.Canvas.MaxX - outCan.Canvas.MinX
			outBoxWdtX := srcPctWdtX * outAllWdtX

			boxOutPctX := float64(srcAllWdtX) / outAllWdtX

			outAllMinX := outCan.Canvas.MinX
			fmt.Printf("%v\n", srcBoxMinX-srcAllMinX)
			fmt.Printf("%v\n", el.Min-srcCan.MinX)
			outBoxMinX := (float64(srcBoxMinX-srcAllMinX) / float64(boxOutPctX)) + outAllMinX
			outBoxMaxX := (float64(srcBoxMaxX-srcAllMinX) / float64(boxOutPctX)) + outAllMinX

			loc := location{
				SrcAllMinX: srcCan.MinX,
				SrcAllMaxX: srcCan.MaxX,
				SrcAllWdtX: srcCan.MaxX - srcCan.MinX,
				SrcBoxMinX: el.Min,
				SrcBoxMaxX: el.Max,
				SrcBoxWdtX: srcBoxWdtX,
				SrcPctWdtX: srcPctWdtX,
				OutAllMinX: outCan.Canvas.MinX,
				OutAllMaxX: outCan.Canvas.MaxX,
				OutAllWdtX: outCan.Canvas.MaxX - outCan.Canvas.MinX,
				OutBoxMinX: outBoxMinX,
				OutBoxMaxX: outBoxMaxX,
				OutBoxWdtX: outBoxWdtX,
				BoxOutPctX: boxOutPctX,
			}

			if loc.OutBoxMaxX > loc.OutAllMaxX {
				panic("C")
			} else if loc.OutBoxMinX < loc.OutAllMinX {
				panic("D")
			}

			elementId := fmt.Sprintf("AutoBox%03d-%v", idx+elBaseIndex, uid)
			if 1 == 0 {
				requests = append(requests, su.TextBoxRequestsSimple(
					pageId, elementId, el.NameShort, outCan.BoxFgColor, outCan.BoxBgColor,
					loc.OutBoxWdtX, outCan.BoxHeight, loc.OutBoxMinX, rowYWatermark)...)
			} else {
				textBoxInfo.ObjectId = elementId
				textBoxInfo.Text = el.NameShort
				if addAhaLinks {
					textBoxInfo.URL = el.URL
				}
				textBoxInfo.Width = loc.OutBoxWdtX
				textBoxInfo.Height = outCan.BoxHeight
				textBoxInfo.LocationX = loc.OutBoxMinX
				textBoxInfo.LocationY = rowYWatermark

				textBoxInfo.ForegroundColorRgb = su.MustParseRgbColorHex(el.ForegroundColorHex)
				textBoxInfo.BackgroundColorRgb = su.MustParseRgbColorHex(el.BackgroundColorHex)

				reqs, err := textBoxInfo.Requests()
				if err != nil {
					panic(err)
				}
				requests = append(requests, reqs...)
			}

			idx++

			if 1 == 0 {
				//break
				if iRow > 3 {
					//break
				}
			}
		}
		rowYWatermark += outCan.BoxHeight + outCan.BoxMarginBottom
	}

	return requests, nil
}

func getVerticalLinesAndHeadings(rmCfg RoadmapConfig, pageId string, minX, maxX, minY float64, dimensionUnit string, qtrStartDt time.Time, numCells int32) []*slides.Request {
	uid := getUniqueId()

	requests := []*slides.Request{}
	if 1 == 0 {
		fmt.Printf("MINX [%v] MAXX [%v] NUM_CELLS [%v]\n", minX, maxX, numCells)
		panic("Z TEXTBOXES_GO LN_326")
	}
	rng := mathutil.RangeFloat64{
		Min:   minX,
		Max:   maxX,
		Cells: numCells,
	}
	linePrefix := "VertLines"
	lineInfo := su.CreateLineRequestInfo{
		PageId:        pageId,
		LineId:        "",
		ColorHex:      rmCfg.RoadmapFormatting.Line.ColorHex,
		LineCategory:  "STRAIGHT",
		Height:        400 - minY, // 400
		Width:         1.0,
		DimensionUnit: dimensionUnit,
		LocationX:     0.0,
		LocationY:     minY - 30.0,
		DashStyle:     rmCfg.RoadmapFormatting.Line.DashStyle,
		Weight:        1.0,
	}
	lineInfo.Height = 400.0 - lineInfo.LocationY

	textBoxInfo := su.CreateShapeTextBoxRequestInfo{
		PageId:             pageId,
		Height:             20.0,
		DimensionUnit:      dimensionUnit,
		LocationUnit:       dimensionUnit,
		FontBold:           rmCfg.RoadmapFormatting.Column.Heading.FontBold,
		FontItalic:         rmCfg.RoadmapFormatting.Column.Heading.FontItalic,
		FontSize:           rmCfg.RoadmapFormatting.Column.Heading.FontSize,
		FontSizeUnit:       rmCfg.RoadmapFormatting.Column.Heading.FontSizeUnit,
		ForegroundColorHex: rmCfg.RoadmapFormatting.Column.Heading.ForegroundColorHex,
	}

	qtrNow := qtrStartDt

	for i := 0; i < int(rng.Cells); i++ {
		min, max, err := rng.CellMinMax(int32(i))
		if err != nil {
			panic(err)
		}
		fmt.Printf("IDX %v MIN %v MAX %v\n", i, min, max)

		// Add Lines
		if i == 0 { // Build first line
			lineInfo.LineId = fmt.Sprintf("%v%03d%v-%v", linePrefix, i, "start", uid)
			lineInfo.LocationX = min
			lineReqs, err := lineInfo.Requests()
			if err != nil {
				panic(err)
			}
			requests = append(requests, lineReqs...)
		}
		lineInfo.LineId = fmt.Sprintf("%v%03d%v-%v", linePrefix, i, "end", uid)
		lineInfo.LocationX = max
		lineReqs, err := lineInfo.Requests()
		if err != nil {
			panic(err)
		}
		requests = append(requests, lineReqs...)

		// Add Quarter Heading
		textBoxInfo.ObjectId = lineInfo.LineId + "heading"
		textBoxInfo.Text = timeutil.FormatQuarter(qtrNow)
		textBoxInfo.Width = max - min - 2
		textBoxInfo.LocationX = min
		textBoxInfo.LocationY = lineInfo.LocationY
		reqs, err := textBoxInfo.Requests()
		if err != nil {
			panic(err)
		}
		requests = append(requests, reqs...)

		req := CenterRequest(textBoxInfo.ObjectId, "CENTER")
		requests = append(requests, req)

		qtrNow = timeutil.NextQuarter(qtrNow)
	}
	return requests
}

func CenterRequest(objectId, alignment string) *slides.Request {
	return &slides.Request{
		UpdateParagraphStyle: &slides.UpdateParagraphStyleRequest{
			ObjectId: objectId,
			Style: &slides.ParagraphStyle{
				Alignment: alignment,
			},
			Fields: "alignment",
		},
	}
}

type AhaTagFeatures struct {
	Tag      string
	Features map[string]*aha.Feature
}
