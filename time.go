package desk

import (
	"time"
)

const format string = "2006-01-02T15:04:05Z"

type Time struct{ *time.Time }

func (self *Time) UnmarshalJSON(data []byte) error {
	if len(data) <= 2 {
		return nil
	}
	str := string(data[1 : len(data)-1])
	t, err := time.Parse(format, str)
	if err != nil {
		return err
	}

	*self = Time{&t}
	return nil
}
