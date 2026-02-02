package ideas

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/grokify/gocharts/v2/data/table"
	"github.com/grokify/mogo/text/markdown"

	"github.com/grokify/go-aha/v3/oag7/aha"
	"github.com/grokify/go-aha/v3/oag7/features"
)

// Column name constants for idea-feature-release reports.
const (
	ColIdeaID           = "Idea ID"
	ColIdeaRefNum       = "Idea Ref"
	ColIdeaName         = "Idea Name"
	ColIdeaStatus       = "Idea Status"
	ColIdeaVotes        = "Idea Votes"
	ColIdeaCategories   = "Idea Categories"
	ColIdeaCreatedAt    = "Idea Created"
	ColIdeaUpdatedAt    = "Idea Updated"
	ColHasFeature       = "Has Feature"
	ColFeatureID        = "Feature ID"
	ColFeatureRefNum    = "Feature Ref"
	ColFeatureName      = "Feature Name"
	ColFeatureStatus    = "Feature Status"
	ColFeatureStartDate = "Feature Start"
	ColFeatureDueDate   = "Feature Due"
	ColFeatureURL       = "Feature URL"
	ColHasRelease       = "Has Release"
	ColReleaseID        = "Release ID"
	ColReleaseRefNum    = "Release Ref"
	ColReleaseName      = "Release Name"
	ColReleaseDate      = "Release Date"
	ColReleaseReleased  = "Released"
	ColReleaseURL       = "Release URL"
	ColJiraKey          = "Jira Key"
	ColJiraURL          = "Jira URL"
)

// IdeaFeatureReport represents a comprehensive report row combining
// idea, feature, and release information.
type IdeaFeatureReport struct {
	// Idea fields
	IdeaID         string
	IdeaRefNum     string
	IdeaName       string
	IdeaStatus     string
	IdeaVotes      int64
	IdeaCategories []string
	IdeaCreatedAt  time.Time
	IdeaUpdatedAt  time.Time
	IdeaURL        string

	// Feature fields (populated if idea has been promoted to a feature)
	HasFeature       bool
	FeatureID        string
	FeatureRefNum    string
	FeatureName      string
	FeatureStatus    string
	FeatureStartDate string
	FeatureDueDate   string
	FeatureURL       string
	FeatureTags      []string
	FeatureJiraKey   string
	FeatureJiraURL   string

	// Release fields (populated if feature is assigned to a release)
	HasRelease      bool
	ReleaseID       string
	ReleaseRefNum   string
	ReleaseName     string
	ReleaseDate     string
	ReleaseReleased bool
	ReleaseURL      string
}

// IdeaFeatureReportSet holds a collection of IdeaFeatureReport items.
type IdeaFeatureReportSet struct {
	Reports []IdeaFeatureReport
	ByIdea  map[string]*IdeaFeatureReport // keyed by IdeaRefNum
}

// NewIdeaFeatureReportSet creates an empty IdeaFeatureReportSet.
func NewIdeaFeatureReportSet() *IdeaFeatureReportSet {
	return &IdeaFeatureReportSet{
		Reports: []IdeaFeatureReport{},
		ByIdea:  make(map[string]*IdeaFeatureReport),
	}
}

// Add adds a report to the set.
func (s *IdeaFeatureReportSet) Add(r IdeaFeatureReport) {
	s.Reports = append(s.Reports, r)
	s.ByIdea[r.IdeaRefNum] = &s.Reports[len(s.Reports)-1]
}

// Len returns the number of reports.
func (s *IdeaFeatureReportSet) Len() int {
	return len(s.Reports)
}

// ProgressFunc is called during operations to report progress.
// Parameters: current item number, total items, current item name/identifier.
type ProgressFunc func(current, total int, name string)

// ListIdeasRequest contains parameters for listing ideas.
type ListIdeasRequest struct {
	Query              string
	WorkflowStatus     string
	Tag                string
	UserID             string
	CreatedSince       *time.Time
	UpdatedSince       *time.Time
	Page               int32
	PerPage            int32
	Spam               bool
	FetchAll           bool         // If true, automatically paginate through all pages
	Inflate            bool         // If true, fetch full details for each idea via GetIdea endpoint
	ProgressFn         ProgressFunc // Optional callback for progress reporting (used with Inflate)
	FailOnFeatureError bool         // If true, stop processing and return error when feature fetch fails
}

