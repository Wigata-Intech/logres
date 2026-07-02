package slugify

import "strings"

func Slugify(s string) string {
	s = strings.ToLower(s)
	return strings.ReplaceAll(s, " ", "-")
}
