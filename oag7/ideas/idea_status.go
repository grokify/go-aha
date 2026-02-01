package ideas

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/grokify/gocharts/v2/data/table"
	"github.com/grokify/mogo/net/urlutil"
	"github.com/grokify/mogo/text/markdown"
	"github.com/grokify/mogo/type/slicesutil"

	"github.com/grokify/go-aha/v3/oag7/aha"
	"github.com/grokify/go-aha/v3/oag7/features"
	"github.com/grokify/go-aha/v3/reports/ideas"
)

func GetIdeaStatusSet(clt *aha.APIClient, ideaIDs []string, ideasPortalURL, ahaAdminURL string) (*IdeaStatusSet, error) {
	out := NewIdeaStatusSet()
	for _, ideaID := range ideaIDs {
		if id, err := GetIdeaStatus(clt, ideaID, ideasPortalURL, ahaAdminURL); err != nil {
			return nil, err
		} else if id != nil {
			out.Data[id.IdeaRefNum] = *id
		}
	}
	return &out, nil
}

func GetIdeaStatuses(clt *aha.APIClient, ideaIDs []string, ideasPortalURL, ahaAdminURL string) (*IdeaStatuses, error) {
	out := &IdeaStatuses{}
	for _, ideaID := range ideaIDs {
		if is, err := GetIdeaStatus(clt, ideaID, ideasPortalURL, ahaAdminURL); err != nil {
			return nil, err
		} else if is != nil {
			*out = append(*out, *is)
		}
	}
	return out, nil
}

func GetIdeaStatus(clt *aha.APIClient, ideaID, ideasPortalURL, ahaAdminURL string) (*IdeaStatus, error) {
	if clt == nil {
		return nil, errors.New("aha api client cannot be nil")
	}
	ideaID = strings.TrimSpace(ideaID)
	ideasPortalURL = strings.TrimSpace(ideasPortalURL)
	ahaAdminURL = strings.TrimSpace(ahaAdminURL)
	if ideaID == "" {
		return nil, errors.New("idea_id cannot be nil")
	}

	info, resp, err := clt.IdeasAPI.GetIdeaExecute(
		clt.IdeasAPI.GetIdea(context.Background(), ideaID),
	)
	if err != nil {
		return nil, err
	} else if resp.StatusCode > 299 {
		return nil, fmt.Errorf("api status code error (%d)", resp.StatusCode)
	}

	out := IdeaStatus{
		IdeaCreatedAt: info.Idea.GetCreatedAt(),
		IdeaID:        info.Idea.Id,
		IdeaName:      info.Idea.Name,
		IdeaRefNum:    strings.TrimSpace(info.Idea.ReferenceNum),
		IdeaVotes:     int64(info.Idea.GetVotes())}

	if out.IdeaRefNum != "" && ideasPortalURL != "" {
		out.IdeaURLPortal = urlutil.JoinAbsolute(ideasPortalURL, "ideas", out.IdeaRefNum)
	}

	for _, cat := range info.Idea.Categories {
		if name := strings.TrimSpace(cat.Name); name != "" {
			out.IdeaCategories = append(out.IdeaCategories, name)
		}
	}
	if len(out.IdeaCategories) > 1 {
		out.IdeaCategories = slicesutil.Dedupe(out.IdeaCategories)
		sort.Strings(out.IdeaCategories)
	}

	if info.Idea.WorkflowStatus != nil {
		out.IdeaStatusName = *info.Idea.WorkflowStatus.Name
	}

	if info.Idea.Feature != nil {
		promotedFeatureID := *info.Idea.Feature.Id

		finfo, resp, err := clt.FeaturesAPI.GetFeatureExecute(
			clt.FeaturesAPI.GetFeature(context.Background(), promotedFeatureID))
		if err != nil {
			return nil, err
		} else if resp.StatusCode > 299 {
			return nil, fmt.Errorf("api status code error (%d)", resp.StatusCode)
		}
		out.FeatureID = finfo.Feature.Id
		out.FeatureName = finfo.Feature.Name
		out.FeatureRefNum = strings.TrimSpace(finfo.Feature.ReferenceNum)
		if out.FeatureRefNum != "" && ahaAdminURL != "" {
			out.FeatureURLAdmin = urlutil.JoinAbsolute(ahaAdminURL, "features", out.FeatureRefNum)
		}

		if finfo.Feature.WorkflowStatus != nil {
			out.FeatureStatusName = *finfo.Feature.WorkflowStatus.Name
		}
		if finfo.Feature.Release != nil {
			out.ReleaseName = *finfo.Feature.Release.Name
			out.ReleaseDate = *finfo.Feature.Release.ReleaseDate
		}
		fm := features.Feature(*finfo.Feature)
		out.FeatureJiraKey = fm.JiraKey()
		out.FeatureJiraURL = fm.JiraURL()
	}

	return &out, nil
}

type IdeaStatusSet struct {
	Data map[string]IdeaStatus
}

func NewIdeaStatusSet() IdeaStatusSet {
	return IdeaStatusSet{Data: map[string]IdeaStatus{}}
}

type IdeaStatuses []IdeaStatus

func (set IdeaStatusSet) Table() *table.Table {
	t := table.NewTable("")
	t.FormatMap = map[int]string{
		0: table.FormatURL,
		2: table.FormatURL,
		4: table.FormatInt,
		5: table.FormatDate,
		6: table.FormatURL,
		7: table.FormatURL,
		9: table.FormatDate}
	t.Columns = table.Columns{
		ideas.IdeaReference,
		ideas.IdeaCategories,
		ideas.IdeaName,
		ideas.IdeaStatus,
		ideas.IdeaVotes,
		ideas.IdeaCreatedDate,
		"Feature reference",
		"Jira reference",
		"Release name",
		"Release date"}
	for _, id := range set.Data {
		row := []string{
			markdown.Linkify(id.IdeaURLPortal, id.IdeaRefNum),
			strings.Join(id.IdeaCategories, "; "),
			markdown.Linkify(id.IdeaURLPortal, id.IdeaName),
			id.IdeaStatusName,
			strconv.Itoa(int(id.IdeaVotes)),
			id.IdeaCreatedAt.Format(time.DateOnly),
			id.FeatureRefNum,
			id.JiraLinkMarkdown(),
			id.ReleaseName,
			id.ReleaseDate}
		t.Rows = append(t.Rows, row)
	}
	return &t
}

type IdeaStatus struct {
	IdeaID            string
	IdeaCategories    []string
	IdeaCreatedAt     time.Time
	IdeaName          string
	IdeaRefNum        string
	IdeaStatusName    string
	IdeaURLPortal     string
	IdeaVotes         int64
	FeatureID         string
	FeatureRefNum     string
	FeatureName       string
	FeatureStatusName string
	FeatureURLAdmin   string
	FeatureJiraKey    string
	FeatureJiraURL    string
	ReleaseName       string
	ReleaseDate       string
}

func (is IdeaStatus) JiraLinkMarkdown() string {
	return markdown.Linkify(is.FeatureJiraURL, is.FeatureJiraKey)
}
