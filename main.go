package main

import (
	"embed"
	"io/fs"
	"net/http"
	"log"
)

//go:embed html/*
var html embed.FS

func main() {
	htmlContent, err := fs.Sub(html, "html")
	if err != nil {
		log.Fatal(err)
	}

	// serve the html directory by default
	http.Handle("/", http.FileServer(http.FS(htmlContent)))
	http.ListenAndServe(":8080", nil)
}
