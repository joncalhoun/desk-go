package desk

import (
	"log"
	"regexp"
)

var replyIdRegexp = regexp.MustCompile(".*/replies/([0-9]+)$")

type Message struct {
	Id        int    `json:"id"`
	CreatedAt *Time  `json:"created_at"`
	UpdatedAt *Time  `json:"updated_at"`
	Body      string `json:"body"`
}

type RawReplies struct {
	Embedded struct {
		Entries []struct {
			Links struct {
				Case struct {
					Class string `json:"class"`
					Href  string `json:"href"`
				} `json:"case"`
				Customer struct {
					Class string `json:"class"`
					Href  string `json:"href"`
				} `json:"customer"`
				HiddenBy interface{} `json:"hidden_by"`
				Self     struct {
					Class string `json:"class"`
					Href  string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
			Bcc        interface{} `json:"bcc"`
			Body       string      `json:"body"`
			Cc         interface{} `json:"cc"`
			ClientType string      `json:"client_type"`
			CreatedAt  *Time       `json:"created_at"`
			Direction  string      `json:"direction"`
			From       string      `json:"from"`
			Hidden     bool        `json:"hidden"`
			HiddenAt   *Time       `json:"hidden_at"`
			Status     string      `json:"status"`
			Subject    string      `json:"subject"`
			To         string      `json:"to"`
			UpdatedAt  *Time       `json:"updated_at"`
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

func ParseRawReplies(rawReplies *RawReplies) []Message {
	replies := make([]Message, len(rawReplies.Embedded.Entries), int(rawReplies.TotalEntries))

	for i, entry := range rawReplies.Embedded.Entries {
		id, err := ParseId(entry.Links.Self.Href, replyIdRegexp, 1)
		if err != nil {
			log.Printf("Failed to parse Id for Reply: %+v", entry)
		} else {
			replies[i].Id = id
		}
		replies[i].CreatedAt = entry.CreatedAt
		replies[i].UpdatedAt = entry.UpdatedAt
		replies[i].Body = entry.Body
	}

	return replies
}
