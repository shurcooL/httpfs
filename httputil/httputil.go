// Package httputil implements HTTP utility functions for http.FileSystem.
package httputil

import (
	"log"
	"net/http"
	"time"
)

// TODO: Decide if this is a good idea/direction to go in.

// FileHandler is an http.Handler that serves the root of File.
type FileHandler struct {
	File        http.FileSystem
	ContentType string
	Name        string // If ContentType is not set, file extension of Name is used to determine content type.
}

func (h FileHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	f, err := h.File.Open("/")
	if err != nil {
		log.Printf("FileHandler.File.Open('/'): %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()
	if h.ContentType != "" {
		w.Header().Set("Content-Type", h.ContentType)
	}
	http.ServeContent(w, req, h.Name, time.Now(), f)
}
