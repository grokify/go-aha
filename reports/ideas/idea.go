package ideas

import (
	"strconv"
	"strings"
	"time"

	"github.com/grokify/mogo/text/markdown"
)

type Idea struct {
	ID                   string
	Reference            string
	Name                 string
	Status               string
	StatusCategory       string
	Votes                int
	Categories           []string
	CreatedBy            string
	CreatedByEmailDomain string
	VoterEmailDomains    []string
	CreatedDate          *time.Time
	LastActiveDate       *time.Time
	LastStatusChangeDate *time.Time
	LastVoteDate         *time.Time
	SubmittedPortalURL   string
	URLInSubmitedPortal  string
}

func (id Idea) ValueString(field string) string {
	switch field {
	case IdeaCategories:
		return strings.Join(id.Categories, "; ")
	case IdeaName:
		return id.Name
	case IdeaStatus:
		return id.Status
	case IdeaVotes:
		return strconv.Itoa(id.Votes)
	case IdeaURLInSubmitedPortal:
		return id.URLInSubmitedPortal
	case IdeaNameWithPortalURL:
		return markdown.Linkify(id.URLInSubmitedPortal, id.Name)
	default:
		return ""
	}
}

func (id Idea) ValueStrings(fields []string, rowNumber int) []string {
	var out []string
	for _, f := range fields {
		if f == RowNumber {
			out = append(out, strconv.Itoa(rowNumber))
		} else {
			out = append(out, id.ValueString(f))
		}
	}
	return out
}
