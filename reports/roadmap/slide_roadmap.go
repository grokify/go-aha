package roadmap

import (
	"fmt"
	"net/http"
	"time"

	au "github.com/grokify/go-aha/v2/ahautil"
	su "github.com/grokify/gogoogle/slidesutil/v1"
	"github.com/grokify/mogo/errors/errorsutil"
	gs "google.golang.org/api/slides/v1"
)

func CreateRoadmapSlide(googleClient *http.Client, presentationID string, roadmapConfig RoadmapConfig, featureSet *au.FeatureSet) (*gs.BatchUpdatePresentationResponse, error) {
	gss, err := su.NewGoogleSlidesService(googleClient)
	if err != nil {
		return nil, errorsutil.Wrap(err, "CreateRoadmapSlide - slidesutil.NewGoogleSlidesService()")
	}
	psv := gss.PresentationsService

	dt := time.Now().UTC()
	pageId := ""
	newPageId := fmt.Sprintf("Roadmap-%v", dt.Unix())
	titleId := fmt.Sprintf("Roadmap-%v-Title", dt.Unix())

	if 1 == 1 {
		// Create slide with title
		// https://developers.google.com/slides/samples/slides#create_a_new_slide_and_modify_placeholders
		requests := []*gs.Request{
			/*{
				DeleteObject: &gs.DeleteObjectRequest{ObjectId: pageId},
			},*/
			{
				CreateSlide: &gs.CreateSlideRequest{
					ObjectId: newPageId,
					SlideLayoutReference: &gs.LayoutReference{
						PredefinedLayout: "TITLE_ONLY",
					},
					PlaceholderIdMappings: []*gs.LayoutPlaceholderIdMapping{
						{
							LayoutPlaceholder: &gs.Placeholder{
								Type:  "TITLE",
								Index: 0,
							},
							ObjectId: titleId,
						},
					},
				},
			},
			{
				InsertText: &gs.InsertTextRequest{
					ObjectId: titleId,
					Text:     roadmapConfig.Title,
				},
			},
		}
		breq := &gs.BatchUpdatePresentationRequest{
			Requests: requests,
		}
		_, err := psv.BatchUpdate(presentationID, breq).Do() // resu
		if err != nil {
			return nil, errorsutil.Wrap(err, "CreateRoadmapSlide - psv.BatchUpdate(res.PresentationId, breq).Do()")
		}
		pageId = newPageId
		//fmt.Println(resu.PresentationId)
	}

	requests := RoadmapTextBoxRequests(
		roadmapConfig,
		featureSet,
		pageId)

	// Add Disclaimer
	if len(roadmapConfig.DisclaimerText) > 0 {
		disclaimer := internalDisclaimer(pageId, roadmapConfig)
		disclaimerReqs, err := disclaimer.Requests()
		if err != nil {
			return nil, errorsutil.Wrap(err, "CreateRoadmapSlide - disclaimer.Requests()")
		}
		requests = append(requests, disclaimerReqs...)
	}
	breq := &gs.BatchUpdatePresentationRequest{
		Requests: requests,
	}

	resu, err := psv.BatchUpdate(presentationID, breq).Do()
	if err != nil {
		return nil, errorsutil.Wrap(err, "CreateRoadmapSlide - psv.BatchUpdate(res.PresentationId, breq).Do()")
	}
	// fmt.Println(resu.PresentationId)
	return resu, nil
}

func internalDisclaimer(pageID string, roadmapConfig RoadmapConfig) su.CreateShapeTextBoxRequestInfo {
	uid := getUniqueId()
	return su.CreateShapeTextBoxRequestInfo{
		PageID:             pageID,
		ObjectID:           "disclaimer_heading-" + uid,
		Width:              140.0,
		Height:             40.0,
		DimensionUnit:      "PT",
		LocationX:          580.0,
		LocationY:          0.0,
		LocationUnit:       "PT",
		FontSize:           10.0,
		FontSizeUnit:       "PT",
		BackgroundColorHex: roadmapConfig.RoadmapFormatting.Disclaimer.DefaultBackgroundColorHex,
		ForegroundColorHex: roadmapConfig.RoadmapFormatting.Disclaimer.DefaultForegroundColorHex,
		Text:               roadmapConfig.DisclaimerText,
		ParagraphAlignment: "CENTER",
	}
}
