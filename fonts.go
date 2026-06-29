package main

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed assets/fonts
var fonts embed.FS

func fontsHandler() http.Handler {
	sub, _ := fs.Sub(fonts, "assets/fonts")
	return http.StripPrefix("/fonts/", http.FileServer(http.FS(sub)))
}