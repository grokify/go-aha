package ideas

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/grokify/go-aha/v3/oag7/aha"
)

type IdeaFeatureReleaseData struct {
	Idea    *aha.Idea
	Feature *aha.Feature
	Release *aha.Release
}

func (data *IdeaFeatureReleaseData) Metadata() *IdeaFeatureReleaseMetadata {
	out := &IdeaFeatureReleaseMetadata{}
	if data.Idea != nil {
		out.IdeaID = data.Idea.Id
		out.IdeaName = data.Idea.Name
		out.IdeaRefNum = strings.TrimSpace(data.Idea.ReferenceNum)
	}
	if data.Feature != nil {
		out.FeatureID = data.Feature.Id
		out.FeatureName = data.Feature.Name
		out.FeatureRefNum = strings.TrimSpace(data.Feature.ReferenceNum)
	}
	if data.Feature.Release != nil {
		if data.Feature.Release.Id != nil {
			out.ReleaseID = *data.Feature.Release.Id
		}
		if data.Feature.Release.Name != nil {
			out.ReleaseName = *data.Feature.Release.Name
		}
		if data.Feature.Release.ReferenceNum != nil {
			out.ReleaseRefNum = strings.TrimSpace(*data.Feature.Release.ReferenceNum)
		}
		if data.Feature.Release.ReleaseDate != nil {
			out.ReleaseDate = *data.Feature.Release.ReleaseDate
		}
	}
	return out
}

type IdeaFeatureReleaseMetadata struct {
	IdeaID        string
	IdeaName      string
	IdeaRefNum    string
	FeatureID     string
	FeatureName   string
	FeatureRefNum string
	ReleaseID     string
	ReleaseName   string
	ReleaseRefNum string
	ReleaseDate   string
}

func QueryIdeaFeatureReleaseData(client *aha.APIClient, ideaID string) (*IdeaFeatureReleaseData, error) {
	if client == nil {
		return nil, errors.New("aha api client cannot be nil")
	}

	ideaID = strings.TrimSpace(ideaID)
	if ideaID == "" {
		return nil, errors.New("idea_id cannot be empty")
	}

	ideaResp, httpResp, err := client.IdeasAPI.GetIdeaExecute(
		client.IdeasAPI.GetIdea(context.Background(), ideaID),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get idea: %w", err)
	}
	if httpResp.StatusCode > 299 {
		return nil, fmt.Errorf("api status code error getting idea (%d)", httpResp.StatusCode)
	}

	result := &IdeaFeatureReleaseData{
		Idea: ideaResp.Idea,
	}

	if ideaResp.Idea.Feature == nil {
		return result, nil
	}

	featureID := *ideaResp.Idea.Feature.Id
	if featureID == "" {
		return result, nil
	}

	featureResp, httpResp, err := client.FeaturesAPI.GetFeatureExecute(
		client.FeaturesAPI.GetFeature(context.Background(), featureID),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get feature: %w", err)
	}
	if httpResp.StatusCode > 299 {
		return nil, fmt.Errorf("api status code error getting feature (%d)", httpResp.StatusCode)
	}

	result.Feature = featureResp.Feature

	result.Release = featureResp.Feature.Release

	return result, nil
}

func QueryIdeaFeatureReleaseMetadata(client *aha.APIClient, ideaID string) (*IdeaFeatureReleaseMetadata, error) {
	if client == nil {
		return nil, errors.New("aha api client cannot be nil")
	}

	ideaID = strings.TrimSpace(ideaID)
	if ideaID == "" {
		return nil, errors.New("idea_id cannot be empty")
	}

	ideaResp, httpResp, err := client.IdeasAPI.GetIdeaExecute(
		client.IdeasAPI.GetIdea(context.Background(), ideaID),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get idea: %w", err)
	}
	if httpResp.StatusCode > 299 {
		return nil, fmt.Errorf("api status code error getting idea (%d)", httpResp.StatusCode)
	}

	result := &IdeaFeatureReleaseMetadata{
		IdeaID:     ideaResp.Idea.Id,
		IdeaName:   ideaResp.Idea.Name,
		IdeaRefNum: strings.TrimSpace(ideaResp.Idea.ReferenceNum),
	}

	if ideaResp.Idea.Feature == nil {
		return result, nil
	}

	featureID := *ideaResp.Idea.Feature.Id
	if featureID == "" {
		return result, nil
	}

	featureResp, httpResp, err := client.FeaturesAPI.GetFeatureExecute(
		client.FeaturesAPI.GetFeature(context.Background(), featureID),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get feature: %w", err)
	}
	if httpResp.StatusCode > 299 {
		return nil, fmt.Errorf("api status code error getting feature (%d)", httpResp.StatusCode)
	}

	result.FeatureID = featureResp.Feature.Id
	result.FeatureName = featureResp.Feature.Name
	result.FeatureRefNum = strings.TrimSpace(featureResp.Feature.ReferenceNum)

	if featureResp.Feature.Release != nil {
		if featureResp.Feature.Release.Id != nil {
			result.ReleaseID = *featureResp.Feature.Release.Id
		}
		if featureResp.Feature.Release.Name != nil {
			result.ReleaseName = *featureResp.Feature.Release.Name
		}
		if featureResp.Feature.Release.ReferenceNum != nil {
			result.ReleaseRefNum = strings.TrimSpace(*featureResp.Feature.Release.ReferenceNum)
		}
		if featureResp.Feature.Release.ReleaseDate != nil {
			result.ReleaseDate = *featureResp.Feature.Release.ReleaseDate
		}
	}

	return result, nil
}
