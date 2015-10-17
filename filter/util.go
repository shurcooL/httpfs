package filter

import (
	"os"
	"path/filepath"
)

// Extensions returns a filter function that ignores any of the given file
// extensions.
func Extensions(exts ...string) Func {
	return func(fi os.FileInfo, path string) bool {
		for _, ext := range exts {
			if filepath.Ext(path) == ext {
				return true
			}
		}
		return false
	}
}

// Not negates the given filter, such that:
//
//  Not(Extensions(".go", ".html"))
//
// would ignore any file that is not a .go or .html file.
func Not(filter Func) Func {
	return func(fi os.FileInfo, path string) bool {
		return !filter(fi, path)
	}
}

// Combine combines multiple filter functions into one. Effectively it is an
// multiple or operator, that is:
//
//  Combine(Extensions(".go"), Extensions(".html"))
//
// would ignore both .go and .html files.
func Combine(filters ...Func) Func {
	return func(fi os.FileInfo, path string) bool {
		for _, filter := range filters {
			if filter(fi, path) {
				return true
			}
		}
		return false
	}
}
