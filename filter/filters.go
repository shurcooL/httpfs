package filter

import (
	"os"
	"path/filepath"
)

// Extensions returns a filter function that ignores any of the given file
// extensions.
func Extensions(exts ...string) Func {
	return func(path string, _ os.FileInfo) bool {
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
// 	Not(Extensions(".go", ".html"))
//
// would ignore any file that is not a .go or .html file.
func Not(filter Func) Func {
	return func(path string, fi os.FileInfo) bool {
		return !filter(path, fi)
	}
}

// Combine combines multiple filter functions into one. Effectively it is an
// multiple OR operator, that is:
//
// 	Combine(Extensions(".go"), Extensions(".html"))
//
// would ignore both .go and .html files.
func Combine(filters ...Func) Func {
	return func(path string, fi os.FileInfo) bool {
		for _, filter := range filters {
			if filter(path, fi) {
				return true
			}
		}
		return false
	}
}
