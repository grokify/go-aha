package ideasreports

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/grokify/gocharts/v2/data/table"
	"github.com/grokify/gocharts/v2/data/table/excelizeutil"
	"github.com/grokify/mogo/text/markdown"
	"github.com/grokify/mogo/time/timeutil"
	"github.com/grokify/mogo/type/maputil"
	"github.com/grokify/mogo/type/stringsutil"
)

// IdeasSet is meant to read and produce reports for customers that interact with the Aha
// Ideas Portal.
type IdeasSet struct {
	IdeasMap map[int]Idea
}

func (is IdeasSet) Table() *table.Table {
	t := table.NewTable("")
	t.Columns = []string{
		HeaderIdeaName,
		"Idea created by",
		"Idea voted on",
		HeaderIdeaVotes,
		HeaderIdeaStatus,
		HeaderIdeaID,
		HeaderIdeaURL,
		HeaderIdeaCreatedDate,
		HeaderIdeaLastStatusChangeDate,
		HeaderIdeaLastVoteDate,
	}
	t.FormatMap = map[int]string{
		-1: table.FormatString,
		0:  table.FormatURL,
		3:  table.FormatInt,
		5:  table.FormatInt,
		6:  table.FormatURL,
		7:  table.FormatDate,
		8:  table.FormatDate,
		9:  table.FormatDate,
	}
	keys := maputil.Keys(is.IdeasMap)
	for _, key := range keys {
		idea, ok := is.IdeasMap[key]
		if !ok {
			continue
		}
		row := []string{
			markdown.Linkify(idea.URL, idea.Name),
			BoolToYN(idea.Created),
			BoolToYN(idea.Voted),
			strconv.Itoa(idea.Votes),
			idea.Status,
			strconv.Itoa(idea.ID),
			idea.URL,
			dtFormatOrDefault(idea.CreatedDate, time.RFC3339, ""),
			dtFormatOrDefault(idea.LastStatusChangeDate, time.RFC3339, ""),
			dtFormatOrDefault(idea.LastVoteDate, time.RFC3339, ""),
		}
		t.Rows = append(t.Rows, row)
	}
	return &t
}

func dtFormatOrDefault(dt *time.Time, layout, def string) string {
	if dt == nil || dt.IsZero() {
		return def
	} else {
		return dt.Format(layout)
	}
}

func BoolToYN(b bool) string {
	if b {
		return "Y"
	} else {
		return "N"
	}
}

func (is IdeasSet) WriteXLSX(filename string) error {
	t := is.Table()
	return t.WriteXLSX(filename, "Ideas")
}

type Ideas []Idea

type Idea struct {
	ID                   int
	Name                 string
	URL                  string
	Votes                int
	LastStatusChangeDate *time.Time
	CreatedDate          *time.Time
	LastVoteDate         *time.Time
	Status               string
	Created              bool
	Voted                bool
}

const (
	HeaderIdeaID    = "Idea ID"
	HeaderIdeaVotes = "Idea votes"

	HeaderIdeaCreatedDate          = "Idea created date"
	HeaderIdeaLastStatusChangeDate = "Idea last status change date"
	HeaderIdeaLastVoteDate         = "Idea last vote date"

	HeaderIdeaCreatedByEmailDomain = "Idea created by email domain"
	HeaderIdeaName                 = "Idea name"
	HeaderIdeaStatus               = "Idea status"
	HeaderIdeaURL                  = "Idea URL"
	HeaderIdeaVoterEmailDomains    = "Idea voter email domains"
)

func ParseFilesXLSX(filenames []string, emailDomain string) (*IdeasSet, error) {
	is := IdeasSet{IdeasMap: map[int]Idea{}}
	for _, filename := range filenames {
		if f, err := excelizeutil.ReadFile(filename); err != nil {
			return nil, err
		} else if cols, rows, err := f.TableDataIndex(0, 1, true, false); err != nil {
			return nil, err
		} else {
			t := table.NewTable("")
			t.Columns = cols
			t.Rows = rows
			if isx, err := ParseTable(&t, emailDomain); err != nil {
				return nil, err
			} else {
				for k, v := range isx.IdeasMap {
					if _, ok := is.IdeasMap[k]; !ok {
						is.IdeasMap[k] = v
					}
				}
			}
		}
	}
	return &is, nil
}

func ParseTable(t *table.Table, emailDomain string) (*IdeasSet, error) {
	if t == nil {
		return nil, table.ErrTableCannotBeNil
	}
	is := IdeasSet{IdeasMap: map[int]Idea{}}
	for _, row := range t.Rows {
		if rowTry := stringsutil.SliceCondenseSpace(row, true, true); len(rowTry) == 0 {
			continue
		}
		idea := parseRow(t.Columns, row, emailDomain)
		if idea.ID != 0 {
			is.IdeasMap[idea.ID] = idea
		} else {
			return nil, errors.New("idea cannot have 0 id")
		}
	}
	return &is, nil
}

func parseRow(cols table.Columns, row []string, emailDomain string) Idea {
	idea := Idea{
		ID:     cols.MustCellIntOrDefault(HeaderIdeaID, row, 0),
		Votes:  cols.MustCellIntOrDefault(HeaderIdeaVotes, row, 0),
		Name:   cols.MustCellString(HeaderIdeaName, row),
		Status: cols.MustCellString(HeaderIdeaStatus, row),
		URL:    cols.MustCellString(HeaderIdeaURL, row),
	}
	if dt := cols.MustCellTime(HeaderIdeaCreatedDate, timeutil.SQLTimestamp, row); dt != nil {
		idea.CreatedDate = dt
	}
	if dt := cols.MustCellTime(HeaderIdeaLastStatusChangeDate, timeutil.SQLTimestamp, row); dt != nil {
		idea.LastStatusChangeDate = dt
	}
	if dt := cols.MustCellTime(HeaderIdeaLastVoteDate, timeutil.SQLTimestamp, row); dt != nil {
		idea.LastVoteDate = dt
	}
	emailDomain = strings.TrimSpace(emailDomain)
	rx := regexp.MustCompile(`\b` + regexp.QuoteMeta(emailDomain) + `\b`)
	if rx.MatchString(cols.MustCellString(HeaderIdeaCreatedByEmailDomain, row)) {
		idea.Created = true
	}
	if rx.MatchString(cols.MustCellString(HeaderIdeaVoterEmailDomains, row)) {
		idea.Voted = true
	}
	return idea
}
