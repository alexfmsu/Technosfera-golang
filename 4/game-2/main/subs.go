package main

import (
	"strings"
)

func join(sep string, s []string) string {
	var joined []string

	for i := range s {
		joined = append(joined, s[i])
	}

	return strings.Join(joined, sep)
}
