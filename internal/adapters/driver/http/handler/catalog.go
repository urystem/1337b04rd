package handler

import (
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
		slog.Error("exe")
		http.Error(w, "Execution error", http.StatusInternalServerError)
	}
}
