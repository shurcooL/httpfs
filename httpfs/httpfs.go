// Package httpfs implements http.FileSystem using a webdav.FileSystem.
package httpfs

import (
	"net/http"
	"os"

	"golang.org/x/net/webdav"
)

// New returns an http.FileSystem adapter for the provided webdav.FileSystem.
func New(fs webdav.FileSystem) http.FileSystem {
	return &httpFS{fs}
}

type httpFS struct {
	fs webdav.FileSystem
}

func (h *httpFS) Open(name string) (http.File, error) {
	return h.fs.OpenFile(name, os.O_RDONLY, 0)
}
