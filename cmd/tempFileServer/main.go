package main

import (
	"embed"
	"io/fs"
	"net/http"

	spahttpserver "github.com/binuud/watchdog/pkg/spaHttpServer"
)

//go:embed ui/*
var embeddedFiles embed.FS

func main() {
	// 1. Strip the "dist" prefix from the embedded paths
	// so you don't have to include "/dist/" in your URLs.
	distFS, err := fs.Sub(embeddedFiles, "ui")
	if err != nil {
		panic(err)
	}

	// 2. Initialize our SPA handler wrapper
	handler := spahttpserver.SpaFileSystem{
		BaseFS:    distFS,
		IndexPath: "index.html",
	}

	http.Handle("/", handler)

	println("Server running smoothly at http://localhost:9080")
	if err := http.ListenAndServe(":9080", nil); err != nil {
		panic(err)
	}
}
