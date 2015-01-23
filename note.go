package desk

import (
	"log"
	"regexp"
)

var noteIdRegexp = regexp.MustCompile(".*/notes/([0-9]+)$")

type Note struct {
	Id        int
	CreatedAt *Time  `json:"created_at"`
	UpdatedAt *Time  `json:"updated_at"`
	Body      string `json:"body"`
}

// This is used to parse json from the Desk API, and then translated into a more usable []Note.
type RawNotes struct {
	Embedded struct {
		Entries []struct {
			Links struct {
				Case struct {
					Class string `json:"class"`
					Href  string `json:"href"`
				} `json:"case"`
				ErasedBy interface{} `json:"erased_by"`
				Self     struct {
					Class string `json:"class"`
					Href  string `json:"href"`
				} `json:"self"`
				User struct {
					Class string `json:"class"`
					Href  string `json:"href"`
				} `json:"user"`
			} `json:"_links"`
			Body      string `json:"body"`
			CreatedAt *Time  `json:"created_at"`
			ErasedAt  *Time  `json:"erased_at"`
			UpdatedAt *Time  `json:"updated_at"`
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

func ParseRawNotes(rawNotes *RawNotes) []Note {
	notes := make([]Note, len(rawNotes.Embedded.Entries), int(rawNotes.TotalEntries))

	for i, entry := range rawNotes.Embedded.Entries {
		id, err := ParseId(entry.Links.Self.Href, noteIdRegexp, 1)
		if err != nil {
			log.Printf("Failed to parse Id for Note: %+v", entry)
		} else {
			notes[i].Id = id
		}
		notes[i].CreatedAt = entry.CreatedAt
		notes[i].UpdatedAt = entry.UpdatedAt
		notes[i].Body = entry.Body
	}

	return notes
}
