package handler

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed web/*
var webFiles embed.FS

func static() http.Handler {
	subtree, _ := fs.Sub(webFiles, "web")
	return http.FileServer(http.FS(subtree))
}
