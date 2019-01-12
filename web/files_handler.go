package web

import (
	"net/http"
	"strings"
)

// filesHandler represents handler for files directory
// that prevents to show the whole content of directory - allows to show only specified file
// eg. serving directory is "avatars"
// if pass in url /avatars/ - it will return NotFound status
// if pass in url /avatars/some-name.jpg - will return file if it exist
type filesHandler struct {
	path string
}

// ServeHTTP handles the HTTP request.
func (fh *filesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h := http.FileServer(http.Dir(fh.path))
	if strings.HasSuffix(r.URL.Path, "/") || r.URL.Path == "" {
		http.NotFound(w, r)
		return
	}
	h.ServeHTTP(w, r)
}

// NewFilesHandler creates new assetsHandler object that satisfy http.Handler interface
func NewFilesHandler(path string) http.Handler {
	return &filesHandler{
		path: path,
	}
}
