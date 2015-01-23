package desk

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func ParseId(str string, regexp *regexp.Regexp, expectedIndex int) (int, error) {
	m := regexp.FindStringSubmatch(str)
	if m == nil || len(m) <= expectedIndex {
		return -1, errors.New(fmt.Sprintf("Failed to parse ID. String: %s, Regexp: %v", str, regexp))
	}
	id, err := strconv.Atoi(m[expectedIndex])
	return id, err
}
