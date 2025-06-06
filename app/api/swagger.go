package api

import (
	"embed"
	"net/http"
	"path/filepath"
	"runtime"
)

//go:embed swagger/*
var swaggerFiles embed.FS

func (api *Application) SwaggerHandler(w http.ResponseWriter, r *http.Request) {
	// Get the absolute path to the swagger directory
	_, currentFile, _, _ := runtime.Caller(0)
	swaggerDir := filepath.Join(filepath.Dir(currentFile), "swagger")

	// Set content type
	w.Header().Set("Content-Type", "text/html")

	// Serve the Swagger UI HTML file
	http.ServeFile(w, r, filepath.Join(swaggerDir, "swagger-ui.html"))
}

func (api *Application) SwaggerJSONHandler(w http.ResponseWriter, r *http.Request) {
	// Get the absolute path to the swagger directory
	_, currentFile, _, _ := runtime.Caller(0)
	swaggerDir := filepath.Join(filepath.Dir(currentFile), "swagger")

	// Set content type
	w.Header().Set("Content-Type", "application/json")

	// Serve the Swagger JSON file
	http.ServeFile(w, r, filepath.Join(swaggerDir, "swagger.json"))
}

func SwaggerUI(w http.ResponseWriter, r *http.Request) {
	// Get the absolute path to the swagger directory
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(filepath.Dir(filepath.Dir(b)))
	swaggerPath := filepath.Join(basepath, "swagger", "swagger-ui.html")

	// Set content type
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, swaggerPath)
}

func SwaggerJSON(w http.ResponseWriter, r *http.Request) {
	// Get the absolute path to the swagger directory
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(filepath.Dir(filepath.Dir(b)))
	swaggerPath := filepath.Join(basepath, "swagger", "swagger.json")

	// Set content type
	w.Header().Set("Content-Type", "application/json")
	http.ServeFile(w, r, swaggerPath)
}
