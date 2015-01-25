package desk

import (
	"strconv"
	"net/url"
)

type ListParams struct {
	Page, PerPage int
	SortField, SortDirection string
}

func (lp ListParams) UrlValues(urlValues *url.Values) {
	if lp.Page > 0 {
		urlValues.Add("page", strconv.Itoa(lp.Page))
	}

	if lp.PerPage > 0 {
		urlValues.Add("per_page", strconv.Itoa(lp.PerPage))
	}

	if lp.SortField != "" {
		urlValues.Add("sort_field", lp.SortField)
	}

	if lp.SortDirection != "" {
		urlValues.Add("sort_direction", lp.SortDirection)
	}
}
