package spahttpserver

import (
	"io"
	"io/fs"
	"net/http"
	"path/filepath"
	"time"
)

type SpaFileSystem struct {
	BaseFS    fs.FS
	IndexPath string
}

func (s SpaFileSystem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Clean the path to avoid directory traversal vulnerabilities
	cleanedPath := filepath.Clean(r.URL.Path)
	if cleanedPath == "/" || cleanedPath == "." {
		cleanedPath = s.IndexPath
	} else {
		// Strip the leading slash so it matches the embedded FS naming
		cleanedPath = cleanedPath[1:]
	}

	// Try to open the file in the embedded filesystem
	file, err := s.BaseFS.Open(cleanedPath)
	if err != nil {
		// Fallback: If the file is not found, serve the index.html file
		indexFile, err := s.BaseFS.Open(s.IndexPath)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer indexFile.Close()

		// Cast the file explicitly to io.ReadSeeker for http.ServeContent
		if seeker, ok := indexFile.(io.ReadSeeker); ok {
			http.ServeContent(w, r, s.IndexPath, time.Now(), seeker)
			return
		}

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// If the file exists, hand off execution to the standard FileServer
	http.FileServer(http.FS(s.BaseFS)).ServeHTTP(w, r)
}