// GetIdeaFeatureReports fetches ideas and their associated features/releases.
// It returns a comprehensive report set that can be exported to various formats.
// If req.FetchAll is true, it will automatically paginate through all pages.
// If req.Inflate is true, it fetches full details for each idea via GetIdea endpoint.
// If req.ProgressFn is set and Inflate is true, progress is reported during inflation.
func GetIdeaFeatureReports(ctx context.Context, client *aha.APIClient, req ListIdeasRequest) (*IdeaFeatureReportSet, error) {
	if client == nil {
		return nil, fmt.Errorf("aha api client cannot be nil")
	}

	// Phase 1: Collect all ideas from the list endpoint
	allIdeas, err := collectIdeas(ctx, client, req)
	if err != nil {
		return nil, err
	}

	// Phase 2: Process ideas (with optional inflation and progress reporting)
	reportSet := NewIdeaFeatureReportSet()
	total := len(allIdeas)

	for i, idea := range allIdeas {
		var report IdeaFeatureReport

		// If Inflate is true, fetch full idea details via GetIdea endpoint
		if req.Inflate {
			// Report progress if callback is provided
			if req.ProgressFn != nil {
				req.ProgressFn(i+1, total, idea.ReferenceNum)
			}

			inflatedIdea, err := inflateIdea(ctx, client, idea.Id)
			if err != nil {
				// Fall back to list data if inflate fails
				report = buildReportFromIdea(idea)
				report.IdeaName = fmt.Sprintf("%s (inflate error: %v)", idea.Name, err)
			} else {
				report = buildReportFromIdea(*inflatedIdea)
			}
		} else {
			report = buildReportFromIdea(idea)
		}

		// If the idea has a feature reference, fetch full feature details
		if report.HasFeature && report.FeatureID != "" {
			if err := enrichReportWithFeature(ctx, client, &report, report.FeatureID); err != nil {
				if req.FailOnFeatureError {
					return nil, fmt.Errorf("failed to get feature %s for idea %s: %w",
						report.FeatureID, report.IdeaRefNum, err)
				}
				// Log but continue - don't fail the entire report for one feature
				report.FeatureName = fmt.Sprintf("(error: %v)", err)
			}
		}

		reportSet.Add(report)
	}

	return reportSet, nil
}

// collectIdeas fetches ideas from the list endpoint with optional pagination.
func collectIdeas(ctx context.Context, client *aha.APIClient, req ListIdeasRequest) ([]aha.Idea, error) {
	var allIdeas []aha.Idea

	currentPage := req.Page
	if currentPage <= 0 {
		currentPage = 1
	}
	perPage := req.PerPage
	if perPage <= 0 {
		perPage = 30 // default page size
	}

	for {
		// Build the ideas request for this page
		ideasReq := client.IdeasAPI.ListIdeas(ctx)

		if req.Query != "" {
			ideasReq = ideasReq.Q(req.Query)
		}
		if req.WorkflowStatus != "" {
			ideasReq = ideasReq.WorkflowStatus(req.WorkflowStatus)
		}
		if req.Tag != "" {
			ideasReq = ideasReq.Tag(req.Tag)
		}
		if req.UserID != "" {
			ideasReq = ideasReq.UserId(req.UserID)
		}
		if req.CreatedSince != nil {
			ideasReq = ideasReq.CreatedSince(*req.CreatedSince)
		}
		if req.UpdatedSince != nil {
			ideasReq = ideasReq.UpdatedSince(*req.UpdatedSince)
		}
		if req.Spam {
			ideasReq = ideasReq.Spam(true)
		}

		ideasReq = ideasReq.Page(currentPage).PerPage(perPage)

		// Fetch ideas for this page
		ideasResp, resp, err := ideasReq.Execute()
		if err != nil {
			return nil, fmt.Errorf("failed to list ideas (page %d): %w", currentPage, err)
		}
		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("api error: status %d (page %d)", resp.StatusCode, currentPage)
		}

		allIdeas = append(allIdeas, ideasResp.Ideas...)

		// Check if we should continue paginating
		if !req.FetchAll {
			// Single page mode - we're done
			break
		}

		// Check pagination info to see if there are more pages
		if ideasResp.Pagination == nil {
			break
		}

		totalPages := ideasResp.Pagination.GetTotalPages()
		if int64(currentPage) >= totalPages {
			// We've fetched all pages
			break
		}

		// No ideas returned means we're done
		if len(ideasResp.Ideas) == 0 {
			break
		}

		currentPage++
	}

	return allIdeas, nil
}

