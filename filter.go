package main

import (
	"os"
	"regexp"
)

type filterer interface {
	Filter([]byte) []byte
}

func newFilter(mask string) filterer {
	if mask == "" {
		return &dummyFilter{}
	}

	return newRegexpFilter(mask)
}

func newRegexpFilter(mask string) filterer {
	re := regexp.MustCompile("(" + mask + ")")
	colored := os.Getenv("TERM") != ""

	return &regexpFilter{
		colored: colored,
		regexp:  re,
	}
}

type dummyFilter struct{}

func (filter dummyFilter) Filter(message []byte) []byte {
	return message
}

// Highlight matched value with red color.
var highlightReplacement = []byte("\x1b[31m$1\x1b[0m")

type regexpFilter struct {
	colored bool
	regexp  *regexp.Regexp
}

func (filter regexpFilter) Filter(message []byte) []byte {
	if filter.regexp.Match(message) {
		return filter.colorize(message)
	}

	return []byte{}
}

func (filter regexpFilter) colorize(message []byte) []byte {
	if filter.colored {
		return filter.regexp.ReplaceAll(message, highlightReplacement)
	}

	return message
}
