package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"text/template"
)

func (h *handler) Catalog(w http.ResponseWriter, r *http.Request) {
	posts, err := h.use.ListOfPosts(r.Context())
	if err != nil {
		// http.Error(w, "Error parsing template", http.StatusInternalServerError)
		slog.Error(err.Error())
		http.Error(w, "Error loading posts", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("web/templates/catalog.html")
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, posts)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Execution error", http.StatusInternalServerError)
	}
}

func (h *handler) ServeImage(w http.ResponseWriter, r *http.Request) {
	// Получаем имя файла из URL
	imageName := r.PathValue("image")
	fmt.Println("serveimage", imageName)
	if imageName == "" {
		http.Error(w, "missing image name", http.StatusBadRequest)
		return
	}

	// Получаем объект из MinIO
	obj, err := h.use.GetPostImage(r.Context(), imageName)
	if err != nil {
		slog.Error("get object:", err)
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	defer obj.Close()

	w.Header().Set("Content-Type", obj.ConType)

	http.ServeContent(w, r, "", obj.Modified, obj)
}
