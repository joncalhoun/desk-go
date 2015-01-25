package desk

import (
	"log"
	"regexp"
	"net/url"
)

var caseIdRegexp = regexp.MustCompile(".*/cases/([0-9]+)$")

type Case struct {
	Id        int       `json:"id"`
	Subject   string    `json:"subject"`
	CreatedAt *Time     `json:"created_at"`
	UpdatedAt *Time     `json:"updated_at"`
	Message   *Message  `json:"message"`
	Notes     []Note    `json:"notes"`
	Replies   []Message `json:"replies"`
}

type CaseListParams struct {
	ListParams
}

func (c CaseListParams) UrlValues() *url.Values {
	ret := url.Values{}
	c.ListParams.UrlValues(&ret)
	return &ret
}


type CaseSearchParams struct {
	// If q is provided, it will be used exclusively and all other search params will be ignored.
	// See http://dev.desk.com/API/cases/#search for more info
	Q string
	// These are ignored if Q has any value
	// See http://dev.desk.com/API/cases/#search under the heading "Other Search Parameters"
	Options map[string]string
	ListParams
}

func (c CaseSearchParams) UrlValues() *url.Values {
	ret := url.Values{}

	if c.Q != "" {
		ret.Add("q", c.Q)
	} else if len(c.Options) > 0 {
		for key, value := range c.Options {
			ret.Add(key, value)
		}
	}

	c.ListParams.UrlValues(&ret)
	return &ret
}

func (c *CaseSearchParams) AddOption(key, value string) {
	if c.Options == nil {
		c.Options = make(map[string]string)
	}
	c.Options[key] = value
}

type RawCases struct {
	Embedded struct {
		Entries []struct {
			Links struct {
				AssignedGroup struct {
					Class string `json:"class"`
					Href  string `json:"href"`
				} `json:"assigned_group"`
				AssignedUser struct {
					Class string `json:"class"`
					Href  string `json:"href"`
				} `json:"assigned_user"`
				Attachments struct {
					Class string `json:"class"`
					Href  string `json:"href"`
				} `json:"attachments"`
				Customer struct {
					Class string `json:"class"`
					Href  string `json:"href"`
				} `json:"customer"`
				Draft struct {
					Class string `json:"class"`
					Href  string `json:"href"`
				} `json:"draft"`
				History struct {
					Class string `json:"class"`
					Href  string `json:"href"`
				} `json:"history"`
				LockedBy     interface{} `json:"locked_by"`
				MacroPreview struct {
					Class string `json:"class"`
					Href  string `json:"href"`
				} `json:"macro_preview"`
				Message struct {
					Class string `json:"class"`
					Href  string `json:"href"`
				} `json:"message"`
				Notes struct {
					Class string `json:"class"`
					Href  string `json:"href"`
				} `json:"notes"`
				Replies struct {
					Class string `json:"class"`
					Href  string `json:"href"`
				} `json:"replies"`
				Self struct {
					Class string `json:"class"`
					Href  string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
			ActiveAt     *Time       `json:"active_at"`
			Blurb        interface{} `json:"blurb"`
			ChangedAt    *Time       `json:"changed_at"`
			CreatedAt    *Time       `json:"created_at"`
			CustomFields struct {
				Level string `json:"level"`
			} `json:"custom_fields"`
			Description     interface{}   `json:"description"`
			ExternalID      interface{}   `json:"external_id"`
			FirstOpenedAt   *Time         `json:"first_opened_at"`
			FirstResolvedAt *Time         `json:"first_resolved_at"`
			ID              float64       `json:"id"`
			LabelIds        []interface{} `json:"label_ids"`
			Labels          []interface{} `json:"labels"`
			Language        string        `json:"language"`
			LockedUntil     interface{}   `json:"locked_until"`
			OpenedAt        *Time         `json:"opened_at"`
			Priority        float64       `json:"priority"`
			ReceivedAt      *Time         `json:"received_at"`
			ResolvedAt      *Time         `json:"resolved_at"`
			Status          string        `json:"status"`
			Subject         string        `json:"subject"`
			Type            string        `json:"type"`
			UpdatedAt       *Time         `json:"updated_at"`
		} `json:"entries"`
	} `json:"_embedded"`
	Links struct {
		First struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"first"`
		Last struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"last"`
		Next     interface{} `json:"next"`
		Previous interface{} `json:"previous"`
		Self     struct {
			Class string `json:"class"`
			Href  string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Page         float64 `json:"page"`
	TotalEntries float64 `json:"total_entries"`
}

func ParseRawCases(rawCases *RawCases) []Case {
	cases := make([]Case, len(rawCases.Embedded.Entries), int(rawCases.TotalEntries))

	for i, entry := range rawCases.Embedded.Entries {
		id, err := ParseId(entry.Links.Self.Href, caseIdRegexp, 1)
		if err != nil {
			log.Printf("Failed to parse Id for Case: %+v", entry)
		} else {
			cases[i].Id = id
		}
		cases[i].CreatedAt = entry.CreatedAt
		cases[i].UpdatedAt = entry.UpdatedAt
	}

	return cases
}