// GetIdeaFeatureReport fetches a single idea and its associated feature/release.
func GetIdeaFeatureReport(ctx context.Context, client *aha.APIClient, ideaID string) (*IdeaFeatureReport, error) {
	if client == nil {
		return nil, fmt.Errorf("aha api client cannot be nil")
	}

	ideaID = strings.TrimSpace(ideaID)
	if ideaID == "" {
		return nil, fmt.Errorf("idea_id cannot be empty")
	}

	// Fetch the idea
	ideaResp, resp, err := client.IdeasAPI.GetIdea(ctx, ideaID).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get idea: %w", err)
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("api error: status %d", resp.StatusCode)
	}

	report := buildReportFromIdea(*ideaResp.Idea)

	// If the idea has a feature reference, fetch full feature details
	if ideaResp.Idea.Feature != nil && ideaResp.Idea.Feature.Id != nil {
		featureID := *ideaResp.Idea.Feature.Id
		if featureID != "" {
			if err := enrichReportWithFeature(ctx, client, &report, featureID); err != nil {
				return nil, fmt.Errorf("failed to get feature: %w", err)
			}
		}
	}

	return &report, nil
}

// inflateIdea fetches full idea details via the GetIdea endpoint.
// The list endpoint may return abbreviated data; this ensures we get all fields including categories.
func inflateIdea(ctx context.Context, client *aha.APIClient, ideaID string) (*aha.Idea, error) {
	ideaResp, resp, err := client.IdeasAPI.GetIdea(ctx, ideaID).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get idea %s: %w", ideaID, err)
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("api error getting idea %s: status %d", ideaID, resp.StatusCode)
	}
	return ideaResp.Idea, nil
}

// buildReportFromIdea creates a report from an Idea.
func buildReportFromIdea(idea aha.Idea) IdeaFeatureReport {
	report := IdeaFeatureReport{
		IdeaID:        idea.Id,
		IdeaRefNum:    strings.TrimSpace(idea.ReferenceNum),
		IdeaName:      idea.Name,
		IdeaVotes:     int64(idea.GetVotes()),
		IdeaCreatedAt: idea.GetCreatedAt(),
		IdeaUpdatedAt: idea.GetUpdatedAt(),
	}

	// Extract workflow status
	if idea.WorkflowStatus != nil && idea.WorkflowStatus.Name != nil {
		report.IdeaStatus = *idea.WorkflowStatus.Name
	}

	// Extract categories
	for _, cat := range idea.Categories {
		if name := strings.TrimSpace(cat.Name); name != "" {
			report.IdeaCategories = append(report.IdeaCategories, name)
		}
	}
	if len(report.IdeaCategories) > 1 {
		sort.Strings(report.IdeaCategories)
	}

	// Check if idea has a feature reference
	if idea.Feature != nil && idea.Feature.Id != nil && *idea.Feature.Id != "" {
		report.HasFeature = true
		report.FeatureID = idea.Feature.GetId()
		report.FeatureRefNum = idea.Feature.GetReferenceNum()
		report.FeatureName = idea.Feature.GetName()
		report.FeatureURL = idea.Feature.GetUrl()
	}

	return report
}

// enrichReportWithFeature fetches full feature details and enriches the report.
func enrichReportWithFeature(ctx context.Context, client *aha.APIClient, report *IdeaFeatureReport, featureID string) error {
	featureResp, resp, err := client.FeaturesAPI.GetFeature(ctx, featureID).Execute()
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("api error: status %d", resp.StatusCode)
	}

	feature := featureResp.Feature
	report.HasFeature = true
	report.FeatureID = feature.Id
	report.FeatureRefNum = feature.ReferenceNum
	report.FeatureName = feature.Name
	report.FeatureURL = feature.GetUrl()
	report.FeatureStartDate = feature.GetStartDate()
	report.FeatureDueDate = feature.GetDueDate()
	report.FeatureTags = feature.Tags

	// Extract feature workflow status
	if feature.WorkflowStatus != nil && feature.WorkflowStatus.Name != nil {
		report.FeatureStatus = *feature.WorkflowStatus.Name
	}

	// Extract Jira integration info
	fm := features.Feature(*feature)
	report.FeatureJiraKey = fm.JiraKey()
	report.FeatureJiraURL = fm.JiraURL()

	// Extract release info if present
	if feature.Release != nil {
		report.HasRelease = true
		report.ReleaseID = feature.Release.GetId()
		report.ReleaseRefNum = feature.Release.GetReferenceNum()
		report.ReleaseName = feature.Release.GetName()
		report.ReleaseDate = feature.Release.GetReleaseDate()
		report.ReleaseReleased = feature.Release.GetReleased()
		report.ReleaseURL = feature.Release.GetUrl()
	}

	return nil
}

