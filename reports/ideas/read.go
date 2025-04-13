package ideas

import (
	"strings"
	"time"

	"github.com/grokify/gocharts/v2/data/table"
	"github.com/grokify/mogo/net/urlutil"
	"github.com/grokify/mogo/type/stringsutil"
)

func ParseXLSX(filename, submittedPortalURL string) (Ideas, error) {
	ideas := Ideas{}
	tbl, err := table.ReadTableXLSXIndexFile(filename, 0, 1, true)
	if err != nil {
		return ideas, err
	}
	for _, row := range tbl.Rows {
		if try := stringsutil.SliceCondenseSpace(row, true, true); len(try) == 0 {
			continue
		}
		idea := rowToIdea(tbl.Columns, row)
		if strings.TrimSpace(idea.SubmittedPortalURL) == "" && submittedPortalURL != "" {
			idea.SubmittedPortalURL = submittedPortalURL
		}
		if strings.TrimSpace(idea.URLInSubmitedPortal) == "" && submittedPortalURL != "" && idea.Reference != "" {
			idea.URLInSubmitedPortal = urlutil.JoinAbsolute(submittedPortalURL, "ideas", idea.Reference)
		}
		ideas = append(ideas, idea)
	}
	return ideas, nil
}

func rowToIdea(cols table.Columns, row []string) Idea {
	return Idea{
		Categories:           stringsutil.SplitCSVUnescaped(cols.MustCellString(IdeaCategories, row)),
		CreatedBy:            cols.MustCellString(IdeaCreatedBy, row),
		CreatedByEmailDomain: cols.MustCellString(IdeaCreatedByEmailDomain, row),
		Description:          cols.MustCellString(IdeaDescription, row),
		ID:                   cols.MustCellString(IdeaID, row),
		LastActiveDate:       cols.MustCellTime(IdeaLastActiveDate, time.DateTime, row),
		LastStatusChangeDate: cols.MustCellTime(IdeaLastStatusChangeDate, time.DateTime, row),
		LastVoteDate:         cols.MustCellTime(IdeaLastVoteDate, time.DateTime, row),
		Name:                 cols.MustCellString(IdeaName, row),
		Reference:            cols.MustCellString(IdeaReference, row),
		Status:               cols.MustCellString(IdeaStatus, row),
		StatusCategory:       cols.MustCellString(IdeaStatusCategory, row),
		SubmittedPortalURL:   cols.MustCellString(IdeaSubmittedPortalURL, row),
		URLInSubmitedPortal:  cols.MustCellString(IdeaURLInSubmitedPortal, row),
		VoterEmailDomains:    strings.Split(cols.MustCellString(IdeaVoterEmailDomains, row), ","),
		Votes:                cols.MustCellIntOrDefault(IdeaVotes, row, 0),
	}
}
