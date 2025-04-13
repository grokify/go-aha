package ideas

import (
	"sort"
	"strconv"
	"strings"

	"github.com/grokify/gocharts/v2/data/histogram"
	"github.com/grokify/gocharts/v2/data/table"
)

type Ideas []Idea

func (ideas Ideas) HistogramCategories() *histogram.Histogram {
	h := histogram.NewHistogram("")
	for _, idea := range ideas {
		for _, cat := range idea.Categories {
			if cat = strings.TrimSpace(cat); cat != "" {
				h.Add(strings.TrimSpace(cat), 1)
			}
		}
	}
	return h
}

func (ideas Ideas) HistogramStatus() *histogram.Histogram {
	h := histogram.NewHistogram("")
	for _, idea := range ideas {
		h.Add(strings.TrimSpace(idea.Status), 1)
	}
	return h
}

func (ideas Ideas) SortVotes(desc bool) {
	if desc {
		sort.Slice(ideas, func(i, j int) bool {
			return ideas[i].Votes > ideas[j].Votes
		})
	} else {
		sort.Slice(ideas, func(i, j int) bool {
			return ideas[i].Votes < ideas[j].Votes
		})
	}
}

// Percentile returns the precentile of the vote count supplied.
func (ideas Ideas) Percentile(voteCount int) (float32, error) {
	h := histogram.NewHistogram("")
	for _, id := range ideas {
		votes := id.Votes
		h.Add(strconv.Itoa(votes), 1)
	}
	return h.Percentile(voteCount)
}

func (ideas Ideas) HistogramStatuses() map[string]int {
	out := map[string]int{}
	for _, idea := range ideas {
		out[idea.StatusCategory]++
	}
	return out
}

func (ideas Ideas) Table(fields []string, fmtMap map[int]string) *table.Table {
	tbl := table.NewTable("")
	tbl.Columns = fields
	tbl.FormatMap = fmtMap
	for i, f := range fields {
		switch f {
		case IdeaURLInSubmitedPortal:
			tbl.FormatMap[i] = table.FormatURL
		case IdeaVotes:
			tbl.FormatMap[i] = table.FormatInt
		}
	}
	for i, id := range ideas {
		tbl.Rows = append(tbl.Rows, id.ValueStrings(fields, i+1))
	}
	return &tbl
}