// Table returns a gocharts table.Table representation of the report set.
// This can be exported to XLSX, Markdown, CSV, and other formats.
func (s *IdeaFeatureReportSet) Table() *table.Table {
	t := table.NewTable("Idea Feature Report")
	t.FormatMap = map[int]string{
		4:  table.FormatInt,  // Votes
		6:  table.FormatDate, // Idea Created
		14: table.FormatURL,  // Feature URL
		19: table.FormatDate, // Release Date
		21: table.FormatURL,  // Release URL
	}

	t.Columns = []string{
		ColIdeaRefNum,
		ColIdeaName,
		ColIdeaStatus,
		ColIdeaCategories,
		ColIdeaVotes,
		ColHasFeature,
		ColIdeaCreatedAt,
		ColFeatureRefNum,
		ColFeatureName,
		ColFeatureStatus,
		ColFeatureStartDate,
		ColFeatureDueDate,
		ColHasRelease,
		ColReleaseName,
		ColReleaseDate,
		ColReleaseReleased,
		ColJiraKey,
	}

	for _, r := range s.Reports {
		row := []string{
			r.IdeaRefNum,
			r.IdeaName,
			r.IdeaStatus,
			strings.Join(r.IdeaCategories, "; "),
			strconv.FormatInt(r.IdeaVotes, 10),
			boolToYesNo(r.HasFeature),
			formatDate(r.IdeaCreatedAt),
			r.FeatureRefNum,
			r.FeatureName,
			r.FeatureStatus,
			r.FeatureStartDate,
			r.FeatureDueDate,
			boolToYesNo(r.HasRelease),
			r.ReleaseName,
			r.ReleaseDate,
			boolToYesNo(r.ReleaseReleased),
			r.FeatureJiraKey,
		}
		t.Rows = append(t.Rows, row)
	}

	return &t
}

// TableCompact returns a compact table with essential columns only.
func (s *IdeaFeatureReportSet) TableCompact() *table.Table {
	t := table.NewTable("Idea Feature Report (Compact)")
	t.FormatMap = map[int]string{
		3: table.FormatInt, // Votes
	}

	t.Columns = []string{
		ColIdeaRefNum,
		ColIdeaName,
		ColIdeaStatus,
		ColIdeaVotes,
		ColFeatureRefNum,
		ColFeatureStatus,
		ColReleaseName,
		ColReleaseDate,
	}

	for _, r := range s.Reports {
		row := []string{
			r.IdeaRefNum,
			r.IdeaName,
			r.IdeaStatus,
			strconv.FormatInt(r.IdeaVotes, 10),
			r.FeatureRefNum,
			r.FeatureStatus,
			r.ReleaseName,
			r.ReleaseDate,
		}
		t.Rows = append(t.Rows, row)
	}

	return &t
}

// TableWithLinks returns a table with markdown links for URLs.
func (s *IdeaFeatureReportSet) TableWithLinks(ideaPortalBaseURL, featureBaseURL string) *table.Table {
	t := table.NewTable("Idea Feature Report")
	t.FormatMap = map[int]string{
		1: table.FormatURL, // Idea Name (linked)
		3: table.FormatInt, // Votes
		5: table.FormatURL, // Feature Name (linked)
		9: table.FormatURL, // Jira (linked)
	}

	t.Columns = []string{
		ColIdeaRefNum,
		ColIdeaName,
		ColIdeaStatus,
		ColIdeaVotes,
		ColFeatureRefNum,
		ColFeatureName,
		ColFeatureStatus,
		ColReleaseName,
		ColReleaseDate,
		ColJiraKey,
	}

	for _, r := range s.Reports {
		ideaLink := r.IdeaName
		if ideaPortalBaseURL != "" && r.IdeaRefNum != "" {
			ideaURL := strings.TrimRight(ideaPortalBaseURL, "/") + "/ideas/" + r.IdeaRefNum
			ideaLink = markdown.Linkify(ideaURL, r.IdeaName)
		}

		featureLink := r.FeatureName
		if featureBaseURL != "" && r.FeatureRefNum != "" {
			featureURL := strings.TrimRight(featureBaseURL, "/") + "/features/" + r.FeatureRefNum
			featureLink = markdown.Linkify(featureURL, r.FeatureName)
		}

		jiraLink := r.FeatureJiraKey
		if r.FeatureJiraURL != "" && r.FeatureJiraKey != "" {
			jiraLink = markdown.Linkify(r.FeatureJiraURL, r.FeatureJiraKey)
		}

		row := []string{
			r.IdeaRefNum,
			ideaLink,
			r.IdeaStatus,
			strconv.FormatInt(r.IdeaVotes, 10),
			r.FeatureRefNum,
			featureLink,
			r.FeatureStatus,
			r.ReleaseName,
			r.ReleaseDate,
			jiraLink,
		}
		t.Rows = append(t.Rows, row)
	}

	return &t
}

