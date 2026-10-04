package staticFileServer

// binu@dronasys.com

import (
	"io/fs"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/sirupsen/logrus"
)

func LogRequestHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logrus.Infof("AI File Serving file %s %s %s\n", r.RemoteAddr, r.Method, r.URL)

		fileName := path.Base(r.URL.Path)
		if !isSafeFileName(fileName) {
			logrus.Errorf("Static server, Invalid file name %s %s", fileName, r.URL.Path)
			http.Error(w, "Invalid file name", http.StatusBadRequest)
			return
		}
		h.ServeHTTP(w, r)

	})
}

func LogRequestSPAHandler(h http.Handler, staticFileServer fs.FS, staticDir string) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logrus.Infof("SPA Request file  %s %s (%s%s)\n", r.Method, r.RemoteAddr, staticDir, r.URL.Path)

		// Normalize root
		// if r.URL.Path == "" || r.URL.Path == "/" {
		// 	logrus.Infof("Serving index.html for root path")
		// 	r2 := r.Clone(r.Context())
		// 	r2.URL = &url.URL{Path: "/index.html"}
		// 	h.ServeHTTP(w, r2)
		// 	return
		// }

		// Clean path to avoid simple traversal
		cleanPath := path.Clean(r.URL.Path)
		if cleanPath == "" || cleanPath == "." || cleanPath == "/" {
			logrus.Infof("Normalized path to root, serving index.html [%s]", cleanPath)
			r2 := r.Clone(r.Context())
			r2.URL = &url.URL{Path: "/index.html"}
			h.ServeHTTP(w, r2)
			return
		}

		// Try to open the file in the embedded FS
		f, err := staticFileServer.Open(cleanPath)
		if err != nil {
			// File doesn't exist: SPA fallback to index.html
			logrus.Infof("File not found in embed FS, serving index.html: %s, %v", cleanPath, err)
			r2 := r.Clone(r.Context())
			r2.URL = &url.URL{Path: "/index.html"}
			h.ServeHTTP(w, r2)
			return
		}
		_ = f.Close()

		// Optional: safety check on base name
		fileName := path.Base(cleanPath)
		if !isSafeFileName(fileName) {
			logrus.Errorf("SPA Static server: invalid file name %q (path=%s)", fileName, cleanPath)
			http.Error(w, "Invalid file name", http.StatusBadRequest)
			return
		}

		h.ServeHTTP(w, r)

	})
}

func isSafeFileName(fileName string) bool {
	validExtensions := []string{
		".", // for index.html
		".js", ".css", ".html",
		".ico", ".jpg", ".png", ".jpeg",
		".mp3", ".wav", ".flac",
		".ttf", ".woff2"}
	if strings.ContainsAny(fileName, "/\\?%*:|\"<>") {
		return false
	}
	fileName = strings.ToLower(fileName)
	for _, ext := range validExtensions {
		if strings.HasSuffix(fileName, ext) {
			return true
		}
	}
	return false
}

// gw http Server mux
// urlPath - the url path, which the file server listents to eg: /, /static, /assets, /public
// staticDir - actual directory from which the files are served
func SPAStaticServePath(httpMux *http.ServeMux, urlPath string, embedFS fs.FS, staticDir string) {

	// Optional: Use fs.Sub to strip the "ui" prefix
	// so files are served directly from the root URL "/"
	staticContent, err := fs.Sub(embedFS, staticDir)
	if err != nil {
		panic(err)
	}

	// dump all embed files
	entries, err := fs.ReadDir(staticContent, ".")
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		logrus.Printf("Static file %s", entry)
	}

	// urlPath should end with slash
	if urlPath[len(urlPath)-1] != '/' {
		urlPath += "/"
	}

	// Create a handler for serving static files
	staticFileServer := http.FileServer(http.FS(staticContent))
	// handle SPA, if file is not found serve index.html
	//spaHandler := LogRequestSPAHandler(staticFileServer, staticContent, staticDir)

	// Strip the URL prefix and serve files
	// httpMux.Handle(urlPath, http.StripPrefix(urlPath, spaHandler))
	httpMux.Handle(urlPath, staticFileServer)
	// httpMux.Handle(urlPath, http.StripPrefix(urlPath, staticFileServer))
	// log.Println("Server is listening on port ", listenAdd)
	logrus.Infof("Server directory %s -> %s", urlPath, staticDir)

}

// gw http Server mux
// serves Ai images and other generated resources
// SPA is not active here
// urlPath - the url path, which the file server listents to eg: /, /static, /assets, /public
// staticDir - actual directory from which the files are served
func StaticServeAiPath(gw *http.ServeMux, urlPath string, staticDir string) {

	// urlPath should end with slash
	if urlPath[len(urlPath)-1] != '/' {
		urlPath += "/"
	}

	// Create a handler for serving static files
	staticHandler := http.FileServer(http.Dir(staticDir))
	loggedHandler := LogRequestHandler(staticHandler)

	// Strip the URL prefix and serve files
	gw.Handle(urlPath, http.StripPrefix(urlPath, loggedHandler))

	// log.Println("Server is listening on port ", listenAdd)
	logrus.Infof("Server directory %s -> %s", urlPath, staticDir)

}
