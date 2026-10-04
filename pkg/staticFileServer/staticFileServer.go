package staticFileServer

// binu@dronasys.com

import (
	"net/http"
	"path"
	"strings"

	"github.com/sirupsen/logrus"
)

func logRequestHandler(h http.Handler) http.Handler {
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

func logRequestSPAHandler(h http.Handler, staticDir string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logrus.Infof("SPA Serving file %s %s %s\n", r.RemoteAddr, r.Method, r.URL)

		_, err := http.Dir(staticDir).Open(r.URL.Path)
		if err != nil {
			// If file doesn't exist, serve index.html (for SPA routing)

			// double check if trying to access other folders with special characters
			if strings.ContainsAny(r.URL.Path, "/\\?%*:|\"<>") {
				logrus.Errorf("Serving file not found, path contains special char %s\n", r.URL.Path)
				http.ServeFile(w, r, staticDir+"/index.html")

				return
			}
			logrus.Errorf("Serving file not found  %s\n", r.URL.Path)
			http.ServeFile(w, r, staticDir+"/index.html")
			return
		}

		fileName := path.Base(r.URL.Path)
		if !isSafeFileName(fileName) {
			logrus.Errorf("SPA Static server, Invalid file name %s %s", fileName, r.URL.Path)
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
func SPAStaticServePath(gw *http.ServeMux, urlPath string, staticDir string) {

	// urlPath should end with slash
	if urlPath[len(urlPath)-1] != '/' {
		urlPath += "/"
	}

	// Create a handler for serving static files
	staticHandler := http.FileServer(http.Dir(staticDir))
	// handle SPA, if file is not found server index.html
	loggedHandler := logRequestSPAHandler(staticHandler, staticDir)

	// Strip the URL prefix and serve files
	gw.Handle(urlPath, http.StripPrefix(urlPath, loggedHandler))

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
	loggedHandler := logRequestHandler(staticHandler)

	// Strip the URL prefix and serve files
	gw.Handle(urlPath, http.StripPrefix(urlPath, loggedHandler))

	// log.Println("Server is listening on port ", listenAdd)
	logrus.Infof("Server directory %s -> %s", urlPath, staticDir)

}
