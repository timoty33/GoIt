package dev

import (
	dbs "github.com/bmatcuk/doublestar/v4"
)

func shouldIgnore(path string, ignores []string) bool {
	var match bool
	for _, ignore := range ignores {
		match, _ = dbs.PathMatch(ignore, path)
	}
	return match
}
