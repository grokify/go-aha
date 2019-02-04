// Go example that covers:
// Quickstart: https://developers.google.com/slides/quickstart/go
// Basic writing: adding a text box to slide: https://developers.google.com/slides/samples/writing
// Using SDK: https://github.com/google/google-api-go-client/blob/master/slides/v1/slides-gen.go
// Creating and Managing Presentations https://developers.google.com/slides/how-tos/presentations
// Adding Shapes and Text to a Slide: https://developers.google.com/slides/how-tos/add-shape#example
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/grokify/gotilla/fmt/fmtutil"
	tu "github.com/grokify/gotilla/time/timeutil"
	ou "github.com/grokify/oauth2more"
	oug "github.com/grokify/oauth2more/google"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
	"google.golang.org/api/slides/v1"

	aha "github.com/grokify/go-aha/client"
	"github.com/grokify/gocharts/data/roadmap"
	su "github.com/grokify/googleutil/slidesutil/v1"
)

var (
	GoogleSlideUnitPoint = "PT"
)

func NewClient(forceNewToken bool) (*http.Client, error) {
	conf, err := oug.ConfigFromEnv(oug.ClientSecretEnv,
		[]string{slides.DriveScope, slides.PresentationsScope})
	if err != nil {
		return nil, err
	}

	tokenFile := "slides.googleapis.com-go-quickstart.json"
	tokenStore, err := ou.NewTokenStoreFileDefault(tokenFile, true, 0700)
	if err != nil {
		return nil, err
	}

	return ou.NewClientWebTokenStore(context.Background(), conf, tokenStore, forceNewToken)
}

type CanvasFloat64 struct {
	MinX float64
	MinY float64
	MaxX float64
	MaxY float64
}

func (c64 *CanvasFloat64) ThisX(this, min, max float64) (float64, error) {
	if min > max {
		return 0.0, fmt.Errorf("Min (%v) is larger than max (%v)", min, max)
	} else if this < min || this > max {
		return 0.0, fmt.Errorf("This (%v) is not within min,max (%v, %v)", this, min, max)
	}
	diff := max - min
	plus := this - min
	pct := float64(plus) / float64(diff)
	diffCan := c64.MaxX - c64.MinX
	thisPlus := pct * diffCan
	thisX := c64.MinX + thisPlus
	return thisX, nil
}

func ReadFeatures(featuresPath string) (map[string]aha.Feature, error) {
	featuresMap := map[string]aha.Feature{}
	bytes, err := ioutil.ReadFile(featuresPath)
	if err != nil {
		return featuresMap, err
	}
	err = json.Unmarshal(bytes, &featuresMap)
	return featuresMap, err
}

type SlideCanvasInfo struct {
	BoxFgColor      *slides.RgbColor
	BoxBgColor      *slides.RgbColor
	BoxHeight       float64
	BoxMarginBottom float64
	Canvas          CanvasFloat64
}

/*
	locYHeight := boxHeight + 5.0
	for i, itemText := range items {
		elementId := fmt.Sprintf("item%v", i)
		locYThis := locY + locYHeight*float64(i)
		requests = append(requests, su.TextBoxRequestsSimple(
			pageId, elementId, itemText, fgColor, bgColor,
			boxWidth, boxHeight, locX, locYThis)...)
	}

GivenCanMinX,CanMaxX
Want Item X

	SrcCan - MinX,MaxX (int64)
	Item   - MinX,MaxX (int64)
	Want   - MinX,MaxX (float64)
*/

type Location struct {
	SrcAllMinX int64
	SrcAllMaxX int64
	SrcAllWdtX int64
	SrcBoxMinX int64
	SrcBoxMaxX int64
	SrcBoxWdtX int64
	SrcPctWdtX float64
	OutAllMinX float64
	OutAllMaxX float64
	OutAllWdtX float64
	OutBoxMinX float64
	OutBoxMaxX float64
	OutBoxWdtX float64
	BoxOutPctX float64
}

func GoogleSlideDrawRoadmap(pageId string, srcCan roadmap.Canvas, outCan SlideCanvasInfo) ([]*slides.Request, error) {
	requests := []*slides.Request{}
	err := srcCan.InflateItems()
	if err != nil {
		return requests, err
	}
	srcCan.BuildRows()

	idx := 0
	rowYWatermark := outCan.Canvas.MinY

	for i, row := range srcCan.Rows {
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

			loc := Location{
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

			fmtutil.PrintJSON(loc)
			if loc.OutBoxMaxX > loc.OutAllMaxX {
				panic("C")
			} else if loc.OutBoxMinX < loc.OutAllMinX {
				panic("D")
			}
			//panic("Z")
			elementId := fmt.Sprintf("AutoBox%03d", idx)
			requests = append(requests, su.TextBoxRequestsSimple(
				pageId, elementId, el.Name, outCan.BoxFgColor, outCan.BoxBgColor,
				loc.OutBoxWdtX, outCan.BoxHeight, loc.OutBoxMinX, rowYWatermark)...)
			idx++
			//break
			if i > 3 {
				//break
			}
		}
		rowYWatermark += outCan.BoxHeight + outCan.BoxMarginBottom
	}

	return requests, nil
}

