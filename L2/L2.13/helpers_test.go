package main

import (
	"strings"
)

func parseFieldSpecMust(spec string) fieldSet {
	fs, err := parseFields(spec)
	if err != nil {
		panic(err)
	}
	return fs
}

func join(parts ...string) string { return strings.Join(parts, "\n") }
