package ideas

import (
	"regexp"
	"slices"
	"strings"

	"github.com/grokify/mogo/type/strslices"
)

func (ideas Ideas) FilterByCategories(categoriesIncl []string, caseInsensive bool) Ideas {
	var out Ideas
	for _, idea := range ideas {
		if strslices.MatchAny(categoriesIncl, idea.Categories, caseInsensive, true) {
			out = append(out, idea)
		}
	}
	return out
}

func (ideas Ideas) FilterByEmailDomain(domain string, inclCreated, inclVoted bool) Ideas {
	var out Ideas
	domain = strings.TrimSpace(domain)
	if domain == "" {
		return ideas
	}
	for _, idea := range ideas {
		if inclCreated && idea.CreatedByEmailDomain == domain {
			out = append(out, idea)
		} else if inclVoted && slices.Index(idea.VoterEmailDomains, domain) >= 0 {
			out = append(out, idea)
		}
	}
	return out
}

func (ideas Ideas) FilterByNameDescKeyword(keyword string, inclName, inclDesc, caseInsensitive bool) Ideas {
	var out Ideas
	if caseInsensitive {
		keyword = strings.ToLower(keyword)
	}
	for _, idea := range ideas {
		if inclName {
			name := idea.Name
			if caseInsensitive {
				name = strings.ToLower(name)
			}
			if strings.Contains(name, keyword) {
				out = append(out, idea)
				continue
			}
		}
		if inclDesc {
			desc := idea.Description
			if caseInsensitive {
				desc = strings.ToLower(desc)
			}
			if strings.Contains(desc, keyword) {
				out = append(out, idea)
				continue
			}
		}
	}
	return out
}

func (ideas Ideas) FilterByNameDescRegexp(rx *regexp.Regexp, inclName, inclDesc bool) Ideas {
	var out Ideas
	if rx == nil {
		return out
	}
	for _, idea := range ideas {
		if inclName {
			if rx.MatchString(idea.Name) {
				out = append(out, idea)
				continue
			}
		}
		if inclDesc {
			if rx.MatchString(idea.Description) {
				out = append(out, idea)
				continue
			}
		}
	}
	return out
}

func (ideas Ideas) FilterByStatusCategory(incl []string, caseInsensitive bool) Ideas {
	var out Ideas
	if len(incl) == 0 {
		return out
	}
	if caseInsensitive {
		incl = strslices.ToLower(incl)
	}
	for _, idea := range ideas {
		statusCat := idea.StatusCategory
		if caseInsensitive {
			statusCat = strings.ToLower(statusCat)
		}
		if slices.Contains(incl, statusCat) {
			out = append(out, idea)
		}
	}
	return out
}