func main() {
	forceNewToken := true

	featuresPath := "/Users/john.wang/jwdev/JGo/gopath/src/github.com/grokify/go-aha/examples/get_features_full/_features.json"

	//featuersMap := map[string]aha.Feature{}
	filterArr := []string{"rmglip", "rmcc", "rmcpaas", "rmeco", "rmreq"}
	filterMap := map[string]int{"rmglip": 1, "rmcc": 1, "rmcpaas": 1, "rmeco": 1, "rmreq": 1}

	featuresMap, err := ReadFeatures(featuresPath)
	if err != nil {
		panic(err)
	}
	fmtutil.PrintJSON(featuresMap)

	featuresMap2 := map[string]map[string]aha.Feature{}

FEATS:
	for id, feat := range featuresMap {
		for _, tagTry := range feat.Tags {
			tagTry = strings.ToLower(tagTry)
			if _, ok := filterMap[tagTry]; ok {
				if _, ok2 := featuresMap2[tagTry]; !ok2 {
					featuresMap2[tagTry] = map[string]aha.Feature{}
				}
				featuresMap2[tagTry][id] = feat
				continue FEATS
			}
		}
	}

	featureMapsArr := []map[string]aha.Feature{}
	for _, filter := range filterArr {
		if fMap, ok := featuresMap2[filter]; ok {
			featureMapsArr = append(featureMapsArr, fMap)
		}
	}

	fmtutil.PrintJSON(featuresMap2)
	//panic("ZZ")

	itemsRM := []roadmap.Item{}

	qtrMin := int32(20174)
	qtrMax := int32(20184)

	//SetMinMaxQuarter(qtrMin, qtrMax int32) error {

	can := roadmap.Canvas{}

	err = can.SetMinMaxQuarter(qtrMin, qtrMax)
	if err != nil {
		panic(err)
	}
	can.SetRange(200)

	rng, err := can.Range.CellRange()
	if err != nil {
		panic(err)
	}
	//panic("AA")
	fmtutil.PrintJSON(can)
	fmt.Printf("Cell Range %v", rng)
	//panic("A")
	for _, feat := range featuresMap {
		item := roadmap.Item{
			MinTime: tu.ParseOrZero(tu.RFC3339FullDate, feat.StartDate),
			MaxTime: tu.ParseOrZero(tu.RFC3339FullDate, feat.DueDate),
			Name:    feat.Name,
		}
		itemsRM = append(itemsRM, item)
		can.AddItem(item)
	}
	//fmtutil.PrintJSON(itemsRM)
	err = can.InflateItems()
	if err != nil {
		panic(err)
	}
	can.BuildRows()
	fmtutil.PrintJSON(can)
	fmt.Println(len(can.Rows))
	//panic("Z")

	err = godotenv.Load()
	if err != nil {
		panic(err)
	}

	client, err := NewClient(forceNewToken)
	if err != nil {
		log.Fatal("Unable to get Client")
	}

	srv, err := slides.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Slides Client %v", err)
	}

	psv := slides.NewPresentationsService(srv)

	t := time.Now().UTC()
	slideName := fmt.Sprintf("GOLANG TEST PRES %v\n", t.Format(time.RFC3339))
	fmt.Printf("Slide Name: %v", slideName)

	pres := &slides.Presentation{Title: slideName}
	res, err := psv.Create(pres).Do()
	if err != nil {
		panic(err)
	}

	fmt.Printf("CREATED Presentation with Id %v\n", res.PresentationId)

	for i, slide := range res.Slides {
		fmt.Printf("- Slide #%d id %v contains %d elements.\n", (i + 1),
			slide.ObjectId,
			len(slide.PageElements))
	}

	pageId := res.Slides[0].ObjectId
	requests := []*slides.Request{}

	fgColor, err := su.ParseRgbColorHex("#ffffff")
	if err != nil {
		panic(err)
	}
	bgColor, err := su.ParseRgbColorHex("#4688f1")
	if err != nil {
		panic(err)
	}

	outCan := CanvasFloat64{
		MinX: 150.0,
		MinY: 70.0,
		MaxX: 700.0,
		MaxY: 500.0,
	}

	//items := []string{"Item #1", "Item #2"}
	//locX := 150.0
	//locY := 50.0
	//boxWidth := 550.0
	//boxHeight := 25.0

	slideCanvas := SlideCanvasInfo{
		BoxFgColor:      fgColor,
		BoxBgColor:      bgColor,
		BoxHeight:       25.0,
		BoxMarginBottom: 5.0,
		Canvas:          outCan,
	}

	requestsRoadmap, err := GoogleSlideDrawRoadmap(pageId, can, slideCanvas)
	if err != nil {
		panic(err)
	}

	requests = requestsRoadmap
	/*

		locYHeight := boxHeight + 5.0
		for i, itemText := range items {
			elementId := fmt.Sprintf("item%v", i)
			locYThis := locY + locYHeight*float64(i)
			requests = append(requests, su.TextBoxRequestsSimple(
				pageId, elementId, itemText, fgColor, bgColor,
				boxWidth, boxHeight, locX, locYThis)...)
		}
	*/

	breq := &slides.BatchUpdatePresentationRequest{
		Requests: requests,
	}

	resu, err := psv.BatchUpdate(res.PresentationId, breq).Do()
	if err != nil {
		panic(err)
	}
	fmt.Println(resu.PresentationId)

	fmt.Println("DONE")
}