// FilterByHasFeature returns a new set containing only reports with/without features.
func (s *IdeaFeatureReportSet) FilterByHasFeature(hasFeature bool) *IdeaFeatureReportSet {
	result := NewIdeaFeatureReportSet()
	for _, r := range s.Reports {
		if r.HasFeature == hasFeature {
			result.Add(r)
		}
	}
	return result
}

// FilterByHasRelease returns a new set containing only reports with/without releases.
func (s *IdeaFeatureReportSet) FilterByHasRelease(hasRelease bool) *IdeaFeatureReportSet {
	result := NewIdeaFeatureReportSet()
	for _, r := range s.Reports {
		if r.HasRelease == hasRelease {
			result.Add(r)
		}
	}
	return result
}

// FilterByIdeaStatus returns a new set filtered by idea status.
func (s *IdeaFeatureReportSet) FilterByIdeaStatus(status string) *IdeaFeatureReportSet {
	result := NewIdeaFeatureReportSet()
	status = strings.ToLower(strings.TrimSpace(status))
	for _, r := range s.Reports {
		if strings.ToLower(r.IdeaStatus) == status {
			result.Add(r)
		}
	}
	return result
}

// FilterByFeatureStatus returns a new set filtered by feature status.
func (s *IdeaFeatureReportSet) FilterByFeatureStatus(status string) *IdeaFeatureReportSet {
	result := NewIdeaFeatureReportSet()
	status = strings.ToLower(strings.TrimSpace(status))
	for _, r := range s.Reports {
		if strings.ToLower(r.FeatureStatus) == status {
			result.Add(r)
		}
	}
	return result
}

// SortByVotes sorts reports by vote count (descending).
func (s *IdeaFeatureReportSet) SortByVotes() {
	sort.Slice(s.Reports, func(i, j int) bool {
		return s.Reports[i].IdeaVotes > s.Reports[j].IdeaVotes
	})
}

// SortByCreatedAt sorts reports by idea creation date (newest first).
func (s *IdeaFeatureReportSet) SortByCreatedAt() {
	sort.Slice(s.Reports, func(i, j int) bool {
		return s.Reports[i].IdeaCreatedAt.After(s.Reports[j].IdeaCreatedAt)
	})
}

// SortByUpdatedAt sorts reports by idea update date (most recently updated first).
func (s *IdeaFeatureReportSet) SortByUpdatedAt() {
	sort.Slice(s.Reports, func(i, j int) bool {
		return s.Reports[i].IdeaUpdatedAt.After(s.Reports[j].IdeaUpdatedAt)
	})
}

// SortBy sorts the report set by the specified field.
// Valid values: "votes" (default), "created", "updated".
// Returns an error for invalid sort options.
func (s *IdeaFeatureReportSet) SortBy(sortOrder string) error {
	switch strings.ToLower(strings.TrimSpace(sortOrder)) {
	case "created", "created_at", "createdat":
		s.SortByCreatedAt()
	case "updated", "updated_at", "updatedat":
		s.SortByUpdatedAt()
	case "votes", "popular", "":
		s.SortByVotes()
	default:
		return fmt.Errorf("invalid sort option %q: use votes, created, or updated", sortOrder)
	}
	return nil
}

// ParseBoolFilter parses common boolean string representations.
// Accepts: "yes", "true", "1" (returns true), "no", "false", "0" (returns false).
// Returns the parsed value and whether the input was valid.
func ParseBoolFilter(value string) (bool, bool) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "yes", "true", "1":
		return true, true
	case "no", "false", "0":
		return false, true
	default:
		return false, false
	}
}

// Helper functions

func boolToYesNo(b bool) string {
	if b {
		return "Yes"
	}
	return "No"
}

func formatDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.DateOnly)
}
